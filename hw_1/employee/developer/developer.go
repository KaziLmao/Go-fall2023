package developer

type Developer struct {
	salary   int
	address  string
	position string
}

func (dev *Developer) SetSalary(salary int) {
	dev.salary = salary
}

func (dev *Developer) GetSalary() int {
	return dev.salary
}

func (dev *Developer) SetPosition(position string) {
	dev.position = position
}

func (dev *Developer) GetPosition() string {
	return dev.position
}

func (dev *Developer) SetAddress(address string) {
	dev.address = address
}

func (dev *Developer) GetAddress() string {
	return dev.address
}
