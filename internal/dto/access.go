package dto

type AccessCreateRequest struct {
	Module *string `json:"module" query:"module"`
	Option *string `json:"option" query:"option"`
}
