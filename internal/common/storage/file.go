package storage

import (
	"encoding/gob"
	"log"
	"os"
	"path"
)

func NewFile[T any](p string) Archiver[T] {
	if p == "" {
		p = path.Join(os.TempDir(), "packapi.data")
	}

	return &file[T]{p}
}

type file[T any] struct {
	path string
}

func (f *file[T]) Load() (T, error) {
	var res T
	df, err := os.Open(f.path)
	if err != nil {
		return res, err
	}

	dec := gob.NewDecoder(df)
	err = dec.Decode(&res)
	log.Printf("loaded data: %+v", res)

	return res, err
}

func (f *file[T]) LoadOrDefault(data T) T {
	d, err := f.Load()
	if err != nil {
		return data
	}

	return d
}

func (f *file[T]) Save(data T) error {
	wf, err := os.OpenFile(f.path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	log.Printf("saving data: %+v", data)
	enc := gob.NewEncoder(wf)

	return enc.Encode(data)
}
