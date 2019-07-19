// +build !ignore_autogenerated

// Copyright (c) 2019 Sylabs, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.JobResults":     schema_operator_apis_wlm_v1alpha1_JobResults(ref),
		"github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.SlurmJob":       schema_operator_apis_wlm_v1alpha1_SlurmJob(ref),
		"github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.SlurmJobSpec":   schema_operator_apis_wlm_v1alpha1_SlurmJobSpec(ref),
		"github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.SlurmJobStatus": schema_operator_apis_wlm_v1alpha1_SlurmJobStatus(ref),
		"github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.WlmJob":         schema_operator_apis_wlm_v1alpha1_WlmJob(ref),
		"github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.WlmJobSpec":     schema_operator_apis_wlm_v1alpha1_WlmJobSpec(ref),
		"github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.WlmJobStatus":   schema_operator_apis_wlm_v1alpha1_WlmJobStatus(ref),
		"github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.WlmResources":   schema_operator_apis_wlm_v1alpha1_WlmResources(ref),
	}
}

func schema_operator_apis_wlm_v1alpha1_JobResults(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JobResults is a schema for results collection.",
				Properties: map[string]spec.Schema{
					"mount": {
						SchemaProps: spec.SchemaProps{
							Description: "Mount is a directory where job results will be stored. After results collection all job generated files can be found in Mount/<SlurmJob.Name> directory.",
							Ref:         ref("k8s.io/api/core/v1.Volume"),
						},
					},
					"from": {
						SchemaProps: spec.SchemaProps{
							Description: "From is a path to the results to be collected from a Slurm cluster.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"mount", "from"},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.Volume"},
	}
}

func schema_operator_apis_wlm_v1alpha1_SlurmJob(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SlurmJob is the Schema for the slurmjobs API.",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.SlurmJobSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.SlurmJobStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.SlurmJobSpec", "github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.SlurmJobStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_operator_apis_wlm_v1alpha1_SlurmJobSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SlurmJobSpec defines the desired state of SlurmJob",
				Properties: map[string]spec.Schema{
					"batch": {
						SchemaProps: spec.SchemaProps{
							Description: "Batch is a script that will be submitted to a Slurm cluster as a batch job.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"nodeSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeSelector is a selector which must be true for the SlurmJob to fit on a node. Selector which must match a node's labels for the SlurmJob to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"results": {
						SchemaProps: spec.SchemaProps{
							Description: "Results may be specified for an optional results collection step. When specified, after job is completed all results will be downloaded from Slurm cluster with respect to this configuration.",
							Ref:         ref("github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.JobResults"),
						},
					},
				},
				Required: []string{"batch"},
			},
		},
		Dependencies: []string{
			"github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.JobResults"},
	}
}

func schema_operator_apis_wlm_v1alpha1_SlurmJobStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SlurmJobStatus defines the observed state of a SlurmJob.",
				Properties: map[string]spec.Schema{
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status reflects job status, e.g running, succeeded.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"status"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_operator_apis_wlm_v1alpha1_WlmJob(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WlmJob is the Schema for the wlm jobs API.",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.WlmJobSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.WlmJobStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.WlmJobSpec", "github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.WlmJobStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_operator_apis_wlm_v1alpha1_WlmJobSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WlmJobSpec defines the desired state of WlmJob",
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Image name to start as a job",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Description: "Resources describes required for a job resources",
							Ref:         ref("github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.WlmResources"),
						},
					},
					"nodeSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeSelector is a selector which must be true for the WlmJob to fit on a node. Selector which must match a node's labels for the WlmJob to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"results": {
						SchemaProps: spec.SchemaProps{
							Description: "Results may be specified for an optional results collection step. When specified, after job is completed all results will be downloaded from WLM cluster with respect to this configuration.",
							Ref:         ref("github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.JobResults"),
						},
					},
				},
				Required: []string{"image"},
			},
		},
		Dependencies: []string{
			"github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.JobResults", "github.com/sylabs/wlm-operator/pkg/operator/apis/wlm/v1alpha1.WlmResources"},
	}
}

func schema_operator_apis_wlm_v1alpha1_WlmJobStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WlmJobStatus defines the observed state of a WlmJob.",
				Properties: map[string]spec.Schema{
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status reflects job status, e.g running, succeeded.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"status"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_operator_apis_wlm_v1alpha1_WlmResources(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WlmResources is a schema for wlm resources.",
				Properties: map[string]spec.Schema{
					"nodes": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
					"cpuPerNode": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
					"memPerNode": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
					"wallTime": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}
