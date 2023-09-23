package manager

type Manager struct {
	salary   int
	address  string
	position string
}

func (m *Manager) SetSalary(salary int) {
	m.salary = salary
}

func (m *Manager) GetSalary() int {
	return m.salary
}

func (m *Manager) SetPosition(position string) {
	m.position = position
}

func (m *Manager) GetPosition() string {
	return m.position
}

func (m *Manager) SetAddress(address string) {
	m.address = address
}

func (m *Manager) GetAddress() string {
	return m.address
}
