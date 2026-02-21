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
	client := Client{
		Name: "test-client-2",
		Key:  "key-client-2",
		IP:   "10.10.10.22/24",
	}
	err = config.AddClient(client)
	require.NoError(t, err)

}
