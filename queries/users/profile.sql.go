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