package queries

// 게시글 전체 리스트 조회
var SelectAllPosts = `
	SELECT
		p.post_seq, 
		p.post_title, 
		p.post_contents, 
		IFNULL(c.category_name, 'NULL') as category_name,
		IFNULL(u.user_name, 'unknown') as user_name,
		p.is_pinned,
		p.viewed,
		p.reg_date, 
		p.mod_date
	FROM
		post_table AS p
	LEFT JOIN user_table AS u ON u.user_id = p.user_id
	LEFT JOIN category_table AS c ON c.post_seq = p.post_seq
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
		IFNULL(c.category_name, 'NULL') category_name,
		IFNULL(u.user_name, 'unknown') as user_name,
		p.post_seq,
		p.post_title,
		p.post_contents,
		p.viewed,
		p.reg_date,
		p.mod_date
	FROM
		post_table p
	LEFT JOIN user_table u ON u.user_id = p.user_id
	LEFT JOIN tag_table t ON t.post_seq = p.post_seq
	LEFT JOIN category_table AS c ON c.post_seq = p.post_seq
	WHERE
		tags LIKE ? AND
		p.post_status = 1
	ORDER BY p.reg_date DESC
	LIMIT ?
	OFFSET ?;
`

// 태그 이름을 통해 게시글 가져오기
var SelectPostByCategory = `
	SELECT
		t.tags as tags,
		IFNULL(c.category_name, 'NULL') category_name,
		IFNULL(u.user_name, 'unknown') as user_name,
		p.post_seq,
		p.post_title,
		p.post_contents,
		p.viewed,
		p.reg_date,
		p.mod_date
	FROM
		post_table p
	LEFT JOIN user_table u ON u.user_id = p.user_id
	LEFT JOIN tag_table t ON t.post_seq = p.post_seq
	LEFT JOIN category_table AS c ON c.post_seq = p.post_seq
	WHERE
		c.category_name LIKE ? AND
		p.post_status = 1
	ORDER BY p.reg_date DESC
	LIMIT ?
	OFFSET ?;
`