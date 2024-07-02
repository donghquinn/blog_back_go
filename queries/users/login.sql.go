package queries

var SelectUserInfo = `
	SELECT
		user_id, user_password, user_status, blog_owner
	FROM
		user_table
	WHERE
		user_email = ?
	;
`

var InsertSessionData = `
	INSERT INTO session_table
	SET
		user_id = ?
`