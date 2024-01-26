package accessmodel

type AccessModel interface {
	IsAccessValid(current AccessType) bool
}
