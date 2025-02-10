package models

// RegisterInfo - submit registration info
type RegisterInfo struct {
	Fullname string
	Status   string
	Username string
	Password string
}

// LoginInfo
type LoginInfo struct {
	Username string
	Password string
}

// LoginResponse
type LoginResponse struct {
	Success bool
	UUID    string
	Message string
}
