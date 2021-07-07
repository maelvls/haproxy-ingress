/*
Copyright 2019 The HAProxy Ingress Controller Authors.

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

package types

import (
	"net"
	"time"

	api "k8s.io/api/core/v1"
)

// Cache ...
type Cache interface {
	ExternalNameLookup(externalName string) ([]net.IP, error)
	GetService(serviceName string) (*api.Service, error)
	GetEndpoints(service *api.Service) (*api.Endpoints, error)
	GetTerminatingPods(service *api.Service) ([]*api.Pod, error)
	GetPod(podName string) (*api.Pod, error)
	GetTLSSecretPath(defaultNamespace, secretName string) (CrtFile, error)
	GetCASecretPath(defaultNamespace, secretName string) (ca, crl File, err error)
	GetDHSecretPath(defaultNamespace, secretName string) (File, error)
	GetSecretContent(defaultNamespace, secretName, keyName string) ([]byte, error)
}

// File ...
type File struct {
	Filename string
	SHA1Hash string
}

// CrtFile ...
type CrtFile struct {
	Filename   string
	SHA1Hash   string
	CommonName string
	NotAfter   time.Time
}
