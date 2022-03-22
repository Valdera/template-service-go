package module

import (
	"context"
	"test/config"
	"test/core/entity"
	"test/core/repository"
)

type TestUseCase interface {
	GetTest(ctx context.Context) (*entity.Test, error)
}

type testUseCase struct {
	testRepo repository.TestRepository
	cfg      config.Config
}

func NewTestUsecase(cfg config.Config, testRepo repository.TestRepository) TestUseCase {
	return &testUseCase{
		cfg:      cfg,
		testRepo: testRepo,
	}
}

func (tc *testUseCase) GetTest(ctx context.Context) (*entity.Test, error) {
	msg, err := tc.testRepo.GetTest(ctx)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
