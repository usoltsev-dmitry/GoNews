package storage

import (
	"GoNews/packages/storage/postgres/dbconfig"
	"context"
	"encoding/json"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// База данных
type DB struct {
	pool *pgxpool.Pool
}

// Публикация, получаемая из RSS.
type Post struct {
	ID      int       // id записи
	Title   string    // заголовок публикации
	Content string    // содержание публикации
	PubDate time.Time // время публикации
	Link    string    // ссылка на источник
}

func New() (*DB, error) {
	// Получаем строку соединения с БД
	connStr := dbconfig.ConnString()

	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	db := DB{
		pool: pool,
	}
	return &db, nil
}

// Добавляет список публикаций
func (s *DB) AddPosts(posts []Post) error {
	tx, err := s.pool.Begin(context.Background())
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
func (s *DB) GetPosts(n int) ([]Post, error) {

	rows, err := s.pool.Query(context.Background(), `
	SELECT p.id,
		   p.title,
		   p.content,
		   p.pubdate,
		   p.link
	FROM get_posts($1) p;
	`,
		n,
	)
	if err != nil {
		return nil, err
	}

	posts := make([]Post, 0, n) // Заранее выделяем память под n публикаций

	for rows.Next() {
		var t Post
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&t.PubDate,
			&t.Link,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, t)
	}
	return posts, rows.Err()
}
