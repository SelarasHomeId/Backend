package response

type Meta struct {
	Success bool    `json:"success" default:"true"`
	Message string  `json:"message" default:"true"`
	Info    *string `json:"info"`
	ErrorId *string `json:"error_id"`
}
