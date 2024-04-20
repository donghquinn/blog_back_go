package quries


var GetAllPosts = `
	SELECT
		post_title, post_content
	FROM
		post_table
	ORDER BY
		reg_date ASC
	LIMIT
		?, ?
`