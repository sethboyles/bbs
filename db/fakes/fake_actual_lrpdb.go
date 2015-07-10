// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/bbs"
	"github.com/cloudfoundry-incubator/bbs/db"
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/pivotal-golang/lager"
)

type FakeActualLRPDB struct {
	ActualLRPGroupsStub        func(filter models.ActualLRPFilter, logger lager.Logger) (*models.ActualLRPGroups, *bbs.Error)
	actualLRPGroupsMutex       sync.RWMutex
	actualLRPGroupsArgsForCall []struct {
		filter models.ActualLRPFilter
		logger lager.Logger
	}
	actualLRPGroupsReturns struct {
		result1 *models.ActualLRPGroups
		result2 *bbs.Error
	}
	ActualLRPGroupsByProcessGuidStub        func(processGuid string, logger lager.Logger) (*models.ActualLRPGroups, *bbs.Error)
	actualLRPGroupsByProcessGuidMutex       sync.RWMutex
	actualLRPGroupsByProcessGuidArgsForCall []struct {
		processGuid string
		logger      lager.Logger
	}
	actualLRPGroupsByProcessGuidReturns struct {
		result1 *models.ActualLRPGroups
		result2 *bbs.Error
	}
	ActualLRPGroupByProcessGuidAndIndexStub        func(processGuid string, index int32, logger lager.Logger) (*models.ActualLRPGroup, *bbs.Error)
	actualLRPGroupByProcessGuidAndIndexMutex       sync.RWMutex
	actualLRPGroupByProcessGuidAndIndexArgsForCall []struct {
		processGuid string
		index       int32
		logger      lager.Logger
	}
	actualLRPGroupByProcessGuidAndIndexReturns struct {
		result1 *models.ActualLRPGroup
		result2 *bbs.Error
	}
}

func (fake *FakeActualLRPDB) ActualLRPGroups(filter models.ActualLRPFilter, logger lager.Logger) (*models.ActualLRPGroups, *bbs.Error) {
	fake.actualLRPGroupsMutex.Lock()
	fake.actualLRPGroupsArgsForCall = append(fake.actualLRPGroupsArgsForCall, struct {
		filter models.ActualLRPFilter
		logger lager.Logger
	}{filter, logger})
	fake.actualLRPGroupsMutex.Unlock()
	if fake.ActualLRPGroupsStub != nil {
		return fake.ActualLRPGroupsStub(filter, logger)
	} else {
		return fake.actualLRPGroupsReturns.result1, fake.actualLRPGroupsReturns.result2
	}
}

func (fake *FakeActualLRPDB) ActualLRPGroupsCallCount() int {
	fake.actualLRPGroupsMutex.RLock()
	defer fake.actualLRPGroupsMutex.RUnlock()
	return len(fake.actualLRPGroupsArgsForCall)
}

func (fake *FakeActualLRPDB) ActualLRPGroupsArgsForCall(i int) (models.ActualLRPFilter, lager.Logger) {
	fake.actualLRPGroupsMutex.RLock()
	defer fake.actualLRPGroupsMutex.RUnlock()
	return fake.actualLRPGroupsArgsForCall[i].filter, fake.actualLRPGroupsArgsForCall[i].logger
}

func (fake *FakeActualLRPDB) ActualLRPGroupsReturns(result1 *models.ActualLRPGroups, result2 *bbs.Error) {
	fake.ActualLRPGroupsStub = nil
	fake.actualLRPGroupsReturns = struct {
		result1 *models.ActualLRPGroups
		result2 *bbs.Error
	}{result1, result2}
}

func (fake *FakeActualLRPDB) ActualLRPGroupsByProcessGuid(processGuid string, logger lager.Logger) (*models.ActualLRPGroups, *bbs.Error) {
	fake.actualLRPGroupsByProcessGuidMutex.Lock()
	fake.actualLRPGroupsByProcessGuidArgsForCall = append(fake.actualLRPGroupsByProcessGuidArgsForCall, struct {
		processGuid string
		logger      lager.Logger
	}{processGuid, logger})
	fake.actualLRPGroupsByProcessGuidMutex.Unlock()
	if fake.ActualLRPGroupsByProcessGuidStub != nil {
		return fake.ActualLRPGroupsByProcessGuidStub(processGuid, logger)
	} else {
		return fake.actualLRPGroupsByProcessGuidReturns.result1, fake.actualLRPGroupsByProcessGuidReturns.result2
	}
}

func (fake *FakeActualLRPDB) ActualLRPGroupsByProcessGuidCallCount() int {
	fake.actualLRPGroupsByProcessGuidMutex.RLock()
	defer fake.actualLRPGroupsByProcessGuidMutex.RUnlock()
	return len(fake.actualLRPGroupsByProcessGuidArgsForCall)
}

func (fake *FakeActualLRPDB) ActualLRPGroupsByProcessGuidArgsForCall(i int) (string, lager.Logger) {
	fake.actualLRPGroupsByProcessGuidMutex.RLock()
	defer fake.actualLRPGroupsByProcessGuidMutex.RUnlock()
	return fake.actualLRPGroupsByProcessGuidArgsForCall[i].processGuid, fake.actualLRPGroupsByProcessGuidArgsForCall[i].logger
}

func (fake *FakeActualLRPDB) ActualLRPGroupsByProcessGuidReturns(result1 *models.ActualLRPGroups, result2 *bbs.Error) {
	fake.ActualLRPGroupsByProcessGuidStub = nil
	fake.actualLRPGroupsByProcessGuidReturns = struct {
		result1 *models.ActualLRPGroups
		result2 *bbs.Error
	}{result1, result2}
}

func (fake *FakeActualLRPDB) ActualLRPGroupByProcessGuidAndIndex(processGuid string, index int32, logger lager.Logger) (*models.ActualLRPGroup, *bbs.Error) {
	fake.actualLRPGroupByProcessGuidAndIndexMutex.Lock()
	fake.actualLRPGroupByProcessGuidAndIndexArgsForCall = append(fake.actualLRPGroupByProcessGuidAndIndexArgsForCall, struct {
		processGuid string
		index       int32
		logger      lager.Logger
	}{processGuid, index, logger})
	fake.actualLRPGroupByProcessGuidAndIndexMutex.Unlock()
	if fake.ActualLRPGroupByProcessGuidAndIndexStub != nil {
		return fake.ActualLRPGroupByProcessGuidAndIndexStub(processGuid, index, logger)
	} else {
		return fake.actualLRPGroupByProcessGuidAndIndexReturns.result1, fake.actualLRPGroupByProcessGuidAndIndexReturns.result2
	}
}

func (fake *FakeActualLRPDB) ActualLRPGroupByProcessGuidAndIndexCallCount() int {
	fake.actualLRPGroupByProcessGuidAndIndexMutex.RLock()
	defer fake.actualLRPGroupByProcessGuidAndIndexMutex.RUnlock()
	return len(fake.actualLRPGroupByProcessGuidAndIndexArgsForCall)
}

func (fake *FakeActualLRPDB) ActualLRPGroupByProcessGuidAndIndexArgsForCall(i int) (string, int32, lager.Logger) {
	fake.actualLRPGroupByProcessGuidAndIndexMutex.RLock()
	defer fake.actualLRPGroupByProcessGuidAndIndexMutex.RUnlock()
	return fake.actualLRPGroupByProcessGuidAndIndexArgsForCall[i].processGuid, fake.actualLRPGroupByProcessGuidAndIndexArgsForCall[i].index, fake.actualLRPGroupByProcessGuidAndIndexArgsForCall[i].logger
}

func (fake *FakeActualLRPDB) ActualLRPGroupByProcessGuidAndIndexReturns(result1 *models.ActualLRPGroup, result2 *bbs.Error) {
	fake.ActualLRPGroupByProcessGuidAndIndexStub = nil
	fake.actualLRPGroupByProcessGuidAndIndexReturns = struct {
		result1 *models.ActualLRPGroup
		result2 *bbs.Error
	}{result1, result2}
}

var _ db.ActualLRPDB = new(FakeActualLRPDB)