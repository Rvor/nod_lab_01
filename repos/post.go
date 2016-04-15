package repos

import (
	"log"

	_ "github.com/lib/pq"
	m "nhaoday.com/models"
)

func PostList() (m.Posts, error) {
	strSql := "select id, title, description, metadata, tags, createdon from post order by id desc"
	rows, err := db.Query(strSql)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var posts m.Posts
	for rows.Next() {
		var post m.Post
		err = rows.Scan(&post.Id, &post.Title, &post.Description, &post.Metadata, &post.Tags, &post.CreatedOn)
		if err != nil {
			log.Fatal(err)
			return posts, err
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return posts, err
	}
	return posts, err
}

func FindPostById(Id int) (m.Post, error) {
	strSql := "select id, title, description, metadata, tags, createdon from post where id = $1"
	row := db.QueryRow(strSql, Id)
	var post m.Post
	err := row.Scan(&post.Id, &post.Title, &post.Description, &post.Metadata, &post.Tags, &post.CreatedOn)
	if err != nil {
		return post, err
	}
	return post, nil
}

func AddPost(post *m.Post) error {
	strSql := "insert into post(title, description, metadata, tags) values ($1, $2, $3, $4) returning id"
	var insertedId int
	err := db.QueryRow(strSql, post.Title, post.Description, post.Metadata, post.Tags).Scan(&insertedId)
	if err != nil {
		return err
	}
	post.Id = int(insertedId)
	return nil
}

func UpdatePost(id int, p m.Post) error {
	return nil
}
