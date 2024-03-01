package model

import "math/big"

type UpdateRequest struct {
	TaskId *big.Int `json:"taskid"`
}

type UpdateResponse struct {
	ErrorString string `json:"error,omitempty"`
	Success     bool   `json:"success"`
}
