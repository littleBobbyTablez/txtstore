package txtstore

import "github.com/google/uuid"

var store = make(map[string]string)

func Write(s string) string {
	id := uuid.NewString()
	store[id] = s
	return id
}

func Read(id string) string {
	return store[id]
}
