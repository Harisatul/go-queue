package domain

type Repository interface {
	Save(number int) (int, error)
	GetAllListOfNumber() (list []int, err error)
}

type repository struct {
	listOfQueue *[]int
}

func (r *repository) GetAllListOfNumber() ([]int, error) {
	return *r.listOfQueue, nil
}

func (r *repository) Save(number int) (int, error) {
	*r.listOfQueue = append(*r.listOfQueue, number)
	return number, nil
}

func NewRepository(list *[]int) Repository {
	return &repository{list}
}
