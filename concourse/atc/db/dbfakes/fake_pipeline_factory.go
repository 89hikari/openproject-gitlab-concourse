// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/db"
)

type FakePipelineFactory struct {
	AllPipelinesStub        func() ([]db.Pipeline, error)
	allPipelinesMutex       sync.RWMutex
	allPipelinesArgsForCall []struct {
	}
	allPipelinesReturns struct {
		result1 []db.Pipeline
		result2 error
	}
	allPipelinesReturnsOnCall map[int]struct {
		result1 []db.Pipeline
		result2 error
	}
	PipelinesToScheduleStub        func() ([]db.Pipeline, error)
	pipelinesToScheduleMutex       sync.RWMutex
	pipelinesToScheduleArgsForCall []struct {
	}
	pipelinesToScheduleReturns struct {
		result1 []db.Pipeline
		result2 error
	}
	pipelinesToScheduleReturnsOnCall map[int]struct {
		result1 []db.Pipeline
		result2 error
	}
	VisiblePipelinesStub        func([]string) ([]db.Pipeline, error)
	visiblePipelinesMutex       sync.RWMutex
	visiblePipelinesArgsForCall []struct {
		arg1 []string
	}
	visiblePipelinesReturns struct {
		result1 []db.Pipeline
		result2 error
	}
	visiblePipelinesReturnsOnCall map[int]struct {
		result1 []db.Pipeline
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePipelineFactory) AllPipelines() ([]db.Pipeline, error) {
	fake.allPipelinesMutex.Lock()
	ret, specificReturn := fake.allPipelinesReturnsOnCall[len(fake.allPipelinesArgsForCall)]
	fake.allPipelinesArgsForCall = append(fake.allPipelinesArgsForCall, struct {
	}{})
	fake.recordInvocation("AllPipelines", []interface{}{})
	fake.allPipelinesMutex.Unlock()
	if fake.AllPipelinesStub != nil {
		return fake.AllPipelinesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.allPipelinesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePipelineFactory) AllPipelinesCallCount() int {
	fake.allPipelinesMutex.RLock()
	defer fake.allPipelinesMutex.RUnlock()
	return len(fake.allPipelinesArgsForCall)
}

func (fake *FakePipelineFactory) AllPipelinesCalls(stub func() ([]db.Pipeline, error)) {
	fake.allPipelinesMutex.Lock()
	defer fake.allPipelinesMutex.Unlock()
	fake.AllPipelinesStub = stub
}

func (fake *FakePipelineFactory) AllPipelinesReturns(result1 []db.Pipeline, result2 error) {
	fake.allPipelinesMutex.Lock()
	defer fake.allPipelinesMutex.Unlock()
	fake.AllPipelinesStub = nil
	fake.allPipelinesReturns = struct {
		result1 []db.Pipeline
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineFactory) AllPipelinesReturnsOnCall(i int, result1 []db.Pipeline, result2 error) {
	fake.allPipelinesMutex.Lock()
	defer fake.allPipelinesMutex.Unlock()
	fake.AllPipelinesStub = nil
	if fake.allPipelinesReturnsOnCall == nil {
		fake.allPipelinesReturnsOnCall = make(map[int]struct {
			result1 []db.Pipeline
			result2 error
		})
	}
	fake.allPipelinesReturnsOnCall[i] = struct {
		result1 []db.Pipeline
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineFactory) PipelinesToSchedule() ([]db.Pipeline, error) {
	fake.pipelinesToScheduleMutex.Lock()
	ret, specificReturn := fake.pipelinesToScheduleReturnsOnCall[len(fake.pipelinesToScheduleArgsForCall)]
	fake.pipelinesToScheduleArgsForCall = append(fake.pipelinesToScheduleArgsForCall, struct {
	}{})
	fake.recordInvocation("PipelinesToSchedule", []interface{}{})
	fake.pipelinesToScheduleMutex.Unlock()
	if fake.PipelinesToScheduleStub != nil {
		return fake.PipelinesToScheduleStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.pipelinesToScheduleReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePipelineFactory) PipelinesToScheduleCallCount() int {
	fake.pipelinesToScheduleMutex.RLock()
	defer fake.pipelinesToScheduleMutex.RUnlock()
	return len(fake.pipelinesToScheduleArgsForCall)
}

func (fake *FakePipelineFactory) PipelinesToScheduleCalls(stub func() ([]db.Pipeline, error)) {
	fake.pipelinesToScheduleMutex.Lock()
	defer fake.pipelinesToScheduleMutex.Unlock()
	fake.PipelinesToScheduleStub = stub
}

func (fake *FakePipelineFactory) PipelinesToScheduleReturns(result1 []db.Pipeline, result2 error) {
	fake.pipelinesToScheduleMutex.Lock()
	defer fake.pipelinesToScheduleMutex.Unlock()
	fake.PipelinesToScheduleStub = nil
	fake.pipelinesToScheduleReturns = struct {
		result1 []db.Pipeline
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineFactory) PipelinesToScheduleReturnsOnCall(i int, result1 []db.Pipeline, result2 error) {
	fake.pipelinesToScheduleMutex.Lock()
	defer fake.pipelinesToScheduleMutex.Unlock()
	fake.PipelinesToScheduleStub = nil
	if fake.pipelinesToScheduleReturnsOnCall == nil {
		fake.pipelinesToScheduleReturnsOnCall = make(map[int]struct {
			result1 []db.Pipeline
			result2 error
		})
	}
	fake.pipelinesToScheduleReturnsOnCall[i] = struct {
		result1 []db.Pipeline
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineFactory) VisiblePipelines(arg1 []string) ([]db.Pipeline, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.visiblePipelinesMutex.Lock()
	ret, specificReturn := fake.visiblePipelinesReturnsOnCall[len(fake.visiblePipelinesArgsForCall)]
	fake.visiblePipelinesArgsForCall = append(fake.visiblePipelinesArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	fake.recordInvocation("VisiblePipelines", []interface{}{arg1Copy})
	fake.visiblePipelinesMutex.Unlock()
	if fake.VisiblePipelinesStub != nil {
		return fake.VisiblePipelinesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.visiblePipelinesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePipelineFactory) VisiblePipelinesCallCount() int {
	fake.visiblePipelinesMutex.RLock()
	defer fake.visiblePipelinesMutex.RUnlock()
	return len(fake.visiblePipelinesArgsForCall)
}

func (fake *FakePipelineFactory) VisiblePipelinesCalls(stub func([]string) ([]db.Pipeline, error)) {
	fake.visiblePipelinesMutex.Lock()
	defer fake.visiblePipelinesMutex.Unlock()
	fake.VisiblePipelinesStub = stub
}

func (fake *FakePipelineFactory) VisiblePipelinesArgsForCall(i int) []string {
	fake.visiblePipelinesMutex.RLock()
	defer fake.visiblePipelinesMutex.RUnlock()
	argsForCall := fake.visiblePipelinesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePipelineFactory) VisiblePipelinesReturns(result1 []db.Pipeline, result2 error) {
	fake.visiblePipelinesMutex.Lock()
	defer fake.visiblePipelinesMutex.Unlock()
	fake.VisiblePipelinesStub = nil
	fake.visiblePipelinesReturns = struct {
		result1 []db.Pipeline
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineFactory) VisiblePipelinesReturnsOnCall(i int, result1 []db.Pipeline, result2 error) {
	fake.visiblePipelinesMutex.Lock()
	defer fake.visiblePipelinesMutex.Unlock()
	fake.VisiblePipelinesStub = nil
	if fake.visiblePipelinesReturnsOnCall == nil {
		fake.visiblePipelinesReturnsOnCall = make(map[int]struct {
			result1 []db.Pipeline
			result2 error
		})
	}
	fake.visiblePipelinesReturnsOnCall[i] = struct {
		result1 []db.Pipeline
		result2 error
	}{result1, result2}
}

func (fake *FakePipelineFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allPipelinesMutex.RLock()
	defer fake.allPipelinesMutex.RUnlock()
	fake.pipelinesToScheduleMutex.RLock()
	defer fake.pipelinesToScheduleMutex.RUnlock()
	fake.visiblePipelinesMutex.RLock()
	defer fake.visiblePipelinesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePipelineFactory) recordInvocation(key string, args []interface{}) {
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

var _ db.PipelineFactory = new(FakePipelineFactory)
