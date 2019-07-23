/*
Copyright (C) 2018 Synopsys, Inc.

Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements. See the NOTICE file
distributed with this work for additional information
regarding copyright ownership. The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied. See the License for the
specific language governing permissions and limitations
under the License.
*/

package api

import (
	"github.com/blackducksoftware/horizon/pkg/components"
	routev1 "github.com/openshift/api/route/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// Route defines the route component
type Route struct {
	Namespace          string
	Name               string
	Kind               string
	ServiceName        string
	PortName           string
	Labels             map[string]string
	TLSTerminationType routev1.TLSTerminationType
}

// ComponentList defines the list of components for an app
type ComponentList struct {
	ReplicationControllers    []*components.ReplicationController
	Services                  []*components.Service
	ConfigMaps                []*components.ConfigMap
	ServiceAccounts           []*components.ServiceAccount
	ClusterRoleBindings       []*components.ClusterRoleBinding
	ClusterRoles              []*components.ClusterRole
	RoleBindings              []*components.RoleBinding
	Roles                     []*components.Role
	Deployments               []*components.Deployment
	Secrets                   []*components.Secret
	PersistentVolumeClaims    []*components.PersistentVolumeClaim
	Routes                    []*Route
	CustomResourceDefinitions []*components.CustomResourceDefinition
}

// GetKubeInterfaces returns a list of kube components as interfaces
func (clist *ComponentList) GetKubeInterfaces() []interface{} {
	components := []interface{}{}
	for _, rc := range clist.ReplicationControllers {
		components = append(components, rc.ReplicationController)
	}
	for _, svc := range clist.Services {
		components = append(components, svc.Service)
	}
	for _, cm := range clist.ConfigMaps {
		components = append(components, cm.ConfigMap)
	}
	for _, sa := range clist.ServiceAccounts {
		components = append(components, sa.ServiceAccount)
	}
	for _, crb := range clist.ClusterRoleBindings {
		components = append(components, crb.ClusterRoleBinding)
	}
	for _, cr := range clist.ClusterRoles {
		components = append(components, cr.ClusterRole)
	}
	for _, d := range clist.Deployments {
		components = append(components, d.Deployment)
	}
	for _, sec := range clist.Secrets {
		components = append(components, sec.Secret)
	}
	for _, pvc := range clist.PersistentVolumeClaims {
		components = append(components, pvc.PersistentVolumeClaim)
	}
	return components
}

func (clist *ComponentList) Filter(filter string) (*ComponentList, error) {
	components := &ComponentList{}

	labelSelector, err := metav1.ParseToLabelSelector(filter)
	if err != nil {
		return nil, err
	}
	selector, err := metav1.LabelSelectorAsSelector(labelSelector)
	if err != nil {
		return nil, err
	}
	if clist.ReplicationControllers != nil {
		for _, rc := range clist.ReplicationControllers {
			if rc != nil && selector.Matches(labels.Set(rc.Labels)) {
				components.ReplicationControllers = append(components.ReplicationControllers, rc)
			}
		}
	}

	if clist.Services != nil {
		for _, svc := range clist.Services {
			if svc != nil && selector.Matches(labels.Set(svc.Labels)) {
				components.Services = append(components.Services, svc)
			}
		}
	}

	if clist.ConfigMaps != nil {
		for _, cm := range clist.ConfigMaps {
			if cm != nil && selector.Matches(labels.Set(cm.Labels)) {
				components.ConfigMaps = append(components.ConfigMaps, cm)
			}
		}
	}

	if clist.ServiceAccounts != nil {
		for _, sa := range clist.ServiceAccounts {
			if sa != nil && selector.Matches(labels.Set(sa.Labels)) {
				components.ServiceAccounts = append(components.ServiceAccounts, sa)
			}
		}
	}

	if clist.ClusterRoleBindings != nil {
		for _, crb := range clist.ClusterRoleBindings {
			if crb != nil && selector.Matches(labels.Set(crb.Labels)) {
				components.ClusterRoleBindings = append(components.ClusterRoleBindings, crb)
			}
		}
	}

	if clist.ClusterRoles != nil {
		for _, cr := range clist.ClusterRoles {
			if cr != nil && selector.Matches(labels.Set(cr.Labels)) {
				components.ClusterRoles = append(components.ClusterRoles, cr)
			}
		}
	}

	if clist.Deployments != nil {
		for _, d := range clist.Deployments {
			if d != nil && selector.Matches(labels.Set(d.Labels)) {
				components.Deployments = append(components.Deployments, d)
			}
		}
	}

	if clist.Secrets != nil {
		for _, sec := range clist.Secrets {
			if sec != nil && selector.Matches(labels.Set(sec.Labels)) {
				components.Secrets = append(components.Secrets, sec)
			}
		}
	}

	if clist.PersistentVolumeClaims != nil {
		for _, pvc := range clist.PersistentVolumeClaims {
			if pvc != nil && selector.Matches(labels.Set(pvc.Labels)) {
				components.PersistentVolumeClaims = append(components.PersistentVolumeClaims, pvc)
			}
		}
	}
	return components, nil
}
