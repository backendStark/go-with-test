package fuzzing

import (
	"fmt"
	"testing"
)

func Test_ParseUserInput_Success(t *testing.T) {
	testCases := []struct {
		input  string
		isWork bool
	}{
		{"dmitrii:30", true},
		{"raul :31", true},
		{"nikita", false},
		{"tanya:30:somethingnew", false},
		{"lev:abv", false},
		{"volodya:-45", false},
		{"kostya:0", false},
		{":100", false},
		{":49", false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Test for string %q", tc.input), func(t *testing.T) {
			res, err := ParseUserInput(tc.input)
			if tc.isWork && err != nil {
				t.Fatalf("should work, but got err %v", err)
			}

			if tc.isWork == false && err == nil {
				t.Fatalf("should return error, but got result %v", res)
			}
		})
	}
}
