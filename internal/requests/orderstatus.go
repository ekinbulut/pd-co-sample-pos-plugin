package requests

type OrderStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
