package myhttp

import (
	"io"
	"log"
)

func WriteAll(w io.Writer, buf []byte) {
	w.Write(buf)
}

func WriteAll1(w io.Writer, buf []byte) error {
	_, err := w.Write(buf)
	if nil != err {
		log.Println("unable to write:", err)
		return err
	}

	return nil
}
