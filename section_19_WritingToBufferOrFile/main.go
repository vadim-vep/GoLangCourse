package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.first))
	return err
}

func main() {
	p := person{
		first: "James",
	}

	var b bytes.Buffer

	f, err := os.Create("C:/Users/vadim.vepritskiy/Documents/go_lang_course/" +
		"section_19_WritingToBufferOrFile/test.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	p.writeOut(f)
	p.writeOut(&b)
	fmt.Println(b.String())
}
