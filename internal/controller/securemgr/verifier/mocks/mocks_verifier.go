/*******************************************************************************
 * Copyright 2020 Samsung Electronics All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *******************************************************************************/

// Code generated by MockGen. DO NOT EDIT.
// Source: controller/securemgr/verifier (interfaces: Conf)

// Package mocks is a generated GoMock package.
package mocks

import (
	verifier "github.com/lf-edge/edge-home-orchestration-go/internal/controller/securemgr/verifier"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockVerifierConf is a mock of Conf interface.
type MockVerifierConf struct {
	ctrl     *gomock.Controller
	recorder *MockVerifierConfMockRecorder
}

// MockVerifierConfMockRecorder is the mock recorder for MockVerifierConf.
type MockVerifierConfMockRecorder struct {
	mock *MockVerifierConf
}

// NewMockVerifierConf creates a new mock instance.
func NewMockVerifierConf(ctrl *gomock.Controller) *MockVerifierConf {
	mock := &MockVerifierConf{ctrl: ctrl}
	mock.recorder = &MockVerifierConfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVerifierConf) EXPECT() *MockVerifierConfMockRecorder {
	return m.recorder
}

// RequestVerifierConf mocks base method.
func (m *MockVerifierConf) RequestVerifierConf(arg0 verifier.RequestVerifierConf) verifier.ResponseVerifierConf {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestVerifierConf", arg0)
	ret0, _ := ret[0].(verifier.ResponseVerifierConf)
	return ret0
}

// RequestVerifierConf indicates an expected call of RequestVerifierConf.
func (mr *MockVerifierConfMockRecorder) RequestVerifierConf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestVerifierConf", reflect.TypeOf((*MockVerifierConf)(nil).RequestVerifierConf), arg0)
}
