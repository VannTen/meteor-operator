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

package v1alpha1

const (

	// There is at least one ImageStreamTag available in the corresponding ImageStream.
	Ready = "Ready"

	// Reasons for (False -> True):
	// - ImagePushed
	// Reasons for (True -> False):
	// - ImageStreamDeleted
	// - ImageTagDeleted

	// The ImageStreamTag corresponding to the current CRE observedGeneration exists.
	ImageUpToDate = "ImageUpToDate"

	// Reasons for (False -> True):
	// - ImagePushed
	// Reasons for (True -> False):
	// - ImageStreamDeleted
	// - ImageTagDeleted
	// - NewCreVersion

	// The last pipeline building the image errored out.
	LatestBuildSucceeded = "LatestBuildSucceeded"

	// Reasons for (True -> Unknown || False -> Unknown)
	// - NewPipelineRun
	// Reasons for (Unknown -> True):
	// - PipelineRunSucceeded
	// Reasons for (Unknown -> False):
	// - PipelineRunFailed
	// More detailed causes of failure for pipelines should use the metav1.Condition.Message field
)
