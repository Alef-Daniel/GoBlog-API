package post_repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/goblog-api/internal/domain/entities"
	"github.com/goblog-api/internal/domain/repositories"
)

type PostRepository struct {
	conn *sql.DB
}

func (p *PostRepository) CreatePost(ctx context.Context, post *entities.Post) (*entities.Post, error) {
	query := `INSERT INTO post (id, title, categoryID,author,content) VALUES (?, ?, ?, ?, ?)`
	stmt, err := p.conn.Prepare(query)
	if err != nil {
		return nil, repositories.ErrRepository
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, post.ID, post.Title, post.CategoryID, post.Author, post.Content)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			if mysqlErr.Number == 1062 {
				return nil, repositories.ErrAlreadyExists
			}
		}
		return nil, repositories.ErrRepository
	}

	post, err = p.GetByID(ctx, post.ID)
	if err != nil {
		return nil, err
	}

	return post, nil

}

func (p *PostRepository) GetByID(ctx context.Context, id int64) (*entities.Post, error) {
	query := `SELECT id,createdAt,updateAt,deleteAt, title, categoryID, author, content FROM post WHERE id=? AND deleted_at IS NULL`
	stmt, err := p.conn.Prepare(query)
	if err != nil {
		return nil, repositories.ErrNotFound
	}
	defer stmt.Close()
	var post entities.Post
	err = stmt.QueryRowContext(ctx, id).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.DeletedAt,
		&post.Title,
		&post.CategoryID,
		&post.Author,
		&post.Content)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repositories.ErrNotFound
		}
		return nil, repositories.ErrNotFound
	}

	return &post, nil
}

func (p *PostRepository) GetAllPost(ctx context.Context) (*[]entities.Post, error) {
	query := `SELECT id,createdAt,updateAt,deleteAt, title, categoryID, author, content FROM post WHERE deleted_at IS NULL`
	var posts []entities.Post
	rows, err := p.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, repositories.ErrRetrieve
	}

	defer rows.Close()

	for rows.Next() {
		var post entities.Post
		err = rows.Scan(
			&post.ID,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.DeletedAt,
			&post.Title,
			&post.CategoryID,
			&post.Author,
			&post.Content,
		)
		if err != nil {
			return nil, repositories.ErrRetrieve
		}

		posts = append(posts, post)
	}

	return &posts, nil
}

func (p *PostRepository) UpdatePost(ctx context.Context, post *entities.Post) (*entities.Post, error) {
	query := `UPDATE post SET title=?, categoryID=?, author=?, content=? WHERE id=?`
	stmt, err := p.conn.Prepare(query)
	if err != nil {
		return nil, repositories.ErrUpdate
	}
	defer stmt.Close()

	res, err := stmt.Exec(post.Title, post.CategoryID, post.Author, post.Content, post.ID)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			if mysqlErr.Number == 1062 {
				return nil, repositories.ErrAlreadyExists
			} else if mysqlErr.Number == 1452 {
				return nil, repositories.ErrForeignKeyViolation
			}
		}
		return nil, repositories.ErrUpdate
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return nil, repositories.ErrNotFound
	}

	post, err = p.GetByID(ctx, post.ID)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *PostRepository) DeletePost(ctx context.Context, post *entities.Post) error {
	query := `UPDATE post SET deletedAt=CURRENT_TIMESTAMP WHERE id=? AND deletedAt IS NULL`
	stmt, err := p.conn.Prepare(query)
	if err != nil {
		return repositories.ErrDelete
	}

	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, post.ID)
	if err != nil {
		return repositories.ErrDelete
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return repositories.ErrNotFound
	}

	return nil

}

func NewPostRepository(conn *sql.DB) *PostRepository {
	return &PostRepository{
		conn: conn,
	}
}
