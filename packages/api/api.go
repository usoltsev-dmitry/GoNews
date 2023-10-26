package api

import (
	"GoNews/packages/storage"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Программный интерфейс сервера GoNews
type API struct {
	db     storage.Interface
	router *mux.Router
}

// Конструктор объекта API
func New(db storage.Interface) *API {
	api := API{
		db: db,
	}
	api.router = mux.NewRouter()
	api.endpoints()
	return &api
}

// Получение маршрутизатора запросов
func (api *API) Router() *mux.Router {
	return api.router
}

// Регистрация обработчиков API
func (api *API) endpoints() {
	api.router.Use(api.HeadersMiddleware)
	api.router.HandleFunc("/posts/{n}", api.getPostsHandler).Methods(http.MethodGet)
	api.router.HandleFunc("/posts", api.addPostsHandler).Methods(http.MethodPost)
}

// Обработчик получения списка из n публикаций
func (api *API) getPostsHandler(w http.ResponseWriter, r *http.Request) {
	// Считывание параметра {n} из пути запроса.
	s := mux.Vars(r)["n"]
	n, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Получение данных из БД
	posts, err := api.db.GetPosts(n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка данных клиенту в формате JSON
	json.NewEncoder(w).Encode(posts)
}

// Обработчик добавление списка публикаций
func (api *API) addPostsHandler(w http.ResponseWriter, r *http.Request) {
	var posts []storage.Post
	err := json.NewDecoder(r.Body).Decode(&posts)
	if err != nil {
		http.Error(w, "Ошибка декодирования JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = api.db.AddPosts(posts)
	if err != nil {
		http.Error(w, "Ошибка добавления публикаций: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка клиенту статуса успешного выполнения запроса
	w.WriteHeader(http.StatusCreated)
}

func (api *API) HeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
