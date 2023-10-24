package storage

// Post - публикация.
type Post struct {
	ID       int
	Title    string
	Content  string
	PostTime int64
	Link     string
}

// Interface задаёт контракт на работу с БД.
type Interface interface {
	GetPosts(n int) ([]Post, error) // Возвращает список из n публикаций
	AddPost(Post) (int, error)      // Добавляет новую публикацию
}
