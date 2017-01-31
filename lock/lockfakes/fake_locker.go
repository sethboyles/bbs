// This file was generated by counterfeiter
package lockfakes

import (
	"sync"

	"code.cloudfoundry.org/bbs/lock"
	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/lager"
)

type FakeLocker struct {
	LockStub        func(logger lager.Logger, lock models.Lock) error
	lockMutex       sync.RWMutex
	lockArgsForCall []struct {
		logger lager.Logger
		lock   models.Lock
	}
	lockReturns struct {
		result1 error
	}
	ReleaseLockStub        func(logger lager.Logger, lock models.Lock) error
	releaseLockMutex       sync.RWMutex
	releaseLockArgsForCall []struct {
		logger lager.Logger
		lock   models.Lock
	}
	releaseLockReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLocker) Lock(logger lager.Logger, lock models.Lock) error {
	fake.lockMutex.Lock()
	fake.lockArgsForCall = append(fake.lockArgsForCall, struct {
		logger lager.Logger
		lock   models.Lock
	}{logger, lock})
	fake.recordInvocation("Lock", []interface{}{logger, lock})
	fake.lockMutex.Unlock()
	if fake.LockStub != nil {
		return fake.LockStub(logger, lock)
	} else {
		return fake.lockReturns.result1
	}
}

func (fake *FakeLocker) LockCallCount() int {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	return len(fake.lockArgsForCall)
}

func (fake *FakeLocker) LockArgsForCall(i int) (lager.Logger, models.Lock) {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	return fake.lockArgsForCall[i].logger, fake.lockArgsForCall[i].lock
}

func (fake *FakeLocker) LockReturns(result1 error) {
	fake.LockStub = nil
	fake.lockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLocker) ReleaseLock(logger lager.Logger, lock models.Lock) error {
	fake.releaseLockMutex.Lock()
	fake.releaseLockArgsForCall = append(fake.releaseLockArgsForCall, struct {
		logger lager.Logger
		lock   models.Lock
	}{logger, lock})
	fake.recordInvocation("ReleaseLock", []interface{}{logger, lock})
	fake.releaseLockMutex.Unlock()
	if fake.ReleaseLockStub != nil {
		return fake.ReleaseLockStub(logger, lock)
	} else {
		return fake.releaseLockReturns.result1
	}
}

func (fake *FakeLocker) ReleaseLockCallCount() int {
	fake.releaseLockMutex.RLock()
	defer fake.releaseLockMutex.RUnlock()
	return len(fake.releaseLockArgsForCall)
}

func (fake *FakeLocker) ReleaseLockArgsForCall(i int) (lager.Logger, models.Lock) {
	fake.releaseLockMutex.RLock()
	defer fake.releaseLockMutex.RUnlock()
	return fake.releaseLockArgsForCall[i].logger, fake.releaseLockArgsForCall[i].lock
}

func (fake *FakeLocker) ReleaseLockReturns(result1 error) {
	fake.ReleaseLockStub = nil
	fake.releaseLockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLocker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	fake.releaseLockMutex.RLock()
	defer fake.releaseLockMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeLocker) recordInvocation(key string, args []interface{}) {
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

var _ lock.Locker = new(FakeLocker)