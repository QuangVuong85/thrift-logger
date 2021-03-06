//
// Created by vuongdq85 on 24/07/2020.
//

#include "gen-cpp/Logger.h"
#include <thrift/protocol/TBinaryProtocol.h>
#include <thrift/server/TSimpleServer.h>
#include <thrift/transport/TServerSocket.h>
#include <thrift/transport/TBufferTransports.h>

#include <chrono>
#include <fstream>
#include <string>
#include <iostream>

using namespace ::apache::thrift;
using namespace ::apache::thrift::protocol;
using namespace ::apache::thrift::transport;
using namespace ::apache::thrift::server;

using boost::shared_ptr;

using namespace LoggerCpp;

class LoggerHandler : virtual public LoggerIf {
public:
    LoggerHandler() {
        // Your initialization goes here
    }

    void timestamp(const std::string &filename) {
        std::fstream file;
        file.open(filename.c_str(), std::fstream::out | std::fstream::app);
        if (file.is_open()) {
            std::time_t now;
            now = std::chrono::system_clock::to_time_t(std::chrono::system_clock::now());
            file << std::ctime(&now);
            file.close();
        } else {
            LoggerException err;
            err.error_code = 1;
            err.error_description = std::string("Could not open file ") + filename;
            throw err;
        }
    }

    void get_last_log_entry(std::string &_return, const std::string &filename) {
        std::fstream file;

        file.open(filename.c_str(), std::fstream::in);
        if (file.is_open()) {
            std::string line;
            std::string lastline;
            while (std::getline(file, line))
                lastline = line;
            _return = lastline;
            file.close();
        } else {
            LoggerException err;
            err.error_code = 1;
            err.error_description = std::string("Could not open file ") + filename;
            throw err;
        }
    }

    void write_log(const std::string &filename, const std::string &message) {
        std::fstream file;
        file.open(filename.c_str(), std::fstream::out | std::fstream::app);
        if (file.is_open()) {
            file << message << std::endl;
            file.close();
        } else {
            LoggerException err;
            err.error_code = 1;
            err.error_description = std::string("Could not open file ") + filename;
            throw err;
        }
    }

    int32_t get_log_size(const std::string &filename) {
        int fs = 0;
        std::fstream file;
        file.open(filename.c_str(), std::fstream::in | std::fstream::ate);
        if (file.is_open()) {
            fs = file.tellg();
            file.close();
        } else {
            LoggerException err;
            err.error_code = 1;
            err.error_description = std::string("Could not open file ") + filename;
            throw err;
        }
        return fs;
    }

};

int main(int argc, char **argv) {
    int port = 9090;
    std::shared_ptr<LoggerHandler> handler(new LoggerHandler());
    std::shared_ptr<TProcessor> processor(new LoggerProcessor(handler));
    std::shared_ptr<TServerTransport> serverTransport(new TServerSocket(port));
    std::shared_ptr<TTransportFactory> transportFactory(new TBufferedTransportFactory());
    std::shared_ptr<TProtocolFactory> protocolFactory(new TBinaryProtocolFactory());

    TSimpleServer server(processor, serverTransport, transportFactory, protocolFactory);
    server.serve();
    return 0;
}

