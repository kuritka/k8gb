package ingress

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

	k8gbv1beta1 "github.com/k8gb-io/k8gb/api/v1beta1"
	"github.com/k8gb-io/k8gb/controllers/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetServers(t *testing.T) {
	var tests = []struct {
		name            string
		ingressFile     string
		expectedServers []*k8gbv1beta1.Server
	}{
		{
			name:        "single server",
			ingressFile: "../testdata/ingress_referenced.yaml",
			expectedServers: []*k8gbv1beta1.Server{
				{
					Host: "ingress-referenced.cloud.example.com",
					Services: []*k8gbv1beta1.NamespacedName{
						{
							Name:      "ingress-referenced",
							Namespace: "test-gslb",
						},
					},
				},
			},
		},
		{
			name:        "multiple servers",
			ingressFile: "./testdata/ingress_multiple_servers.yaml",
			expectedServers: []*k8gbv1beta1.Server{
				{
					Host: "h1.cloud.example.com",
					Services: []*k8gbv1beta1.NamespacedName{
						{
							Name:      "s1",
							Namespace: "test-gslb",
						},
					},
				},
				{
					Host: "h2.cloud.example.com",
					Services: []*k8gbv1beta1.NamespacedName{
						{
							Name:      "ss1",
							Namespace: "test-gslb",
						},
						{
							Name:      "ss2",
							Namespace: "test-gslb",
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// arrange
			ingress := utils.FileToIngress(test.ingressFile)
			resolver := ReferenceResolver{
				ingress: ingress,
			}

			// act
			servers, err := resolver.GetServers()
			assert.NoError(t, err)

			// assert
			assert.Equal(t, test.expectedServers, servers)
		})
	}
}

func TestGetGslbExposedIPs(t *testing.T) {
	var tests = []struct {
		name          string
		annotations   map[string]string
		ingressYaml   string
		expectedIPs   []string
		expectedError bool
	}{
		{
			name:          "no exposed IPs",
			annotations:   map[string]string{},
			ingressYaml:   "./testdata/ingress_no_ips.yaml",
			expectedIPs:   []string{},
			expectedError: false,
		},
		{
			name:          "single exposed IP",
			annotations:   map[string]string{},
			ingressYaml:   "../testdata/ingress_referenced.yaml",
			expectedIPs:   []string{"10.0.0.1"},
			expectedError: false,
		},
		{
			name:          "multiple exposed IPs",
			annotations:   map[string]string{},
			ingressYaml:   "./testdata/ingress_multiple_ips.yaml",
			expectedIPs:   []string{"10.0.0.1", "10.0.0.2"},
			expectedError: false,
		},
		{
			name:          "annotation with no exposed IPs",
			annotations:   map[string]string{"k8gb.io/exposed-ip-addresses": ""},
			ingressYaml:   "./testdata/ingress_multiple_ips.yaml",
			expectedIPs:   []string{},
			expectedError: true,
		},
		{
			name:          "annotation with single exposed IP",
			annotations:   map[string]string{"k8gb.io/exposed-ip-addresses": "185.199.110.153"},
			ingressYaml:   "./testdata/ingress_multiple_ips.yaml",
			expectedIPs:   []string{"185.199.110.153"},
			expectedError: false,
		},
		{
			name:          "annotation with multiple exposed IPs",
			annotations:   map[string]string{"k8gb.io/exposed-ip-addresses": "185.199.110.153,185.199.109.153"},
			ingressYaml:   "./testdata/ingress_multiple_ips.yaml",
			expectedIPs:   []string{"185.199.110.153", "185.199.109.153"},
			expectedError: false,
		},
		{
			name:          "annotation with invalid IP",
			annotations:   map[string]string{"k8gb.io/exposed-ip-addresses": "192.169.0.test"},
			ingressYaml:   "./testdata/ingress_multiple_ips.yaml",
			expectedIPs:   []string{},
			expectedError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			ingress := utils.FileToIngress(tt.ingressYaml)
			resolver := ReferenceResolver{
				ingress: ingress,
			}

			// act
			IPs, err := resolver.GetGslbExposedIPs(tt.annotations, []utils.DNSServer{})
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// assert
			assert.Equal(t, tt.expectedIPs, IPs)
		})
	}
}
