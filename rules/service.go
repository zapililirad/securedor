package rules

import (
	"context"

	"github.com/zapililirad/securedor"
	"github.com/zapililirad/securedor/accessmodel"
)

type RuleService struct {
	repository RuleRepository
}

func NewRuleService(repository RuleRepository) *RuleService {
	return &RuleService{repository: repository}
}

func NewRule(actor securedor.SecurityPrincipal, object securedor.SecurityPrincipal, accesses []accessmodel.AccessType) (*Rule, error) {
	if len(accesses) < 1 {
		return nil, securedor.ErrIncorrectAccessType
	}

	if !actor.CanBeAnActor() {
		return nil, securedor.ErrIncorrectActor
	}

	as := []string{}
	for _, a := range accesses {
		as = append(as, string(a))
	}

	return &Rule{
		ActorID:  actor.GetSecurityID(),
		ObjectID: object.GetSecurityID(),
		Accesses: as,
	}, nil
}

func (s *RuleService) AddRule(ctx context.Context, rule *Rule) error {
	// dto := RuleToDTO(rule)

	// err := s.repository.Add(ctx, dto)
	// if err != nil {
	// 	return err
	// }

	// rule.id = dto.ID

	return s.repository.Add(ctx, rule)
}

func (s *RuleService) DeleteRule(ctx context.Context, rule *Rule) error {
	// dto := RuleToDTO(rule)

	return s.repository.Delete(ctx, rule)
}

func (s *RuleService) GetAllRules(ctx context.Context) ([]*Rule, error) {

	// dtos, err := s.repository.GetAllRules(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// rules := make([]*Rule, 0)

	// for _, dto := range dtos {
	// 	rules = append(rules, DTOToRule(dto))
	// }

	return s.repository.GetAllRules(ctx)
}

func (s *RuleService) UpdateRule(ctx context.Context, rule *Rule) error {
	return s.repository.Update(ctx, rule)
}
