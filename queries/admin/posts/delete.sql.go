package queries

// 게글 삭제
var DeletePost = `
	UPDATE post_table
	SET
		post_status = ?
	WHERE
		post_seq = ? AND
		blog_owner = ?
`

var DeleteTags = `
	UPDATE tag_table
	SET
		tag_status = ?
	WHERE post_seq = ?
		AND blog_owner = ?
`

var DeleteCategory = `
	UPDATE category_table
	SET
		category_status = ?
	WHERE post_seq = ?
		AND blog_owner = ?
`