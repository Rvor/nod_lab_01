package models

import (
	"nhaoday.com/common"
	"time"
)

type (
	Post struct {
		Id          int          `json:"id"`
		Title       string       `json:"title"`
		Description string       `json:"description"`
		Metadata    common.JSONB `json:"metadata" sql:"type:jsonb"`
		CreatedOn   time.Time    `json:"createdon"`
		Tags        common.JSONB `json:"tags" sql:"type:jsonb"`
	}

	Posts []Post
)
