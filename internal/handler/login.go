package handler

func Login(body string) {

}

type LoginData struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
