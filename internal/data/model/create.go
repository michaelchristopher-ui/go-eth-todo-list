package model

type CreateRequest struct {
	TaskName string `json:"taskname"`
}

type CreateResponse struct {
	ErrorString string `json:"error,omitempty"`
	Success     bool   `json:"success"`
}
