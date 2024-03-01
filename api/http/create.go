package http

import (
	"errors"
	"go-eth-todo-list/api/common"
	"go-eth-todo-list/internal/data/model"
	"go-eth-todo-list/internal/pkg/abi/todolist"
	"go-eth-todo-list/internal/pkg/constant"
	internalErr "go-eth-todo-list/internal/pkg/errors"
	"go-eth-todo-list/internal/pkg/util"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo"
)

type CreateIntegrator struct {
	conn   *todolist.Todolist
	auth   *bind.TransactOpts
	client *ethclient.Client
}

func CreateApi(e *echo.Echo) {
	integrator := CreateIntegrator{
		conn:   common.Conn,
		auth:   common.Auth,
		client: common.Client,
	}

	e.POST("/create", integrator.CreateHandler)
}

func (integrator CreateIntegrator) CreateHandler(ctx echo.Context) error {
	//Get value from parameter
	req := model.CreateRequest{}
	err := ctx.Bind(&req)

	if req.TaskName == "" {
		return ctx.JSON(400, model.CreateResponse{
			ErrorString: errors.New("task name cannot be empty").Error(),
			Success:     false,
		})
	}

	if err != nil {
		return ctx.JSON(400, model.CreateResponse{
			ErrorString: internalErr.ErrInvalidFieldFormat("invalid query parameter", err).Error(),
			Success:     false,
		})
	}

	//Create task
	nonce, _ := util.GetNonceAndKeys(integrator.client, constant.GetWalletKey())
	integrator.auth.Nonce = big.NewInt(int64(nonce))
	_, err = integrator.conn.CreateTask(integrator.auth, req.TaskName)
	if err != nil {
		return err
	}

	//Have the frontend reload
	return ctx.JSON(205, model.CreateResponse{
		Success: true,
	})
}
