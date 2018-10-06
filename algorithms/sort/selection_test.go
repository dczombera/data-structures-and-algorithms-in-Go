package sort

import "testing"

type testCase struct {
	data []MyInt
	want []MyInt
}

var testCases = []testCase{
	{[]MyInt{5, 4, 3, 2, 1}, []MyInt{1, 2, 3, 4, 5}},
	{[]MyInt{42, -42}, []MyInt{-42, 42}},
	{[]MyInt{3, 5, 3, 5, -5, -3, 0, -3}, []MyInt{-5, -3, -3, 0, 3, 3, 5, 5}},
	{[]MyInt{-1, 0, -1}, []MyInt{-1, -1, 0}},
	{[]MyInt{}, []MyInt{}},
}

func TestSelection(t *testing.T) {
	for _, tc := range testCases {
		Selection(tc.data)
		err := false
		for i := 0; i < len(tc.data); i++ {
			if tc.data[i] != tc.want[i] {
				err = true
			}
		}
		if err {
			t.Errorf("Got %v, want %v", tc.data, tc.want)
		}
	}
}
