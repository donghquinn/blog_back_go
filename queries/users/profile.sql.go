package queries

var InsertUpdateProfileInfo = `
	INSERT INTO user_table
	(
		user_id, 
		profile_seq, 
		background_seq,
		preferred_color,
		title,
		sns_instagram,
		github_url,
		personal_url,
		memo
	) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?
	) ON DUPLICATE KEY UPDATE 
		profile_seq = VALUES(profile_seq),
		background_seq = VALUES(background_seq),
		preferred_color = VALUES(preferred_color),
		title = VALUES(title),
		sns_instagram = VALUES(sns_instagram),
		github_url = VALUES(github_url),
		personal_url = VALUES(personal_url),
		memo = VALUES(memo);
`

var UpdateProfileColor = `
		UPDATE user_table
		SET
			preferred_color = ?
		WHERE
			user_id = ?;
`

var UpdateTitle = `
		UPDATE user_table
		SET
			title = ?
		WHERE
			user_id = ?
`

var SelectUserProfile = `
		SELECT
			f.*,
			u.user_id,
			u.user_email,
			u.user_name,
			u.profile_seq,
			u.background_seq,
			u.preferred_color as color,
			u.title,
			u.github_url,
			u.personal_url,
			u.memo
		FROM (
			SELECT
				file_format,
				file_type,
				target_purpose,
				target_id,
				object_name
			FROM file_table
			WHERE
				target_id = ? AND
				target_purpose = ( ?, ? )
		) f
		LEFT JOIN user_table u ON u.user_id = f.target_id
		WHERE
			user_status = 1 AND
			user_id = ?
		GROUP BY u.user_id
		;

`