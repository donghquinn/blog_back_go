package queries

// 게시글 전체 리스트 조회
var SelectUnPinnedPosts = `
	SELECT
		p.post_seq, 
		p.post_title, 
		p.post_contents, 
		c.category_name,
		IFNULL(u.user_name, 'unknown') as user_name,
		p.is_pinned,
		p.viewed,
		p.reg_date, 
		p.mod_date
	FROM
		post_table AS p
	LEFT JOIN user_table AS u ON u.user_id = p.user_id
	LEFT JOIN category_table AS c ON c.post_seq = p.post_seq
	WHERE p.post_status = 1
		AND p.is_pinned = 0
	ORDER BY
		p.mod_date DESC
	LIMIT ?
	OFFSET ?;
`

var SelectAllPinnedPosts = `
	SELECT
		p.post_seq, 
		p.post_title, 
		p.post_contents, 
		c.category_name,
		IFNULL(u.user_name, 'unknown') as user_name,
		p.is_pinned,
		p.viewed,
		p.reg_date, 
		p.mod_date
	FROM
		post_table AS p
	LEFT JOIN user_table AS u ON u.user_id = p.user_id
	LEFT JOIN category_table AS c ON c.post_seq = p.post_seq
	WHERE p.post_status = 1
		AND p.is_pinned = 1
	ORDER BY
		p.mod_date DESC
	LIMIT ?
	OFFSET ?;
`


var SelectPinnedPosts = `
	SELECT
		p.post_seq, 
		p.post_title, 
		p.post_contents, 
		c.category_name,
		IFNULL(u.user_name, 'unknown') as user_name,
		p.is_pinned,
		p.viewed,
		p.reg_date, 
		p.mod_date
	FROM
		post_table AS p
	LEFT JOIN user_table AS u ON u.user_id = p.user_id
	LEFT JOIN category_table AS c ON c.post_seq = p.post_seq
	WHERE p.post_status = 1
		AND p.is_pinned = 1
	ORDER BY
		p.mod_date DESC
	LIMIT 5;
`

var SelectUnPinnedPostCount = `
	SELECT COUNT(*) AS count
	FROM post_table
	WHERE post_status = 1
		AND is_pinned = 0
`

var SelectPinnedPostCount = `
	SELECT COUNT(*) AS count
	FROM post_table
	WHERE post_status = 1
		AND is_pinned = 1
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
		t.tags LIKE ? AND
		p.post_status = 1
	ORDER BY p.mod_date DESC
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
		category_name LIKE ? AND
		p.post_status = 1
	ORDER BY p.mod_date DESC
	LIMIT ?
	OFFSET ?;
`



var SelectAllCategories = `

`