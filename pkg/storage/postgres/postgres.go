package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"GoNews/pkg/storage"
)

// Хранилище данных
type Storage struct {
	db *pgxpool.Pool
}

// Конструктор
func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

// Возвращает список всех публикаций
func (s *Storage) GetPosts() ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(), `
	SELECT p.id,
		   p.author_id,
		   p.author,
		   p.title,
		   p.content,
		   p.created_at,
		   p.updated_at
	FROM get_posts() p;
	`,
	)
	if err != nil {
		return nil, err
	}
	var posts []storage.Post
	for rows.Next() {
		var t storage.Post
		err = rows.Scan(
			&t.ID,
			&t.AuthorID,
			&t.Author,
			&t.Title,
			&t.Content,
			&t.Created,
			&t.Updated,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, t)
	}
	return posts, rows.Err()
}

// Добавляет нового автора и возвращает его id
func (s *Storage) AddAuthor(t storage.Post) (int64, error) {
	var id int64
	err := s.db.QueryRow(context.Background(), `
	SELECT add_author($1);
		`,
		t.Author,
	).Scan(&id)
	return id, err
}

// Добавляет новую статью и возвращает её id
func (s *Storage) AddPost(t storage.Post) (int64, error) {
	var id int64
	err := s.db.QueryRow(context.Background(), `
	SELECT add_post($1, $2, $3);
		`,
		t.AuthorID,
		t.Title,
		t.Content,
	).Scan(&id)
	return id, err
}

// Обновляет статью
func (s *Storage) UpdatePost(t storage.Post) (bool, error) {
	var is_updated bool
	err := s.db.QueryRow(context.Background(), `
	SELECT update_post($1, $2, $3);
		`,
		t.AuthorID,
		t.Title,
		t.Content,
	).Scan(&is_updated)
	return is_updated, err
}

// Удаляет статью
func (s *Storage) DeletePost(t storage.Post) (bool, error) {
	var is_deleted bool
	err := s.db.QueryRow(context.Background(), `
	SELECT delete_post($1);
		`,
		t.ID,
	).Scan(&is_deleted)
	return is_deleted, err
}
