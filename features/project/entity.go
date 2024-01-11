package project

// struct user gorm model
type User struct {
	Name        string
	Email       string
	Address     string
	PhoneNumber string
	Role        string
}

// struct task model
type Task struct {
	Task      string
	ProjectID uint
	Status    string
}

// struct project model
type Core struct {
	ID          uint
	Name        string
	UserID      uint
	Description string
	User        User
	Task        *Task
}

// interface untuk Data Layer
type ProjectDataInterface interface {
	Insert(input Core) error
	SelectAll() ([]Core, error)
	SelectByProjecttID(id int) ([]Core, error)
	Update(id int, input Core) error
	Delete(id int) error
}

// interface untuk Service Layer
type ProjectServiceInterface interface {
	Create(input Core) error
	SelectAllAll() ([]Core, error)
	SelectByProjecttID(id int) ([]Core, error)
	Update(id int, input Core) error
	Delete(id int) error
}
