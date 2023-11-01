package rss

import (
	"testing"
)

func TestParse(t *testing.T) {
	feed, err := Parse("http://static.feed.rbc.ru/rbc/logical/footer/news.rss")
	if err != nil {
		t.Fatal(err)
	}
	if len(feed) == 0 {
		t.Fatal("данные не раскодированы")
	}
	t.Logf("получено %d новостей\n%+v", len(feed), feed)
}
