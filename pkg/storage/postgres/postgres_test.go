package postgres

import (
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/postgres/dbconfig"
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

/*
Тесты postgres являются интеграционными тестами без проверок возвращаемых значений и
необходимы только для проверки базовой работоспособности вызовов в БД
*/
func TestStorage_AddPost(t *testing.T) {

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
		t storage.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Тест записи новой публикации",
			fields: fields{
				db: db,
			},
			args: args{
				t: storage.Post{
					Title:    "Effective Go",
					Content:  "Go is a new language. Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from programs written in its relatives. A straightforward translation of a C++ or Java program into Go is unlikely to produce a satisfactory result—Java programs are written in Java, not Go. On the other hand, thinking about the problem from a Go perspective could produce a successful but quite different program. In other words, to write Go well, it's important to understand its properties and idioms. It's also important to know the established conventions for programming in Go, such as naming, formatting, program construction, and so on, so that programs you write will be easy for other Go programmers to understand.",
					PostTime: 1634847432,
					Link:     "https://gonews.xyz",
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
				fmt.Printf("Storage.AddPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

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
