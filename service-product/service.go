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

	CreateProduct(product Product) (Product, error)
	GetProducts() ([]Product, error)
}

type service struct {
	repository Repository
}

func (s service) Version() Version {
	return Version{Version: "v0.1"}
}

func (s service) CreateProduct(product Product) (Product, error) {
	return s.repository.Save(product)
}

func (s service) GetProducts() ([]Product, error) {
	return s.repository.GetProducts()
}

// NewService creates a user service with necessary dependencies.
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
