package services

type BaseService struct {
}

func (s *BaseService) Rules(size int) []string {
	return make([]string, size)
}

func (s *BaseService) Validate(data []string) bool {
	if data != nil {
		return true
	}
	return false
}
