package controller

type HttpResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}
