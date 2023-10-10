package my_example

type EmployeeInfo struct {
	Name    string
	Surname string
}

type EmployeeAddress struct {
	Address string
}

type EmployeePostion struct {
	JobTitle string
}

func (e EmployeeInfo) GetEmployeeInfo() (string, string) {
	return e.Name, e.Surname
}

func (e EmployeeAddress) GetEmployeeAddress() string {
	return e.Address
}

func (e EmployeePostion) GetEmployeePosition() string {
	return e.JobTitle
}
