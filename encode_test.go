package csvparser

import (
	"strconv"
	"testing"
)

type Book struct {
	ID     int     `csv:"id"`
	Title  string  `csv:"book_title"`
	Price  float32 `csv:"price"`
	Author string  `csv:"name_of_author"`
}

func TestMarshal(t *testing.T) {
	books := make([]Book, 0)
	book1 := Book{ID: 101, Title: "Hands on go programming", Price: 120.45, Author: "Prithvi"}
	books = append(books, book1)
	book2 := Book{ID: 102, Title: "Hands on rust programming", Price: 490.95, Author: "Parthiv"}
	books = append(books, book2)
	data := Marshal(books)
	actualLen := len(data)
	wantLen := 3
	if actualLen != wantLen {
		t.Errorf("Marshal lenght = %q, want %q", actualLen, wantLen)
	}

	actualTitleTag := data[0][1]
	wantTitleTag := "book_title"
	if actualTitleTag != wantTitleTag {
		t.Errorf("Marshal lenght = %q, want %q", actualTitleTag, wantTitleTag)
	}

	actualBook1Price := data[1][2]
	price := 120.44999694824219
	wantBook1Price := strconv.FormatFloat(price, 'E', -1, 64)
	if actualBook1Price != wantBook1Price {
		t.Errorf("Marshal lenght = %q, want %q", actualBook1Price, wantBook1Price)
	}

	actualBook2Author := data[2][3]

	wantBook2Author := "Parthiv"
	if actualBook2Author != wantBook2Author {
		t.Errorf("Marshal lenght = %q, want %q", actualBook2Author, wantBook2Author)
	}
}
