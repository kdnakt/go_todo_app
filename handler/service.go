package handler

import (
	"context"

	"github.com/kdnakt/go_todo_app/entity"
)

type ListTasksService interface {
	ListTasks(ctx context.Context) (entity.Tasks, error)
}

type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*entity.Task, error)
}
