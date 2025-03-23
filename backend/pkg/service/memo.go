package service

import (
	"backend/gen/memo"
	"backend/internal/model"
	"backend/internal/repository"
	"context"
	"time"
)

type MemoService struct {
	memoRepository repository.IMemoRepository
}

func NewMemoService(memoRepository repository.IMemoRepository) *MemoService {
	return &MemoService{memoRepository: memoRepository}
}

func (s *MemoService) Create(ctx context.Context, m *memo.CreatePayload) (*memo.MemoResult, error) {
	memoModel := &model.Memo{
		Content:   m.Content,
		TodoID:    uint(m.TodoID),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	createdMemo, err := s.memoRepository.Create(memoModel)
	if err != nil {
		return nil, err
	}
	return &memo.MemoResult{
		ID:        int(createdMemo.ID),
		Content:   createdMemo.Content,
		TodoID:    int(createdMemo.TodoID),
		CreatedAt: createdMemo.CreatedAt.Format(time.RFC3339),
		UpdatedAt: createdMemo.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *MemoService) List(ctx context.Context, m *memo.ListPayload) (*memo.MemoCollection, error) {
	memos, err := s.memoRepository.FindByTodoId(m.TodoID)
	if err != nil {
		return nil, err
	}

	var memoResults []*memo.MemoResult
	for _, memoItem := range memos {
		memoResults = append(memoResults, &memo.MemoResult{
			ID:        int(memoItem.ID),
			Content:   memoItem.Content,
			TodoID:    int(memoItem.TodoID),
			CreatedAt: memoItem.CreatedAt.Format(time.RFC3339),
			UpdatedAt: memoItem.UpdatedAt.Format(time.RFC3339),
		})
	}
	return &memo.MemoCollection{
		Items: memoResults,
	}, nil
}

func (s *MemoService) Get(ctx context.Context, m *memo.GetPayload) (*memo.MemoResult, error) {
	foundMemo, err := s.memoRepository.FindById(m.MemoID)
	if err != nil {
		return nil, err
	}
	return &memo.MemoResult{
		ID:        int(foundMemo.ID),
		Content:   foundMemo.Content,
		TodoID:    int(foundMemo.TodoID),
		CreatedAt: foundMemo.CreatedAt.Format(time.RFC3339),
		UpdatedAt: foundMemo.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *MemoService) Update(ctx context.Context, m *memo.UpdatePayload) (*memo.MemoResult, error) {
	memoModel := &model.Memo{
		ID:      uint(m.MemoID),
		Content: m.Content,
		TodoID:  uint(m.TodoID),
	}
	updatedMemo, err := s.memoRepository.Update(memoModel)
	if err != nil {
		return nil, err
	}
	return &memo.MemoResult{
		ID:        int(updatedMemo.ID),
		Content:   updatedMemo.Content,
		TodoID:    int(updatedMemo.TodoID),
		CreatedAt: updatedMemo.CreatedAt.Format(time.RFC3339),
		UpdatedAt: updatedMemo.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *MemoService) Delete(ctx context.Context, m *memo.DeletePayload) error {
	err := s.memoRepository.Delete(m.MemoID)
	if err != nil {
		return err
	}
	return nil
}
