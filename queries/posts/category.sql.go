package queries

var SelectAllCategories = `
	SELECT DISTINCT(category_name) AS category_name
	FROM category_table
`