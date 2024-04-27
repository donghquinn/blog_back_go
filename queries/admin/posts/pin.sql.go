package queries

// 고정 게시글 여부 업데이트
var UpdatePinPost = `
	UPDATE post_table
	SET
		is_pinned = ?
	WHERE
		post_seq = ?
`
