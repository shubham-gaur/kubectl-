package helper

import (
	"testing"
)

func TestExecKubectlCmd(t *testing.T) {
	expectedOp := "Test Successful"
	args := []string{"-n", "Test", "Successful"}
	op := execute("echo", args...)
	if expectedOp != op {
		t.Fatalf("Output: [%v] doesnot match with Expected output: [%v]", op, expectedOp)
	}
}
