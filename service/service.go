package service

import (
	"domain/model"
	"regexp"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CheckDomainService(model model.Domain, ch chan bool) {
	regExp := regexp.MustCompile(`^(www)\.(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})$`)

	ch <- regExp.MatchString(model.Url)
}
