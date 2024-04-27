package queries

// 게시글 전체 리스트 조회
var SelectAllPosts = `
	SELECT
		p.post_seq, 
		p.post_title, 
		p.post_contents, 
		c.category_name,
		c.category_seq,
		u.user_id, 
		u.user_name,
		p.is_pinned,
		p.viewed,
		p.reg_date, 
		p.mod_date
	FROM
		post_table AS p
	LEFT JOIN user_table AS u ON u.user_id = p.user_id
	LEFT JOIN category_table AS c ON c.category_seq = p.category_seq
	WHERE
		p.post_status = 1
	ORDER BY
		p.is_pinned DESC, p.reg_date DESC
	LIMIT ?
	OFFSET ?;
`

// 태그 이름을 통해 게시글 가져오기
var SelectPostByTags = `
	SELECT
		t.tags as tags,
		p.post_seq,
		p.post_title,
		p.viewed,
		p.reg_date,
		p.mod_date
	FROM
		post_table p
	LEFT JOIN tag_table t ON t.post_seq = p.post_seq
	WHERE
		tags LIKE ? AND
		p.post_status = 1
	ORDER BY reg_date DESC
	LIMIT ?
	OFFSET ?;
`