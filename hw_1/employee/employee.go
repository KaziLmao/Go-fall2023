package employee

type Employee interface {
	SetSalary(salary int)
	SetPosition(position string)
	SetAddress(address string)
	GetSalary() int
	GetPosition() string
	GetAddress() string
}
