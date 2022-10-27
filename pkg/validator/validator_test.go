package validator

import (
	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"jobsearch-demos/kanban/pkg/randomize"
	"testing"
)

type testStruct struct {
	Uid                uuid.UUID `validate:"required"`
	Str                string    `validate:"required" default:"randomString"`
	Integer64          int64     `validate:"required" default:"64"`
	Integer32          int32     `validate:"required" default:"32"`
	Integer8           int8      `validate:"required" default:"8"`
	FloatingPointNum32 float32   `validate:"required" default:"32.32"`
	FloatingPointNum64 float64   `validate:"required" default:"64.64"`
	inexportedField    string    `default:"not_set"`
}

// TestWithDefaultsValidator_ValidateWithDefaults tests if the omitted struct elements gained their default values
func TestWithDefaultsValidator_ValidateWithDefaults(t *testing.T) {
	testStructWithDefaults := testStruct{
		Uid:       randomize.RandomUUID4(),
		Str:       randomize.RandomString(10),
		Integer8:  2,
		Integer32: 3,
		//integer64: omit, so that default value is set
		FloatingPointNum32: 32.32,
		//floatingPointNum64: omit, so that default value is set
	}

	validatorWithDefaults := NewWithDefaultsValidator()
	err := validatorWithDefaults.ValidateWithDefaults(&testStructWithDefaults)

	require.NoError(t, err)
	require.EqualValues(t, 64, testStructWithDefaults.Integer64)
	require.EqualValues(t, 64.64, testStructWithDefaults.FloatingPointNum64)

	// check that not exported field did not get default
	require.EqualValues(t, "", testStructWithDefaults.inexportedField)
	require.NotEqual(t, "not_set", testStructWithDefaults.inexportedField)
}
