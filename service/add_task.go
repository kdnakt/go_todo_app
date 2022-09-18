package service

import (
	"context"
	"fmt"

	"github.com/kdnakt/go_todo_app/auth"
	"github.com/kdnakt/go_todo_app/entity"
	"github.com/kdnakt/go_todo_app/store"
)

type AddTask struct {
	DB   store.Execer
	Repo TaskAdder
}

func (at *AddTask) AddTask(ctx context.Context, title string) (*entity.Task, error) {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user_id not found")
	}
	t := &entity.Task{
		UserID: id,
		Title:  title,
		Status: entity.TaskStatusTodo,
	}
	err := at.Repo.AddTask(ctx, at.DB, t)
	if err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}
	return t, nil
}
