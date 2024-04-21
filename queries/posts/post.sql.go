package queries


var GetAllPosts = `
	SELECT
		post_title, post_content, user_id, reg_date, mod_date
	FROM
		post_table
	ORDER BY
		reg_date ASC
	LIMIT
		?, ?
`