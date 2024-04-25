package queries

var InsertPostImageData = `
 	INSERT INTO file_table 
		(file_type, target_id, target_table, target_purpose, file_size, object_name, file_format)
    VALUES (?, ?, ?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE 
		file_type = VALUES(file_type),
		file_format = VALUES(file_format),
		target_seq = VALUES(target_seq),
		target_id = VALUES(target_id),
		target_table = VALUES(target_table),
		target_purpose = VALUES(target_purpose),
		file_size = VALUES(file_size),
		object_name = VALUES(object_name)
`

var InsertProfileImageData = `
	INSERT INTO file_table 
		(file_type, target_id, target_table, target_purpose, file_size, object_name, file_format)
    VALUES (?, ?, ?, ?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE 
		file_type = VALUES(file_type),
		file_format = VALUES(file_format),
		target_seq = VALUES(target_seq),
		target_id = VALUES(target_id),
		target_table = VALUES(target_table),
		target_purpose = VALUES(target_purpose),
		file_size = VALUES(file_size),
		object_name = VALUES(object_name)
`