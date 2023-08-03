package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"TaskApp/config"
	"TaskApp/model"

	"github.com/labstack/echo/v4"
)

func AddTaskController(c echo.Context) error {

	var task model.Task
	c.Bind(&task)

	result := config.DB.Create(&task)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: "Failed create product into database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, model.Response{
		Status:  true,
		Message: "Success add data product",
		Data:    task,
	})
}

func GetDetailTaskController(c echo.Context) error {

	// ==id, _ := strconv.Atoi(c.Param("id"))

	task := model.Task{}

	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Berhasil",
		Data:    task,
	})
}

// Fungsi handler untuk menghapus tugas berdasarkan ID
func DeleteTaskController(c echo.Context) error {
	// Ambil ID tugas dari URL parameter
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: "Invalid task ID",
			Data:    nil,
		})
	}

	err = config.DB.Delete(&model.Task{}, taskID).Error
	if err != nil {
		return err
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: fmt.Sprintf("Gagal menghapus tugas dengan ID %d", taskID),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: fmt.Sprintf("Tugas dengan ID %d berhasil dihapus", taskID),
		Data:    nil,
	})
}

func GetTaskController(c echo.Context) error {

	var dataTask []model.Task

	result := config.DB.Find(&dataTask)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: "Failed get product from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Berhasil",
		Data:    dataTask,
	})
}

func UpdateTaskController(c echo.Context) error {

	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: "Invalid task ID",
			Data:    nil,
		})
	}

	var task model.Task
	c.Bind(&task)
	task.ID = taskID
	result := config.DB.Save(&task)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: "failed to update",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, model.Response{
		Status:  true,
		Message: "Success update",
		Data:    task,
	})
}
