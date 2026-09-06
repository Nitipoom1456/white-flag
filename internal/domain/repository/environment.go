package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/white-flag/internal/domain/entity"
	"gorm.io/gorm"
)

type EnvironmentRepository interface {
	Create(ctx context.Context, env entity.Environment) error
	FindByAppID(ctx context.Context, appID uuid.UUID) ([]entity.Environment, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type Environment struct {
	DB *gorm.DB
}

func NewEnvironmentRepository(db *gorm.DB) EnvironmentRepository {
	return &Environment{DB: db}
}

func (e *Environment) Create(ctx context.Context, env entity.Environment) error {
	result := e.DB.WithContext(ctx).Create(&env)
	return result.Error
}

func (e *Environment) FindByAppID(ctx context.Context, appID uuid.UUID) ([]entity.Environment, error) {
	var envs []entity.Environment
	result := e.DB.WithContext(ctx).Where("app_id = ?", appID).Find(&envs)
	return envs, result.Error
}

func (e *Environment) Delete(ctx context.Context, id uuid.UUID) error {
	result := e.DB.WithContext(ctx).Delete(&entity.Environment{ID: id})
	return result.Error
}
