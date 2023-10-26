package postgres

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v4/pgxpool"

	"GoNews/packages/storage"
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

// Добавляет список публикаций
func (s *Storage) AddPosts(posts []storage.Post) error {
	tx, err := s.db.Begin(context.Background())
	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	// Преобразуем список публикаций в JSON
	postsJSON, err := json.Marshal(posts)
	if err != nil {
		return err
	}

	// Вызываем процедуру добавления публикаций
	_, err = tx.Exec(context.Background(), `CALL add_posts($1);`, postsJSON)
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
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
