package users

import "sync"

type User struct {
	Name string `json:"name"`
}

var CacheMutex sync.RWMutex
var UserCache = make(map[int]User)
