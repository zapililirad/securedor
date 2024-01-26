package rules

import (
	"context"
)

type RuleRepository interface {
	GetAllRules(ctx context.Context) ([]*Rule, error)
	Add(ctx context.Context, rule *Rule) error
	Update(ctx context.Context, rule *Rule) error
	Delete(ctx context.Context, rule *Rule) error
}

// type RuleDTO struct {
// 	ID       string
// 	Actor    string
// 	Object   string
// 	Accesses []string
// }

// func RuleToDTO(rule *Rule) *RuleDTO {
// 	acesses := []string{}
// 	for _, a := range rule.accesses {
// 		acesses = append(acesses, a.GetAccessType())
// 	}

// 	return &RuleDTO{
// 		Actor:    rule.actor.GetSecurityID(),
// 		Object:   rule.object.GetSecurityID(),
// 		Accesses: acesses,
// 	}
// }

// func DTOToRule(dto *RuleDTO) *Rule {
// 	accesses := []accessmodel.AccessType{}

// 	for _, a := range dto.Accesses {
// 		accesses = append(accesses, accessmodel.NewAccessType(a))
// 	}

// 	return &Rule{
// 		id:       dto.ID,
// 		actor:    access.NewMetaSecurityPrincipal(dto.Actor),
// 		object:   access.NewMetaSecurityPrincipal(dto.Object),
// 		accesses: accesses,
// 	}
// }
