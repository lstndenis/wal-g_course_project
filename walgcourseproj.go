package main

import (
    "fmt"
    "github.com/wal-g/tracelog"
    "net"
    "os"
)

// Функция для отправки информации о версии WAL-G
func sendWALGVersionInfo(destination string) {
    versionInfo := "WAL-G Version: 3.0.1"  // Здесь должна быть актуальная версия WAL-G
    conn, err := net.Dial("tcp", destination)
    if err != nil {
        tracelog.ErrorLogger.Fatalf("Failed to connect to %s: %v", destination, err)
    }
    defer conn.Close()

    _, err = fmt.Fprintf(conn, "%s\n", versionInfo)
    if err != nil {
        tracelog.ErrorLogger.Fatalf("Failed to send WAL-G version info to %s: %v", destination, err)
    }

    tracelog.InfoLogger.Printf("Sent WAL-G version info to %s", destination)
}

// Функция для приема информации о версии WAL-G
func receiveWALGVersionInfo(port int) {
    listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
    if err != nil {
        tracelog.ErrorLogger.Fatalf("Failed to start listener on port %d: %v", port, err)
    }
    defer listener.Close()

    conn, err := listener.Accept()
    if err != nil {
        tracelog.ErrorLogger.Fatalf("Failed to accept incoming connection: %v", err)
    }
    defer conn.Close()

    // Читаем версию WAL-G, отправленную другим узлом
    buf := make([]byte, 1024)
    n, err := conn.Read(buf)
    if err != nil {
        tracelog.ErrorLogger.Fatalf("Failed to read WAL-G version info: %v", err)
    }

    versionInfo := string(buf[:n])
    tracelog.InfoLogger.Printf("Received WAL-G version info: %s", versionInfo)
}

func main() {
    // Пример запуска обмена информацией о версии WAL-G
    go sendWALGVersionInfo("192.168.1.2:12345")  // Замените на адрес и порт другого узла
    receiveWALGVersionInfo(12345)
}
