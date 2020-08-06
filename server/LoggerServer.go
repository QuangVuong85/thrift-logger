package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"thrift-logger/gen-go/loggerservice"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
)

const (
	NetWorkAddr = "127.0.0.1:9090"
)

type LoggerHandler struct{}

func (this *LoggerHandler) Timestamp(ctx context.Context, filename string) error {

	now := []byte(time.Now().String())

	err := ioutil.WriteFile(filename, now, 0775)
	if err != nil {
		return err
	}

	return nil

}

func (this *LoggerHandler) GetLastLogEntry(ctx context.Context, filename string) (string, error) {
	err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return "", nil
}

func (this *LoggerHandler) WriteLog(ctx context.Context, filename string, message string) error {
	err := ioutil.WriteFile(filename, []byte(message), 0775)
	if err != nil {
		return err
	}

	return nil
}

func (this *LoggerHandler) GetLogSize(ctx context.Context, filename string) (int32, error) {
	var fs int32 = 0

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return 0, err
	}

	info, err := file.Stat()
	if err != nil {
		return 0, err
	}

	fs = int32(info.Size())

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
