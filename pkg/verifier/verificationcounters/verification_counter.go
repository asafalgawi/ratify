package verificationcounters

import (
	"context"
	"sync"

	ctxutils "github.com/ratify-project/ratify/internal/context"
)

type VerifierCounter struct {
	mu                        sync.Mutex
	subjectVerificationCounts map[string]int
}

func GetFromContext(ctx context.Context, verifierName string) *VerifierCounter {
	if ctx == nil {
		return nil
	}

	counterMap := ctxutils.GetVerificationCountersFromContext(ctx)
	if counterMap == nil {
		return nil
	}

	counter, ok := counterMap[verifierName]
	if !ok {
		return nil
	}

	return counter.(*VerifierCounter)
}

func (vc *VerifierCounter) Increment(subject string) {
	vc.mu.Lock()
	defer vc.mu.Unlock()
	vc.subjectVerificationCounts[subject]++
}

func (vc *VerifierCounter) Get(subject string) int {
	vc.mu.Lock()
	defer vc.mu.Unlock()
	return vc.subjectVerificationCounts[subject]
}
