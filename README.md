# csvparser

The csvparser is a library to parse cvs file. It can marshal slice of struct to csv and vice-versa. It also has encoder and decoder.

## Examples

### Marshal

#### Write in standard output(terminal/consule)
```
package main

import (
	"fmt"
	"os"

	"github.com/Prithvipal/csvparser"
)

type Book struct {
	ID     int     `csv:"id"`
	Title  string  `csv:"book_title"`
	Price  float32 `csv:"price"`
	Author string  `csv:"name_of_author"`
}

func main() {
	b := []Book{
		{
			ID:     101,
			Title:  "Let us C",
			Price:  1232,
			Author: "Auth one",
		},
		{
			ID:     102,
			Title:  "Hands on go programming",
			Price:  334,
			Author: "Prithvipal Singh",
		},
	}
	enc := csvparser.NewEncoder(os.Stdout)
	err := enc.Encode(b)
	if err != nil {
		fmt.Println(err)
	}

}

```

**Output**
```
id,book_title,price,name_of_author
101,Let us C,1.232E+03,Auth one
102,Hands on go programming,3.34E+02,Prithvipal Singh
```
