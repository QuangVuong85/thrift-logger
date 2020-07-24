package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"thrift-logger/gen-go/loggerservice"

	"github.com/apache/thrift/lib/go/thrift"
)

const (
	NetWorkAddr = "127.0.0.1:9090"
)

type LoggerHandler struct {
}

func (this *LoggerHandler) Timestamp(ctx context.Context, filename string) (err error) {
	// file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE)
	// if err != nil {
	// 	return err
	// }

	// if file.Open() {
	// 	now := time.Now().String()
	// 	_, err := f.WriteString(now)
	// 	if err != nil {
	// 		defer file.Close()
	// 		return err
	// 	}

	// 	fmt.Println(now)

	// 	defer file.Close()
	// } else {
	// 	err := loggerservice.LoggerException{}
	// 	err.ErrorCode = 1
	// 	err.ErrorDescription = "Could not open file " + filename
	// 	return err
	// }

	return nil

}

func (this *LoggerHandler) GetLastLogEntry(ctx context.Context, filename string) (r string, err error) {
	// file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE)
	// if err != nil {
	// 	return "", err
	// }

	// if file.Open() {
	// 	last := nil
	// 	///

	// 	defer file.Close()
	// } else {
	// 	err := loggerservice.LoggerException{}
	// 	err.ErrorCode = 1
	// 	err.ErrorDescription = "Could not open file " + filename
	// 	return "", err
	// }

	return "", nil
}

func (this *LoggerHandler) WriteLog(ctx context.Context, filename string, message string) (err error) {
	// file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE)
	// if err != nil {
	// 	return err
	// }

	// if file.Open() {
	// 	fmt.Println(message)

	// 	defer file.Close()
	// } else {
	// 	err := loggerservice.LoggerException{}
	// 	err.ErrorCode = 1
	// 	err.ErrorDescription = "Could not open file " + filename
	// 	return err
	// }

	return nil
}

func (this *LoggerHandler) GetLogSize(ctx context.Context, filename string) (r int32, err error) {

	var fs int32 = 0

	// file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE)
	// if err != nil {
	// 	return 0, err
	// }

	// if file.Open() {
	// 	fs = file.GetLogSize()

	// 	defer file.Close()
	// } else {
	// 	err := loggerservice.LoggerException{}
	// 	err.ErrorCode = 1
	// 	err.ErrorDescription = "Could not open file " + filename
	// 	return nil, err
	// }

	return fs, nil
}

func main() {
	transportFactory := thrift.NewTBufferedTransportFactory(1024)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	serverTransport, err := thrift.NewTServerSocket(NetWorkAddr)
	if err != nil {
		log.Println(err, ": error")
		os.Exit(1)
	}

	handler := &LoggerHandler{}
	processor := loggerservice.NewLoggerProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Server has started\nListening...")
	server.Serve()

}
