// Code generated by counterfeiter. DO NOT EDIT.
package metricsfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/bbs/metrics"
)

type FakeDBStats struct {
	OpenConnectionsStub        func() int
	openConnectionsMutex       sync.RWMutex
	openConnectionsArgsForCall []struct {
	}
	openConnectionsReturns struct {
		result1 int
	}
	openConnectionsReturnsOnCall map[int]struct {
		result1 int
	}
	WaitCountStub        func() int64
	waitCountMutex       sync.RWMutex
	waitCountArgsForCall []struct {
	}
	waitCountReturns struct {
		result1 int64
	}
	waitCountReturnsOnCall map[int]struct {
		result1 int64
	}
	WaitDurationStub        func() time.Duration
	waitDurationMutex       sync.RWMutex
	waitDurationArgsForCall []struct {
	}
	waitDurationReturns struct {
		result1 time.Duration
	}
	waitDurationReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDBStats) OpenConnections() int {
	fake.openConnectionsMutex.Lock()
	ret, specificReturn := fake.openConnectionsReturnsOnCall[len(fake.openConnectionsArgsForCall)]
	fake.openConnectionsArgsForCall = append(fake.openConnectionsArgsForCall, struct {
	}{})
	fake.recordInvocation("OpenConnections", []interface{}{})
	openConnectionsStubCopy := fake.OpenConnectionsStub
	fake.openConnectionsMutex.Unlock()
	if openConnectionsStubCopy != nil {
		return openConnectionsStubCopy()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.openConnectionsReturns
	return fakeReturns.result1
}

func (fake *FakeDBStats) OpenConnectionsCallCount() int {
	fake.openConnectionsMutex.RLock()
	defer fake.openConnectionsMutex.RUnlock()
	return len(fake.openConnectionsArgsForCall)
}

func (fake *FakeDBStats) OpenConnectionsCalls(stub func() int) {
	fake.openConnectionsMutex.Lock()
	defer fake.openConnectionsMutex.Unlock()
	fake.OpenConnectionsStub = stub
}

func (fake *FakeDBStats) OpenConnectionsReturns(result1 int) {
	fake.openConnectionsMutex.Lock()
	defer fake.openConnectionsMutex.Unlock()
	fake.OpenConnectionsStub = nil
	fake.openConnectionsReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeDBStats) OpenConnectionsReturnsOnCall(i int, result1 int) {
	fake.openConnectionsMutex.Lock()
	defer fake.openConnectionsMutex.Unlock()
	fake.OpenConnectionsStub = nil
	if fake.openConnectionsReturnsOnCall == nil {
		fake.openConnectionsReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.openConnectionsReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeDBStats) WaitCount() int64 {
	fake.waitCountMutex.Lock()
	ret, specificReturn := fake.waitCountReturnsOnCall[len(fake.waitCountArgsForCall)]
	fake.waitCountArgsForCall = append(fake.waitCountArgsForCall, struct {
	}{})
	fake.recordInvocation("WaitCount", []interface{}{})
	waitCountStubCopy := fake.WaitCountStub
	fake.waitCountMutex.Unlock()
	if waitCountStubCopy != nil {
		return waitCountStubCopy()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.waitCountReturns
	return fakeReturns.result1
}

func (fake *FakeDBStats) WaitCountCallCount() int {
	fake.waitCountMutex.RLock()
	defer fake.waitCountMutex.RUnlock()
	return len(fake.waitCountArgsForCall)
}

func (fake *FakeDBStats) WaitCountCalls(stub func() int64) {
	fake.waitCountMutex.Lock()
	defer fake.waitCountMutex.Unlock()
	fake.WaitCountStub = stub
}

func (fake *FakeDBStats) WaitCountReturns(result1 int64) {
	fake.waitCountMutex.Lock()
	defer fake.waitCountMutex.Unlock()
	fake.WaitCountStub = nil
	fake.waitCountReturns = struct {
		result1 int64
	}{result1}
}

func (fake *FakeDBStats) WaitCountReturnsOnCall(i int, result1 int64) {
	fake.waitCountMutex.Lock()
	defer fake.waitCountMutex.Unlock()
	fake.WaitCountStub = nil
	if fake.waitCountReturnsOnCall == nil {
		fake.waitCountReturnsOnCall = make(map[int]struct {
			result1 int64
		})
	}
	fake.waitCountReturnsOnCall[i] = struct {
		result1 int64
	}{result1}
}

func (fake *FakeDBStats) WaitDuration() time.Duration {
	fake.waitDurationMutex.Lock()
	ret, specificReturn := fake.waitDurationReturnsOnCall[len(fake.waitDurationArgsForCall)]
	fake.waitDurationArgsForCall = append(fake.waitDurationArgsForCall, struct {
	}{})
	fake.recordInvocation("WaitDuration", []interface{}{})
	waitDurationStubCopy := fake.WaitDurationStub
	fake.waitDurationMutex.Unlock()
	if waitDurationStubCopy != nil {
		return waitDurationStubCopy()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.waitDurationReturns
	return fakeReturns.result1
}

func (fake *FakeDBStats) WaitDurationCallCount() int {
	fake.waitDurationMutex.RLock()
	defer fake.waitDurationMutex.RUnlock()
	return len(fake.waitDurationArgsForCall)
}

func (fake *FakeDBStats) WaitDurationCalls(stub func() time.Duration) {
	fake.waitDurationMutex.Lock()
	defer fake.waitDurationMutex.Unlock()
	fake.WaitDurationStub = stub
}

func (fake *FakeDBStats) WaitDurationReturns(result1 time.Duration) {
	fake.waitDurationMutex.Lock()
	defer fake.waitDurationMutex.Unlock()
	fake.WaitDurationStub = nil
	fake.waitDurationReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeDBStats) WaitDurationReturnsOnCall(i int, result1 time.Duration) {
	fake.waitDurationMutex.Lock()
	defer fake.waitDurationMutex.Unlock()
	fake.WaitDurationStub = nil
	if fake.waitDurationReturnsOnCall == nil {
		fake.waitDurationReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.waitDurationReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeDBStats) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.openConnectionsMutex.RLock()
	defer fake.openConnectionsMutex.RUnlock()
	fake.waitCountMutex.RLock()
	defer fake.waitCountMutex.RUnlock()
	fake.waitDurationMutex.RLock()
	defer fake.waitDurationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDBStats) recordInvocation(key string, args []interface{}) {
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

var _ metrics.DBStats = new(FakeDBStats)
