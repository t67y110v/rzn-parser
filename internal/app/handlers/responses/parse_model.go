package responses

type ParserResult struct {
	Result string `json:"result"`
}

type Error struct {
	Message string `json:"message"`
}
