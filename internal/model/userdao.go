package model

type DefaultUserModel struct {
}

func (m *DefaultUserModel) FindById(id string) (*User, error) {
	return nil, nil
}

func (m *DefaultUserModel) FindByStudentId(studentId string) (*User, error) {
	return nil, nil
}

func (m *DefaultUserModel) Unable(id string) error {
	return nil
}