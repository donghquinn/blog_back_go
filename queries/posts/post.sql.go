package queries


var SelectAllPosts = `
	SELECT
		post_title, post_contents, user_id, reg_date, mod_date
	FROM
		post_table
	WHERE
		post_status = 1
	ORDER BY
		reg_date ASC
	LIMIT ?
	OFFSET ?
`

var SelectSpecificPostContents = `
	SELECT
		p.post_title, p.post_contents, u.user_id, u.user_name, f.object_name, f.file_format, p.reg_date, p.mod_date
	FROM
		post_table AS p
	LEFT JOIN user_table AS u ON u.user_id = p.user_id
	LEFT JOIN file_table AS f ON f.target_seq = p.post_seq
	WHERE
		p.post_status = 1 AND
		p.post_seq = ?
	;
`

var InsertPost = `
	INSERT INTO post_table
	SET
		user_id = ?,
		post_title = ?,
		post_contents = ?
`

var InsertUpdatePostImage = `
	UPDATE file_table SET 
		target_seq = ?
	WHERE
		file_seq = ?
`