package controller

import (
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

	// id, _ := strconv.Atoi(c.Param("id"))

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

	// Hapus tugas dari sumber data, misalnya dari database
	err = model.DeleteTaskController(taskID) // Contoh fungsi untuk menghapus data tugas dari database

	if err != nil {
		// Jika terjadi error saat menghapus tugas
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: "Gagal menghapus tugas dengan ID %d",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Tugas dengan ID %d berhasil dihapus",
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
