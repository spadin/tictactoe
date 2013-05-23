package io

type errCallback func()

type Getter interface {
	GetString() string
	GetInt() int
	GetIntWithCallback(errCallback) int
}
