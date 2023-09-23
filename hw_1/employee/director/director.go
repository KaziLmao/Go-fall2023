package director

type Director struct {
	salary   int
	address  string
	position string
}

func (dr *Director) SetSalary(salary int) {
	dr.salary = salary
}

func (dr *Director) GetSalary() int {
	return dr.salary
}

func (dr *Director) SetPosition(position string) {
	dr.position = position
}

func (dr *Director) GetPosition() string {
	return dr.position
}

func (dr *Director) SetAddress(address string) {
	dr.address = address
}

func (dr *Director) GetAddress() string {
	return dr.address
}
