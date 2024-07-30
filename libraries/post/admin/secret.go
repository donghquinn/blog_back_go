package post

import (
	"github.com/donghquinn/blog_back_go/libraries/database"
	queries "github.com/donghquinn/blog_back_go/queries/admin/posts"
)

func ChangeToSecretPost(postSeq string) error {
	dbconn, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		return dbErr
	}

	_, insertErr := database.InsertQuery(dbconn, queries.ChangeToSecretPostQuery, postSeq)

	if insertErr != nil {
		return insertErr
	}

	defer dbconn.Close()

	return nil
}

func ChangeToNotSecretPost(postSeq string) error {
	dbconn, dbErr := database.InitDatabaseConnection()

	if dbErr != nil {
		return dbErr
	}

	_, insertErr := database.InsertQuery(dbconn, queries.ChangeToNotSecretPostQuery, postSeq)

	if insertErr != nil {
		return insertErr
	}

	defer dbconn.Close()

	return nil
}