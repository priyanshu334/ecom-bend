package user

import (
	"gorm.io/gorm"
)

type Service interface {
	GetOrCreateProfile(userID uint) (*UserProfile, error)
	UpdateProfile(userID uint, fullName, phone string) (*UserProfile, error)
	ListAddresses(userID uint) ([]Address, error)
	AddAddress(userID uint, address Address) error
	RemoveAddress(userID, addressID uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetOrCreateProfile(userID uint) (*UserProfile, error) {
	profile, err := s.repo.GetProfileByUserID(userID)
	if err == nil {
		return profile, nil
	}
	if err != gorm.ErrRecordNotFound {
		return nil, err

	}
	newProfile := &UserProfile{
		UserID: userID,
	}
	if err := s.repo.CreateProfile(newProfile); err != nil {
		return nil, err
	}
	return newProfile, nil
}

func (s *service) UpdateProfile(userID uint, fullName, phone string) (*UserProfile, error) {
	profile, err := s.GetOrCreateProfile(userID)
	if err != nil {
		return nil, err
	}
	profile.FullName = fullName
	profile.Phone = phone

	if err := s.repo.UpdateProfile(profile); err != nil {
		return nil, err
	}
	return profile, nil
}

func (s *service) AddAddress(userID uint, address Address) error {
	address.UserID = userID
	return s.repo.CreateAddress(&address)

}

func (s *service) ListAddresses(userID uint) ([]Address, error) {
	return s.repo.GetAddress(userID)
}
func (s *service) RemoveAddress(userID, addressID uint) error {
	return s.repo.DeleteAddress(userID, addressID)

}
