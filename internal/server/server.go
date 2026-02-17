package server

func startServer(localPort string, localTunAddr string) {
	conn, err := serverUDP(localPort)
	if err != nil {
		panic(err)
	}
	c := 0
	ifce, err := createTunInterface(localTunAddr)
	if err != nil {
		panic(err)
	}	
}