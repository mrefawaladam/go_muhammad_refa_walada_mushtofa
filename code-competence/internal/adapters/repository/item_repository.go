package repository

import (
	"capston-lms/internal/entity"

	"gorm.io/gorm"
)

type ItemRepository struct {
	DB *gorm.DB
}

func (repo ItemRepository) GetAllItems() ([]entity.Items, error) {
	var items []entity.Items
	result := repo.DB.Preload("Category", "deleted_at IS NULL").Find(&items)
	return items, result.Error
}

func (repo ItemRepository) GetItem(id int) (entity.Items, error) {
	var item entity.Items
	result := repo.DB.First(&item, id)
	return item, result.Error
}

func (repo ItemRepository) CreateItem(item entity.Items) error {
	result := repo.DB.Create(&item)
	return result.Error
}

func (repo ItemRepository) UpdateItem(id int, item entity.Items) error {
	result := repo.DB.Model(&item).Where("id = ?", id).Updates(&item)
	return result.Error
}

func (repo ItemRepository) GetItemByName(name string) (entity.Items, error) {
	var item entity.Items
	result := repo.DB.Find(&item, "item_name = ?", name)
	return item, result.Error
}
func (repo ItemRepository) GetItemByCategory(categoryID int) (entity.Items, error) {
	var item entity.Items
	result := repo.DB.Find(&item, "category_id = ?", categoryID)
	// log.Printf("Incoming request:", item)

	return item, result.Error
}

func (repo ItemRepository) DeleteItems(id int) error {
	result := repo.DB.Delete(&entity.Items{}, id)
	return result.Error
}

func (repo ItemRepository) FindItem(id int) error {
	result := repo.DB.First(&entity.Items{}, id)
	return result.Error
}
