package repository

import (
	e "main/domain/errors"
	"main/domain/model"
	"time"

	"github.com/lib/pq"
)

func (s *Store) AddPost(classID int, createTime time.Time, newPost *model.PostCreate) (int, error) {
	var id int
	if err := s.db.QueryRow(
		`INSERT INTO posts (classID, text, attaches, createTime)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id;`,
		classID, newPost.Text, (*pq.StringArray)(&newPost.Attaches), createTime,
	).Scan(&id); err != nil {
		return 0, e.StacktraceError(err)
	}

	return int(id), nil
}

func (s *Store) DeletePost(id int) error {
	_, err := s.db.Exec(
		`DELETE FROM posts WHERE id = $1;`,
		id,
	)

	if err != nil {
		return e.StacktraceError(err)
	}

	return nil
}

func (s *Store) GetClassFeed(classID int) (*model.Feed, error) {
	rows, err := s.db.Query(
		`SELECT id, text, attaches, createTime FROM posts WHERE classID = $1;`,
		classID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	posts := []*model.Post{}
	for rows.Next() {
		var tmpPost model.Post

		if err := rows.Scan(
			&tmpPost.ID, &tmpPost.Text,
			(*pq.StringArray)(&tmpPost.Attaches), &tmpPost.CreateTime,
		); err != nil {
			return nil, e.StacktraceError(err)
		}

		posts = append(posts, &tmpPost)
	}

	return &model.Feed{Posts: posts}, nil
}
