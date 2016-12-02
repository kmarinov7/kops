/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package validation

import (
	"testing"
	"time"
	//k8sapi "k8s.io/kubernetes/pkg/api"
	k8sapi "k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/client/clientset_generated/release_1_5/fake"
)

//func TestBuildNodeAPIAdapter(t *testing.T) {
//
//}
//
//func TestGetReadySchedulableNodes(t *testing.T) {
//
//}

func TestWaitForNodeToBeReady(t *testing.T) {
	conditions := []k8sapi.NodeCondition{{Type: "Ready", Status: "True"}}
	nodeName := "node-foo"
	nodeAA := setupNodeAA(t, conditions, nodeName)

	test, err := nodeAA.WaitForNodeToBeReady(nodeName)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if test != true {
		t.Fatalf("unexpected error WaitForNodeToBeReady Failed: %v", test)
	}
}

func TestWaitForNodeToBeNotReady(t *testing.T) {
	conditions := []k8sapi.NodeCondition{{Type: "Ready", Status: "False"}}
	nodeName := "node-foo"
	nodeAA := setupNodeAA(t, conditions, nodeName)

	test, err := nodeAA.WaitForNodeToBeNotReady(nodeName)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if test != true {
		t.Fatalf("unexpected error WaitForNodeToBeNotReady Failed: %v", test)
	}
}

//func TestIsNodeConditionUnset(t *testing.T) {
//
//}

func setupNodeAA(t *testing.T, conditions []k8sapi.NodeCondition, nodeName string) *NodeAPIAdapter {
	node := &k8sapi.Node{
		ObjectMeta: k8sapi.ObjectMeta{Name: nodeName},
		Spec:       k8sapi.NodeSpec{Unschedulable: false},
		Status:     k8sapi.NodeStatus{Conditions: conditions},
	}

	c := fake.NewSimpleClientset(node)
	//c.Validate(t, response, err)
	nodeAA, err := NewNodeAPIAdapter(c, time.Duration(10)*time.Millisecond)
	if err != nil {
		t.Fatalf("unexpected error building NodeAPIAdapter: %v", err)
	}
	return nodeAA
}
