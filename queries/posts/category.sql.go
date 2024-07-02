package queries

var SelectAllCategories = `
	SELECT DISTINCT(c.category_name) AS category_name
	FROM category_table c
	LEFT JOIN post_table p ON p.post_seq = c.post_seq
	WHERE
		p.blog_owner = ?
`