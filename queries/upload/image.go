package queries

var InsertPostImageData = `
 	INSERT INTO file_table 
		(file_type, target_id, target_table, file_size, object_name, file_format)
    VALUES (?, ?, ?, ?, ?, ?)
`

var InsertProfileImageData = `
	INSERT INTO file_table 
		(file_type, target_id, target_table, file_size, object_name, file_format)
    VALUES (?, ?, ?, ?, ?, ?)
`