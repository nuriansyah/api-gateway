package repository

import (
	"database/sql"
	"time"
)

type User struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Role    string  `json:"role"`
	Nrp     string  `json:"nrp"`
	Prodi   string  `json:"prodi"`
	Program string  `json:"institute"`
	Company *string `json:"major"`
	Batch   *int    `json:"batch"`
	Avatar  *string `json:"avatar"`
}

type Posts struct {
	ID            int            `db:"id"`
	AuthorID      int            `db:"author_id"`
	AuthorName    string         `db:"author_name"`
	AuthorRole    string         `db:"author_role"`
	AuthorAvatar  sql.NullString `db:"author_avatar"`
	AuthorProgram sql.NullString `db:"author_program"`
	AuthorCompany sql.NullString `db:"author_company"`
	AuthorBatch   sql.NullInt32  `db:"author_batch"`
	Title         string         `db:"title"`
	Description   string         `db:"description"`
	CreatedAt     time.Time      `db:"created_at"`
}
type Comment struct {
	ID         int        `json:"id"`
	PostID     int        `json:"post_id"`
	Comment    string     `json:"comment"`
	CreatedAt  *time.Time `json:"created_at"`
	AuthorID   int        `json:"author_id"`
	AuthorName string     `json:"author_name" db:"author_name"`
	IsAuthor   bool       `json:"is_author"`
}
