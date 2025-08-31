package handler

import (
	"net/http"
	"strconv"

	"github.com/NickSarychev/todo-app"
	"github.com/gin-gonic/gin"
)

// @Summary Create item
// @Security ApiKeyAuth
// @Tags items
// @Description create todo item
// @ID create-item
// @Accept  json
// @Produce  json
// @Param id path int true "List ID"
// @Param input body todo.TodoItem true "item info"
// @Success 200 {integer} map[string]int"id of created item"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/lists/{id}/items [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	var input todo.TodoItem

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get all items
// @Security ApiKeyAuth
// @Tags items
// @Description get all todo items for a list
// @ID get-all-items
// @Accept  json
// @Produce  json
// @Param id path int true "List ID"
// @Success 200 {array} todo.TodoItem
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/lists/{id}/items [get]
func (h *Handler) getAllItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, items)
}

// @Summary Get item by id
// @Security ApiKeyAuth
// @Tags items
// @Description get todo item by id
// @ID get-item-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Success 200 {object} todo.TodoItem
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/items/{id} [get]
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

// @Summary Update item
// @Security ApiKeyAuth
// @Tags items
// @Description update todo item
// @ID update-item
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Param input body todo.UpdateItemInput true "update info"
// @Success 200 {object} StatusResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/items/{id} [put]
func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	var input todo.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoItem.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}

// @Summary Delete item
// @Security ApiKeyAuth
// @Tags items
// @Description delete todo item
// @ID delete-item
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Success 200 {object} StatusResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/items/{id} [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, StatusResponse{Status: "ok"})

}
