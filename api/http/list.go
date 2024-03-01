package http

import (
	"go-eth-todo-list/api/common"
	"go-eth-todo-list/internal/data/model"
	"go-eth-todo-list/internal/pkg/abi/todolist"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/labstack/echo"
)

type ListIntegrator struct {
	conn *todolist.Todolist
}

func ListApi(e *echo.Echo) {
	integrator := ListIntegrator{
		conn: common.Conn,
	}
	e.POST("/", integrator.ListHandler)
}

func (integrator ListIntegrator) ListHandler(ctx echo.Context) error {
	count, err := integrator.conn.TaskCount(&bind.CallOpts{})
	if err != nil {
		return err
	}
	one := big.NewInt(1)
	resp := model.ListResponse{}
	for i := big.NewInt(1); i.Cmp(count) <= 0; i.Add(i, one) {
		task, err := integrator.conn.Tasks(&bind.CallOpts{}, i)
		if err != nil {
			return err
		}
		resp.Tasks = append(resp.Tasks, model.ListTasksResponse{
			TaskId:    task.Id,
			Content:   task.Content,
			Completed: task.Completed,
		})
	}
	return ctx.JSON(200, resp)
}
