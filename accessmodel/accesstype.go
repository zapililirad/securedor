package accessmodel

type AccessType string

func (a AccessType) GetAccessType() string {
	return string(a)
}

func NewAccessType(name string) AccessType {
	return AccessType(name)
}
