package http

import (
	"go-eth-todo-list/api/common"
	"go-eth-todo-list/internal/data/model"
	"go-eth-todo-list/internal/pkg/abi/todolist"
	"go-eth-todo-list/internal/pkg/constant"
	"go-eth-todo-list/internal/pkg/errors"
	"go-eth-todo-list/internal/pkg/util"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo"
)

type UpdateIntegrator struct {
	conn   *todolist.Todolist
	auth   *bind.TransactOpts
	client *ethclient.Client
}

func UpdateApi(e *echo.Echo) {
	integrator := UpdateIntegrator{
		conn:   common.Conn,
		auth:   common.Auth,
		client: common.Client,
	}

	e.POST("/update", integrator.UpdateHandler)
}

func (integrator UpdateIntegrator) UpdateHandler(ctx echo.Context) error {
	req := model.UpdateRequest{}
	err := ctx.Bind(&req)

	if err != nil {
		return errors.ErrInvalidFieldFormat("invalid query parameter", err)
	}

	if req.TaskId.Cmp(big.NewInt(0)) == 0 {
		return ctx.JSON(400, model.UpdateResponse{
			ErrorString: "task id cannot be empty",
			Success:     false,
		})
	}

	//Toggle task as completed
	nonce, _ := util.GetNonceAndKeys(integrator.client, constant.WalletDemoKey)
	integrator.auth.Nonce = big.NewInt(int64(nonce))
	integrator.auth.GasPrice = big.NewInt(493945075) //Block's base fee per gas
	_, err = integrator.conn.ToggleCompleted(integrator.auth, req.TaskId)
	if err != nil {
		return err
	}

	//Have the frontend reload
	return ctx.JSON(205, model.UpdateResponse{
		Success: true,
	})
}
