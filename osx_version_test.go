package osxversion

import "testing"

func TestOSXVersion(t *testing.T)  {
	osx := new(OSXVersion)
	osx.Query()

	expect := "Mac OS X"
	if osx.Name == expect {
		t.Errorf("Expect '%s', but got '%s'", expect, osx.Name)
	}
}