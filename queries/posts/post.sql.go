package queries


var GetAllPosts = `
	SELECT
		post_title, post_contents, user_id, reg_date, mod_date
	FROM
		post_table
	ORDER BY
		reg_date ASC
	LIMIT ?
	OFFSET ?
`

var InsertPost = `
	INSERT INTO post_table
	SET
		user_id = ?,
		post_title = ?,
		post_contents = ?
`