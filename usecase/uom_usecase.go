package usecase

import (
	"clean-code/model"
	"clean-code/model/dto"
	"clean-code/repository"
	"fmt"
)

type UomUseCase interface {
	CreateNew(payload model.Uom) error
	GetById(id string) (model.Uom, error)
	GetAll() ([]model.Uom, error)
	GetByName(name string) ([]model.Uom, error)
	Update(payload model.Uom) error
	Delete(id string) error
	Paging(payload dto.PageRequest) ([]model.Uom, dto.Paging, error)
}

type uomUseCase struct {
	repo repository.UomRepository
}

// Paging implements UomUseCase.
func (u *uomUseCase) Paging(payload dto.PageRequest) ([]model.Uom, dto.Paging, error) {
	return u.repo.Paging(payload)
}

// GetByName implements UomUseCase.
func (u *uomUseCase) GetByName(name string) ([]model.Uom, error) {
	u2, err := u.repo.FindByName(name)
	if err != nil {
		return u2, err
	}
	return u2, nil
}

// CreateNew implements UomUseCase.
func (u *uomUseCase) CreateNew(payload model.Uom) error {
	if payload.ID == "" {
		return fmt.Errorf("id is required")
	}

	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	err := u.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("failed to create new uom: %v", err)
	}

	return nil
}

// Delete implements UomUseCase.
func (u *uomUseCase) Delete(id string) error {
	uom, err := u.GetById(id)
	if err != nil {
		return err
	}

	err = u.repo.DeleteById(uom.ID)
	if err != nil {
		return fmt.Errorf("failed to delete uom: %v", err)
	}

	return nil
}

// GetAll implements UomUseCase.
func (u *uomUseCase) GetAll() ([]model.Uom, error) {
	uom, err := u.repo.FindByAll()
	if err != nil {
		return []model.Uom{}, fmt.Errorf("data is empty")
	}

	return uom, nil
}

// GetById implements UomUseCase.
func (u *uomUseCase) GetById(id string) (model.Uom, error) {
	uom, err := u.repo.FindById(id)
	if err != nil {
		return model.Uom{}, fmt.Errorf("failed to find uom: %v", err)
	}

	return uom, nil
}

// Update implements UomUseCase.
func (u *uomUseCase) Update(payload model.Uom) error {
	if payload.ID == "" {
		return fmt.Errorf("id is required")
	}

	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	_, err := u.repo.FindById(payload.ID)
	if err != nil {
		return fmt.Errorf("id is not found")
	}

	err = u.repo.UpdateById(payload)
	if err != nil {
		return fmt.Errorf("failed to update uom: %v", err)
	}

	return nil
}

func NewUomUseCase(uomRepo repository.UomRepository) UomUseCase {
	return &uomUseCase{repo: uomRepo}
}
