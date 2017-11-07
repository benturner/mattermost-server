// Code generated by mockery v1.0.0

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/mattermost/mattermost-server/model"
import store "github.com/mattermost/mattermost-server/store"

// StatusStore is an autogenerated mock type for the StatusStore type
type StatusStore struct {
	mock.Mock
}

// Get provides a mock function with given fields: userId
func (_m *StatusStore) Get(userId string) store.StoreChannel {
	ret := _m.Called(userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAllFromTeam provides a mock function with given fields: teamId
func (_m *StatusStore) GetAllFromTeam(teamId string) store.StoreChannel {
	ret := _m.Called(teamId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(teamId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetByIds provides a mock function with given fields: userIds
func (_m *StatusStore) GetByIds(userIds []string) store.StoreChannel {
	ret := _m.Called(userIds)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func([]string) store.StoreChannel); ok {
		r0 = rf(userIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetOnline provides a mock function with given fields:
func (_m *StatusStore) GetOnline() store.StoreChannel {
	ret := _m.Called()

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func() store.StoreChannel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetOnlineAway provides a mock function with given fields:
func (_m *StatusStore) GetOnlineAway() store.StoreChannel {
	ret := _m.Called()

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func() store.StoreChannel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetTotalActiveUsersCount provides a mock function with given fields:
func (_m *StatusStore) GetTotalActiveUsersCount() store.StoreChannel {
	ret := _m.Called()

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func() store.StoreChannel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// ResetAll provides a mock function with given fields:
func (_m *StatusStore) ResetAll() store.StoreChannel {
	ret := _m.Called()

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func() store.StoreChannel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// SaveOrUpdate provides a mock function with given fields: status
func (_m *StatusStore) SaveOrUpdate(status *model.Status) store.StoreChannel {
	ret := _m.Called(status)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Status) store.StoreChannel); ok {
		r0 = rf(status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// UpdateLastActivityAt provides a mock function with given fields: userId, lastActivityAt
func (_m *StatusStore) UpdateLastActivityAt(userId string, lastActivityAt int64) store.StoreChannel {
	ret := _m.Called(userId, lastActivityAt)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int64) store.StoreChannel); ok {
		r0 = rf(userId, lastActivityAt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}