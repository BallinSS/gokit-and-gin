package service

// Service 定义服务接口
type Service interface {
	Hello(name string) (string, error)
}

// service 实现服务接口
type service struct{}

// NewService 创建新服务
func NewService() Service {
	return &service{}
}

func (s *service) Hello(name string) (string, error) {
	return "Hello, " + name, nil
}
