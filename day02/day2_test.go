package main

import "testing"

func TestReportIsSafe(t *testing.T) {
	testInput := []struct {
		name     string
		report   []int
		expected bool
	}{
		{
			name:     "Safe without removing any level.",
			report:   []int{7, 6, 4, 2, 1},
			expected: true,
		},
		{
			name:     "Unsafe regardless of which level is removed.",
			report:   []int{1, 2, 7, 8, 9},
			expected: false,
		},
		{
			name:     "Unsafe regardless of which level is removed.",
			report:   []int{9, 7, 6, 2, 1},
			expected: false,
		},
		{
			name:     "Safe by removing the second level, 3.",
			report:   []int{1, 3, 2, 4, 5},
			expected: true,
		},
		{
			name:     "Safe by removing the third level, 4.",
			report:   []int{8, 6, 4, 4, 1},
			expected: true,
		},
		{
			name:     "Safe without removing any level.",
			report:   []int{1, 3, 6, 7, 9},
			expected: true,
		},
		{
			name:     "No changes - unsafe (must be all increasing or all decreasing)",
			report:   []int{1, 1, 1},
			expected: false,
		},
		{
			name:     "One value is always safe",
			report:   []int{1},
			expected: true,
		},
		{
			name:     "Slowly ascending - safe",
			report:   []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "Slowly descending - safe",
			report:   []int{3, 2, 1},
			expected: true,
		},
		{
			name:     "Mixed ascending,descending - unsafe",
			report:   []int{1, 2, 3, 2, 1},
			expected: false,
		},
		{
			name:     "Rapid ascending (+4) - unsafe",
			report:   []int{1, 2, 6, 7},
			expected: false,
		},
		{
			name:     "Almost rapid ascending (+3) - safe",
			report:   []int{1, 2, 5, 6},
			expected: true,
		},
		{
			name:     "Rapid descending (-4) - unsafe",
			report:   []int{7, 6, 2, 1},
			expected: false,
		},
		{
			name:     "Almost rapid descending (-3) - safe",
			report:   []int{5, 2, 1},
			expected: true,
		},
		{
			name:     "Last element is bad - safe",
			report:   []int{1, 2, 3, 4, 5, 6, 1},
			expected: true,
		},
		{
			name:     "First element is bad - safe",
			report:   []int{10, 2, 3, 4, 5, 6},
			expected: true,
		},

		{
			name:     "1 2 3",
			report:   []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "1 2 4",
			report:   []int{1, 2, 4},
			expected: true,
		},
		{
			name:     "1 2 5",
			report:   []int{1, 2, 5},
			expected: true,
		},
		{
			name:     "1 2 6",
			report:   []int{1, 2, 6},
			expected: true,
		},
		{
			name:     "1 1 6",
			report:   []int{1, 1, 6},
			expected: false,
		},
		{
			name:     "10 1 11",
			report:   []int{10, 1, 11},
			expected: false,
		},
		{
			name:     "10 1 11",
			report:   []int{10, 1, 11},
			expected: false,
		},
		{
			name:     "65 66 63 60",
			report:   []int{65, 66, 63, 60},
			expected: true,
		},
	}

	for _, test := range testInput {
		t.Run(test.name, func(t *testing.T) {
			safe := CheckReport(test.report)
			if safe != test.expected {
				t.Fail()
			}
		})
	}
}
