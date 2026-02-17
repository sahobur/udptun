package stun

import (
	"github.com/pion/stun"
)

func GetIPPost() (stun.XORMappedAddress, error) {
	// Parse a STUN URI
	var xorAddr stun.XORMappedAddress

	u, err := stun.ParseURI("stun:stun.l.google.com:19302")
	if err != nil {
		return xorAddr, err
	}

	// Creating a "connection" to STUN server.
	c, err := stun.DialURI(u, &stun.DialConfig{})
	if err != nil {
		return xorAddr, err
	}
	// Building binding request with random transaction id.
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)
	// Sending request to STUN server, waiting for response message.
	if err := c.Do(message, func(res stun.Event) {
		if res.Error != nil {
			panic(res.Error)
		}
		// Decoding XOR-MAPPED-ADDRESS attribute from message.
		if err := xorAddr.GetFrom(res.Message); err != nil {
			panic(err)
		}
	}); err != nil {
		return xorAddr, err
	}
	return xorAddr, nil
}
