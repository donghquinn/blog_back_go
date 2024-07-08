package queries

var SelectAllCategories = `
	SELECT DISTINCT(category_name) AS category_name
	FROM category_table
	WHERE blog_owner = ?
		AND category_status = 1
`