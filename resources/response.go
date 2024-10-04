package resources

type JsonResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type Resource interface {
	mapResource() interface{}
	Make() JsonResponse
	Collections() JsonResponse
}
