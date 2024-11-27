package testhelper

type T interface {
	Errorf(format string, args ...interface{})
	Cleanup(f func())
}
