package naming

// Service 定义了基础服务的抽象接口
type Service interface {
	ServiceID() string
	ServiceName() string
	GetMeta() map[string]string
}

// ServiceRegistration 定义服务注册的抽象接口
type ServiceRegistration interface {
	Service
	PublicAddress() string
	PublicPort() uint64
	GetNamespace() string
	GroupName() string
}

type DefaultService struct {
	Id        string
	Name      string
	Address   string
	Port      uint64
	Namespace string
	Group     string
	Meta      map[string]string
}

func (s DefaultService) ServiceID() string {
	return s.Id
}

func (s DefaultService) ServiceName() string {
	return s.Name
}

func (s DefaultService) GetMeta() map[string]string {
	return s.Meta
}

func (s DefaultService) PublicAddress() string {
	return s.Address
}

func (s DefaultService) PublicPort() uint64 {
	return s.Port
}

func (s DefaultService) GetNamespace() string {
	return s.Namespace
}

func (s DefaultService) GroupName() string {
	return s.Group
}
