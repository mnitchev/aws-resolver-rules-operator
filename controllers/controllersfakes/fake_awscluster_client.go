// Code generated by counterfeiter. DO NOT EDIT.
package controllersfakes

import (
	"context"
	"sync"

	"github.com/aws-resolver-rules-operator/controllers"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/cluster-api-provider-aws/api/v1beta1"
	v1beta1a "sigs.k8s.io/cluster-api/api/v1beta1"
)

type FakeAWSClusterClient struct {
	AddFinalizerStub        func(context.Context, *v1beta1.AWSCluster, string) error
	addFinalizerMutex       sync.RWMutex
	addFinalizerArgsForCall []struct {
		arg1 context.Context
		arg2 *v1beta1.AWSCluster
		arg3 string
	}
	addFinalizerReturns struct {
		result1 error
	}
	addFinalizerReturnsOnCall map[int]struct {
		result1 error
	}
	GetAWSClusterStub        func(context.Context, types.NamespacedName) (*v1beta1.AWSCluster, error)
	getAWSClusterMutex       sync.RWMutex
	getAWSClusterArgsForCall []struct {
		arg1 context.Context
		arg2 types.NamespacedName
	}
	getAWSClusterReturns struct {
		result1 *v1beta1.AWSCluster
		result2 error
	}
	getAWSClusterReturnsOnCall map[int]struct {
		result1 *v1beta1.AWSCluster
		result2 error
	}
	GetClusterStub        func(context.Context, types.NamespacedName) (*v1beta1a.Cluster, error)
	getClusterMutex       sync.RWMutex
	getClusterArgsForCall []struct {
		arg1 context.Context
		arg2 types.NamespacedName
	}
	getClusterReturns struct {
		result1 *v1beta1a.Cluster
		result2 error
	}
	getClusterReturnsOnCall map[int]struct {
		result1 *v1beta1a.Cluster
		result2 error
	}
	GetIdentityStub        func(context.Context, *v1beta1.AWSCluster) (*v1beta1.AWSClusterRoleIdentity, error)
	getIdentityMutex       sync.RWMutex
	getIdentityArgsForCall []struct {
		arg1 context.Context
		arg2 *v1beta1.AWSCluster
	}
	getIdentityReturns struct {
		result1 *v1beta1.AWSClusterRoleIdentity
		result2 error
	}
	getIdentityReturnsOnCall map[int]struct {
		result1 *v1beta1.AWSClusterRoleIdentity
		result2 error
	}
	GetOwnerStub        func(context.Context, *v1beta1.AWSCluster) (*v1beta1a.Cluster, error)
	getOwnerMutex       sync.RWMutex
	getOwnerArgsForCall []struct {
		arg1 context.Context
		arg2 *v1beta1.AWSCluster
	}
	getOwnerReturns struct {
		result1 *v1beta1a.Cluster
		result2 error
	}
	getOwnerReturnsOnCall map[int]struct {
		result1 *v1beta1a.Cluster
		result2 error
	}
	MarkConditionTrueStub        func(context.Context, *v1beta1.AWSCluster, v1beta1a.ConditionType) error
	markConditionTrueMutex       sync.RWMutex
	markConditionTrueArgsForCall []struct {
		arg1 context.Context
		arg2 *v1beta1.AWSCluster
		arg3 v1beta1a.ConditionType
	}
	markConditionTrueReturns struct {
		result1 error
	}
	markConditionTrueReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveFinalizerStub        func(context.Context, *v1beta1.AWSCluster, string) error
	removeFinalizerMutex       sync.RWMutex
	removeFinalizerArgsForCall []struct {
		arg1 context.Context
		arg2 *v1beta1.AWSCluster
		arg3 string
	}
	removeFinalizerReturns struct {
		result1 error
	}
	removeFinalizerReturnsOnCall map[int]struct {
		result1 error
	}
	UnpauseStub        func(context.Context, *v1beta1.AWSCluster, *v1beta1a.Cluster) error
	unpauseMutex       sync.RWMutex
	unpauseArgsForCall []struct {
		arg1 context.Context
		arg2 *v1beta1.AWSCluster
		arg3 *v1beta1a.Cluster
	}
	unpauseReturns struct {
		result1 error
	}
	unpauseReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAWSClusterClient) AddFinalizer(arg1 context.Context, arg2 *v1beta1.AWSCluster, arg3 string) error {
	fake.addFinalizerMutex.Lock()
	ret, specificReturn := fake.addFinalizerReturnsOnCall[len(fake.addFinalizerArgsForCall)]
	fake.addFinalizerArgsForCall = append(fake.addFinalizerArgsForCall, struct {
		arg1 context.Context
		arg2 *v1beta1.AWSCluster
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.AddFinalizerStub
	fakeReturns := fake.addFinalizerReturns
	fake.recordInvocation("AddFinalizer", []interface{}{arg1, arg2, arg3})
	fake.addFinalizerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAWSClusterClient) AddFinalizerCallCount() int {
	fake.addFinalizerMutex.RLock()
	defer fake.addFinalizerMutex.RUnlock()
	return len(fake.addFinalizerArgsForCall)
}

func (fake *FakeAWSClusterClient) AddFinalizerCalls(stub func(context.Context, *v1beta1.AWSCluster, string) error) {
	fake.addFinalizerMutex.Lock()
	defer fake.addFinalizerMutex.Unlock()
	fake.AddFinalizerStub = stub
}

func (fake *FakeAWSClusterClient) AddFinalizerArgsForCall(i int) (context.Context, *v1beta1.AWSCluster, string) {
	fake.addFinalizerMutex.RLock()
	defer fake.addFinalizerMutex.RUnlock()
	argsForCall := fake.addFinalizerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAWSClusterClient) AddFinalizerReturns(result1 error) {
	fake.addFinalizerMutex.Lock()
	defer fake.addFinalizerMutex.Unlock()
	fake.AddFinalizerStub = nil
	fake.addFinalizerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAWSClusterClient) AddFinalizerReturnsOnCall(i int, result1 error) {
	fake.addFinalizerMutex.Lock()
	defer fake.addFinalizerMutex.Unlock()
	fake.AddFinalizerStub = nil
	if fake.addFinalizerReturnsOnCall == nil {
		fake.addFinalizerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addFinalizerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAWSClusterClient) GetAWSCluster(arg1 context.Context, arg2 types.NamespacedName) (*v1beta1.AWSCluster, error) {
	fake.getAWSClusterMutex.Lock()
	ret, specificReturn := fake.getAWSClusterReturnsOnCall[len(fake.getAWSClusterArgsForCall)]
	fake.getAWSClusterArgsForCall = append(fake.getAWSClusterArgsForCall, struct {
		arg1 context.Context
		arg2 types.NamespacedName
	}{arg1, arg2})
	stub := fake.GetAWSClusterStub
	fakeReturns := fake.getAWSClusterReturns
	fake.recordInvocation("GetAWSCluster", []interface{}{arg1, arg2})
	fake.getAWSClusterMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAWSClusterClient) GetAWSClusterCallCount() int {
	fake.getAWSClusterMutex.RLock()
	defer fake.getAWSClusterMutex.RUnlock()
	return len(fake.getAWSClusterArgsForCall)
}

func (fake *FakeAWSClusterClient) GetAWSClusterCalls(stub func(context.Context, types.NamespacedName) (*v1beta1.AWSCluster, error)) {
	fake.getAWSClusterMutex.Lock()
	defer fake.getAWSClusterMutex.Unlock()
	fake.GetAWSClusterStub = stub
}

func (fake *FakeAWSClusterClient) GetAWSClusterArgsForCall(i int) (context.Context, types.NamespacedName) {
	fake.getAWSClusterMutex.RLock()
	defer fake.getAWSClusterMutex.RUnlock()
	argsForCall := fake.getAWSClusterArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAWSClusterClient) GetAWSClusterReturns(result1 *v1beta1.AWSCluster, result2 error) {
	fake.getAWSClusterMutex.Lock()
	defer fake.getAWSClusterMutex.Unlock()
	fake.GetAWSClusterStub = nil
	fake.getAWSClusterReturns = struct {
		result1 *v1beta1.AWSCluster
		result2 error
	}{result1, result2}
}

func (fake *FakeAWSClusterClient) GetAWSClusterReturnsOnCall(i int, result1 *v1beta1.AWSCluster, result2 error) {
	fake.getAWSClusterMutex.Lock()
	defer fake.getAWSClusterMutex.Unlock()
	fake.GetAWSClusterStub = nil
	if fake.getAWSClusterReturnsOnCall == nil {
		fake.getAWSClusterReturnsOnCall = make(map[int]struct {
			result1 *v1beta1.AWSCluster
			result2 error
		})
	}
	fake.getAWSClusterReturnsOnCall[i] = struct {
		result1 *v1beta1.AWSCluster
		result2 error
	}{result1, result2}
}

func (fake *FakeAWSClusterClient) GetCluster(arg1 context.Context, arg2 types.NamespacedName) (*v1beta1a.Cluster, error) {
	fake.getClusterMutex.Lock()
	ret, specificReturn := fake.getClusterReturnsOnCall[len(fake.getClusterArgsForCall)]
	fake.getClusterArgsForCall = append(fake.getClusterArgsForCall, struct {
		arg1 context.Context
		arg2 types.NamespacedName
	}{arg1, arg2})
	stub := fake.GetClusterStub
	fakeReturns := fake.getClusterReturns
	fake.recordInvocation("GetCluster", []interface{}{arg1, arg2})
	fake.getClusterMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAWSClusterClient) GetClusterCallCount() int {
	fake.getClusterMutex.RLock()
	defer fake.getClusterMutex.RUnlock()
	return len(fake.getClusterArgsForCall)
}

func (fake *FakeAWSClusterClient) GetClusterCalls(stub func(context.Context, types.NamespacedName) (*v1beta1a.Cluster, error)) {
	fake.getClusterMutex.Lock()
	defer fake.getClusterMutex.Unlock()
	fake.GetClusterStub = stub
}

func (fake *FakeAWSClusterClient) GetClusterArgsForCall(i int) (context.Context, types.NamespacedName) {
	fake.getClusterMutex.RLock()
	defer fake.getClusterMutex.RUnlock()
	argsForCall := fake.getClusterArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAWSClusterClient) GetClusterReturns(result1 *v1beta1a.Cluster, result2 error) {
	fake.getClusterMutex.Lock()
	defer fake.getClusterMutex.Unlock()
	fake.GetClusterStub = nil
	fake.getClusterReturns = struct {
		result1 *v1beta1a.Cluster
		result2 error
	}{result1, result2}
}

func (fake *FakeAWSClusterClient) GetClusterReturnsOnCall(i int, result1 *v1beta1a.Cluster, result2 error) {
	fake.getClusterMutex.Lock()
	defer fake.getClusterMutex.Unlock()
	fake.GetClusterStub = nil
	if fake.getClusterReturnsOnCall == nil {
		fake.getClusterReturnsOnCall = make(map[int]struct {
			result1 *v1beta1a.Cluster
			result2 error
		})
	}
	fake.getClusterReturnsOnCall[i] = struct {
		result1 *v1beta1a.Cluster
		result2 error
	}{result1, result2}
}

func (fake *FakeAWSClusterClient) GetIdentity(arg1 context.Context, arg2 *v1beta1.AWSCluster) (*v1beta1.AWSClusterRoleIdentity, error) {
	fake.getIdentityMutex.Lock()
	ret, specificReturn := fake.getIdentityReturnsOnCall[len(fake.getIdentityArgsForCall)]
	fake.getIdentityArgsForCall = append(fake.getIdentityArgsForCall, struct {
		arg1 context.Context
		arg2 *v1beta1.AWSCluster
	}{arg1, arg2})
	stub := fake.GetIdentityStub
	fakeReturns := fake.getIdentityReturns
	fake.recordInvocation("GetIdentity", []interface{}{arg1, arg2})
	fake.getIdentityMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAWSClusterClient) GetIdentityCallCount() int {
	fake.getIdentityMutex.RLock()
	defer fake.getIdentityMutex.RUnlock()
	return len(fake.getIdentityArgsForCall)
}

func (fake *FakeAWSClusterClient) GetIdentityCalls(stub func(context.Context, *v1beta1.AWSCluster) (*v1beta1.AWSClusterRoleIdentity, error)) {
	fake.getIdentityMutex.Lock()
	defer fake.getIdentityMutex.Unlock()
	fake.GetIdentityStub = stub
}

func (fake *FakeAWSClusterClient) GetIdentityArgsForCall(i int) (context.Context, *v1beta1.AWSCluster) {
	fake.getIdentityMutex.RLock()
	defer fake.getIdentityMutex.RUnlock()
	argsForCall := fake.getIdentityArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAWSClusterClient) GetIdentityReturns(result1 *v1beta1.AWSClusterRoleIdentity, result2 error) {
	fake.getIdentityMutex.Lock()
	defer fake.getIdentityMutex.Unlock()
	fake.GetIdentityStub = nil
	fake.getIdentityReturns = struct {
		result1 *v1beta1.AWSClusterRoleIdentity
		result2 error
	}{result1, result2}
}

func (fake *FakeAWSClusterClient) GetIdentityReturnsOnCall(i int, result1 *v1beta1.AWSClusterRoleIdentity, result2 error) {
	fake.getIdentityMutex.Lock()
	defer fake.getIdentityMutex.Unlock()
	fake.GetIdentityStub = nil
	if fake.getIdentityReturnsOnCall == nil {
		fake.getIdentityReturnsOnCall = make(map[int]struct {
			result1 *v1beta1.AWSClusterRoleIdentity
			result2 error
		})
	}
	fake.getIdentityReturnsOnCall[i] = struct {
		result1 *v1beta1.AWSClusterRoleIdentity
		result2 error
	}{result1, result2}
}

func (fake *FakeAWSClusterClient) GetOwner(arg1 context.Context, arg2 *v1beta1.AWSCluster) (*v1beta1a.Cluster, error) {
	fake.getOwnerMutex.Lock()
	ret, specificReturn := fake.getOwnerReturnsOnCall[len(fake.getOwnerArgsForCall)]
	fake.getOwnerArgsForCall = append(fake.getOwnerArgsForCall, struct {
		arg1 context.Context
		arg2 *v1beta1.AWSCluster
	}{arg1, arg2})
	stub := fake.GetOwnerStub
	fakeReturns := fake.getOwnerReturns
	fake.recordInvocation("GetOwner", []interface{}{arg1, arg2})
	fake.getOwnerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAWSClusterClient) GetOwnerCallCount() int {
	fake.getOwnerMutex.RLock()
	defer fake.getOwnerMutex.RUnlock()
	return len(fake.getOwnerArgsForCall)
}

func (fake *FakeAWSClusterClient) GetOwnerCalls(stub func(context.Context, *v1beta1.AWSCluster) (*v1beta1a.Cluster, error)) {
	fake.getOwnerMutex.Lock()
	defer fake.getOwnerMutex.Unlock()
	fake.GetOwnerStub = stub
}

func (fake *FakeAWSClusterClient) GetOwnerArgsForCall(i int) (context.Context, *v1beta1.AWSCluster) {
	fake.getOwnerMutex.RLock()
	defer fake.getOwnerMutex.RUnlock()
	argsForCall := fake.getOwnerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAWSClusterClient) GetOwnerReturns(result1 *v1beta1a.Cluster, result2 error) {
	fake.getOwnerMutex.Lock()
	defer fake.getOwnerMutex.Unlock()
	fake.GetOwnerStub = nil
	fake.getOwnerReturns = struct {
		result1 *v1beta1a.Cluster
		result2 error
	}{result1, result2}
}

func (fake *FakeAWSClusterClient) GetOwnerReturnsOnCall(i int, result1 *v1beta1a.Cluster, result2 error) {
	fake.getOwnerMutex.Lock()
	defer fake.getOwnerMutex.Unlock()
	fake.GetOwnerStub = nil
	if fake.getOwnerReturnsOnCall == nil {
		fake.getOwnerReturnsOnCall = make(map[int]struct {
			result1 *v1beta1a.Cluster
			result2 error
		})
	}
	fake.getOwnerReturnsOnCall[i] = struct {
		result1 *v1beta1a.Cluster
		result2 error
	}{result1, result2}
}

func (fake *FakeAWSClusterClient) MarkConditionTrue(arg1 context.Context, arg2 *v1beta1.AWSCluster, arg3 v1beta1a.ConditionType) error {
	fake.markConditionTrueMutex.Lock()
	ret, specificReturn := fake.markConditionTrueReturnsOnCall[len(fake.markConditionTrueArgsForCall)]
	fake.markConditionTrueArgsForCall = append(fake.markConditionTrueArgsForCall, struct {
		arg1 context.Context
		arg2 *v1beta1.AWSCluster
		arg3 v1beta1a.ConditionType
	}{arg1, arg2, arg3})
	stub := fake.MarkConditionTrueStub
	fakeReturns := fake.markConditionTrueReturns
	fake.recordInvocation("MarkConditionTrue", []interface{}{arg1, arg2, arg3})
	fake.markConditionTrueMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAWSClusterClient) MarkConditionTrueCallCount() int {
	fake.markConditionTrueMutex.RLock()
	defer fake.markConditionTrueMutex.RUnlock()
	return len(fake.markConditionTrueArgsForCall)
}

func (fake *FakeAWSClusterClient) MarkConditionTrueCalls(stub func(context.Context, *v1beta1.AWSCluster, v1beta1a.ConditionType) error) {
	fake.markConditionTrueMutex.Lock()
	defer fake.markConditionTrueMutex.Unlock()
	fake.MarkConditionTrueStub = stub
}

func (fake *FakeAWSClusterClient) MarkConditionTrueArgsForCall(i int) (context.Context, *v1beta1.AWSCluster, v1beta1a.ConditionType) {
	fake.markConditionTrueMutex.RLock()
	defer fake.markConditionTrueMutex.RUnlock()
	argsForCall := fake.markConditionTrueArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAWSClusterClient) MarkConditionTrueReturns(result1 error) {
	fake.markConditionTrueMutex.Lock()
	defer fake.markConditionTrueMutex.Unlock()
	fake.MarkConditionTrueStub = nil
	fake.markConditionTrueReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAWSClusterClient) MarkConditionTrueReturnsOnCall(i int, result1 error) {
	fake.markConditionTrueMutex.Lock()
	defer fake.markConditionTrueMutex.Unlock()
	fake.MarkConditionTrueStub = nil
	if fake.markConditionTrueReturnsOnCall == nil {
		fake.markConditionTrueReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.markConditionTrueReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAWSClusterClient) RemoveFinalizer(arg1 context.Context, arg2 *v1beta1.AWSCluster, arg3 string) error {
	fake.removeFinalizerMutex.Lock()
	ret, specificReturn := fake.removeFinalizerReturnsOnCall[len(fake.removeFinalizerArgsForCall)]
	fake.removeFinalizerArgsForCall = append(fake.removeFinalizerArgsForCall, struct {
		arg1 context.Context
		arg2 *v1beta1.AWSCluster
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.RemoveFinalizerStub
	fakeReturns := fake.removeFinalizerReturns
	fake.recordInvocation("RemoveFinalizer", []interface{}{arg1, arg2, arg3})
	fake.removeFinalizerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAWSClusterClient) RemoveFinalizerCallCount() int {
	fake.removeFinalizerMutex.RLock()
	defer fake.removeFinalizerMutex.RUnlock()
	return len(fake.removeFinalizerArgsForCall)
}

func (fake *FakeAWSClusterClient) RemoveFinalizerCalls(stub func(context.Context, *v1beta1.AWSCluster, string) error) {
	fake.removeFinalizerMutex.Lock()
	defer fake.removeFinalizerMutex.Unlock()
	fake.RemoveFinalizerStub = stub
}

func (fake *FakeAWSClusterClient) RemoveFinalizerArgsForCall(i int) (context.Context, *v1beta1.AWSCluster, string) {
	fake.removeFinalizerMutex.RLock()
	defer fake.removeFinalizerMutex.RUnlock()
	argsForCall := fake.removeFinalizerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAWSClusterClient) RemoveFinalizerReturns(result1 error) {
	fake.removeFinalizerMutex.Lock()
	defer fake.removeFinalizerMutex.Unlock()
	fake.RemoveFinalizerStub = nil
	fake.removeFinalizerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAWSClusterClient) RemoveFinalizerReturnsOnCall(i int, result1 error) {
	fake.removeFinalizerMutex.Lock()
	defer fake.removeFinalizerMutex.Unlock()
	fake.RemoveFinalizerStub = nil
	if fake.removeFinalizerReturnsOnCall == nil {
		fake.removeFinalizerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeFinalizerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAWSClusterClient) Unpause(arg1 context.Context, arg2 *v1beta1.AWSCluster, arg3 *v1beta1a.Cluster) error {
	fake.unpauseMutex.Lock()
	ret, specificReturn := fake.unpauseReturnsOnCall[len(fake.unpauseArgsForCall)]
	fake.unpauseArgsForCall = append(fake.unpauseArgsForCall, struct {
		arg1 context.Context
		arg2 *v1beta1.AWSCluster
		arg3 *v1beta1a.Cluster
	}{arg1, arg2, arg3})
	stub := fake.UnpauseStub
	fakeReturns := fake.unpauseReturns
	fake.recordInvocation("Unpause", []interface{}{arg1, arg2, arg3})
	fake.unpauseMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAWSClusterClient) UnpauseCallCount() int {
	fake.unpauseMutex.RLock()
	defer fake.unpauseMutex.RUnlock()
	return len(fake.unpauseArgsForCall)
}

func (fake *FakeAWSClusterClient) UnpauseCalls(stub func(context.Context, *v1beta1.AWSCluster, *v1beta1a.Cluster) error) {
	fake.unpauseMutex.Lock()
	defer fake.unpauseMutex.Unlock()
	fake.UnpauseStub = stub
}

func (fake *FakeAWSClusterClient) UnpauseArgsForCall(i int) (context.Context, *v1beta1.AWSCluster, *v1beta1a.Cluster) {
	fake.unpauseMutex.RLock()
	defer fake.unpauseMutex.RUnlock()
	argsForCall := fake.unpauseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAWSClusterClient) UnpauseReturns(result1 error) {
	fake.unpauseMutex.Lock()
	defer fake.unpauseMutex.Unlock()
	fake.UnpauseStub = nil
	fake.unpauseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAWSClusterClient) UnpauseReturnsOnCall(i int, result1 error) {
	fake.unpauseMutex.Lock()
	defer fake.unpauseMutex.Unlock()
	fake.UnpauseStub = nil
	if fake.unpauseReturnsOnCall == nil {
		fake.unpauseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unpauseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAWSClusterClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addFinalizerMutex.RLock()
	defer fake.addFinalizerMutex.RUnlock()
	fake.getAWSClusterMutex.RLock()
	defer fake.getAWSClusterMutex.RUnlock()
	fake.getClusterMutex.RLock()
	defer fake.getClusterMutex.RUnlock()
	fake.getIdentityMutex.RLock()
	defer fake.getIdentityMutex.RUnlock()
	fake.getOwnerMutex.RLock()
	defer fake.getOwnerMutex.RUnlock()
	fake.markConditionTrueMutex.RLock()
	defer fake.markConditionTrueMutex.RUnlock()
	fake.removeFinalizerMutex.RLock()
	defer fake.removeFinalizerMutex.RUnlock()
	fake.unpauseMutex.RLock()
	defer fake.unpauseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAWSClusterClient) recordInvocation(key string, args []interface{}) {
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

var _ controllers.AWSClusterClient = new(FakeAWSClusterClient)
