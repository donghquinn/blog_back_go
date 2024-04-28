package queries

var InsertUpdatePost = `
	INSERT INTO post_table
	(post_title, post_contents, category_seq, is_pinned)
	VALUES
	(?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE 
		post_title = VALUES(post_title),
		post_contents = VALUES(post_contents),
		category_seq = VALUES(category_seq),
		is_pinned = VALUES(is_pinned)
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