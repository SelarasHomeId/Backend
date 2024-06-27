package dto

type TestResponse struct {
	Message string `json:"message"`
}

type TestGomailRequest struct {
	Recipient string `json:"recipient"`
}
