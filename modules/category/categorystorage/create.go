package categorystorage

import (
	"context"

	"example.com/modules/category/categorymodel"
)

func (s *sqlStore) Create(ctx context.Context, data *categorymodel.CategoryCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
