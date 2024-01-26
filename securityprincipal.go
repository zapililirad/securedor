package securedor

type SecurityPrincipal interface {
	GetSecurityID() string
	CanBeAnActor() bool
}

type MetaSecurityPrincipal struct {
	id string
}

func NewMetaSecurityPrincipal(id string) *MetaSecurityPrincipal {
	return &MetaSecurityPrincipal{id}
}

func (p *MetaSecurityPrincipal) GetSecurityID() string {
	return p.id
}

func (p *MetaSecurityPrincipal) CanBeAnActor() bool {
	return true
}
