// Code generated by counterfeiter. DO NOT EDIT.
package dbadminfakes

import (
	"sync"

	"github.com/app-sre/dba-operator/pkg/dbadmin"
)

type FakeDbAdmin struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	ConstraintWillFailStub        func(dbadmin.TableName, dbadmin.ConstraintType, ...string) (bool, error)
	constraintWillFailMutex       sync.RWMutex
	constraintWillFailArgsForCall []struct {
		arg1 dbadmin.TableName
		arg2 dbadmin.ConstraintType
		arg3 []string
	}
	constraintWillFailReturns struct {
		result1 bool
		result2 error
	}
	constraintWillFailReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	GetNextIDsStub        func([]dbadmin.TableName) (map[dbadmin.TableName]uint64, error)
	getNextIDsMutex       sync.RWMutex
	getNextIDsArgsForCall []struct {
		arg1 []dbadmin.TableName
	}
	getNextIDsReturns struct {
		result1 map[dbadmin.TableName]uint64
		result2 error
	}
	getNextIDsReturnsOnCall map[int]struct {
		result1 map[dbadmin.TableName]uint64
		result2 error
	}
	GetSchemaVersionStub        func() (string, error)
	getSchemaVersionMutex       sync.RWMutex
	getSchemaVersionArgsForCall []struct {
	}
	getSchemaVersionReturns struct {
		result1 string
		result2 error
	}
	getSchemaVersionReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetTableSizeEstimatesStub        func([]dbadmin.TableName) (map[dbadmin.TableName]uint64, error)
	getTableSizeEstimatesMutex       sync.RWMutex
	getTableSizeEstimatesArgsForCall []struct {
		arg1 []dbadmin.TableName
	}
	getTableSizeEstimatesReturns struct {
		result1 map[dbadmin.TableName]uint64
		result2 error
	}
	getTableSizeEstimatesReturnsOnCall map[int]struct {
		result1 map[dbadmin.TableName]uint64
		result2 error
	}
	IsBlockingIndexCreationStub        func(dbadmin.TableName, dbadmin.IndexType, ...string) (bool, error)
	isBlockingIndexCreationMutex       sync.RWMutex
	isBlockingIndexCreationArgsForCall []struct {
		arg1 dbadmin.TableName
		arg2 dbadmin.IndexType
		arg3 []string
	}
	isBlockingIndexCreationReturns struct {
		result1 bool
		result2 error
	}
	isBlockingIndexCreationReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	ListUsernamesStub        func(string) ([]string, error)
	listUsernamesMutex       sync.RWMutex
	listUsernamesArgsForCall []struct {
		arg1 string
	}
	listUsernamesReturns struct {
		result1 []string
		result2 error
	}
	listUsernamesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	SelectFloatStub        func(string) (float64, error)
	selectFloatMutex       sync.RWMutex
	selectFloatArgsForCall []struct {
		arg1 string
	}
	selectFloatReturns struct {
		result1 float64
		result2 error
	}
	selectFloatReturnsOnCall map[int]struct {
		result1 float64
		result2 error
	}
	VerifyUnusedAndDeleteCredentialsStub        func(string) error
	verifyUnusedAndDeleteCredentialsMutex       sync.RWMutex
	verifyUnusedAndDeleteCredentialsArgsForCall []struct {
		arg1 string
	}
	verifyUnusedAndDeleteCredentialsReturns struct {
		result1 error
	}
	verifyUnusedAndDeleteCredentialsReturnsOnCall map[int]struct {
		result1 error
	}
	WriteCredentialsStub        func(string, string) error
	writeCredentialsMutex       sync.RWMutex
	writeCredentialsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	writeCredentialsReturns struct {
		result1 error
	}
	writeCredentialsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDbAdmin) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.closeReturns
	return fakeReturns.result1
}

func (fake *FakeDbAdmin) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeDbAdmin) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeDbAdmin) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDbAdmin) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDbAdmin) ConstraintWillFail(arg1 dbadmin.TableName, arg2 dbadmin.ConstraintType, arg3 ...string) (bool, error) {
	fake.constraintWillFailMutex.Lock()
	ret, specificReturn := fake.constraintWillFailReturnsOnCall[len(fake.constraintWillFailArgsForCall)]
	fake.constraintWillFailArgsForCall = append(fake.constraintWillFailArgsForCall, struct {
		arg1 dbadmin.TableName
		arg2 dbadmin.ConstraintType
		arg3 []string
	}{arg1, arg2, arg3})
	fake.recordInvocation("ConstraintWillFail", []interface{}{arg1, arg2, arg3})
	fake.constraintWillFailMutex.Unlock()
	if fake.ConstraintWillFailStub != nil {
		return fake.ConstraintWillFailStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.constraintWillFailReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDbAdmin) ConstraintWillFailCallCount() int {
	fake.constraintWillFailMutex.RLock()
	defer fake.constraintWillFailMutex.RUnlock()
	return len(fake.constraintWillFailArgsForCall)
}

func (fake *FakeDbAdmin) ConstraintWillFailCalls(stub func(dbadmin.TableName, dbadmin.ConstraintType, ...string) (bool, error)) {
	fake.constraintWillFailMutex.Lock()
	defer fake.constraintWillFailMutex.Unlock()
	fake.ConstraintWillFailStub = stub
}

func (fake *FakeDbAdmin) ConstraintWillFailArgsForCall(i int) (dbadmin.TableName, dbadmin.ConstraintType, []string) {
	fake.constraintWillFailMutex.RLock()
	defer fake.constraintWillFailMutex.RUnlock()
	argsForCall := fake.constraintWillFailArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDbAdmin) ConstraintWillFailReturns(result1 bool, result2 error) {
	fake.constraintWillFailMutex.Lock()
	defer fake.constraintWillFailMutex.Unlock()
	fake.ConstraintWillFailStub = nil
	fake.constraintWillFailReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeDbAdmin) ConstraintWillFailReturnsOnCall(i int, result1 bool, result2 error) {
	fake.constraintWillFailMutex.Lock()
	defer fake.constraintWillFailMutex.Unlock()
	fake.ConstraintWillFailStub = nil
	if fake.constraintWillFailReturnsOnCall == nil {
		fake.constraintWillFailReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.constraintWillFailReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeDbAdmin) GetNextIDs(arg1 []dbadmin.TableName) (map[dbadmin.TableName]uint64, error) {
	var arg1Copy []dbadmin.TableName
	if arg1 != nil {
		arg1Copy = make([]dbadmin.TableName, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.getNextIDsMutex.Lock()
	ret, specificReturn := fake.getNextIDsReturnsOnCall[len(fake.getNextIDsArgsForCall)]
	fake.getNextIDsArgsForCall = append(fake.getNextIDsArgsForCall, struct {
		arg1 []dbadmin.TableName
	}{arg1Copy})
	fake.recordInvocation("GetNextIDs", []interface{}{arg1Copy})
	fake.getNextIDsMutex.Unlock()
	if fake.GetNextIDsStub != nil {
		return fake.GetNextIDsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getNextIDsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDbAdmin) GetNextIDsCallCount() int {
	fake.getNextIDsMutex.RLock()
	defer fake.getNextIDsMutex.RUnlock()
	return len(fake.getNextIDsArgsForCall)
}

func (fake *FakeDbAdmin) GetNextIDsCalls(stub func([]dbadmin.TableName) (map[dbadmin.TableName]uint64, error)) {
	fake.getNextIDsMutex.Lock()
	defer fake.getNextIDsMutex.Unlock()
	fake.GetNextIDsStub = stub
}

func (fake *FakeDbAdmin) GetNextIDsArgsForCall(i int) []dbadmin.TableName {
	fake.getNextIDsMutex.RLock()
	defer fake.getNextIDsMutex.RUnlock()
	argsForCall := fake.getNextIDsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDbAdmin) GetNextIDsReturns(result1 map[dbadmin.TableName]uint64, result2 error) {
	fake.getNextIDsMutex.Lock()
	defer fake.getNextIDsMutex.Unlock()
	fake.GetNextIDsStub = nil
	fake.getNextIDsReturns = struct {
		result1 map[dbadmin.TableName]uint64
		result2 error
	}{result1, result2}
}

func (fake *FakeDbAdmin) GetNextIDsReturnsOnCall(i int, result1 map[dbadmin.TableName]uint64, result2 error) {
	fake.getNextIDsMutex.Lock()
	defer fake.getNextIDsMutex.Unlock()
	fake.GetNextIDsStub = nil
	if fake.getNextIDsReturnsOnCall == nil {
		fake.getNextIDsReturnsOnCall = make(map[int]struct {
			result1 map[dbadmin.TableName]uint64
			result2 error
		})
	}
	fake.getNextIDsReturnsOnCall[i] = struct {
		result1 map[dbadmin.TableName]uint64
		result2 error
	}{result1, result2}
}

func (fake *FakeDbAdmin) GetSchemaVersion() (string, error) {
	fake.getSchemaVersionMutex.Lock()
	ret, specificReturn := fake.getSchemaVersionReturnsOnCall[len(fake.getSchemaVersionArgsForCall)]
	fake.getSchemaVersionArgsForCall = append(fake.getSchemaVersionArgsForCall, struct {
	}{})
	fake.recordInvocation("GetSchemaVersion", []interface{}{})
	fake.getSchemaVersionMutex.Unlock()
	if fake.GetSchemaVersionStub != nil {
		return fake.GetSchemaVersionStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getSchemaVersionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDbAdmin) GetSchemaVersionCallCount() int {
	fake.getSchemaVersionMutex.RLock()
	defer fake.getSchemaVersionMutex.RUnlock()
	return len(fake.getSchemaVersionArgsForCall)
}

func (fake *FakeDbAdmin) GetSchemaVersionCalls(stub func() (string, error)) {
	fake.getSchemaVersionMutex.Lock()
	defer fake.getSchemaVersionMutex.Unlock()
	fake.GetSchemaVersionStub = stub
}

func (fake *FakeDbAdmin) GetSchemaVersionReturns(result1 string, result2 error) {
	fake.getSchemaVersionMutex.Lock()
	defer fake.getSchemaVersionMutex.Unlock()
	fake.GetSchemaVersionStub = nil
	fake.getSchemaVersionReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeDbAdmin) GetSchemaVersionReturnsOnCall(i int, result1 string, result2 error) {
	fake.getSchemaVersionMutex.Lock()
	defer fake.getSchemaVersionMutex.Unlock()
	fake.GetSchemaVersionStub = nil
	if fake.getSchemaVersionReturnsOnCall == nil {
		fake.getSchemaVersionReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getSchemaVersionReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeDbAdmin) GetTableSizeEstimates(arg1 []dbadmin.TableName) (map[dbadmin.TableName]uint64, error) {
	var arg1Copy []dbadmin.TableName
	if arg1 != nil {
		arg1Copy = make([]dbadmin.TableName, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.getTableSizeEstimatesMutex.Lock()
	ret, specificReturn := fake.getTableSizeEstimatesReturnsOnCall[len(fake.getTableSizeEstimatesArgsForCall)]
	fake.getTableSizeEstimatesArgsForCall = append(fake.getTableSizeEstimatesArgsForCall, struct {
		arg1 []dbadmin.TableName
	}{arg1Copy})
	fake.recordInvocation("GetTableSizeEstimates", []interface{}{arg1Copy})
	fake.getTableSizeEstimatesMutex.Unlock()
	if fake.GetTableSizeEstimatesStub != nil {
		return fake.GetTableSizeEstimatesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getTableSizeEstimatesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDbAdmin) GetTableSizeEstimatesCallCount() int {
	fake.getTableSizeEstimatesMutex.RLock()
	defer fake.getTableSizeEstimatesMutex.RUnlock()
	return len(fake.getTableSizeEstimatesArgsForCall)
}

func (fake *FakeDbAdmin) GetTableSizeEstimatesCalls(stub func([]dbadmin.TableName) (map[dbadmin.TableName]uint64, error)) {
	fake.getTableSizeEstimatesMutex.Lock()
	defer fake.getTableSizeEstimatesMutex.Unlock()
	fake.GetTableSizeEstimatesStub = stub
}

func (fake *FakeDbAdmin) GetTableSizeEstimatesArgsForCall(i int) []dbadmin.TableName {
	fake.getTableSizeEstimatesMutex.RLock()
	defer fake.getTableSizeEstimatesMutex.RUnlock()
	argsForCall := fake.getTableSizeEstimatesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDbAdmin) GetTableSizeEstimatesReturns(result1 map[dbadmin.TableName]uint64, result2 error) {
	fake.getTableSizeEstimatesMutex.Lock()
	defer fake.getTableSizeEstimatesMutex.Unlock()
	fake.GetTableSizeEstimatesStub = nil
	fake.getTableSizeEstimatesReturns = struct {
		result1 map[dbadmin.TableName]uint64
		result2 error
	}{result1, result2}
}

func (fake *FakeDbAdmin) GetTableSizeEstimatesReturnsOnCall(i int, result1 map[dbadmin.TableName]uint64, result2 error) {
	fake.getTableSizeEstimatesMutex.Lock()
	defer fake.getTableSizeEstimatesMutex.Unlock()
	fake.GetTableSizeEstimatesStub = nil
	if fake.getTableSizeEstimatesReturnsOnCall == nil {
		fake.getTableSizeEstimatesReturnsOnCall = make(map[int]struct {
			result1 map[dbadmin.TableName]uint64
			result2 error
		})
	}
	fake.getTableSizeEstimatesReturnsOnCall[i] = struct {
		result1 map[dbadmin.TableName]uint64
		result2 error
	}{result1, result2}
}

func (fake *FakeDbAdmin) IsBlockingIndexCreation(arg1 dbadmin.TableName, arg2 dbadmin.IndexType, arg3 ...string) (bool, error) {
	fake.isBlockingIndexCreationMutex.Lock()
	ret, specificReturn := fake.isBlockingIndexCreationReturnsOnCall[len(fake.isBlockingIndexCreationArgsForCall)]
	fake.isBlockingIndexCreationArgsForCall = append(fake.isBlockingIndexCreationArgsForCall, struct {
		arg1 dbadmin.TableName
		arg2 dbadmin.IndexType
		arg3 []string
	}{arg1, arg2, arg3})
	fake.recordInvocation("IsBlockingIndexCreation", []interface{}{arg1, arg2, arg3})
	fake.isBlockingIndexCreationMutex.Unlock()
	if fake.IsBlockingIndexCreationStub != nil {
		return fake.IsBlockingIndexCreationStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.isBlockingIndexCreationReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDbAdmin) IsBlockingIndexCreationCallCount() int {
	fake.isBlockingIndexCreationMutex.RLock()
	defer fake.isBlockingIndexCreationMutex.RUnlock()
	return len(fake.isBlockingIndexCreationArgsForCall)
}

func (fake *FakeDbAdmin) IsBlockingIndexCreationCalls(stub func(dbadmin.TableName, dbadmin.IndexType, ...string) (bool, error)) {
	fake.isBlockingIndexCreationMutex.Lock()
	defer fake.isBlockingIndexCreationMutex.Unlock()
	fake.IsBlockingIndexCreationStub = stub
}

func (fake *FakeDbAdmin) IsBlockingIndexCreationArgsForCall(i int) (dbadmin.TableName, dbadmin.IndexType, []string) {
	fake.isBlockingIndexCreationMutex.RLock()
	defer fake.isBlockingIndexCreationMutex.RUnlock()
	argsForCall := fake.isBlockingIndexCreationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDbAdmin) IsBlockingIndexCreationReturns(result1 bool, result2 error) {
	fake.isBlockingIndexCreationMutex.Lock()
	defer fake.isBlockingIndexCreationMutex.Unlock()
	fake.IsBlockingIndexCreationStub = nil
	fake.isBlockingIndexCreationReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeDbAdmin) IsBlockingIndexCreationReturnsOnCall(i int, result1 bool, result2 error) {
	fake.isBlockingIndexCreationMutex.Lock()
	defer fake.isBlockingIndexCreationMutex.Unlock()
	fake.IsBlockingIndexCreationStub = nil
	if fake.isBlockingIndexCreationReturnsOnCall == nil {
		fake.isBlockingIndexCreationReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.isBlockingIndexCreationReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeDbAdmin) ListUsernames(arg1 string) ([]string, error) {
	fake.listUsernamesMutex.Lock()
	ret, specificReturn := fake.listUsernamesReturnsOnCall[len(fake.listUsernamesArgsForCall)]
	fake.listUsernamesArgsForCall = append(fake.listUsernamesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ListUsernames", []interface{}{arg1})
	fake.listUsernamesMutex.Unlock()
	if fake.ListUsernamesStub != nil {
		return fake.ListUsernamesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listUsernamesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDbAdmin) ListUsernamesCallCount() int {
	fake.listUsernamesMutex.RLock()
	defer fake.listUsernamesMutex.RUnlock()
	return len(fake.listUsernamesArgsForCall)
}

func (fake *FakeDbAdmin) ListUsernamesCalls(stub func(string) ([]string, error)) {
	fake.listUsernamesMutex.Lock()
	defer fake.listUsernamesMutex.Unlock()
	fake.ListUsernamesStub = stub
}

func (fake *FakeDbAdmin) ListUsernamesArgsForCall(i int) string {
	fake.listUsernamesMutex.RLock()
	defer fake.listUsernamesMutex.RUnlock()
	argsForCall := fake.listUsernamesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDbAdmin) ListUsernamesReturns(result1 []string, result2 error) {
	fake.listUsernamesMutex.Lock()
	defer fake.listUsernamesMutex.Unlock()
	fake.ListUsernamesStub = nil
	fake.listUsernamesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeDbAdmin) ListUsernamesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.listUsernamesMutex.Lock()
	defer fake.listUsernamesMutex.Unlock()
	fake.ListUsernamesStub = nil
	if fake.listUsernamesReturnsOnCall == nil {
		fake.listUsernamesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listUsernamesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeDbAdmin) SelectFloat(arg1 string) (float64, error) {
	fake.selectFloatMutex.Lock()
	ret, specificReturn := fake.selectFloatReturnsOnCall[len(fake.selectFloatArgsForCall)]
	fake.selectFloatArgsForCall = append(fake.selectFloatArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SelectFloat", []interface{}{arg1})
	fake.selectFloatMutex.Unlock()
	if fake.SelectFloatStub != nil {
		return fake.SelectFloatStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.selectFloatReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDbAdmin) SelectFloatCallCount() int {
	fake.selectFloatMutex.RLock()
	defer fake.selectFloatMutex.RUnlock()
	return len(fake.selectFloatArgsForCall)
}

func (fake *FakeDbAdmin) SelectFloatCalls(stub func(string) (float64, error)) {
	fake.selectFloatMutex.Lock()
	defer fake.selectFloatMutex.Unlock()
	fake.SelectFloatStub = stub
}

func (fake *FakeDbAdmin) SelectFloatArgsForCall(i int) string {
	fake.selectFloatMutex.RLock()
	defer fake.selectFloatMutex.RUnlock()
	argsForCall := fake.selectFloatArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDbAdmin) SelectFloatReturns(result1 float64, result2 error) {
	fake.selectFloatMutex.Lock()
	defer fake.selectFloatMutex.Unlock()
	fake.SelectFloatStub = nil
	fake.selectFloatReturns = struct {
		result1 float64
		result2 error
	}{result1, result2}
}

func (fake *FakeDbAdmin) SelectFloatReturnsOnCall(i int, result1 float64, result2 error) {
	fake.selectFloatMutex.Lock()
	defer fake.selectFloatMutex.Unlock()
	fake.SelectFloatStub = nil
	if fake.selectFloatReturnsOnCall == nil {
		fake.selectFloatReturnsOnCall = make(map[int]struct {
			result1 float64
			result2 error
		})
	}
	fake.selectFloatReturnsOnCall[i] = struct {
		result1 float64
		result2 error
	}{result1, result2}
}

func (fake *FakeDbAdmin) VerifyUnusedAndDeleteCredentials(arg1 string) error {
	fake.verifyUnusedAndDeleteCredentialsMutex.Lock()
	ret, specificReturn := fake.verifyUnusedAndDeleteCredentialsReturnsOnCall[len(fake.verifyUnusedAndDeleteCredentialsArgsForCall)]
	fake.verifyUnusedAndDeleteCredentialsArgsForCall = append(fake.verifyUnusedAndDeleteCredentialsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("VerifyUnusedAndDeleteCredentials", []interface{}{arg1})
	fake.verifyUnusedAndDeleteCredentialsMutex.Unlock()
	if fake.VerifyUnusedAndDeleteCredentialsStub != nil {
		return fake.VerifyUnusedAndDeleteCredentialsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.verifyUnusedAndDeleteCredentialsReturns
	return fakeReturns.result1
}

func (fake *FakeDbAdmin) VerifyUnusedAndDeleteCredentialsCallCount() int {
	fake.verifyUnusedAndDeleteCredentialsMutex.RLock()
	defer fake.verifyUnusedAndDeleteCredentialsMutex.RUnlock()
	return len(fake.verifyUnusedAndDeleteCredentialsArgsForCall)
}

func (fake *FakeDbAdmin) VerifyUnusedAndDeleteCredentialsCalls(stub func(string) error) {
	fake.verifyUnusedAndDeleteCredentialsMutex.Lock()
	defer fake.verifyUnusedAndDeleteCredentialsMutex.Unlock()
	fake.VerifyUnusedAndDeleteCredentialsStub = stub
}

func (fake *FakeDbAdmin) VerifyUnusedAndDeleteCredentialsArgsForCall(i int) string {
	fake.verifyUnusedAndDeleteCredentialsMutex.RLock()
	defer fake.verifyUnusedAndDeleteCredentialsMutex.RUnlock()
	argsForCall := fake.verifyUnusedAndDeleteCredentialsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDbAdmin) VerifyUnusedAndDeleteCredentialsReturns(result1 error) {
	fake.verifyUnusedAndDeleteCredentialsMutex.Lock()
	defer fake.verifyUnusedAndDeleteCredentialsMutex.Unlock()
	fake.VerifyUnusedAndDeleteCredentialsStub = nil
	fake.verifyUnusedAndDeleteCredentialsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDbAdmin) VerifyUnusedAndDeleteCredentialsReturnsOnCall(i int, result1 error) {
	fake.verifyUnusedAndDeleteCredentialsMutex.Lock()
	defer fake.verifyUnusedAndDeleteCredentialsMutex.Unlock()
	fake.VerifyUnusedAndDeleteCredentialsStub = nil
	if fake.verifyUnusedAndDeleteCredentialsReturnsOnCall == nil {
		fake.verifyUnusedAndDeleteCredentialsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.verifyUnusedAndDeleteCredentialsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDbAdmin) WriteCredentials(arg1 string, arg2 string) error {
	fake.writeCredentialsMutex.Lock()
	ret, specificReturn := fake.writeCredentialsReturnsOnCall[len(fake.writeCredentialsArgsForCall)]
	fake.writeCredentialsArgsForCall = append(fake.writeCredentialsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("WriteCredentials", []interface{}{arg1, arg2})
	fake.writeCredentialsMutex.Unlock()
	if fake.WriteCredentialsStub != nil {
		return fake.WriteCredentialsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.writeCredentialsReturns
	return fakeReturns.result1
}

func (fake *FakeDbAdmin) WriteCredentialsCallCount() int {
	fake.writeCredentialsMutex.RLock()
	defer fake.writeCredentialsMutex.RUnlock()
	return len(fake.writeCredentialsArgsForCall)
}

func (fake *FakeDbAdmin) WriteCredentialsCalls(stub func(string, string) error) {
	fake.writeCredentialsMutex.Lock()
	defer fake.writeCredentialsMutex.Unlock()
	fake.WriteCredentialsStub = stub
}

func (fake *FakeDbAdmin) WriteCredentialsArgsForCall(i int) (string, string) {
	fake.writeCredentialsMutex.RLock()
	defer fake.writeCredentialsMutex.RUnlock()
	argsForCall := fake.writeCredentialsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDbAdmin) WriteCredentialsReturns(result1 error) {
	fake.writeCredentialsMutex.Lock()
	defer fake.writeCredentialsMutex.Unlock()
	fake.WriteCredentialsStub = nil
	fake.writeCredentialsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDbAdmin) WriteCredentialsReturnsOnCall(i int, result1 error) {
	fake.writeCredentialsMutex.Lock()
	defer fake.writeCredentialsMutex.Unlock()
	fake.WriteCredentialsStub = nil
	if fake.writeCredentialsReturnsOnCall == nil {
		fake.writeCredentialsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeCredentialsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDbAdmin) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.constraintWillFailMutex.RLock()
	defer fake.constraintWillFailMutex.RUnlock()
	fake.getNextIDsMutex.RLock()
	defer fake.getNextIDsMutex.RUnlock()
	fake.getSchemaVersionMutex.RLock()
	defer fake.getSchemaVersionMutex.RUnlock()
	fake.getTableSizeEstimatesMutex.RLock()
	defer fake.getTableSizeEstimatesMutex.RUnlock()
	fake.isBlockingIndexCreationMutex.RLock()
	defer fake.isBlockingIndexCreationMutex.RUnlock()
	fake.listUsernamesMutex.RLock()
	defer fake.listUsernamesMutex.RUnlock()
	fake.selectFloatMutex.RLock()
	defer fake.selectFloatMutex.RUnlock()
	fake.verifyUnusedAndDeleteCredentialsMutex.RLock()
	defer fake.verifyUnusedAndDeleteCredentialsMutex.RUnlock()
	fake.writeCredentialsMutex.RLock()
	defer fake.writeCredentialsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDbAdmin) recordInvocation(key string, args []interface{}) {
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

var _ dbadmin.DbAdmin = new(FakeDbAdmin)