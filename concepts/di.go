package concepts

import "fmt"

type UserRepository interface {
	GetUser(id int) string
}

type MySqlRepo struct{}
type PostgresRepo struct{}

func (m *MySqlRepo) GetUser(id int) string {
	return "User Returned from the mysql db"
}

func (m *PostgresRepo) GetUser(id int) string {
	return "User Returned from the Postgres db"
}

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) PrintUser(id int) {
	fmt.Println(s.repo.GetUser(id))
}
