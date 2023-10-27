package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Body    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
}

func getPagination(c *gin.Context) (int, int, string, string, string, time.Time, time.Time) {
	pageQ := c.Query("page")
	sizeQ := c.Query("size")
	filter := c.Query("filter")
	sortOrder := c.Query("sortOrder")
	sortOrderName := c.Query("sortOrderName")

	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	var page int
	var size int

	var err error
	if page, err = strconv.Atoi(pageQ); err != nil {
		page = 1
	}

	if size, err = strconv.Atoi(sizeQ); err != nil {
		size = 10
	}

	if sortOrderName == "" {
		sortOrderName = "created_at"
	}

	if sortOrder == "" {
		sortOrder = "desc"
	}

	layout := "2006-01-02 00:00"

	var _starDate time.Time
	var _endDate time.Time

	if startDate != "" {
		_starDate, err = time.Parse(layout, startDate)
		if err != nil {
			fmt.Println("Erro", err)
		}
	}

	if endDate != "" {
		_endDate, err = time.Parse(layout, endDate)

		fmt.Println("xiiiiiii", _starDate, _endDate)
	}

	return page, size, filter, sortOrder, sortOrderName, _starDate, _endDate
}

func errorResponse(err error) *Response {
	return &Response{
		Body:    nil,
		Success: false,
		Message: err.Error(),
	}
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type HTTPUnauthorized struct {
	Code    int    `json:"code" example:"401"`
	Message string `json:"message" example:"Unauthorized"`
}

func successResponse(data interface{}) *Response {
	return &Response{
		Body:    data,
		Success: true,
	}
}
