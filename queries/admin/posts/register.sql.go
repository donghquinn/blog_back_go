package queries

// 게시글 등록
var InsertPost = `
	INSERT INTO post_table
	SET
		user_id = ?,
		post_title = ?,
		post_contents = ?,
		is_pinned = ?,
		category_seq = NULLIF(?, "");
`

// 게시글 태그 등록
var InsertTag = `
	INSERT INTO tag_table
	SET
		post_seq = ?,
		tags = ?
`

var InsertCategory  = `
	INSERT INTO category_table
	SET
		category_name = ?
`

// 게시글 이미지 등록
var InsertUpdatePostImage = `
	UPDATE file_table SET 
		target_seq = ?
	WHERE
		file_seq = ?;
`