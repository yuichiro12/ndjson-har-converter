package main

import (
	"bufio"
	"encoding/json"
	"github.com/yuichiro12/har"
	"io"
	"log"
	"os"
)

func main() {
	opt := os.Args[1]
	fn := os.Args[2]
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	switch (opt) {
	case "h2e":
		if err = Har2Entries(f, os.Stdout); err != nil {
			log.Fatal(err)
		}
	case "e2h":
		if err := Entries2Har(f, os.Stdout); err != nil {
			log.Fatal(err)
		}
	}
}

func Har2Entries(r io.Reader, w io.Writer) error {
	l, err := har.NewLog(r)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(w)
	for _, e := range l.Entries {
		if err := enc.Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func Entries2Har(r io.Reader, w io.Writer) error {
	var ets []har.Entry
	br := bufio.NewReader(r)
	for {
		b, err := br.ReadBytes(byte('\n'))
		if err == io.EOF || string(b) == "" {
			break
		}
		var et har.Entry
		if err := json.Unmarshal(b, &et); err != nil {
			return err
		}
		ets = append(ets, et)
	}
	l := &har.Log{
		Version: "1.2",
		Creator: &har.Creator{
			Name:    "harconv",
			Version: "0.0.0",
		},
		Entries: ets,
	}
	if err := har.Dump(w, l); err != nil {
		return err
	}

	return nil
}