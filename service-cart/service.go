package cart

var (
	emptyBody = EmptyBody{}
)

type EmptyBody struct{}

type Service interface {
	Version() string
}

type service struct {
}

func (s service) Version() string {
	return "v0.1"
}

// NewService creates a user service with necessary dependencies.
func NewService() Service {
	return &service{}
}
