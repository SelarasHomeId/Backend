package dto

type ContactCreateRequest struct {
	Name    string `json:"name" query:"name"`
	Email   string `json:"email" query:"email"`
	Phone   string `json:"phone" query:"phone"`
	Message string `json:"message" query:"message"`
}
