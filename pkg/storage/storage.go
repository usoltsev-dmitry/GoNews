package storage

// Post - публикация.
type Post struct {
	ID          int
	Title       string
	Content     string
	AuthorID    int
	AuthorName  string
	CreatedAt   int64
	PublishedAt int64
}

// Interface задаёт контракт на работу с БД.
type Interface interface {
	Posts() ([]Post, error) // получение всех публикаций
	AddPost(Post) error     // создание новой публикации
	UpdatePost(Post) error  // обновление публикации
	DeletePost(Post) error  // удаление публикации по ID
}
