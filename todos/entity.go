package todos

type Todo struct {
	ID       int    `json:"id" example:"1"`
	Name     string `json:"name" example:"Makan"`
	Complete bool   `json:"complete" example:"false"`
}

type CreateTodo struct {
	Name     string `json:"name" example:"Makan"`
	Complete bool   `json:"complete" example:"false"`
}

type Todos []Todo
