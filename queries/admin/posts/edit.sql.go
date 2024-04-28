package queries

var UpdateEditPost = `
	UPDATE post_table
	SET
		post_title = ?,
		post_contents = ?,
		is_pinned = ?
	WHERE
		post_seq = ?
`

var InsertUpdateImage = `
	INSERT INTO file_table
	(file_type, file_format, file_size, object_name)
	VALUES (?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE 
		file_type = VALUES(file_type),
		file_format = VALUES(file_format),
		file_size = VALUES(file_size)
		object_name = VALUES(object_name)
`

var InsertUpdateCategory = `
	INSERT INTO category_table
	SET
		category_name = ?
	ON DUPLICATE KEY UPDATE
		category_name = VALUES(category_name)
`

var DeletePostCategory = `
	DELETE 
	FROM category_table
	WHERE post_seq = ?
`

var DeletePostTag = `
	DELETE
	FROM tag_table
	WHERE post_seq = ?
`