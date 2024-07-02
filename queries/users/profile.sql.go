package queries

var InsertUpdateProfileInfo = `
	UPDATE user_table
	SET
		user_name = ?,
		preferred_color = ?,
		title = ?,
		sns_instagram = ?,
		github_url = ?,
		personal_url = ?,
		memo = ?
	WHERE
		user_id = ? AND
		blog_owner = ?
`

var UpdateProfileColor = `
		UPDATE user_table
		SET
			preferred_color = ?
		WHERE
			user_id = ? AND 
			blog_owner = ?
`

var UpdateTitle = `
		UPDATE user_table
		SET
			title = ?
		WHERE
			user_id = ? AND
			blog_owner = ?
`

var SelectUserProfile = `
	SELECT
		user_id,
		user_email,
		user_name,
		preferred_color as color,
		title,
		sns_instagram,
		github_url,
		personal_url,
		memo
	FROM user_table
	WHERE
		user_status = 1 AND
		blog_owner = ?
	;
`

var SelectUserProfileByUserId = `
	SELECT
		user_id,
		user_email,
		user_name,
		preferred_color as color,
		title,
		sns_instagram,
		github_url,
		personal_url,
		memo
	FROM user_table
	WHERE
		user_status = 1 AND
		user_id = ?
	;
`


var SelectUserProfileProfileAndBackground = `
	SELECT
		file_format,
		file_type,
		target_purpose,
		target_id,
		object_name
	FROM file_table
	WHERE
		target_id = ? AND
		target_purpose IN ( ?, ? )
`