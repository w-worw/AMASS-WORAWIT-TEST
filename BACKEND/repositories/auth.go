package repositories

import (
	"amass-test/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	CheckUsernameExists(username string) (bool, error)
	GetByUsername(username string) (models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Register(user models.User) (models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *authRepository) CheckUsernameExists(username string) (bool, error) {
	var count int64
	if err := r.db.Model(&models.User{}).Where("username = ? and deleted_at is null", username).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *authRepository) GetByUsername(username string) (models.User, error) {
	var user models.User
	if err := r.db.Where("username = ? and deleted_at is null", username).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
