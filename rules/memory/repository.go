package memory

import (
	"context"
	"math/rand"
	"time"

	"github.com/zapililirad/securedor/rules"
)

const letters = "abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type MemoryRule struct {
	actor    string
	object   string
	accesses []string
}

type MemoryRuleRepository struct {
	storage map[string]*MemoryRule
}

func NewMemoryRuleRepository() *MemoryRuleRepository {
	return &MemoryRuleRepository{
		storage: make(map[string]*MemoryRule, 0),
	}
}

func (r *MemoryRuleRepository) Add(ctx context.Context, rule *rules.Rule) error {
	rule.ID = randomString(10)
	r.storage[rule.ID] = &MemoryRule{
		actor:    rule.ActorID,
		object:   rule.ObjectID,
		accesses: rule.Accesses,
	}

	return nil
}

func (r *MemoryRuleRepository) Update(ctx context.Context, rule *rules.Rule) error {
	r.storage[rule.ID] = &MemoryRule{
		actor:    rule.ActorID,
		object:   rule.ObjectID,
		accesses: rule.Accesses,
	}

	return nil
}

func (r *MemoryRuleRepository) Delete(ctx context.Context, rule *rules.Rule) error {
	_, ok := r.storage[rule.ID]
	if !ok {
		return rules.ErrRuleNotFound
	}

	delete(r.storage, rule.ID)

	return nil
}

func (r *MemoryRuleRepository) GetAllRules(ctx context.Context) ([]*rules.Rule, error) {
	rulesOutput := []*rules.Rule{}

	for k := range r.storage {
		rulesOutput = append(rulesOutput, &rules.Rule{
			ID:       k,
			ActorID:  r.storage[k].actor,
			ObjectID: r.storage[k].object,
			Accesses: r.storage[k].accesses,
		})
	}

	return rulesOutput, nil
}
