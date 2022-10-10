package apimodels

type ReqRegistration struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

type ResRegistration struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
