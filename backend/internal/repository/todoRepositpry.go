package repository

import (
	"backend/internal/model"

	"gorm.io/gorm"
)

type ITodoRepository interface {
	Create(todo *model.Todo) (*model.Todo, error)
	FindAll() ([]*model.Todo, error)
	FindById(id int) (*model.Todo, error)
	Update(todo *model.Todo) (*model.Todo, error)
	Delete(id int) error
	FindByUserId(userId int) ([]*model.Todo, error)
}

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) ITodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(todo *model.Todo) (*model.Todo, error) {
	if err := r.db.Create(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoRepository) FindAll() ([]*model.Todo, error) {
	var todos []*model.Todo
	if err := r.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoRepository) FindById(id int) (*model.Todo, error) {
	var todo model.Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) Update(todo *model.Todo) (*model.Todo, error) {
	if err := r.db.Save(todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoRepository) Delete(id int) error {
	if err := r.db.Delete(&model.Todo{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) FindByUserId(userId int) ([]*model.Todo, error) {
	var todos []*model.Todo
	if err := r.db.Where("user_id = ?", userId).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
