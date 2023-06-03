package shared

type Service struct {
	Subject string
	Name    string
	Version string
	Fn      Wrapper
}

var (
	Registry = []Service{}
)

func RegisterService(s Service) {
	Registry = append(Registry, s)
}
