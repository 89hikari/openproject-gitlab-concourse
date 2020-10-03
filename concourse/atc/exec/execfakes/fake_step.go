// Code generated by counterfeiter. DO NOT EDIT.
package execfakes

import (
	"context"
	"sync"

	"github.com/concourse/concourse/atc/exec"
)

type FakeStep struct {
	RunStub        func(context.Context, exec.RunState) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 context.Context
		arg2 exec.RunState
	}
	runReturns struct {
		result1 error
	}
	runReturnsOnCall map[int]struct {
		result1 error
	}
	SucceededStub        func() bool
	succeededMutex       sync.RWMutex
	succeededArgsForCall []struct {
	}
	succeededReturns struct {
		result1 bool
	}
	succeededReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStep) Run(arg1 context.Context, arg2 exec.RunState) error {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 context.Context
		arg2 exec.RunState
	}{arg1, arg2})
	fake.recordInvocation("Run", []interface{}{arg1, arg2})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.runReturns
	return fakeReturns.result1
}

func (fake *FakeStep) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeStep) RunCalls(stub func(context.Context, exec.RunState) error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = stub
}

func (fake *FakeStep) RunArgsForCall(i int) (context.Context, exec.RunState) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	argsForCall := fake.runArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStep) RunReturns(result1 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStep) RunReturnsOnCall(i int, result1 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStep) Succeeded() bool {
	fake.succeededMutex.Lock()
	ret, specificReturn := fake.succeededReturnsOnCall[len(fake.succeededArgsForCall)]
	fake.succeededArgsForCall = append(fake.succeededArgsForCall, struct {
	}{})
	fake.recordInvocation("Succeeded", []interface{}{})
	fake.succeededMutex.Unlock()
	if fake.SucceededStub != nil {
		return fake.SucceededStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.succeededReturns
	return fakeReturns.result1
}

func (fake *FakeStep) SucceededCallCount() int {
	fake.succeededMutex.RLock()
	defer fake.succeededMutex.RUnlock()
	return len(fake.succeededArgsForCall)
}

func (fake *FakeStep) SucceededCalls(stub func() bool) {
	fake.succeededMutex.Lock()
	defer fake.succeededMutex.Unlock()
	fake.SucceededStub = stub
}

func (fake *FakeStep) SucceededReturns(result1 bool) {
	fake.succeededMutex.Lock()
	defer fake.succeededMutex.Unlock()
	fake.SucceededStub = nil
	fake.succeededReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeStep) SucceededReturnsOnCall(i int, result1 bool) {
	fake.succeededMutex.Lock()
	defer fake.succeededMutex.Unlock()
	fake.SucceededStub = nil
	if fake.succeededReturnsOnCall == nil {
		fake.succeededReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.succeededReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeStep) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.succeededMutex.RLock()
	defer fake.succeededMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStep) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ exec.Step = new(FakeStep)
