package requests

type ParserLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Path     string `json:"path"`
}
