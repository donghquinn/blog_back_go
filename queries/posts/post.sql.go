package queries

// 게시글 전체 리스트 조회
var SelectAllPosts = `
	SELECT
		p.post_seq, 
		p.post_title, 
		p.post_contents, 
		u.user_id, 
		u.user_name,
		p.is_pinned,
		p.viewed,
		p.reg_date, 
		p.mod_date
	FROM
		post_table AS p
	LEFT JOIN user_table AS u ON u.user_id = p.user_id
	WHERE
		p.post_status = 1
	ORDER BY
		reg_date ASC
	LIMIT ?
	OFFSET ?;
`

// 특정 게시글 조회
var SelectSpecificPostContents = `
	SELECT
		p.post_seq,
		p.post_title, 
		p.post_contents, 
		p.post_status, 
		t.tag_name,
		u.user_id, 
		u.user_name,
		p.viewed, 
		p.is_pinned, 
		p.reg_date, 
		p.mod_date
	FROM
		post_table AS p
	LEFT JOIN user_table AS u ON u.user_id = p.user_id
	LEFT JOIN tag_table AS t ON t.post_seq = p.post_seq
	WHERE
		p.post_status = 1 AND
		p.post_seq = ?
	;
`

// 조회수 업데이트
var UpdateViewCount = `
	UPDATE post_table SET
		viewed = viewed + 1
	WHERE
		post_seq = ?;
`

// 전체 이미지 데이터
var SelectImageData = `
	SELECT
		object_name, file_format, target_seq
	FROM
		file_table
	WHERE
		target_seq = ?;
`

// 게시글 등록
var InsertPost = `
	INSERT INTO post_table
	SET
		user_id = ?,
		post_title = ?,
		post_contents = ?,
		is_pinned = ?;
`

// 게시글 태그 등록
var InsertTag = `
	INSERT INTO tag_table
	SET
		post_seq = ?,
		tag_name = ?
`

// 게시글 이미지 등록
var InsertUpdatePostImage = `
	UPDATE file_table SET 
		target_seq = ?
	WHERE
		file_seq = ?;
`