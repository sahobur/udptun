package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig_Load(t *testing.T) {
	config := &Config{}
	err := config.Load()
	require.NoError(t, err)
	fmt.Printf("%v", config)
}
