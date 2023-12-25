package response

type TaskResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Note   string `json:"note"`
	Status string `json:"status"`
	Date   string `json:"date"`
}
