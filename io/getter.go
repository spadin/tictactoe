package io

type errCallback func()

type Getter interface {
	GetString() (input string)
	GetInt() (input int)
	GetIntWithCallback(cb errCallback) (input int)
}
