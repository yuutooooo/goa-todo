package service

import (
	"backend/gen/todo"
	"backend/internal/model"
	"backend/internal/repository"
	"context"
	"time"
)

type TodoService struct {
	todoRepository repository.ITodoRepository
}

func NewTodoService(todoRepository repository.ITodoRepository) *TodoService {
	return &TodoService{todoRepository: todoRepository}
}

func (s *TodoService) Create(ctx context.Context, t *todo.CreatePayload) (*todo.TodoResult, error) {
	todoModel := &model.Todo{
		Title:       t.Title,
		Description: t.Description,
		UserID:      uint(t.UserID),
		Completed:   t.Completed,
	}
	createdTodo, err := s.todoRepository.Create(todoModel)
	if err != nil {
		return nil, err
	}
	return &todo.TodoResult{
		ID:          int(createdTodo.ID),
		Title:       createdTodo.Title,
		Description: createdTodo.Description,
		CreatedAt:   createdTodo.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   createdTodo.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *TodoService) List(ctx context.Context, t *todo.ListPayload) (*todo.TodoCollection, error) {
	todos, err := s.todoRepository.FindByUserId(t.UserID)
	if err != nil {
		return nil, err
	}

	// TodoResultの配列に変換
	var todoResults []*todo.TodoResult
	for _, t := range todos {
		todoResults = append(todoResults, &todo.TodoResult{
			ID:          int(t.ID),
			UserID:      int(t.UserID),
			Title:       t.Title,
			Description: t.Description,
			Completed:   t.Completed,
			CreatedAt:   t.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   t.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &todo.TodoCollection{
		Items: todoResults,
	}, nil
}

func (s *TodoService) Get(ctx context.Context, p *todo.GetPayload) (*todo.TodoResult, error) {
	foundTodo, err := s.todoRepository.FindById(p.TodoID)
	if err != nil {
		return nil, err
	}
	return &todo.TodoResult{
		ID:          int(foundTodo.ID),
		UserID:      int(foundTodo.UserID),
		Title:       foundTodo.Title,
		Description: foundTodo.Description,
		Completed:   foundTodo.Completed,
		CreatedAt:   foundTodo.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   foundTodo.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *TodoService) Update(ctx context.Context, p *todo.UpdatePayload) (*todo.TodoResult, error) {
	todoModel := &model.Todo{
		ID:          uint(p.TodoID),
		UserID:      uint(p.UserID),
		Title:       *p.Title,
		Description: *p.Description,
		Completed:   *p.Completed,
	}
	updatedTodo, err := s.todoRepository.Update(todoModel)
	if err != nil {
		return nil, err
	}
	return &todo.TodoResult{
		ID:          int(updatedTodo.ID),
		UserID:      int(updatedTodo.UserID),
		Title:       updatedTodo.Title,
		Description: updatedTodo.Description,
		Completed:   updatedTodo.Completed,
		CreatedAt:   updatedTodo.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   updatedTodo.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *TodoService) Delete(ctx context.Context, p *todo.DeletePayload) error {
	err := s.todoRepository.Delete(p.TodoID)
	if err != nil {
		return err
	}
	return nil
}
