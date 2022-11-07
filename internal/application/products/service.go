package products

import "fmt"

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (s *Service) Save() {
	fmt.Println("calling service madafaka")
}
