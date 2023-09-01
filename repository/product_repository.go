package repository

import (
	"clean-code/model"
	"database/sql"
)

type ProductRepository interface {
	Save(product model.Product) error
	FindById(id string) (model.Product, error)
	FindByAll() ([]model.Product, error)
	FindByName(name string) ([]model.Product, error)
	UpdateById(product model.Product) error
	DeleteById(id string) error
}

type productRepository struct {
	db *sql.DB
}

// FindByName implements ProductRepository.
func (p *productRepository) FindByName(name string) ([]model.Product, error) {
	rows, err := p.db.Query(`SELECT p.id, p.name, p.price, u.id, u.name FROM product AS p JOIN uom AS u ON u.id = p.uom_id WHERE p.name ILIKE $1;`, "%"+name+"%")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	productArr := []model.Product{}

	for rows.Next() {
		modelProduct := model.Product{}
		err = rows.Scan(&modelProduct.ID, &modelProduct.Name, &modelProduct.Price, &modelProduct.Uom.ID, &modelProduct.Uom.Name)
		if err != nil {
			return nil, err
		}

		productArr = append(productArr, modelProduct)
	}

	return productArr, nil
}

// DeleteById implements ProductRepository.
func (p *productRepository) DeleteById(id string) error {
	_, err := p.db.Exec("DELETE FROM product WHERE id=$1;", id)
	if err != nil {
		return err
	}

	return nil
}

// FindByAll implements ProductRepository.
func (p *productRepository) FindByAll() ([]model.Product, error) {
	row, err := p.db.Query("SELECT p.id, p.name, p.price, u.id, u.name FROM product AS p JOIN uom AS u ON u.id = p.uom_id;")
	if err != nil {
		return nil, err
	}

	defer row.Close()

	productArr := []model.Product{}

	for row.Next() {
		modelProduct := model.Product{}
		err := row.Scan(&modelProduct.ID, &modelProduct.Name, &modelProduct.Price, &modelProduct.Uom.ID, &modelProduct.Uom.Name)
		if err != nil {
			return nil, err
		}
		productArr = append(productArr, modelProduct)
	}

	return productArr, nil
}

// FindById implements ProductRepository.
func (p *productRepository) FindById(id string) (model.Product, error) {
	row := p.db.QueryRow("SELECT  p.id, p.name, p.price, u.id, u.name FROM product AS p JOIN uom AS u ON u.id = p.uom_id WHERE p.id=$1;", id)

	product := model.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Uom.ID, &product.Uom.Name)
	if err != nil {
		return product, err
	}

	return product, nil
}

// Save implements ProductRepository.
func (p *productRepository) Save(product model.Product) error {
	_, err := p.db.Exec("INSERT INTO product VALUES ($1, $2, $3, $4);", product.ID, product.Name, product.Price, product.Uom.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateById implements ProductRepository.
func (p *productRepository) UpdateById(product model.Product) error {
	_, err := p.db.Exec("UPDATE product SET name=$1, price=$2, uom_id=$3 WHERE id=$4;", product.Name, product.Price, product.Uom.ID, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}
