package queries

var InsertSignupUser = `
	INSERT INTO user_table
	SET
		user_id = ?,
		user_email = ?,
		user_password = ?,
		user_name = ?,
		blog_owner = ?
`