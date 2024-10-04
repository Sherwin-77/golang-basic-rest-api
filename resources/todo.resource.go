package resources

import "github.com/sherwin-77/golang-basic-rest-api/models"

type TodoResource struct {
}

func (tr *TodoResource) mapResource(todo models.Todo) interface{} {
	return todo
}

func (tr *TodoResource) Make(todo models.Todo) JsonResponse {
	return JsonResponse{
		Data:    tr.mapResource(todo),
		Message: "Success",
	}
}

func (tr *TodoResource) Collections(todos []models.Todo) JsonResponse {
	var data = []interface{}{}
	for _, todo := range todos {
		data = append(data, tr.mapResource(todo))
	}

	return JsonResponse{
		Data:    data,
		Message: "Success",
	}
}
