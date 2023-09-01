package repository

import (
	"clean-code/model"
	"database/sql"
)

type UomRepository interface {
	Save(uom model.Uom) error
	FindById(id string) (model.Uom, error)
	FindByAll() ([]model.Uom, error)
	FindByName(name string) ([]model.Uom, error)
	UpdateById(uom model.Uom) error
	DeleteById(id string) error
}

type uomRepository struct {
	db *sql.DB
}

// FindByName implements UomRepository.
func (u *uomRepository) FindByName(name string) ([]model.Uom, error) {
	r, err := u.db.Query(`SELECT id, name FROM uom WHERE name ILIKE $1;`, "%"+name+"%")
	if err != nil {
		return nil, err
	}

	defer r.Close()

	uomArr := []model.Uom{}

	for r.Next() {
		uomModel := model.Uom{}
		err = r.Scan(&uomModel.ID, &uomModel.Name)
		if err != nil {
			return nil, err
		}

		uomArr = append(uomArr, uomModel)
	}

	return uomArr, nil
}

// UpdateById implements UomRepository.
func (u *uomRepository) UpdateById(uom model.Uom) error {
	_, err := u.db.Exec("UPDATE uom SET name=$1 WHERE id=$2;", uom.Name, uom.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteById implements UomRepository.
func (u *uomRepository) DeleteById(id string) error {
	_, err := u.db.Exec("DELETE FROM uom WHERE id=$1;", id)
	if err != nil {
		return err
	}

	return nil
}

// FindByAll implements UomRepository.
func (u *uomRepository) FindByAll() ([]model.Uom, error) {
	row, err := u.db.Query("SELECT id, name FROM uom;")
	if err != nil {
		return nil, err
	}

	defer row.Close()
	uomArr := []model.Uom{}

	for row.Next() {
		modelUom := model.Uom{}
		err := row.Scan(&modelUom.ID, &modelUom.Name)
		if err != nil {
			return nil, err
		}
		uomArr = append(uomArr, modelUom)
	}

	return uomArr, nil
}

// FindById implements UomRepository.
func (u *uomRepository) FindById(id string) (model.Uom, error) {
	row := u.db.QueryRow("SELECT id, name FROM uom WHERE id=$1;", id)

	uom := model.Uom{}
	err := row.Scan(&uom.ID, &uom.Name)
	if err != nil {
		return uom, err
	}

	return uom, nil
}

// Save implements UomRepository.
func (u *uomRepository) Save(uom model.Uom) error {
	_, err := u.db.Exec("INSERT INTO uom VALUES ($1, $2);", uom.ID, uom.Name)
	if err != nil {
		return err
	}

	return nil
}

func NewUomRepository(db *sql.DB) UomRepository {
	return &uomRepository{db: db}
}
