package task

// struct task model
type Task struct {
	ID        uint
	Task      string
	ProjectID uint
	Status    string
}

// interface untuk Data Layer
type TaskDataInterface interface {
	Insert(input Task) error
	Update(id int, input Task) error
	Delete(id int) error
}

// interface untuk Service Layer
type TaskServiceInterface interface {
	Create(input Task) error
	Update(id int, input Task) error
	Delete(id int) error
}
