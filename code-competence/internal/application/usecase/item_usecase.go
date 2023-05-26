package usecase

import (
	"capston-lms/internal/adapters/repository"
	"capston-lms/internal/entity"
)

type ItemUseCase struct {
	Repo repository.ItemRepository
}

func (usecase ItemUseCase) GetAlItem() ([]entity.Items, error) {
	items, err := usecase.Repo.GetAllItems()
	return items, err
}

func (usecase ItemUseCase) GetItem(id int) (entity.Items, error) {
	item, err := usecase.Repo.GetItem(id)
	return item, err
}
func (usecase ItemUseCase) GetItemByCategory(id int) (entity.Items, error) {
	item, err := usecase.Repo.GetItemByCategory(id)
	return item, err
}
func (usecase ItemUseCase) GetItemByName(name string) (entity.Items, error) {
	item, err := usecase.Repo.GetItemByName(name)
	return item, err
}
func (usecase ItemUseCase) CreateItem(item entity.Items) error {
	err := usecase.Repo.CreateItem(item)
	return err
}

func (usecase ItemUseCase) UpdateItem(id int, item entity.Items) error {
	err := usecase.Repo.UpdateItem(id, item)
	return err
}

func (usecase ItemUseCase) DeleteIetem(id int) error {
	err := usecase.Repo.DeleteItems(id)
	return err
}
func (usecase ItemUseCase) FindItem(id int) error {
	err := usecase.Repo.FindItem(id)
	return err
}
