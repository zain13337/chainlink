// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	log "github.com/smartcontractkit/chainlink/core/services/log"
	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/core/store/models"
)

// Listener is an autogenerated mock type for the Listener type
type Listener struct {
	mock.Mock
}

// HandleLog provides a mock function with given fields: b
func (_m *Listener) HandleLog(b log.Broadcast) {
	_m.Called(b)
}

// IsV2Job provides a mock function with given fields:
func (_m *Listener) IsV2Job() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// JobID provides a mock function with given fields:
func (_m *Listener) JobID() models.JobID {
	ret := _m.Called()

	var r0 models.JobID
	if rf, ok := ret.Get(0).(func() models.JobID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.JobID)
		}
	}

	return r0
}

// JobIDV2 provides a mock function with given fields:
func (_m *Listener) JobIDV2() int32 {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}
