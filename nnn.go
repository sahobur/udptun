// package main

// import (
// 	"fmt"
// 	"log"
// 	"net"
// 	"time"

// 	"github.com/ccding/go-stun/stun"
// )

// // Получаем публичный IP:порт через STUN
// func getPublicAddr(stunServer string) (*stun.MappedAddress, error) {
// 	conn, err := net.Dial("udp", stunServer)
// 	if err != nil {
// 		return nil, fmt.Errorf("не удалось подключиться к STUN-серверу: %w", err)
// 	}
// 	defer conn.Close()

// 	request := stun.NewBindingRequest()
// 	_, err = request.WriteTo(conn)
// 	if err != nil {
// 		return nil, fmt.Errorf("ошибка отправки STUN-запроса: %w", err)
// 	}

// 	response := stun.NewResponse()
// 	err = response.ReadFrom(conn)
// 	if err != nil {
// 		return nil, fmt.Errorf("ошибка получения STUN-ответа: %w", err)
// 	}

// 	mappedAddr := response.GetMappedAddress()
// 	return mappedAddr, nil
// }

// // Запускаем UDP-сервер на локальном порту
// func startUDPServer(localPort int) (*net.UDPConn, error) {
// 	addr := &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: localPort}
// 	conn, err := net.ListenUDP("udp", addr)
// 	if err != nil {
// 		return nil, fmt.Errorf("не удалось запустить UDP-сервер: %w", err)
// 	}
// 	return conn, nil
// }

// // Отправляем сообщение на удалённый адрес
// func sendMessage(conn *net.UDPConn, remoteAddr string, msg string) error {
// 	remoteUDPAddr, err := net.ResolveUDPAddr("udp", remoteAddr)
// 	if err != nil {
// 		return fmt.Errorf("неверный адрес: %w", err)
// 	}

// 	_, err = conn.WriteTo([]byte(msg), remoteUDPAddr)
// 	return err
// }

// // Ждём сообщения (блокирующий вызов)
// func receiveMessages(conn *net.UDPConn) {
// 	buffer := make([]byte, 1024)
// 	for {
// 		n, remoteAddr, err := conn.ReadFromUDP(buffer)
// 		if err != nil {
// 			log.Printf("Ошибка при приёме: %v", err)
// 			continue
// 		}
// 		fmt.Printf("Получено от %s: %s\n", remoteAddr, string(buffer[:n]))
// 	}
// }

// // clinetn a
// func main_A() {
// 	const (
// 		localPort      = 50000
// 		stunServer     = "stun.l.google.com:19302"
// 		peerPublicIP   = "203.0.113.20" // Публичный IP Стороны B (вручную или через сигнальный сервер)
// 		peerPublicPort = 50001          // Публичный порт Стороны B
// 	)

// 	// Шаг 1: Получаем свой публичный адрес через STUN
// 	mappedAddr, err := getPublicAddr(stunServer)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Сторона A: мой публичный адрес %s:%d\n", mappedAddr.IP, mappedAddr.Port)

// 	// Шаг 2: Запускаем UDP-сервер
// 	udpConn, err := startUDPServer(localPort)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer udpConn.Close()

// 	// Шаг 3: Отправляем приветствие на публичный адрес Стороны B
// 	peerAddr := fmt.Sprintf("%s:%d", peerPublicIP, peerPublicPort)
// 	err = sendMessage(udpConn, peerAddr, "Привет от Стороны A!")
// 	if err != nil {
// 		log.Printf("Не удалось отправить сообщение: %v", err)
// 	} else {
// 		fmt.Printf("Отправлено на %s\n", peerAddr)
// 	}

// 	// Шаг 4: Ждём ответов
// 	fmt.Println("Ожидание сообщений...")
// 	receiveMessages(udpConn)
// }

// // cline b
// func main_B() {
// 	const (
// 		localPort      = 50001
// 		stunServer     = "stun.l.google.com:19302"
// 		peerPublicIP   = "203.0.113.10" // Публичный IP Стороны A
// 		peerPublicPort = 50000          // Публичный порт Стороны A
// 	)

// 	// Шаг 1: Получаем свой публичный адрес через STUN
// 	mappedAddr, err := getPublicAddr(stunServer)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Сторона B: мой публичный адрес %s:%d\n", mappedAddr.IP, mappedAddr.Port)

// 	// Шаг 2: Запускаем UDP-сервер
// 	udpConn, err := startUDPServer(localPort)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer udpConn.Close()

// 	// Шаг 3: Ждём сообщений от Стороны A
// 	fmt.Println("Ожидание сообщений...")
// 	go receiveMessages(udpConn)

// 	// Шаг 4: Через 5 секунд отвечаем
// 	time.Sleep(5 * time.Second)
// 	peerAddr := fmt.Sprintf("%s:%d", peerPublicIP, peerPublicPort)
// 	err = sendMessage(udpConn, peerAddr, "Ответ от Стороны B!")
// 	if err != nil {
// 		log.Printf("Не удалось отправить ответ: %v", err)
// 	} else {
// 		fmt.Printf("Ответ отправлен на %s\n", peerAddr)
// 	}

// 	// Держим соединение открытым
// 	select {}
// }
