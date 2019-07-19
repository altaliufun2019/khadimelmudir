package authentication

type loginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
	Msg   string `json:"msg"`
	Name  string `json:"name"`
}

type signUpCredits struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type signUpResponse struct {
	Token string `json:"token"`
	Msg   string `json:"msg"`
}
