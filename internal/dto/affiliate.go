package dto

type AffiliateCreateRequest struct {
	Name      string `json:"name" query:"name"`
	Email     string `json:"email" query:"email"`
	Phone     string `json:"phone" query:"phone"`
	Instagram string `json:"instagram" query:"instagram"`
	Tiktok    string `json:"tiktok" query:"tiktok"`
}
