package product

var (
	emptyBody = EmptyBody{}
)

type EmptyBody struct{}

type Version struct {
	Version string `json:"version"`
}

type Service interface {
	Version() Version
}

type service struct {
}

func (s service) Version() Version {
	return Version{Version: "v0.1"}
}

// NewService creates a user service with necessary dependencies.
func NewService() Service {
	return &service{}
}
