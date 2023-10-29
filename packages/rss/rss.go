package rss

import (
	"GoNews/packages/storage"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"time"
)

type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Chanel  Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Link        string `xml:"link"`
}

// Parse читает rss-поток и возвращет массив раскодированных публикаций.
func Parse(url string) ([]storage.Post, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Ошибка при получении данных: %v", err)
		return nil, err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Ошибка при чтении тела ответа: %v", err)
		return nil, err
	}
	var f Feed
	err = xml.Unmarshal(b, &f)
	if err != nil {
		log.Printf("Ошибка при парсинге XML: %v", err)
		return nil, err
	}
	var data []storage.Post
	for _, item := range f.Chanel.Items {
		var p storage.Post
		p.Title = item.Title
		p.Content = item.Description
		p.Link = item.Link

		// Список возможных форматов даты
		formats := []string{
			"Mon, 2 Jan 2006 15:04:05 MST",
			"Mon, 2 Jan 2006 15:04:05 -0700",
		}

		var t time.Time
		is_parsed := false // Флаг, указывающий на успешный парсинг
		for _, format := range formats {
			t, err = time.Parse(format, item.PubDate)
			if err == nil {
				is_parsed = true // Устанавливаем флаг в true, если парсинг успешен
				break            // Прерываем цикл, если парсинг прошел успешно
			}
		}

		if !is_parsed {
			log.Printf("Ошибка при парсинге времени для '%s'", item.Title)
			continue // Пропускаем элемент, если время не удалось распарсить
		}

		p.PubDate = t

		data = append(data, p)
	}
	return data, nil
}
