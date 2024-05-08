package logging

import (
    "fmt"
    "log"
    "net"
    "runtime"
)

func SetUpLogging(ip string, port string) {
    conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", ip, port))
    if err != nil {
        log.Fatalf("Failed to connect to log server: %v", err)
    }

    log.SetOutput(conn)
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func Info(message string) {
    _, file, line, _ := runtime.Caller(1)
    log.Printf("[INFO] %s:%d: %s", file, line, message)
}

func Error(err error) {
    _, file, line, _ := runtime.Caller(1)
    log.Printf("[ERROR] %s:%d: %s", file, line, err.Error())
}