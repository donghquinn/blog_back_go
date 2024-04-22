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

var InsertUpdatePostImage = `
	INSERT INTO file_table (target_seq)
	VALUES (?)
	ON DUPLICATE UPDATE
		file_seq = VALUES(file_seq),
		file_type = VALUES(file_type),
		target_seq = VALUES(target_seq),
		target_id = VALUES(target_id),
		target_table = VALUES(target_table),
		file_size = VALUES(file_size),
		version_id = VALUES(version_id),
		reg_date = VALUES(reg_date)
`