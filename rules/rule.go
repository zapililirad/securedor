package rules

import (
	"github.com/zapililirad/securedor"
)

type Rule struct {
	ID       string
	ActorID  string
	ObjectID string
	Accesses []string
}

// func (r *Rule) GetAccesses() []accessmodel.AccessType {
// 	return r.Accesses
// }

func FilterRulesByActor(rules []*Rule, actor securedor.SecurityPrincipal) []*Rule {
	output := []*Rule{}
	for _, rule := range rules {
		if actor.GetSecurityID() == rule.ActorID {
			output = append(output, rule)
		}
	}

	return output
}

func FilterRulesByObject(rules []*Rule, object securedor.SecurityPrincipal) []*Rule {
	output := []*Rule{}
	for _, rule := range rules {
		if object.GetSecurityID() == rule.ObjectID {
			output = append(output, rule)
		}
	}

	return output
}
