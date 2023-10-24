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

func (s *Storage) Close() {
	s.db.Close()
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

// Добавляет новую публикацию
func (s *Storage) AddPost(t storage.Post) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
	SELECT add_post($1, $2, $3, $4);
		`,
		t.Title,
		t.Content,
		t.PostTime,
		t.Link,
	).Scan(&id)
	return id, err
}

// Возвращает список из n публикаций
func (s *Storage) GetPosts(n int) ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(), `
	SELECT p.id,
		   p.title,
		   p.content,
		   p.post_time,
		   p.link
	FROM get_posts($1) p;
	`,
		n,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := make([]storage.Post, 0, n) // Заранее выделяем память под 10 публикаций

	for rows.Next() {
		var t storage.Post
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&t.PostTime,
			&t.Link,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, t)
	}
	return posts, rows.Err()
}
