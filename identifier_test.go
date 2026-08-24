package utils

import (
	"testing"

	"github.com/fossoreslp/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewId(t *testing.T) {
	t.Parallel()

	id := NewId()
	assert.Len(t, id, 36)

	_, err := uuid.Parse(id)
	require.NoError(t, err)
}

func TestNewV7Id(t *testing.T) {
	t.Parallel()

	id := NewV7Id()
	assert.Len(t, id, 36)

	_, err := uuid.Parse(id)
	require.NoError(t, err)
}

func TestGeneratedIdsAreUnique(t *testing.T) {
	t.Parallel()

	assert.NotEqual(t, NewId(), NewId())
	assert.NotEqual(t, NewV7Id(), NewV7Id())
}

func TestIsId(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{"v4", uuid.NewV4().String(), true},
		{"v7", uuid.NewV7().String(), true},
		{"garbage", "xxx", false},
		{"empty", "", false},
		{"short", "12345678-1234-1234-1234-12345678", false},
		{"malformed-hex", "zzzzzzzz-zzzz-zzzz-zzzz-zzzzzzzzzzzz", false},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, IsId(tc.input))
		})
	}
}
