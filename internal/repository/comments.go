package repository

import "database/sql"

type CommentsRepository struct {
	db *sql.DB
}

func NewCommentsRepository(db *sql.DB) *CommentsRepository {
	return &CommentsRepository{db: db}
}
func (c CommentsRepository) SelectAllCommentsByID(userID, postID int) ([]Comment, error) {
	sqlStmt := `
	SELECT comments.id, content, name
	FROM comments
	JOIN users ON comments.author_id = users.id
	WHERE comments.post_id = $1;
`
	rows, err := c.db.Query(sqlStmt, userID, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//ch := make(chan []Comment)
	//errCh := make(chan error)

	comments := []Comment{}
	for rows.Next() {
		var comment Comment

		err = rows.Scan(
			&comment.ID,
			&comment.AuthorID,
			&comment.AuthorName,
			&comment.PostID,
			&comment.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		if comment.AuthorID == userID {
			comment.IsAuthor = true
		}

		comments = append(comments, comment)
	}
	return comments, nil
}
func (c *CommentsRepository) FetchCommentAuthorID(commentID int) (int, error) {
	sqlStmt := `
	select author_id from comments where id = $1
`
	var authorID int
	err := c.db.QueryRow(sqlStmt, commentID).Scan(&authorID)
	switch err {
	case sql.ErrNoRows:
		return 0, nil
	case nil:
		return authorID, nil
	default:
		return 0, err
	}
}
func (c *CommentsRepository) InsertComment(comment Comment) (int64, error) {
	sqlStmt := `
	INSERT INTO comments (post_id,author_id,comment,created_at) VALUES ($1,$2,$3,$4)
`
	var id int
	err := c.db.QueryRow(sqlStmt, comment.PostID, comment.AuthorID, comment.CreatedAt).Scan(&id)

	return int64(id), err
}
func (c CommentsRepository) CountComment(postID int) (int, error) {
	sqlStmt := `SELECT COUNT(*) FROM comments WHERE post_id = $1;`
	result := c.db.QueryRow(sqlStmt, postID)

	var totalLike int
	err := result.Scan(&totalLike)
	if err != nil {
		return 0, err
	}

	return totalLike, nil
}
