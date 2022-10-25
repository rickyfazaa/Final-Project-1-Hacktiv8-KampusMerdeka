package response

import "Project1/todos"

type SuccessGet struct {
	Status string      `json:"status" example:"success"`
	Todos  todos.Todos `json:"todos"`
}

type SuccessCreate struct {
	Status   string     `json:"status" example:"success"`
	Messages string     `json:"messages" example:"Success add new todos"`
	Data     todos.Todo `json:"data"`
}

type SuccessDelete struct {
	Status   string `json:"status" example:"success"`
	Messages string `json:"messages" example:"Success delete todos with id 2"`
}

type SuccessUpdate struct {
	Status   string `json:"status" example:"success"`
	Messages string `json:"messages" example:"Success update todos with id 2"`
}
