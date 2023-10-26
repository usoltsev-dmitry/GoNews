package main

import (
	"GoNews/packages/api"
	"GoNews/packages/storage"
	"GoNews/packages/storage/postgres"
	"log"
	"net/http"
	//"GoNews/pkg/storage/memdb"
)

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	// Создаём объект сервера.
	var srv server

	// Реляционная БД PostgreSQL.
	db, err := postgres.New("host=172.22.0.2 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	/*
		// БД в памяти.
		db2 := memdb.New()
	*/

	// Инициализируем хранилище сервера конкретной БД.
	srv.db = db

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	// Запускаем веб-сервер на порту 8080 на всех интерфейсах.
	// Предаём серверу маршрутизатор запросов,
	// поэтому сервер будет все запросы отправлять на маршрутизатор.
	// Маршрутизатор будет выбирать нужный обработчик.
	if err := http.ListenAndServe(":8080", srv.api.Router()); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
