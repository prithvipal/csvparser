# csvparser

The csvparser is a library to parse cvs file. It can marshal slice of struct to csv and vice-versa. It also has encoder and decoder.

### Marshal

```
package main

import (
	"fmt"

	"github.com/Prithvipal/csvparser"
)

type Book struct {
	ID     int     `csv:"id"`
	Title  string  `csv:"book_title"`
	Price  float32 `csv:"price"`
	Author string  `csv:"name_of_author"`
}

func main() {
	books := make([]Book, 0)
	book1 := Book{ID: 101, Title: "Let us C", Price: 120.45, Author: "Prithvi"}
	books = append(books, book1)
	book2 := Book{ID: 102, Title: "Let us Go", Price: 490.95, Author: "Parthiv"}
	books = append(books, book2)
	data := csvparser.Marshal(books)
	fmt.Println(data)
}
```

**Output**
```
[[id book_title price name_of_author] 
[101 Let us C 1.2044999694824219E+02 Prithvi] 
[102 Let us Go 4.9095001220703125E+02 Parthiv]]
```
