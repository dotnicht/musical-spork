package gormrepo

import (
	"context"
	"errors"

	"example.com/modmonolith/internal/modules/users/domain"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repo { return &Repo{db: db} }

func (r *Repo) AutoMigrate(ctx context.Context) error {
	return r.db.WithContext(ctx).AutoMigrate(&userModel{})
}

func (r *Repo) Create(ctx context.Context, u *domain.User) error {
	m := userModel{
		ID:        string(u.ID()),
		Email:     u.Email(),
		Name:      u.Name(),
		CreatedAt: u.CreatedAt(),
		UpdatedAt: u.UpdatedAt(),
	}
	return r.db.WithContext(ctx).Create(&m).Error
}

func (r *Repo) GetByID(ctx context.Context, id domain.UserID) (*domain.User, error) {
	var m userModel
	err := r.db.WithContext(ctx).First(&m, "id = ?", string(id)).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return domain.RehydrateUser(domain.UserID(m.ID), m.Email, m.Name, m.CreatedAt, m.UpdatedAt), nil
}

func (r *Repo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var m userModel
	err := r.db.WithContext(ctx).First(&m, "email = ?", email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return domain.RehydrateUser(domain.UserID(m.ID), m.Email, m.Name, m.CreatedAt, m.UpdatedAt), nil
}

func (r *Repo) List(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	var ms []userModel
	err := r.db.WithContext(ctx).
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Find(&ms).Error
	if err != nil {
		return nil, err
	}
	out := make([]*domain.User, 0, len(ms))
	for _, m := range ms {
		out = append(out, domain.RehydrateUser(domain.UserID(m.ID), m.Email, m.Name, m.CreatedAt, m.UpdatedAt))
	}
	return out, nil
}

func (r *Repo) Update(ctx context.Context, u *domain.User) error {
	updates := map[string]any{
		"email":      u.Email(),
		"name":       u.Name(),
		"updated_at": u.UpdatedAt(),
	}
	return r.db.WithContext(ctx).
		Model(&userModel{}).
		Where("id = ?", string(u.ID())).
		Updates(updates).Error
}

func (r *Repo) Delete(ctx context.Context, id domain.UserID) error {
	return r.db.WithContext(ctx).Delete(&userModel{}, "id = ?", string(id)).Error
}
