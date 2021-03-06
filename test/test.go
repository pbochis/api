package test

import (
	"strconv"

	"github.com/coduno/api/model"

	"golang.org/x/net/context"
)

type Tester int

type TesterFunc func(ctx context.Context, params map[string]string, sub model.KeyedSubmission) error

const (
	Simple Tester = 1 + iota
	Junit
	maxTester
)

var testers = make([]TesterFunc, maxTester)

// RegisterTester registers a function to be called lated via Call. It is
// usually called from the init function of a package that contains a Tester.
func RegisterTester(t Tester, f TesterFunc) {
	if t >= maxTester {
		panic("test: RegisterTester of unknown tester function")
	}
	testers[t] = f
}

// Call looks up a registered Resulter and calls it.
func (t Tester) Call(ctx context.Context, params map[string]string, sub model.KeyedSubmission) error {
	if t > 0 && t < maxTester {
		f := testers[t]
		if f != nil {
			return f(ctx, params, sub)
		}
	}
	panic("test: requested tester function #" + strconv.Itoa(int(t)) + " is unavailable")
}
