package storage

type Archiver[T any] interface {
	Load() (T, error)
	LoadOrDefault(data T) T
	Save(data T) error
}
