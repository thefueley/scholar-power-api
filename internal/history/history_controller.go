package history

import (
	"context"
	"errors"
)

var (
	ErrHistoryNotFound = errors.New("workout history not found")
)

type History struct {
	ID        string
	Date      string
	Duration  string
	Notes     string
	PlanID    string
	AthleteID string
}

type HistoryStore interface {
	CreateHistory(context.Context, History) error
	GetHistory(context.Context, string) ([]History, error)
	UpdateHistory(context.Context, History) error
	DeleteHistory(context.Context, string) error
}

type HistoryService struct {
	Store HistoryStore
}

func NewHistoryService(store HistoryStore) *HistoryService {
	return &HistoryService{
		Store: store,
	}
}

func (hs *HistoryService) CreateHistory(ctx context.Context, h History) error {
	err := hs.Store.CreateHistory(ctx, h)
	if err != nil {
		return err
	}
	return nil
}

func (hs *HistoryService) GetHistory(ctx context.Context, id string) ([]History, error) {
	h, err := hs.Store.GetHistory(ctx, id)
	if err != nil {
		return []History{}, ErrHistoryNotFound
	}
	return h, nil
}

func (hs *HistoryService) UpdateHistory(ctx context.Context, h History) error {
	err := hs.Store.UpdateHistory(ctx, h)
	if err != nil {
		return err
	}
	return nil
}

func (hs *HistoryService) DeleteHistory(ctx context.Context, id string) error {
	err := hs.Store.DeleteHistory(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
