package queries


var SelectAllPosts = `
	SELECT
		p.post_seq, 
		p.post_title, 
		p.post_contents, 
		u.user_id, 
		u.user_name,
		p.reg_date, 
		p.mod_date
	FROM
		post_table AS p
	LEFT JOIN user_table AS u ON u.user_id = p.user_id
	WHERE
		p.post_status = 1
	ORDER BY
		reg_date ASC
	LIMIT ?
	OFFSET ?;
`

var SelectSpecificPostContents = `
	SELECT
		p.post_seq,
		p.post_title, 
		p.post_contents, 
		p.post_status, 
		p.tags, 
		u.user_id, 
		u.user_name,
		p.viewed, 
		p.is_pinned, 
		p.reg_date, 
		p.mod_date
	FROM
		post_table AS p
	LEFT JOIN user_table AS u ON u.user_id = p.user_id
	WHERE
		p.post_status = 1 AND
		p.post_seq = ?
	;
`

var UpdateViewCount = `
	UPDATE post_table SET
		viewed = viewed + 1
	WHERE
		post_seq = ?;
`

var SelectImageData = `
	SELECT
		object_name, file_format, target_seq
	FROM
		file_table
	WHERE
		target_seq = ?;
`

var InsertPost = `
	INSERT INTO post_table
	SET
		user_id = ?,
		post_title = ?,
		post_contents = ?;
`

var InsertUpdatePostImage = `
	UPDATE file_table SET 
		target_seq = ?
	WHERE
		file_seq = ?;
`