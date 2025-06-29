package ec2test

import "testing"

func TestFunction(t *testing.T) {
	t.Run("Should pass if value=value", func(t *testing.T) {
		resp := Function(3)
		if resp != 3 {
			t.Fail()
		}
	})
}
