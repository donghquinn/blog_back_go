package queries


var CreatePostTable = `
	CREATE TABLE IF NOT EXISTS post_table (
		post_seq 		INT(20)		    NOT NULL AUTO_INCREMENT PRIMARY KEY,
		user_id 		VARCHAR(50)		NOT NULL REFERENCES user_table(user_id),
		post_title 		VARCHAR(50)		NOT NULL,
		post_contents 	VARCHAR(5000)	NOT NULL,
		post_status		TINYINT(1)		NOT NULL DEFAULT 1	COMMENT '0: 비활성, 1: 활성, 2: 삭제',
		reg_date 		DATETIME		NOT NULL DEFAULT CURRNET_TIMESTAMP,
		mod_date		DATETIME	    NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

		INDEX post_idx(post_status)
	);
`

var CreateUserTable = `
	CREATE TABLE IF NOT EXISTS user_table (
		user_id 		VARCHAR(50)			NOT NULL	PRIMARY KEY,
		user_email 		VARCHAR(100)		NOT NULL	UNIQUE,
		user_password 	VARCHAR(150)		NOT NULL,
		user_name 		VARCHAR(50)			NOT NULL,
		user_status		TINYINT(1)			NOT NULL 	DEFAULT 1 COMMENT '0: 비활성, 1: 활성, 2: 탈퇴',
		reg_date 		DATETIME			NOT NULL	DEFAULT CURRNET_TIMESTAMP,
		mod_date		DATETIME		    NULL 		DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

		INDEX user_idx(user_status)
	);
`
//	mod_date    DATETIME        NULL        DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
