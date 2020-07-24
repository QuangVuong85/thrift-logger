import sys
sys.path.append('../gen-py')

from LoggerPy import Logger

from thrift import Thrift
from thrift.transport import TSocket, TTransport
from thrift.protocol import TBinaryProtocol

def main():
    try:
        transport = TSocket.TSocket('localhost', 9090)
        transport = TTransport.TBufferedTransport(transport)
        protocol = TBinaryProtocol.TBinaryProtocol(transport)

        client = Logger.Client(protocol)
        transport.open()

        logfile = "logfile.log"

        client.timestamp(logfile)
        print("Logged timestamp to log file")

        client.write_log(logfile, "This is a message that I am writing to the log")
        print("Last line of log file is: %s" % (client.get_last_log_entry(logfile)))
        print("Size of log file is: %d bytes" % client.get_log_size(logfile))

        transport.close()
    except TTransport.TTransportException:
        print ("Error starting client")
    except Thrift.TException as e:
        print("Error: %d %s" % (e.error_code, e.error_description))

if __name__ == '__main__':
    main()
