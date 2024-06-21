package types

type ResponseCategoryResponseType struct {
	Code string `json:"code"`
	Result bool `json:"result"`
	CategoryList []string `json:"categoryList"`
}

type CategoryQueryResult struct {
	CategoryName string
}