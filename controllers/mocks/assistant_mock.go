// Code generated by MockGen. DO NOT EDIT.
// Source: controllers/providers/assistant/assistant.go
//
// Generated by this command:
//
//	mockgen -package=mocks -destination=controllers/mocks/assistant_mock.go -source=controllers/providers/assistant/assistant.go Assistant
//

// Package mocks is a generated GoMock package.
package mocks

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
	reflect "reflect"
	time "time"

	assistant "github.com/k8gb-io/k8gb/controllers/providers/assistant"
	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
	endpoint "sigs.k8s.io/external-dns/endpoint"
)

// MockAssistant is a mock of Assistant interface.
type MockAssistant struct {
	ctrl     *gomock.Controller
	recorder *MockAssistantMockRecorder
}

// MockAssistantMockRecorder is the mock recorder for MockAssistant.
type MockAssistantMockRecorder struct {
	mock *MockAssistant
}

// NewMockAssistant creates a new mock instance.
func NewMockAssistant(ctrl *gomock.Controller) *MockAssistant {
	mock := &MockAssistant{ctrl: ctrl}
	mock.recorder = &MockAssistantMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAssistant) EXPECT() *MockAssistantMockRecorder {
	return m.recorder
}

// CoreDNSExposedIPs mocks base method.
func (m *MockAssistant) CoreDNSExposedIPs() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CoreDNSExposedIPs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CoreDNSExposedIPs indicates an expected call of CoreDNSExposedIPs.
func (mr *MockAssistantMockRecorder) CoreDNSExposedIPs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CoreDNSExposedIPs", reflect.TypeOf((*MockAssistant)(nil).CoreDNSExposedIPs))
}

// GetCoreDNSService mocks base method.
func (m *MockAssistant) GetCoreDNSService() (*v1.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCoreDNSService")
	ret0, _ := ret[0].(*v1.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCoreDNSService indicates an expected call of GetCoreDNSService.
func (mr *MockAssistantMockRecorder) GetCoreDNSService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCoreDNSService", reflect.TypeOf((*MockAssistant)(nil).GetCoreDNSService))
}

// GetExternalTargets mocks base method.
func (m *MockAssistant) GetExternalTargets(host string, extClusterNsNames map[string]string) assistant.Targets {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExternalTargets", host, extClusterNsNames)
	ret0, _ := ret[0].(assistant.Targets)
	return ret0
}

// GetExternalTargets indicates an expected call of GetExternalTargets.
func (mr *MockAssistantMockRecorder) GetExternalTargets(host, extClusterNsNames any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExternalTargets", reflect.TypeOf((*MockAssistant)(nil).GetExternalTargets), host, extClusterNsNames)
}

// InspectTXTThreshold mocks base method.
func (m *MockAssistant) InspectTXTThreshold(fqdn string, splitBrainThreshold time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectTXTThreshold", fqdn, splitBrainThreshold)
	ret0, _ := ret[0].(error)
	return ret0
}

// InspectTXTThreshold indicates an expected call of InspectTXTThreshold.
func (mr *MockAssistantMockRecorder) InspectTXTThreshold(fqdn, splitBrainThreshold any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectTXTThreshold", reflect.TypeOf((*MockAssistant)(nil).InspectTXTThreshold), fqdn, splitBrainThreshold)
}

// RemoveEndpoint mocks base method.
func (m *MockAssistant) RemoveEndpoint(endpointName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveEndpoint", endpointName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveEndpoint indicates an expected call of RemoveEndpoint.
func (mr *MockAssistantMockRecorder) RemoveEndpoint(endpointName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveEndpoint", reflect.TypeOf((*MockAssistant)(nil).RemoveEndpoint), endpointName)
}

// SaveDNSEndpoint mocks base method.
func (m *MockAssistant) SaveDNSEndpoint(namespace string, i *endpoint.DNSEndpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveDNSEndpoint", namespace, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveDNSEndpoint indicates an expected call of SaveDNSEndpoint.
func (mr *MockAssistantMockRecorder) SaveDNSEndpoint(namespace, i any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveDNSEndpoint", reflect.TypeOf((*MockAssistant)(nil).SaveDNSEndpoint), namespace, i)
}
