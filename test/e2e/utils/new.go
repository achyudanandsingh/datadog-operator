// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2019 Datadog, Inc.

package utils

import (
	datadoghqv1alpha1 "github.com/DataDog/datadog-operator/pkg/apis/datadoghq/v1alpha1"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewDatadogAgentDeploymentOptions used to provide creation options to the NewDatadogAgentDeployment function
type NewDatadogAgentDeploymentOptions struct {
	ExtraLabels         map[string]string
	ExtraAnnotations    map[string]string
	ClusterAgentEnabled bool
	UseEDS              bool
}

var (
	pullPolicy = v1.PullIfNotPresent
)

// NewDatadogAgentDeployment returns new DatadogAgentDeployment instance with is config hash
func NewDatadogAgentDeployment(ns, name, image string, options *NewDatadogAgentDeploymentOptions) *datadoghqv1alpha1.DatadogAgentDeployment {
	ad := &datadoghqv1alpha1.DatadogAgentDeployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
		},
	}
	ad.Spec = datadoghqv1alpha1.DatadogAgentDeploymentSpec{
		Credentials: datadoghqv1alpha1.AgentCredentials{
			APIKey: "adflkajdflkjalkcmlkdjacsf",
			AppKey: "sgfggtdhfghfghfghfgbdfdgs",
		},
		Agent: &datadoghqv1alpha1.DatadogAgentDeploymentSpecAgentSpec{
			Image:              datadoghqv1alpha1.ImageConfig{},
			Config:             datadoghqv1alpha1.NodeAgentConfig{},
			DeploymentStrategy: &datadoghqv1alpha1.DaemonSetDeploymentcStrategy{},
			Apm:                datadoghqv1alpha1.APMSpec{},
			Log:                datadoghqv1alpha1.LogSpec{},
			Process:            datadoghqv1alpha1.ProcessSpec{},
		},
	}
	ad = datadoghqv1alpha1.DefaultDatadogAgentDeployment(ad)
	ad.Spec.Agent.Image = datadoghqv1alpha1.ImageConfig{
		Name:        image,
		PullPolicy:  &pullPolicy,
		PullSecrets: &[]v1.LocalObjectReference{},
	}
	ad.Spec.Agent.Rbac.Create = datadoghqv1alpha1.NewBoolPointer(true)

	if options != nil {
		if options.UseEDS && ad.Spec.Agent != nil {
			ad.Spec.Agent.UseExtendedDaemonset = &options.UseEDS
		}

		if options.ExtraLabels != nil {
			if ad.Labels == nil {
				ad.Labels = map[string]string{}
			}
			for key, val := range options.ExtraLabels {
				ad.Labels[key] = val
			}
		}

		if options.ExtraAnnotations != nil {
			if ad.Annotations == nil {
				ad.Annotations = map[string]string{}
			}
			for key, val := range options.ExtraAnnotations {
				ad.Annotations[key] = val
			}
		}

		if options.ClusterAgentEnabled {
			ad.Spec.ClusterAgent = &datadoghqv1alpha1.DatadogAgentDeploymentSpecClusterAgentSpec{
				Config: datadoghqv1alpha1.ClusterAgentConfig{},
				Image: datadoghqv1alpha1.ImageConfig{
					Name:        image,
					PullPolicy:  &pullPolicy,
					PullSecrets: &[]v1.LocalObjectReference{},
				},
			}
		}
	}

	return ad
}
