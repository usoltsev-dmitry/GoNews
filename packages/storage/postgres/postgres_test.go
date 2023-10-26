package postgres

import (
	"GoNews/packages/storage"
	"GoNews/packages/storage/postgres/dbconfig"
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

/*
Тесты postgres являются интеграционными тестами без проверок возвращаемых значений и
необходимы только для проверки базовой работоспособности вызовов в БД
*/

func TestStorage_GetPosts(t *testing.T) {
	// Загружаем файл конфигурацию из файла
	config, err := dbconfig.LoadConfig("dbconfig/dbconfig.json")
	if err != nil {
		t.Fatalf("Ошибка загрузки файла конфигурации подключения к БД: %v", err)
	}

	// Получаем строку соединения с БД
	connStr := config.ConnString()

	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		t.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Тест получения списка из 10 последних публикаций",
			fields:  fields{db: db},
			args:    args{n: 10},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			_, err := s.GetPosts(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestStorage_AddPosts(t *testing.T) {
	// Загружаем файл конфигурацию из файла
	config, err := dbconfig.LoadConfig("dbconfig/dbconfig.json")
	if err != nil {
		t.Fatalf("Ошибка загрузки файла конфигурации подключения к БД: %v", err)
	}

	// Получаем строку соединения с БД
	connStr := config.ConnString()

	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		t.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		posts []storage.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Добавление пустого списка публикаций",
			fields: fields{
				db: db,
			},
			args: args{
				posts: []storage.Post{},
			},
			wantErr: false,
		},
		{
			name: "Добавление одной публикации",
			fields: fields{
				db: db,
			},
			args: args{
				posts: []storage.Post{
					{Title: "Test Title", Content: "Test Content", PostTime: 1234567890, Link: "http://example.com"},
				},
			},
			wantErr: false,
		},
		{
			name: "Добавление нескольких публикаций",
			fields: fields{
				db: db,
			},
			args: args{
				posts: []storage.Post{
					{Title: "Title 1", Content: "Content 1", PostTime: 1234567890, Link: "http://example1.com"},
					{Title: "Title 2", Content: "Content 2", PostTime: 1234567891, Link: "http://example2.com"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			if err := s.AddPosts(tt.args.posts); (err != nil) != tt.wantErr {
				t.Errorf("Storage.AddPosts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
