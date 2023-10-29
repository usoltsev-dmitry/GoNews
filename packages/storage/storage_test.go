package storage

import (
	"GoNews/packages/storage/postgres/dbconfig"
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func TestDB_AddPosts(t *testing.T) {
	// Получаем строку соединения с БД
	connStr := dbconfig.ConnString()

	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		t.Fatalf("Ошибка подключения к БД: %v", err)
	}

	type fields struct {
		pool *pgxpool.Pool
	}
	type args struct {
		posts []Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Тест добавления публикаций",
			fields: fields{pool: db},
			args: args{posts: []Post{
				{Title: "Новость 1", Content: "Содержание 1", PubDate: time.Date(2023, 10, 26, 15, 0, 0, 0, time.UTC), Link: "wwww.gonews.xyz"},
				{Title: "Новость 2", Content: "Содержание 2", PubDate: time.Date(2023, 10, 26, 16, 0, 0, 0, time.UTC), Link: "wwww.gonews.xyz"},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DB{
				pool: tt.fields.pool,
			}
			if err := s.AddPosts(tt.args.posts); (err != nil) != tt.wantErr {
				t.Errorf("DB.AddPosts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_GetPosts(t *testing.T) {
	// Получаем строку соединения с БД
	connStr := dbconfig.ConnString()

	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		t.Fatalf("Ошибка подключения к БД: %v", err)
	}

	type fields struct {
		pool *pgxpool.Pool
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
			fields:  fields{pool: db},
			args:    args{n: 10},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DB{
				pool: tt.fields.pool,
			}
			_, err := s.GetPosts(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.GetPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
