package repository

import (
	"context"

	"github.com/white-flag/internal/domain/entity"
	"gorm.io/gorm"
)

type AppRepository interface {
	Create(ctx context.Context, app entity.App) error
	List(ctx context.Context, page, pageSize int) ([]entity.App, int64, error)
}

type App struct {
	DB *gorm.DB
}

func NewAppRepository(db *gorm.DB) AppRepository {
	return &App{DB: db}
}

func (a *App) Create(ctx context.Context, app entity.App) error {
	return a.DB.WithContext(ctx).Create(&app).Error
}

func (a *App) List(ctx context.Context, offset, limit int) ([]entity.App, int64, error) {
	var apps []entity.App
	if err := a.DB.WithContext(ctx).Offset(offset).Limit(limit).Find(&apps).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	if err := a.DB.WithContext(ctx).Model(&entity.App{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return apps, total, nil
}
