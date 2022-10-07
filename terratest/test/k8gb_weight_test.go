package test

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

	"k8gbterratest/utils"

	"github.com/stretchr/testify/require"
)

func TestWeightsExistsInLocalDNSEndpoint(t *testing.T) {
	t.Parallel()
	const host = "terratest-roundrobin.cloud.example.com"
	const endpointDNSNameEU = "gslb-ns-eu-cloud.example.com"
	const endpointDNSNameUS = "gslb-ns-us-cloud.example.com"
	const gslbPath = "../examples/roundrobin_weight1.yaml"
	instanceEU, err := utils.NewWorkflow(t, "k3d-test-gslb1", 5053).
		WithGslb(gslbPath, host).
		WithTestApp("eu").
		Start()
	require.NoError(t, err)
	defer instanceEU.Kill()

	instanceUS, err := utils.NewWorkflow(t, "k3d-test-gslb2", 5054).
		WithGslb(gslbPath, host).
		WithTestApp("us").
		Start()
	require.NoError(t, err)
	defer instanceUS.Kill()

	err = instanceEU.WaitForAppIsRunning()
	require.NoError(t, err)
	err = instanceUS.WaitForAppIsRunning()
	require.NoError(t, err)

	err = instanceEU.WaitForExternalDNSEndpointExists()
	require.NoError(t, err)
	err = instanceUS.WaitForExternalDNSEndpointExists()
	require.NoError(t, err)

	err = instanceEU.WaitForLocalDNSEndpointExists()
	require.NoError(t, err)
	err = instanceUS.WaitForLocalDNSEndpointExists()
	require.NoError(t, err)

	err = instanceEU.Resources().WaitForExternalDNSEndpointHasTargets(endpointDNSNameEU)
	require.NoError(t, err)
	err = instanceUS.Resources().WaitForExternalDNSEndpointHasTargets(endpointDNSNameUS)
	require.NoError(t, err)

	epExternalEU, err := instanceEU.Resources().GetExternalDNSEndpoint().GetEndpointByName(endpointDNSNameEU)
	require.NoError(t, err, "missing EU endpoint %s", endpointDNSNameEU)
	epExternalUS, err := instanceUS.Resources().GetExternalDNSEndpoint().GetEndpointByName(endpointDNSNameUS)
	require.NoError(t, err, "missing US endpoint %s", endpointDNSNameUS)
	t.Logf("ExternalDNS targets: EU: %v; US: %v", epExternalEU.Targets, epExternalUS.Targets)
	expectedTargets := append(epExternalEU.Targets, epExternalUS.Targets...)

	err = instanceUS.WaitForLocalDNSEndpointHasTargets(expectedTargets)
	require.NoError(t, err, "US expectedTargets %v", expectedTargets)
	err = instanceEU.WaitForLocalDNSEndpointHasTargets(expectedTargets)
	require.NoError(t, err, "EU expectedTargets %v", expectedTargets)

	for _, instance := range []*utils.Instance{instanceEU, instanceUS} {
		ep, err := instance.Resources().GetLocalDNSEndpoint().GetEndpointByName(host)
		require.NoError(t, err, "missing endpoint", host)
		// check all labels are correct
		require.Equal(t, "roundRobin", ep.Labels["strategy"])
		require.NotEqual(t, ep.Labels["weight-eu-0-5"], ep.Labels["weight-eu-1-5"])
		require.NotEqual(t, ep.Labels["weight-us-0-5"], ep.Labels["weight-us-1-5"])
		// check all targets are correct
		for _, v := range epExternalEU.Targets {
			require.True(t, Contains(v, ep.Targets))
		}
		for _, v := range epExternalEU.Targets {
			require.True(t, Contains(v, ep.Targets))
		}
	}
}

func Contains(str string, values []string) bool {
	for _, v := range values {
		if str == v {
			return true
		}
	}
	return false
}
