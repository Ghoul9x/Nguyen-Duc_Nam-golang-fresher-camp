package categorybiz

import (
	"context"
	"errors"

	"example.com/modules/category/categorymodel"
)

type CreateCategoryStore interface {
	Create(ctx context.Context, data *categorymodel.CategoryCreate) error
}

type createCategoryBiz struct {
	store CreateCategoryStore
}

func NewCreateCategoryBiz(store CreateCategoryStore) *createCategoryBiz {
	return &createCategoryBiz{store: store}
}

func (biz *createCategoryBiz) CreateCategory(ctx context.Context, data *categorymodel.CategoryCreate) error {
	if data.Name == "" {
		return errors.New("category name cannot be blank")
	}
	err := biz.store.Create(ctx, data)

	return err
}
