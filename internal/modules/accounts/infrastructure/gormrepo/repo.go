package gormrepo

import (
	"context"
	"errors"

	"example.com/modmonolith/internal/modules/accounts/domain"
	"gorm.io/gorm"
)

type Repo struct{ db *gorm.DB }

func New(db *gorm.DB) *Repo { return &Repo{db: db} }

func (r *Repo) AutoMigrate(ctx context.Context) error {
	return r.db.WithContext(ctx).AutoMigrate(&accountModel{})
}

func (r *Repo) Create(ctx context.Context, a *domain.Account) error {
	m := accountModel{
		ID: string(a.ID()),
		UserID: a.UserID(),
		Label: a.Label(),
		CreatedAt: a.CreatedAt(),
		UpdatedAt: a.UpdatedAt(),
	}
	return r.db.WithContext(ctx).Create(&m).Error
}

func (r *Repo) GetByID(ctx context.Context, id domain.AccountID) (*domain.Account, error) {
	var m accountModel
	err := r.db.WithContext(ctx).First(&m, "id = ?", string(id)).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { return nil, nil }
	if err != nil { return nil, err }
	return domain.RehydrateAccount(domain.AccountID(m.ID), m.UserID, m.Label, m.CreatedAt, m.UpdatedAt), nil
}

func (r *Repo) ListByUser(ctx context.Context, userID string, limit, offset int) ([]*domain.Account, error) {
	var ms []accountModel
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Find(&ms).Error
	if err != nil { return nil, err }

	out := make([]*domain.Account, 0, len(ms))
	for _, m := range ms {
		out = append(out, domain.RehydrateAccount(domain.AccountID(m.ID), m.UserID, m.Label, m.CreatedAt, m.UpdatedAt))
	}
	return out, nil
}

func (r *Repo) Update(ctx context.Context, a *domain.Account) error {
	updates := map[string]any{
		"label": a.Label(),
		"updated_at": a.UpdatedAt(),
	}
	return r.db.WithContext(ctx).Model(&accountModel{}).
		Where("id = ?", string(a.ID())).
		Updates(updates).Error
}

func (r *Repo) Delete(ctx context.Context, id domain.AccountID) error {
	return r.db.WithContext(ctx).Delete(&accountModel{}, "id = ?", string(id)).Error
}
