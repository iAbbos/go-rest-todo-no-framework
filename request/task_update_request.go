package request

type TaskUpdateRequest struct {
	ID     int
	Name   string `validate:"required min=1,max=100" json:"name"`
	Note   string `json:"note"`
	Status string `validate:"required" json:"status"`
	Date   string `validate:"required" json:"date"`
}
