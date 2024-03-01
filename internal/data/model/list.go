package model

import "math/big"

type ListResponse struct {
	Tasks []ListTasksResponse
}

type ListTasksResponse struct {
	TaskId    *big.Int `json:"taskid"`
	Content   string   `json:"string"`
	Completed bool     `json:"completed"`
}
