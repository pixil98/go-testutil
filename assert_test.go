package testutil

import (
	"errors"
	"testing"
)

type mockT struct {
	testing.TB
	errors []string
}

func (m *mockT) Errorf(format string, args ...interface{}) {
	m.errors = append(m.errors, format)
}

func TestAssertEqual(t *testing.T) {
	type person struct {
		Name string
		Age  int
	}

	tests := map[string]struct {
		actual      interface{}
		expected    interface{}
		shouldError bool
	}{
		"equal integers": {
			actual:      42,
			expected:    42,
			shouldError: false,
		},
		"unequal integers": {
			actual:      42,
			expected:    43,
			shouldError: true,
		},
		"equal strings": {
			actual:      "hello",
			expected:    "hello",
			shouldError: false,
		},
		"unequal strings": {
			actual:      "hello",
			expected:    "world",
			shouldError: true,
		},
		"equal structs": {
			actual:      person{"Alice", 30},
			expected:    person{"Alice", 30},
			shouldError: false,
		},
		"unequal structs": {
			actual:      person{"Alice", 30},
			expected:    person{"Bob", 30},
			shouldError: true,
		},
		"equal slices": {
			actual:      []int{1, 2, 3},
			expected:    []int{1, 2, 3},
			shouldError: false,
		},
		"unequal slices": {
			actual:      []int{1, 2, 3},
			expected:    []int{1, 2, 4},
			shouldError: true,
		},
		"equal nil values": {
			actual:      nil,
			expected:    nil,
			shouldError: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			mock := &mockT{TB: t}
			AssertEqual(mock, "test value", tt.actual, tt.expected)
			if tt.shouldError && len(mock.errors) == 0 {
				t.Error("expected an error but got none")
			}
			if !tt.shouldError && len(mock.errors) != 0 {
				t.Errorf("expected no errors but got %d", len(mock.errors))
			}
		})
	}
}

func TestAssertErrorContains(t *testing.T) {
	tests := map[string]struct {
		err         error
		pieces      []string
		shouldError bool
	}{
		"single match": {
			err:         errors.New("connection failed: timeout"),
			pieces:      []string{"timeout"},
			shouldError: false,
		},
		"multiple matches": {
			err:         errors.New("connection failed: timeout"),
			pieces:      []string{"connection", "timeout"},
			shouldError: false,
		},
		"no match": {
			err:         errors.New("connection failed"),
			pieces:      []string{"timeout"},
			shouldError: true,
		},
		"partial match fails": {
			err:         errors.New("failed to connect to database: access denied"),
			pieces:      []string{"connect", "timeout", "denied"},
			shouldError: true,
		},
		"nil error with empty string": {
			err:         nil,
			pieces:      []string{""},
			shouldError: false,
		},
		"nil error with non-empty string": {
			err:         nil,
			pieces:      []string{"expected error"},
			shouldError: true,
		},
		"empty pieces": {
			err:         errors.New("some error"),
			pieces:      []string{},
			shouldError: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			mock := &mockT{TB: t}
			AssertErrorContains(mock, tt.err, tt.pieces...)
			if tt.shouldError && len(mock.errors) == 0 {
				t.Error("expected an error but got none")
			}
			if !tt.shouldError && len(mock.errors) != 0 {
				t.Errorf("expected no errors but got %d", len(mock.errors))
			}
		})
	}
}
