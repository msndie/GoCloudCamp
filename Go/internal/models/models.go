package models

type ConfigEntity struct {
	service string
	data    []PropertyEntity
}

type PropertyEntity struct {
	key   string
	value string
}
