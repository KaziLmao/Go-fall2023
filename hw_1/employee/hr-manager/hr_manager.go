package hr_manager

type HrManager struct {
	salary   int
	address  string
	position string
}

func (hr *HrManager) SetSalary(salary int) {
	hr.salary = salary
}

func (hr *HrManager) GetSalary() int {
	return hr.salary
}

func (hr *HrManager) SetPosition(position string) {
	hr.position = position
}

func (hr *HrManager) GetPosition() string {
	return hr.position
}

func (hr *HrManager) SetAddress(address string) {
	hr.address = address
}

func (hr *HrManager) GetAddress() string {
	return hr.address
}
