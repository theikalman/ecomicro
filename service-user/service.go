package user

var (
	emptyBody = EmptyBody{}
)

type EmptyBody struct{}

type Version struct {
	Version string `json:"version"`
}

type Service interface {
	Version() Version
	Signup(u User) (User, error)
}

type service struct {
	repository Repository
}

func (s service) Version() Version {
	return Version{Version: "v0.1"}
}

func (s service) Signup(u User) (User, error) {
	return s.repository.Save(u)
}

// NewService creates a user service with necessary dependencies.
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
