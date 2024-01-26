package simpleaccessmodel

import "github.com/zapililirad/securedor/accessmodel"

type SimpleAccessModel struct {
	required SimpleAccessType
}

func (m *SimpleAccessModel) IsAccessValid(current accessmodel.AccessType) bool {
	if current == accessmodel.AccessType(m.required) {
		return true
	} else if current == accessmodel.AccessType(ReadWrite) {
		return true
	} else {
		return false
	}
}

func NewSimpleAccessModel(require SimpleAccessType) *SimpleAccessModel {
	return &SimpleAccessModel{require}
}
