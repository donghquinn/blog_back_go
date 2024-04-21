package queries


var CreatePostTable = `
	CREATE TABLE IF NOT EXISTS post_table (
		post_seq 		INT(20)		    NOT NULL AUTO_INCREMENT PRIMARY KEY,
		user_id 		VARCHAR(50)		NOT NULL REFERENCES user_table(user_id),
		post_title 		VARCHAR(50)		NOT NULL,
		post_contents 	TEXT			NOT NULL,
		post_status		TINYINT(1)		NOT NULL DEFAULT 1	COMMENT '0: 비활성, 1: 활성, 2: 삭제',
		reg_date 		DATETIME		NOT NULL DEFAULT CURRENT_TIMESTAMP,
		mod_date		DATETIME	    NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

		INDEX post_idx(post_status)
	);
`

var CreateUserTable = `
	CREATE TABLE IF NOT EXISTS user_table (
		user_id 		VARCHAR(50)			NOT NULL	PRIMARY KEY,
		user_email 		VARCHAR(100)		NOT NULL	UNIQUE,
		user_password 	VARCHAR(200)		NOT NULL,
		user_name 		VARCHAR(100)		NOT NULL,
		user_status		TINYINT(1)			NOT NULL 	DEFAULT 1 COMMENT '0: 비활성, 1: 활성, 2: 탈퇴',
		reg_date 		DATETIME			NOT NULL	DEFAULT CURRENT_TIMESTAMP,
		mod_date		DATETIME		    NULL 		DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);
`

var CreateFileTable = `
	CREATE TABLE IF NOT EXISTS file_table (
		file_seq 		BIGINT(20) 			NOT NULL 	AUTO_INCREMENT PRIMARY KEY,
		file_type		TINYINT(1)			NOT NULL 	COMMENT '1 - 이미지, 2 - 비디오',
		target_seq 		INT(20)				NULL	 	COMMENT '이미지 타겟 포스트 시퀀스' REFERENCES post_table(post_seq),
		target_id		VARCHAR(50)			NOT NULL	COMMENT '이미지 타겟 유저 id - 프로필 / 업로드 유저' REFERENCES user_table(user_id),
		target_table 	VARCHAR(20)			NOT NULL,
		file_size		INT(20)				NOT NULL,
		version_id 		VARCHAR(50)			NOT NULL,
		reg_date 		DATETIME			NOT NULL	DEFAULT CURRENT_TIMESTAMP,

		INDEX target_idx(target_table)
	)
`

//	mod_date    DATETIME        NULL        DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
