package dns

/*
Copyright 2022 The k8gb Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Generated by GoLic, for more details see: https://github.com/AbsaOSS/golic
*/

import (
	"testing"

	"github.com/k8gb-io/k8gb/controllers/utils"

	"github.com/k8gb-io/k8gb/controllers/depresolver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestFactoryInfoblox(t *testing.T) {
	// arrange
	client := fake.NewClientBuilder().WithScheme(scheme.Scheme).WithRuntimeObjects([]runtime.Object{}...).Build()
	customConfig := defaultConfig
	customConfig.EdgeDNSType = depresolver.DNSTypeInfoblox
	// act
	f, err := NewDNSProviderFactory(client, customConfig)
	require.NoError(t, err)
	provider := f.Provider()
	// assert
	assert.NotNil(t, provider)
	assert.Equal(t, "*InfobloxProvider", utils.GetType(provider))
	assert.Equal(t, "Infoblox", provider.String())
}

func TestFactoryExternal(t *testing.T) {
	// arrange
	client := fake.NewClientBuilder().WithScheme(scheme.Scheme).WithRuntimeObjects([]runtime.Object{}...).Build()
	customConfig := defaultConfig
	customConfig.EdgeDNSType = depresolver.DNSTypeExternal
	// act
	f, err := NewDNSProviderFactory(client, customConfig)
	require.NoError(t, err)
	provider := f.Provider()
	// assert
	assert.NotNil(t, provider)
	assert.Equal(t, "*ExternalDNSProvider", utils.GetType(provider))
	assert.Equal(t, "EXTDNS", provider.String())
}

func TestFactoryNoEdgeDNS(t *testing.T) {
	// arrange
	client := fake.NewClientBuilder().WithScheme(scheme.Scheme).WithRuntimeObjects([]runtime.Object{}...).Build()
	customConfig := defaultConfig
	customConfig.EdgeDNSType = depresolver.DNSTypeNoEdgeDNS
	// act
	f, err := NewDNSProviderFactory(client, customConfig)
	require.NoError(t, err)
	provider := f.Provider()
	// assert
	assert.Equal(t, "*EmptyDNSProvider", utils.GetType(provider))
	assert.Equal(t, "EMPTY", provider.String())
}

func TestFactoryNilClient(t *testing.T) {
	// arrange
	customConfig := defaultConfig
	customConfig.EdgeDNSType = depresolver.DNSTypeNoEdgeDNS
	// act
	// assert
	_, err := NewDNSProviderFactory(nil, customConfig)
	require.Error(t, err)
}
