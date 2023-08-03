package model

import (
	"gorm.io/gorm"
)

type Task struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	Judul          string `json:"judul"`
	Deskripsi      string `json:"deskripsi"`
	Prioritas      string `json:"prioritas"`
	TanggalTenggat string `json:"tanggal_tenggat"`
	Status         string `json:"status"`
}

type TaskManager interface {
	AddTask(task *Task) error
	UpdateTaskStatus(id int) error
	GetTask() ([]Task, error)
	GetTaskByID(id int) (Task, error)
	DeleteTaskByID(id int) error
}

// mengakses data dari database menggunakan GORM
type TaskRepository struct {
	DB *gorm.DB
}

// GetTask implements TaskManager.
func (*TaskRepository) GetTask() ([]Task, error) {
	panic("unimplemented")
}

// akses database
func NewTaskManager(DB *gorm.DB) TaskManager {
	return &TaskRepository{DB: DB}
}

func (r *TaskRepository) AddTask(task *Task) error {
	return r.DB.Create(task).Error
}

func (r *TaskRepository) UpdateTaskStatus(id int) error {
	task := Task{}
	err := r.DB.First(&task, id).Error
	if err != nil {
		return err
	}

	task.Status = "Done"
	return r.DB.Save(&task).Error
}

func (r *TaskRepository) GetTasks() ([]Task, error) {
	var task []Task
	err := r.DB.Find(&task).Error
	return task, err
}

func (r *TaskRepository) GetTaskByID(id int) (Task, error) {
	task := Task{}
	err := r.DB.First(&task, id).Error
	return task, err
}

// DeleteTaskByID menghapus tugas berdasarkan ID dari database.
func (r *TaskRepository) DeleteTaskByID(id int) error {
	task := Task{}
	err := r.DB.First(&task, id).Error
	if err != nil {
		return err
	}

	err = r.DB.Delete(&task).Error
	if err != nil {
		return err
	}

	return nil
}
