package todos

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

var allTodos = []Todo{}

// GetTodos godoc
// @Summary Get All Data Todos
// @Description Get All Todos
// @Tags Todo
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessGet
// @Router /todos [get]
func GetTodos(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "success",
		"todos":  allTodos,
	})
}

// GetTodosByID godoc
// @Summary Get Todos By ID
// @Description Get Todos By ID
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path int32 true "ID Todo"
// @Success 200 {object} response.SuccessGet
// @Router /todos/{id} [get]
func GetTodosByID(c *gin.Context) {
	stringId := c.Param("id")
	Id, err := strconv.Atoi(stringId)

	if err != nil {
		c.JSON(400, gin.H{
			"status": "error",
			"error":  err,
		})
		return
	}

	for i, todo := range allTodos {
		if todo.ID == Id {
			c.JSON(200, gin.H{
				"status": "success",
				"todos":  allTodos[i],
			})
			break
		} else if i == len(allTodos)-1 {
			c.JSON(400, gin.H{
				"status":   "error",
				"messages": fmt.Sprintf("ID Not found %v", Id),
			})
			return
		}
	}
}

// AddTodo godoc
// @Summary Post New Data Todos
// @Description Post New Data Todos
// @Tags Todo
// @Accept json
// @Produce json
// @Param todo body todos.CreateTodo true "Todos"
// @Success 200 {object} response.SuccessCreate
// @Router /todos [post]
func AddTodos(c *gin.Context) {
	inputTodos := Todo{}

	err := c.ShouldBindJSON(&inputTodos)

	if err != nil {
		c.JSON(400, gin.H{
			"status": "error",
			"error":  err,
		})
		return
	}

	// get ID from last Todos
	lengthTodos := len(allTodos)

	if lengthTodos == 0 {
		inputTodos.ID = 1
	} else {
		lastTodos := allTodos[lengthTodos-1]
		inputTodos.ID = lastTodos.ID + 1
	}

	allTodos = append(allTodos, inputTodos)

	c.JSON(200, gin.H{
		"status":   "success",
		"messages": "Success add new todos",
		"data":     inputTodos,
	})
}

// DeleteTodo godoc
// @Summary Delete Data Todos
// @Description Delete Data Todos
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path int32 true "ID Todo"
// @Success 200 {object} response.SuccessDelete
// @Router /todos/{id} [delete]
func DeleteTodos(c *gin.Context) {
	stringId := c.Param("id")
	Id, err := strconv.Atoi(stringId)

	if err != nil {
		c.JSON(400, gin.H{
			"status": "error",
			"error":  err,
		})
		return
	}

	for i, todo := range allTodos {
		if todo.ID == Id {
			allTodos = append(allTodos[:i], allTodos[i+1:]...)
			break
		}
		if i == len(allTodos)-1 {
			c.JSON(400, gin.H{
				"status":   "error",
				"messages": fmt.Sprintf("ID Not found %v", Id),
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"status":   "success",
		"messages": fmt.Sprintf("Success delete todos with id %v", Id),
	})
}

// UpdateTodo godoc
// @Summary Put Edit Data Todos
// @Description Edit Data Todos
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path int32 true "ID Todo"
// @Param todo body todos.CreateTodo true "Update"
// @Success 200 {object} response.SuccessUpdate
// @Router /todos/{id} [put]
func UpdateTodos(c *gin.Context) {
	stringId := c.Param("id")
	Id, err := strconv.Atoi(stringId)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "error",
			"error":  err,
		})
		return
	}

	var updatedTodo Todo
	err = c.ShouldBindJSON(&updatedTodo)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "error",
			"error":  err,
		})
		return
	}

	for i, todo := range allTodos {
		if todo.ID == Id {
			updatedTodo.ID = Id
			temp := allTodos[i+1:]
			allTodos = append(allTodos[:i], updatedTodo)
			allTodos = append(allTodos, temp...)
			break
		}
		if i == len(allTodos)-1 {
			c.JSON(400, gin.H{
				"status":   "error",
				"messages": fmt.Sprintf("ID not found %v", Id),
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"status":   "success",
		"messages": fmt.Sprintf("Success update todos with id %v", Id),
	})
}
