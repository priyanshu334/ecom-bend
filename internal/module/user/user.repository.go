package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetProfileByUserID(userID uint) (*UserProfile, error)
	CreateProfile(profile *UserProfile) error
	UpdateProfile(profile *UserProfile) error

	GetAddress(userID uint) ([]Address, error)
	GetAddressByID(userID, addressID uint) (*Address, error)
	CreateAddress(address *Address) error
	DeleteAddress(userID, addressID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetProfileByUserID(userID uint) (*UserProfile, error) {
	var profile UserProfile

	err := r.db.Where("user_id=?", userID).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *repository) CreateProfile(profile *UserProfile) error {
	return r.db.Create(profile).Error
}

func (r *repository) UpdateProfile(profile *UserProfile) error {
	return r.db.Save(profile).Error
}

func (r *repository) GetAddress(userID uint) ([]Address, error) {
	var address []Address
	err := r.db.Where("user_id=?", userID).Find(&address).Error
	return address, err
}

func (r *repository) GetAddressByID(userID, addressID uint) (*Address, error) {
	var address Address
	err := r.db.Where("id=? AND user_id=?", addressID, userID).First(&address).Error

	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (r *repository) CreateAddress(address *Address) error {
	return r.db.Create(address).Error
}

func (r *repository) DeleteAddress(userID, addressID uint) error {
	return r.db.Where("id=? AND user_id=?", addressID, userID).Delete(&Address{}).Error
}
