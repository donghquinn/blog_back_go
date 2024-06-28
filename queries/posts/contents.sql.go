package queries

// 특정 게시글 조회
var SelectSpecificPostContents = `
	SELECT p.post_seq, p.post_title, p.post_contents, p.post_status,
		IFNULL(t.tags, "NULL") AS tags, 
		IFNULL(c.category_name, "NULL") AS category_name, 
		u.user_name, p.viewed, p.is_pinned, p.reg_date, p.mod_date
	FROM
		post_table AS p
	INNER JOIN user_table AS u ON u.user_id = p.user_id
	INNER JOIN category_table AS c ON c.post_seq = p.post_seq
	INNER JOIN tag_table AS t ON t.post_seq = p.post_seq
	WHERE p.post_seq = ?
		AND p.post_status = 1
	;
`

// 게시글의 태그 쿼리
var SelectPostTags =`
		SELECT
			tags
		FROM
			tag_table
		WHERE
			post_seq = ?
`

// 조회수 업데이트
var UpdateViewCount = `
	UPDATE post_table SET
		viewed = viewed + 1
	WHERE
		post_seq = ?;
`

// 이미지 데이터
var SelectImageData = `
	SELECT
		object_name, file_format, target_purpose, file_seq
	FROM
		file_table
	WHERE
		file_seq = ? AND
		target_purpose = ?
	;
`
