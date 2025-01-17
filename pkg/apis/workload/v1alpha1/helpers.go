/*
Copyright 2022 The KCP Authors.

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

package v1alpha1

import (
	"crypto/sha256"
	"math/big"

	"github.com/kcp-dev/logicalcluster/v2"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetResourceState returns the state of the resource for the given sync target, and
// whether the state value is a valid state. A missing label is considered invalid.
func GetResourceState(obj metav1.Object, syncTargetKey string) (state ResourceState, valid bool) {
	value, found := obj.GetLabels()[ClusterResourceStateLabelPrefix+syncTargetKey]
	return ResourceState(value), found && (value == "" || ResourceState(value) == ResourceStateSync)
}

// ToSyncTargetKey hashes the SyncTarget workspace and the SyncTarget name to a string that is used to idenfity
// in a unique way the synctarget in annotations/labels/finalizers.
func ToSyncTargetKey(syncTargetWorkspace logicalcluster.Name, syncTargetName string) string {
	hash := sha256.Sum224([]byte(syncTargetWorkspace.String() + syncTargetName))
	base62hash := toBase62(hash)
	return base62hash
}

func toBase62(hash [28]byte) string {
	var i big.Int
	i.SetBytes(hash[:])
	return i.Text(62)
}
