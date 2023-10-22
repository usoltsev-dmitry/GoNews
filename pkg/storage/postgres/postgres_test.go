package postgres

import (
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/postgres/dbconfig"
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Тесты postgres являются интеграционными без проверок возвращаемых значений и полезны только для проверки базовой работоспособности вызовов в БД

func TestStorage_AddPost(t *testing.T) {

	// Загружаем файл конфигурацию из файла
	config, err := dbconfig.LoadConfig("dbconfig/dbconfig.json")
	if err != nil {
		t.Fatalf("Failed to load DB config: %v", err)
	}

	// Получаем строку соединения с БД
	connStr := config.ConnString()

	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}
	defer db.Close()

	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		t storage.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Add valid post",
			fields: fields{
				db: db,
			},
			args: args{
				t: storage.Post{
					Title:    "Test Title",
					Content:  "Test Content",
					PostTime: 1634847432,
					Link:     "http://testlink.com",
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
			_, err := s.AddPost(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.AddPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestStorage_GetPosts(t *testing.T) {
	// Загружаем файл конфигурацию из файла
	config, err := dbconfig.LoadConfig("dbconfig/dbconfig.json")
	if err != nil {
		t.Fatalf("Failed to load DB config: %v", err)
	}

	// Получаем строку соединения с БД
	connStr := config.ConnString()

	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
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
			name:    "Test fetching 1 posts",
			fields:  fields{db: db},
			args:    args{n: 1},
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
