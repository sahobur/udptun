package server

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateKey(t *testing.T) {
	require.True(t, isAlphanumericRegexp("ClientKey001"))
}
