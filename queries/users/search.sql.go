package queries

var SelectUserEmail = `
	SELECT
		user_email
	FROM
		user_table
	WHERE
		user_name = ? AND
		user_status = 1
	;
`

var SelectUserPassword = `
	SELECT
		user_password
	FROM
		user_table
	WHERE
		user_name = ? AND
		user_email = ? AND
		user_status = 1
	;
`