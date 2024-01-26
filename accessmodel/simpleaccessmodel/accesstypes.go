package simpleaccessmodel

import "github.com/zapililirad/securedor/accessmodel"

type SimpleAccessType accessmodel.AccessType

const (
	Read      = SimpleAccessType("read")
	Write     = SimpleAccessType("write")
	ReadWrite = SimpleAccessType("readwrite")
)
