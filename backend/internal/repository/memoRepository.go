package repository

import (
	"backend/internal/model"

	"gorm.io/gorm"
)

type IMemoRepository interface {
	Create(memo *model.Memo) (*model.Memo, error)
	FindAll() ([]*model.Memo, error)
	FindById(id int) (*model.Memo, error)
	Update(memo *model.Memo) (*model.Memo, error)
	Delete(id int) error
	FindByTodoId(todoId int) ([]*model.Memo, error)
}

type MemoRepository struct {
	db *gorm.DB
}

func NewMemoRepository(db *gorm.DB) IMemoRepository {
	return &MemoRepository{db: db}
}

func (r *MemoRepository) Create(memo *model.Memo) (*model.Memo, error) {
	if err := r.db.Create(memo).Error; err != nil {
		return nil, err
	}
	return memo, nil
}

func (r *MemoRepository) FindAll() ([]*model.Memo, error) {
	var memos []*model.Memo
	if err := r.db.Find(&memos).Error; err != nil {
		return nil, err
	}
	return memos, nil
}

func (r *MemoRepository) FindById(id int) (*model.Memo, error) {
	var memo model.Memo
	if err := r.db.First(&memo, id).Error; err != nil {
		return nil, err
	}
	return &memo, nil
}

func (r *MemoRepository) Update(memo *model.Memo) (*model.Memo, error) {
	if err := r.db.Save(memo).Error; err != nil {
		return nil, err
	}
	return memo, nil
}

func (r *MemoRepository) Delete(id int) error {
	if err := r.db.Delete(&model.Memo{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *MemoRepository) FindByTodoId(todoId int) ([]*model.Memo, error) {
	var memos []*model.Memo
	if err := r.db.Where("todo_id = ?", todoId).Find(&memos).Error; err != nil {
		return nil, err
	}
	return memos, nil
}
