package domain

type Service interface {
	Find(code string) (*Employee, error)
	Store(product *Employee) error
	Update(product *Employee) error
	FindAll() ([]*Employee, error)
	Delete(code string) error
}

//Repository ...
type Repository interface {
	Find(code string) (*Employee, error)
	Store(product *Employee) error
	Update(product *Employee) error
	FindAll() ([]*Employee, error)
	Delete(code string) error
}

type service struct {
	employee Repository
}

//NewProductService ...
func NewEmployeeService(employee Repository) Service {
	return &service{employee: employee}
}

func (s *service) Find(noUser string) (*Employee, error) {
	return s.employee.Find(noUser)
}

func (s *service) Store(product *Employee) error {
	return s.employee.Store(product)
}
func (s *service) Update(product *Employee) error {
	return s.employee.Update(product)
}

func (s *service) FindAll() ([]*Employee, error) {
	return s.employee.FindAll()
}

func (s *service) Delete(code string) error {

	return s.employee.Delete(code)
}
