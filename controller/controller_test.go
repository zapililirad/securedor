package controller

import (
	"context"
	"testing"

	"github.com/zapililirad/securedor"
	"github.com/zapililirad/securedor/accessmodel"
	"github.com/zapililirad/securedor/accessmodel/simpleaccessmodel"
	"github.com/zapililirad/securedor/rules"
	"github.com/zapililirad/securedor/rules/memory"
)

func TestController_ValidateAccess(t *testing.T) {
	testRuleService := rules.NewRuleService(memory.NewMemoryRuleRepository())

	// simpleRequireWrite := simpleaccessmodel.NewSimpleAccessModel(simpleaccessmodel.Write)
	simpleRequireRead := simpleaccessmodel.NewSimpleAccessModel(simpleaccessmodel.Read)
	// simpleRequireReadWrite := simpleaccessmodel.NewSimpleAccessModel(simpleaccessmodel.ReadWrite)

	testObject := securedor.NewMetaSecurityPrincipal("object")

	testActorRead := securedor.NewMetaSecurityPrincipal("actorRead")
	testActorWrite := securedor.NewMetaSecurityPrincipal("actorWrite")
	testActorReadWrite := securedor.NewMetaSecurityPrincipal("actorReadWrite")
	testActorReadAndWrite := securedor.NewMetaSecurityPrincipal("actorReadAndWrite")

	testRuleRead, _ := rules.NewRule(testActorRead, testObject, []accessmodel.AccessType{accessmodel.AccessType(simpleaccessmodel.Read)})
	testRuleWrite, _ := rules.NewRule(testActorWrite, testObject, []accessmodel.AccessType{accessmodel.AccessType(simpleaccessmodel.Write)})
	testRuleReadWrite, _ := rules.NewRule(testActorReadWrite, testObject, []accessmodel.AccessType{accessmodel.AccessType(simpleaccessmodel.ReadWrite)})
	testRuleReadAndWrite, _ := rules.NewRule(testActorReadAndWrite, testObject, []accessmodel.AccessType{accessmodel.AccessType(simpleaccessmodel.Read), accessmodel.AccessType(simpleaccessmodel.Write)})

	testRuleService.AddRule(context.Background(), testRuleRead)
	testRuleService.AddRule(context.Background(), testRuleWrite)
	testRuleService.AddRule(context.Background(), testRuleReadWrite)
	testRuleService.AddRule(context.Background(), testRuleReadAndWrite)

	type fields struct {
		ruleService *rules.RuleService
		accessModel accessmodel.AccessModel
	}
	type args struct {
		actor  securedor.SecurityPrincipal
		object securedor.SecurityPrincipal
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Actor with read permission tries to access object with required read permissions",
			fields: fields{
				ruleService: testRuleService,
				accessModel: simpleRequireRead,
			},
			args: args{
				actor:  testActorRead,
				object: testObject,
			},
			wantErr: false,
		},
		{
			name: "Actor with write permission tries to access object with required read permissions",
			fields: fields{
				ruleService: testRuleService,
				accessModel: simpleRequireRead,
			},
			args: args{
				actor:  testActorWrite,
				object: testObject,
			},
			wantErr: true,
		},
		{
			name: "Actor with readwrite permission tries to access object with required read permissions",
			fields: fields{
				ruleService: testRuleService,
				accessModel: simpleRequireRead,
			},
			args: args{
				actor:  testActorReadWrite,
				object: testObject,
			},
			wantErr: false,
		},
		{
			name: "Actor with read and write permission tries to access object with required read permissions",
			fields: fields{
				ruleService: testRuleService,
				accessModel: simpleRequireRead,
			},
			args: args{
				actor:  testActorReadAndWrite,
				object: testObject,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				ruleService: tt.fields.ruleService,
				accessModel: tt.fields.accessModel,
			}
			if err := c.ValidateAccess(tt.args.actor, tt.args.object); (err != nil) != tt.wantErr {
				t.Errorf("Controller.ValidateAccess() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
