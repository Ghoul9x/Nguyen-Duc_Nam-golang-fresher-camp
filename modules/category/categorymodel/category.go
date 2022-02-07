package categorymodel

type Category struct {
	Id    int    `json:"id" gorm:"column:id;"`
	Name  string `json:"name" gorm:"column:name;"`
	Image string `json:"image" gorm:"column:image;"`
}

func (Category) TableName() string {
	return "category"
}

type CategoryUpdate struct {
	Name  *string `json:"name" gorm:"column:name;"`
	Image *string `json:"image" gorm:"column:image;"`
}

func (CategoryUpdate) TableName() string {
	return Category{}.TableName()
}

type CategoryCreate struct {
	Id    int    `json:"id" gorm:"column:id;"`
	Name  string `json:"name" gorm:"column:name;"`
	Image string `json:"image" gorm:"column:image;"`
}

func (CategoryCreate) TableName() string {
	return Category{}.TableName()
}
