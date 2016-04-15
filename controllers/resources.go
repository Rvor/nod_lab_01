package controllers

import (
	m "nhaoday.com/models"
)

type (
	PostsResource struct {
		Data m.Posts `json:"data"`
	}
)
