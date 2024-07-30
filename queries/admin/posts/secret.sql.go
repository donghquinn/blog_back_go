package queries

var ChangeToSecretPostQuery = `
	UPDATE post_table
	SET is_secret = 1
	WHERE post_seq = ?
`

var ChangeToNotSecretPostQuery = `
	UPDATE post_table
	SET is_secret = 0
	WHERE post_seq = ?
`