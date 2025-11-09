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

func FuzzParseUserInput(f *testing.F) {
	f.Add("dmitrii:30")
	f.Add("raul :31")
	f.Add("nikita")
	f.Add("tanya:30:somethingnew")
	f.Add("lev:abv")
	f.Add("volodya:-45")
	f.Add("kostya:0")
	f.Add(":100")
	f.Add(":49")

	f.Fuzz(func(t *testing.T, input string) {
		res, err := ParseUserInput(input)

		if err == nil && res == nil {
			t.Errorf("got nil result with no error for input %q", input)
		}

		if res != nil && res.Username == "" {
			t.Errorf("got empty username in valid result for input %q", res)
		}
	})
}
