package common

import "log"

type DefaultSaver struct {
}

func (s *DefaultSaver) Save(ct string) {
	log.Println(ct)
}
