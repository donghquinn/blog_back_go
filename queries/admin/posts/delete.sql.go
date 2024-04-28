package queries

// 게글 삭제
var DeletePost = `
	UPDATE post_table
	SET
		post_status = ?
	WHERE
		post_seq = ?
`
