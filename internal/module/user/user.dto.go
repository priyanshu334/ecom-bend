package user

type ProfileResponse struct {
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
}

type UpdateProfileRequest struct {
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
}

type AddressResponse struct {
	ID         uint   `json:"id"`
	Label      string `json:"label"`
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
	IsDefault  bool   `json:"is_default"`
}

type CreateAddressRequest struct {
	Label      string `json:"label"`
	Line1      string `json:"line1" validate:"required"`
	Line2      string `json:"line2"`
	City       string `json:"city" validate:"required"`
	State      string `json:"state" validate:"required"`
	PostalCode string `json:"postal_code" validate:"required"`
	Country    string `json:"country" validate:"required"`
}
