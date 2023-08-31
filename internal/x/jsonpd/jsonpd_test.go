package jsonpd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonPD(t *testing.T) {
	type testCase struct {
		input    string
		expected string
	}

	testCases := []testCase{
		{
			`{"$and":[{"field1":{"$eq":"value"}},{"field2":{"$eq":"value"}}]}`,
			`column->'field1' = '"value"'::jsonb AND column->'field2' = '"value"'::jsonb`,
		},
		{
			`{"$or":[{"field1":{"$eq":"value"}},{"$not":{"field2":{"$eq":"value"}}}]}`,
			`column->'field1' = '"value"'::jsonb OR NOT (column->'field2' = '"value"'::jsonb)`,
		},
		{
			`{"$and":[{"arrayfield":{"$contains":"value"}}]}`,
			`column->'arrayfield' @> '"value"'::jsonb`,
		},
	}

	for _, tc := range testCases {
		exec(t, tc.input, tc.expected)
	}
}

func exec(t *testing.T, input string, expected string) {
	j, err := Parse([]byte(input))
	assert.NoError(t, err)

	sql, err := j.SQL("column")
	assert.NoError(t, err)

	assert.Equal(t, expected, sql)
}
