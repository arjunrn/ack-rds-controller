// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package db_proxy

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if !reflect.DeepEqual(a.ko.Spec.Auth, b.ko.Spec.Auth) {
		delta.Add("Spec.Auth", a.ko.Spec.Auth, b.ko.Spec.Auth)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DebugLogging, b.ko.Spec.DebugLogging) {
		delta.Add("Spec.DebugLogging", a.ko.Spec.DebugLogging, b.ko.Spec.DebugLogging)
	} else if a.ko.Spec.DebugLogging != nil && b.ko.Spec.DebugLogging != nil {
		if *a.ko.Spec.DebugLogging != *b.ko.Spec.DebugLogging {
			delta.Add("Spec.DebugLogging", a.ko.Spec.DebugLogging, b.ko.Spec.DebugLogging)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.EngineFamily, b.ko.Spec.EngineFamily) {
		delta.Add("Spec.EngineFamily", a.ko.Spec.EngineFamily, b.ko.Spec.EngineFamily)
	} else if a.ko.Spec.EngineFamily != nil && b.ko.Spec.EngineFamily != nil {
		if *a.ko.Spec.EngineFamily != *b.ko.Spec.EngineFamily {
			delta.Add("Spec.EngineFamily", a.ko.Spec.EngineFamily, b.ko.Spec.EngineFamily)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.IdleClientTimeout, b.ko.Spec.IdleClientTimeout) {
		delta.Add("Spec.IdleClientTimeout", a.ko.Spec.IdleClientTimeout, b.ko.Spec.IdleClientTimeout)
	} else if a.ko.Spec.IdleClientTimeout != nil && b.ko.Spec.IdleClientTimeout != nil {
		if *a.ko.Spec.IdleClientTimeout != *b.ko.Spec.IdleClientTimeout {
			delta.Add("Spec.IdleClientTimeout", a.ko.Spec.IdleClientTimeout, b.ko.Spec.IdleClientTimeout)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Name, b.ko.Spec.Name) {
		delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
	} else if a.ko.Spec.Name != nil && b.ko.Spec.Name != nil {
		if *a.ko.Spec.Name != *b.ko.Spec.Name {
			delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RequireTLS, b.ko.Spec.RequireTLS) {
		delta.Add("Spec.RequireTLS", a.ko.Spec.RequireTLS, b.ko.Spec.RequireTLS)
	} else if a.ko.Spec.RequireTLS != nil && b.ko.Spec.RequireTLS != nil {
		if *a.ko.Spec.RequireTLS != *b.ko.Spec.RequireTLS {
			delta.Add("Spec.RequireTLS", a.ko.Spec.RequireTLS, b.ko.Spec.RequireTLS)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RoleARN, b.ko.Spec.RoleARN) {
		delta.Add("Spec.RoleARN", a.ko.Spec.RoleARN, b.ko.Spec.RoleARN)
	} else if a.ko.Spec.RoleARN != nil && b.ko.Spec.RoleARN != nil {
		if *a.ko.Spec.RoleARN != *b.ko.Spec.RoleARN {
			delta.Add("Spec.RoleARN", a.ko.Spec.RoleARN, b.ko.Spec.RoleARN)
		}
	}
	if !ackcompare.MapStringStringEqual(ToACKTags(a.ko.Spec.Tags), ToACKTags(b.ko.Spec.Tags)) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}
	if !ackcompare.SliceStringPEqual(a.ko.Spec.VPCSecurityGroupIDs, b.ko.Spec.VPCSecurityGroupIDs) {
		delta.Add("Spec.VPCSecurityGroupIDs", a.ko.Spec.VPCSecurityGroupIDs, b.ko.Spec.VPCSecurityGroupIDs)
	}
	if !ackcompare.SliceStringPEqual(a.ko.Spec.VPCSubnetIDs, b.ko.Spec.VPCSubnetIDs) {
		delta.Add("Spec.VPCSubnetIDs", a.ko.Spec.VPCSubnetIDs, b.ko.Spec.VPCSubnetIDs)
	}

	return delta
}
