package test

import "testing"

func TestMain(m *testing.M) {
	m.Run()

	T.svr.Close()
}