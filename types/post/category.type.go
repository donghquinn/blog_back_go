package types

type ResponseCategoryResponseType struct {
	Status       int      `json:"status"`
	Code         string   `json:"code"`
	Result       bool     `json:"result"`
	CategoryList []string `json:"categoryList"`
	Message      string   `json:"message"`
}

type CategoryQueryResult struct {
	CategoryName string
}
