package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string `json:"-"` // do not share password to the user back
}

type Repository interface {
	Save(user User) (User, error)
	GetByID(id uint) (User, error)
}

// SQLite implementation

type sqliteRepository struct {
	db *gorm.DB
}

func (r sqliteRepository) Save(u User) (User, error) {
	if err := r.db.Create(&u).Error; err != nil {
		return User{}, err
	}

	return u, nil
}

func (r sqliteRepository) GetByID(id uint) (User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return user, err
}

func NewSQLiteRepository(db *gorm.DB) Repository {
	return sqliteRepository{
		db: db,
	}
}
