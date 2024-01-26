package controller

import (
	"context"

	"github.com/zapililirad/securedor"
	"github.com/zapililirad/securedor/accessmodel"
	"github.com/zapililirad/securedor/rules"
)

type Controller struct {
	ruleService *rules.RuleService
	accessModel accessmodel.AccessModel
}

func NewController(ruleService *rules.RuleService, accessModel accessmodel.AccessModel) *Controller {
	return &Controller{
		ruleService,
		accessModel,
	}
}

func (c *Controller) ValidateAccess(actor, object securedor.SecurityPrincipal) error {
	r, err := c.ruleService.GetAllRules(context.Background())
	if err != nil {
		return securedor.ErrAccessDenied
	}

	r = rules.FilterRulesByObject(r, object)
	r = rules.FilterRulesByActor(r, actor)
	for _, rule := range r {
		for _, access := range rule.Accesses {
			if c.accessModel.IsAccessValid(accessmodel.AccessType(access)) {
				return nil
			}
		}
	}

	return securedor.ErrAccessDenied
}
