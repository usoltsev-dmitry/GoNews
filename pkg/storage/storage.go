package storage

// Post - публикация.
type Post struct {
	ID       int64
	AuthorID int64
	Author   string
	Title    string
	Content  string
	Created  int64
	Updated  int64
}

// Interface задаёт контракт на работу с БД.
type Interface interface {
	GetPosts() ([]Post, error)     // получение публикации по id публикации
	AddAuthor(Post) (int64, error) // добавление нового автора
	AddPost(Post) (int64, error)   // добавление новой публикации
	UpdatePost(Post) (bool, error) // обновление публикации
	DeletePost(Post) (bool, error) // удаление публикации по ID
}
