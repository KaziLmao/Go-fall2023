package cleaner

type Cleaner struct {
	salary   int
	address  string
	position string
}

func (c *Cleaner) SetSalary(salary int) {
	c.salary = salary
}

func (c *Cleaner) GetSalary() int {
	return c.salary
}

func (c *Cleaner) SetPosition(position string) {
	c.position = position
}

func (c *Cleaner) GetPosition() string {
	return c.position
}

func (c *Cleaner) SetAddress(address string) {
	c.address = address
}

func (c *Cleaner) GetAddress() string {
	return c.address
}
