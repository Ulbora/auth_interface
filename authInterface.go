package authinterface

//Login Login
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ClientID int64  `json:"clientId"`
}

//AuthInterface AuthInterface
type AuthInterface interface {
	UserLogin(login *Login) bool
}

//GO111MODULE=on go mod init github.com/Ulbora/auth_interface
