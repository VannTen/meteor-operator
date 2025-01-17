/*
Copyright 2021, 2022 The Meteor Authors.

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

// SPDX-License-Identifier: Apache-2.0

package cnbi

import (
	pipelinev1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/utils/pointer"
)

var workspaces_const = []pipelinev1beta1.WorkspaceBinding{
	{
		Name: "data",
		VolumeClaimTemplate: &v1.PersistentVolumeClaim{
			Spec: v1.PersistentVolumeClaimSpec{
				AccessModes: []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce},
				Resources: v1.ResourceRequirements{
					Requests: v1.ResourceList{
						v1.ResourceStorage: resource.MustParse("500Mi"),
					},
				},
			},
		},
	},
	{
		Name: "sslcertdir",
		ConfigMap: &v1.ConfigMapVolumeSource{
			LocalObjectReference: v1.LocalObjectReference{
				Name: "openshift-service-ca.crt",
			},
			Items: []v1.KeyToPath{{
				Key:  "service-ca.crt",
				Path: "ca.crt",
			}},
			DefaultMode: pointer.Int32(420),
		},
	},
}
