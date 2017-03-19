package provider

import io "io"
import mock "github.com/stretchr/testify/mock"
import types "github.com/convox/praxis/types"

// MockProvider is an autogenerated mock type for the Provider type
type MockProvider struct {
	mock.Mock
}

// AppCreate provides a mock function with given fields: name
func (_m *MockProvider) AppCreate(name string) (*types.App, error) {
	ret := _m.Called(name)

	var r0 *types.App
	if rf, ok := ret.Get(0).(func(string) *types.App); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.App)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppDelete provides a mock function with given fields: name
func (_m *MockProvider) AppDelete(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppGet provides a mock function with given fields: name
func (_m *MockProvider) AppGet(name string) (*types.App, error) {
	ret := _m.Called(name)

	var r0 *types.App
	if rf, ok := ret.Get(0).(func(string) *types.App); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.App)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppList provides a mock function with given fields:
func (_m *MockProvider) AppList() (types.Apps, error) {
	ret := _m.Called()

	var r0 types.Apps
	if rf, ok := ret.Get(0).(func() types.Apps); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Apps)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildCreate provides a mock function with given fields: app, url, opts
func (_m *MockProvider) BuildCreate(app string, url string, opts types.BuildCreateOptions) (*types.Build, error) {
	ret := _m.Called(app, url, opts)

	var r0 *types.Build
	if rf, ok := ret.Get(0).(func(string, string, types.BuildCreateOptions) *types.Build); ok {
		r0 = rf(app, url, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, types.BuildCreateOptions) error); ok {
		r1 = rf(app, url, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildGet provides a mock function with given fields: app, id
func (_m *MockProvider) BuildGet(app string, id string) (*types.Build, error) {
	ret := _m.Called(app, id)

	var r0 *types.Build
	if rf, ok := ret.Get(0).(func(string, string) *types.Build); ok {
		r0 = rf(app, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildList provides a mock function with given fields: app
func (_m *MockProvider) BuildList(app string) (types.Builds, error) {
	ret := _m.Called(app)

	var r0 types.Builds
	if rf, ok := ret.Get(0).(func(string) types.Builds); ok {
		r0 = rf(app)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Builds)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(app)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildLogs provides a mock function with given fields: app, id
func (_m *MockProvider) BuildLogs(app string, id string) (io.ReadCloser, error) {
	ret := _m.Called(app, id)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string, string) io.ReadCloser); ok {
		r0 = rf(app, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildUpdate provides a mock function with given fields: app, id, opts
func (_m *MockProvider) BuildUpdate(app string, id string, opts types.BuildUpdateOptions) (*types.Build, error) {
	ret := _m.Called(app, id, opts)

	var r0 *types.Build
	if rf, ok := ret.Get(0).(func(string, string, types.BuildUpdateOptions) *types.Build); ok {
		r0 = rf(app, id, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, types.BuildUpdateOptions) error); ok {
		r1 = rf(app, id, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilesDelete provides a mock function with given fields: app, pid, files
func (_m *MockProvider) FilesDelete(app string, pid string, files []string) error {
	ret := _m.Called(app, pid, files)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, []string) error); ok {
		r0 = rf(app, pid, files)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilesUpload provides a mock function with given fields: app, pid, r
func (_m *MockProvider) FilesUpload(app string, pid string, r io.Reader) error {
	ret := _m.Called(app, pid, r)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, io.Reader) error); ok {
		r0 = rf(app, pid, r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// KeyDecrypt provides a mock function with given fields: app, key, data
func (_m *MockProvider) KeyDecrypt(app string, key string, data []byte) ([]byte, error) {
	ret := _m.Called(app, key, data)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string, []byte) []byte); ok {
		r0 = rf(app, key, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, []byte) error); ok {
		r1 = rf(app, key, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KeyEncrypt provides a mock function with given fields: app, key, data
func (_m *MockProvider) KeyEncrypt(app string, key string, data []byte) ([]byte, error) {
	ret := _m.Called(app, key, data)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string, []byte) []byte); ok {
		r0 = rf(app, key, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, []byte) error); ok {
		r1 = rf(app, key, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ObjectFetch provides a mock function with given fields: app, key
func (_m *MockProvider) ObjectFetch(app string, key string) (io.ReadCloser, error) {
	ret := _m.Called(app, key)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string, string) io.ReadCloser); ok {
		r0 = rf(app, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ObjectStore provides a mock function with given fields: app, key, r, opts
func (_m *MockProvider) ObjectStore(app string, key string, r io.Reader, opts types.ObjectStoreOptions) (*types.Object, error) {
	ret := _m.Called(app, key, r, opts)

	var r0 *types.Object
	if rf, ok := ret.Get(0).(func(string, string, io.Reader, types.ObjectStoreOptions) *types.Object); ok {
		r0 = rf(app, key, r, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Object)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, io.Reader, types.ObjectStoreOptions) error); ok {
		r1 = rf(app, key, r, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessGet provides a mock function with given fields: app, pid
func (_m *MockProvider) ProcessGet(app string, pid string) (*types.Process, error) {
	ret := _m.Called(app, pid)

	var r0 *types.Process
	if rf, ok := ret.Get(0).(func(string, string) *types.Process); ok {
		r0 = rf(app, pid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Process)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, pid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessList provides a mock function with given fields: app, opts
func (_m *MockProvider) ProcessList(app string, opts types.ProcessListOptions) (types.Processes, error) {
	ret := _m.Called(app, opts)

	var r0 types.Processes
	if rf, ok := ret.Get(0).(func(string, types.ProcessListOptions) types.Processes); ok {
		r0 = rf(app, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Processes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, types.ProcessListOptions) error); ok {
		r1 = rf(app, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessLogs provides a mock function with given fields: app, pid
func (_m *MockProvider) ProcessLogs(app string, pid string) (io.ReadCloser, error) {
	ret := _m.Called(app, pid)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string, string) io.ReadCloser); ok {
		r0 = rf(app, pid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, pid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessRun provides a mock function with given fields: app, opts
func (_m *MockProvider) ProcessRun(app string, opts types.ProcessRunOptions) (int, error) {
	ret := _m.Called(app, opts)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, types.ProcessRunOptions) int); ok {
		r0 = rf(app, opts)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, types.ProcessRunOptions) error); ok {
		r1 = rf(app, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessStart provides a mock function with given fields: app, opts
func (_m *MockProvider) ProcessStart(app string, opts types.ProcessRunOptions) (string, error) {
	ret := _m.Called(app, opts)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, types.ProcessRunOptions) string); ok {
		r0 = rf(app, opts)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, types.ProcessRunOptions) error); ok {
		r1 = rf(app, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessStop provides a mock function with given fields: app, pid
func (_m *MockProvider) ProcessStop(app string, pid string) error {
	ret := _m.Called(app, pid)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(app, pid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Proxy provides a mock function with given fields: app, pid, port, in
func (_m *MockProvider) Proxy(app string, pid string, port int, in io.Reader) (io.ReadCloser, error) {
	ret := _m.Called(app, pid, port, in)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string, string, int, io.Reader) io.ReadCloser); ok {
		r0 = rf(app, pid, port, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, int, io.Reader) error); ok {
		r1 = rf(app, pid, port, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueueFetch provides a mock function with given fields: app, queue, opts
func (_m *MockProvider) QueueFetch(app string, queue string, opts types.QueueFetchOptions) (map[string]string, error) {
	ret := _m.Called(app, queue, opts)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string, string, types.QueueFetchOptions) map[string]string); ok {
		r0 = rf(app, queue, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, types.QueueFetchOptions) error); ok {
		r1 = rf(app, queue, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueueStore provides a mock function with given fields: app, queue, attrs
func (_m *MockProvider) QueueStore(app string, queue string, attrs map[string]string) error {
	ret := _m.Called(app, queue, attrs)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, map[string]string) error); ok {
		r0 = rf(app, queue, attrs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegistryAdd provides a mock function with given fields: server, username, password
func (_m *MockProvider) RegistryAdd(server string, username string, password string) (*types.Registry, error) {
	ret := _m.Called(server, username, password)

	var r0 *types.Registry
	if rf, ok := ret.Get(0).(func(string, string, string) *types.Registry); ok {
		r0 = rf(server, username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Registry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(server, username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegistryDelete provides a mock function with given fields: server
func (_m *MockProvider) RegistryDelete(server string) error {
	ret := _m.Called(server)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(server)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegistryList provides a mock function with given fields:
func (_m *MockProvider) RegistryList() (types.Registries, error) {
	ret := _m.Called()

	var r0 types.Registries
	if rf, ok := ret.Get(0).(func() types.Registries); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Registries)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReleaseCreate provides a mock function with given fields: app, opts
func (_m *MockProvider) ReleaseCreate(app string, opts types.ReleaseCreateOptions) (*types.Release, error) {
	ret := _m.Called(app, opts)

	var r0 *types.Release
	if rf, ok := ret.Get(0).(func(string, types.ReleaseCreateOptions) *types.Release); ok {
		r0 = rf(app, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Release)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, types.ReleaseCreateOptions) error); ok {
		r1 = rf(app, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReleaseGet provides a mock function with given fields: app, id
func (_m *MockProvider) ReleaseGet(app string, id string) (*types.Release, error) {
	ret := _m.Called(app, id)

	var r0 *types.Release
	if rf, ok := ret.Get(0).(func(string, string) *types.Release); ok {
		r0 = rf(app, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Release)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReleaseList provides a mock function with given fields: app
func (_m *MockProvider) ReleaseList(app string) (types.Releases, error) {
	ret := _m.Called(app)

	var r0 types.Releases
	if rf, ok := ret.Get(0).(func(string) types.Releases); ok {
		r0 = rf(app)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Releases)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(app)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReleasePromote provides a mock function with given fields: app, id
func (_m *MockProvider) ReleasePromote(app string, id string) error {
	ret := _m.Called(app, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(app, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SystemGet provides a mock function with given fields:
func (_m *MockProvider) SystemGet() (*types.System, error) {
	ret := _m.Called()

	var r0 *types.System
	if rf, ok := ret.Get(0).(func() *types.System); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.System)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TableCreate provides a mock function with given fields: app, name, opts
func (_m *MockProvider) TableCreate(app string, name string, opts types.TableCreateOptions) error {
	ret := _m.Called(app, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, types.TableCreateOptions) error); ok {
		r0 = rf(app, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TableGet provides a mock function with given fields: app, table
func (_m *MockProvider) TableGet(app string, table string) (*types.Table, error) {
	ret := _m.Called(app, table)

	var r0 *types.Table
	if rf, ok := ret.Get(0).(func(string, string) *types.Table); ok {
		r0 = rf(app, table)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Table)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, table)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TableList provides a mock function with given fields: app
func (_m *MockProvider) TableList(app string) (types.Tables, error) {
	ret := _m.Called(app)

	var r0 types.Tables
	if rf, ok := ret.Get(0).(func(string) types.Tables); ok {
		r0 = rf(app)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Tables)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(app)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TableRowDelete provides a mock function with given fields: app, table, key, opts
func (_m *MockProvider) TableRowDelete(app string, table string, key string, opts types.TableRowDeleteOptions) error {
	ret := _m.Called(app, table, key, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, types.TableRowDeleteOptions) error); ok {
		r0 = rf(app, table, key, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TableRowGet provides a mock function with given fields: app, table, key, opts
func (_m *MockProvider) TableRowGet(app string, table string, key string, opts types.TableRowGetOptions) (*types.TableRow, error) {
	ret := _m.Called(app, table, key, opts)

	var r0 *types.TableRow
	if rf, ok := ret.Get(0).(func(string, string, string, types.TableRowGetOptions) *types.TableRow); ok {
		r0 = rf(app, table, key, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.TableRow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, types.TableRowGetOptions) error); ok {
		r1 = rf(app, table, key, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TableRowStore provides a mock function with given fields: app, table, attrs
func (_m *MockProvider) TableRowStore(app string, table string, attrs types.TableRow) (string, error) {
	ret := _m.Called(app, table, attrs)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, types.TableRow) string); ok {
		r0 = rf(app, table, attrs)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, types.TableRow) error); ok {
		r1 = rf(app, table, attrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TableRowsDelete provides a mock function with given fields: app, table, key, opts
func (_m *MockProvider) TableRowsDelete(app string, table string, key []string, opts types.TableRowDeleteOptions) error {
	ret := _m.Called(app, table, key, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, []string, types.TableRowDeleteOptions) error); ok {
		r0 = rf(app, table, key, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TableRowsGet provides a mock function with given fields: app, table, key, opts
func (_m *MockProvider) TableRowsGet(app string, table string, key []string, opts types.TableRowGetOptions) (types.TableRows, error) {
	ret := _m.Called(app, table, key, opts)

	var r0 types.TableRows
	if rf, ok := ret.Get(0).(func(string, string, []string, types.TableRowGetOptions) types.TableRows); ok {
		r0 = rf(app, table, key, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.TableRows)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, []string, types.TableRowGetOptions) error); ok {
		r1 = rf(app, table, key, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TableTruncate provides a mock function with given fields: app, table
func (_m *MockProvider) TableTruncate(app string, table string) error {
	ret := _m.Called(app, table)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(app, table)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
