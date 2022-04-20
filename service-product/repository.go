package product

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
}

type Repository interface {
	Save(p Product) (Product, error)
	GetByID(id uint) (Product, error)
	GetProducts() ([]Product, error)
}

// SQLite implementation

type sqliteRepository struct {
	db *gorm.DB
}

func (r sqliteRepository) Save(p Product) (Product, error) {
	if err := r.db.Create(&p).Error; err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r sqliteRepository) GetByID(id uint) (Product, error) {
	var product Product
	err := r.db.First(&product, id).Error
	return product, err
}

func (r sqliteRepository) GetProducts() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	return products, err
}

func NewSQLiteRepository(db *gorm.DB) Repository {
	return sqliteRepository{
		db: db,
	}
}
