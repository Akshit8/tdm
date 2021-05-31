// Code generated by counterfeiter. DO NOT EDIT.
package servicetesting

import (
	"context"
	"sync"

	"github.com/Akshit8/tdm/internal"
	"github.com/Akshit8/tdm/internal/service"
)

type FakeTaskService struct {
	CreateStub        func(context.Context, string, internal.Priority, internal.Dates) (internal.Task, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 internal.Priority
		arg4 internal.Dates
	}
	createReturns struct {
		result1 internal.Task
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 internal.Task
		result2 error
	}
	TaskStub        func(context.Context, string) (internal.Task, error)
	taskMutex       sync.RWMutex
	taskArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	taskReturns struct {
		result1 internal.Task
		result2 error
	}
	taskReturnsOnCall map[int]struct {
		result1 internal.Task
		result2 error
	}
	UpdateStub        func(context.Context, string, string, internal.Priority, internal.Dates, bool) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 internal.Priority
		arg5 internal.Dates
		arg6 bool
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTaskService) Create(arg1 context.Context, arg2 string, arg3 internal.Priority, arg4 internal.Dates) (internal.Task, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 internal.Priority
		arg4 internal.Dates
	}{arg1, arg2, arg3, arg4})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3, arg4})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTaskService) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeTaskService) CreateCalls(stub func(context.Context, string, internal.Priority, internal.Dates) (internal.Task, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeTaskService) CreateArgsForCall(i int) (context.Context, string, internal.Priority, internal.Dates) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeTaskService) CreateReturns(result1 internal.Task, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 internal.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskService) CreateReturnsOnCall(i int, result1 internal.Task, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 internal.Task
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 internal.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskService) Task(arg1 context.Context, arg2 string) (internal.Task, error) {
	fake.taskMutex.Lock()
	ret, specificReturn := fake.taskReturnsOnCall[len(fake.taskArgsForCall)]
	fake.taskArgsForCall = append(fake.taskArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.TaskStub
	fakeReturns := fake.taskReturns
	fake.recordInvocation("Task", []interface{}{arg1, arg2})
	fake.taskMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTaskService) TaskCallCount() int {
	fake.taskMutex.RLock()
	defer fake.taskMutex.RUnlock()
	return len(fake.taskArgsForCall)
}

func (fake *FakeTaskService) TaskCalls(stub func(context.Context, string) (internal.Task, error)) {
	fake.taskMutex.Lock()
	defer fake.taskMutex.Unlock()
	fake.TaskStub = stub
}

func (fake *FakeTaskService) TaskArgsForCall(i int) (context.Context, string) {
	fake.taskMutex.RLock()
	defer fake.taskMutex.RUnlock()
	argsForCall := fake.taskArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTaskService) TaskReturns(result1 internal.Task, result2 error) {
	fake.taskMutex.Lock()
	defer fake.taskMutex.Unlock()
	fake.TaskStub = nil
	fake.taskReturns = struct {
		result1 internal.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskService) TaskReturnsOnCall(i int, result1 internal.Task, result2 error) {
	fake.taskMutex.Lock()
	defer fake.taskMutex.Unlock()
	fake.TaskStub = nil
	if fake.taskReturnsOnCall == nil {
		fake.taskReturnsOnCall = make(map[int]struct {
			result1 internal.Task
			result2 error
		})
	}
	fake.taskReturnsOnCall[i] = struct {
		result1 internal.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskService) Update(arg1 context.Context, arg2 string, arg3 string, arg4 internal.Priority, arg5 internal.Dates, arg6 bool) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 internal.Priority
		arg5 internal.Dates
		arg6 bool
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	stub := fake.UpdateStub
	fakeReturns := fake.updateReturns
	fake.recordInvocation("Update", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.updateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTaskService) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeTaskService) UpdateCalls(stub func(context.Context, string, string, internal.Priority, internal.Dates, bool) error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeTaskService) UpdateArgsForCall(i int) (context.Context, string, string, internal.Priority, internal.Dates, bool) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeTaskService) UpdateReturns(result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskService) UpdateReturnsOnCall(i int, result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTaskService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.taskMutex.RLock()
	defer fake.taskMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTaskService) recordInvocation(key string, args []interface{}) {
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

var _ service.TaskService = new(FakeTaskService)
