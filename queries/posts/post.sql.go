package queries


var SelectAllPosts = `
	SELECT
		post_seq, post_title, post_contents, user_id, reg_date, mod_date
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
		p.post_seq, p.post_title, p.post_contents, p.post_status, u.user_id, u.user_name, f.*, p.reg_date, p.mod_date
	FROM
		(
			SELECT
				object_name, file_format, target_seq
			FROM
				file_table
			WHERE
				target_seq = ?
		) AS f
	LEFT JOIN post_table AS p ON p.post_seq = f.target_seq
	LEFT JOIN user_table AS u ON u.user_id = p.user_id
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