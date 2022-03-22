package repository

import (
	"context"
	"test/config"
	"test/core/entity"
)

type TestRepository interface {
	GetTest(ctx context.Context, message string) (*entity.Test, error)
}

type testRepository struct {
	cfg config.Config
}

func NewTestRepository(cfg config.Config) TestRepository {
	return &testRepository{
		cfg: cfg,
	}
}

func (tr *testRepository) GetTest(ctx context.Context, message string) (*entity.Test, error) {
	return &entity.Test{Message: "hi"}, nil
}
