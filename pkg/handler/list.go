package handler

import (
	"net/http"
	"strconv"

	"github.com/NickSarychev/todo-app"
	"github.com/gin-gonic/gin"
)

// @Summary Create todo List
// @Security ApiKeyAuth
// @Tags lists
// @Description create todo List
// @ID create-list
// @Accept json
// @Produce json
// @Param input body todo.TodoList true "list info"
// @Success 200 {integer} map[string]int "id of created list"
// @Failure 400,404 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Failure default {object} handler.ErrorResponse
// @Router /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get all lists
// @Security ApiKeyAuth
// @Tags lists
// @Description get all todo lists
// @ID get-all-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} handler.GetAllListsResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /api/lists [get]
func (h *Handler) getAllList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetAllListsResponse{
		Data: lists,
	})

}

// @Summary Get list by id
// @Security ApiKeyAuth
// @Tags lists
// @Description get todo list by id
// @ID get-list-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "List ID"
// @Success 200 {object} todo.TodoList
// @Failure 400 {object} handler.ErrorResponse
// @Failure 404 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /api/lists/{id} [get]
func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	items, err := h.services.TodoList.GetById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary Update list
// @Security ApiKeyAuth
// @Tags lists
// @Description update todo list
// @ID update-list
// @Accept  json
// @Produce  json
// @Param id path int true "List ID"
// @Param input body todo.UpdateListInput true "update info"
// @Success 200 {object} handler.StatusResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 404 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /api/lists/{id} [put]
func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	var input todo.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoList.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}

// @Summary Delete list
// @Security ApiKeyAuth
// @Tags lists
// @Description delete todo list
// @ID delete-list
// @Accept  json
// @Produce  json
// @Param id path int true "List ID"
// @Success 200 {object} handler.StatusResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 404 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /api/lists/{id} [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TodoList.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
