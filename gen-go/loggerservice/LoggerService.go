// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package loggerservice

import(
	"bytes"
	"context"
	"reflect"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

// Attributes:
//  - ErrorCode
//  - ErrorDescription
type LoggerException struct {
  ErrorCode int32 `thrift:"error_code,1" db:"error_code" json:"error_code"`
  ErrorDescription string `thrift:"error_description,2" db:"error_description" json:"error_description"`
}

func NewLoggerException() *LoggerException {
  return &LoggerException{}
}


func (p *LoggerException) GetErrorCode() int32 {
  return p.ErrorCode
}

func (p *LoggerException) GetErrorDescription() string {
  return p.ErrorDescription
}
func (p *LoggerException) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *LoggerException)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.ErrorCode = v
}
  return nil
}

func (p *LoggerException)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.ErrorDescription = v
}
  return nil
}

func (p *LoggerException) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("LoggerException"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *LoggerException) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("error_code", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:error_code: ", p), err) }
  if err := oprot.WriteI32(int32(p.ErrorCode)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.error_code (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:error_code: ", p), err) }
  return err
}

func (p *LoggerException) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("error_description", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:error_description: ", p), err) }
  if err := oprot.WriteString(string(p.ErrorDescription)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.error_description (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:error_description: ", p), err) }
  return err
}

func (p *LoggerException) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("LoggerException(%+v)", *p)
}

func (p *LoggerException) Error() string {
  return p.String()
}

type Logger interface {
  // Parameters:
  //  - Filename
  Timestamp(ctx context.Context, filename string) (err error)
  // Parameters:
  //  - Filename
  GetLastLogEntry(ctx context.Context, filename string) (r string, err error)
  // Parameters:
  //  - Filename
  //  - Message
  WriteLog(ctx context.Context, filename string, message string) (err error)
  // Parameters:
  //  - Filename
  GetLogSize(ctx context.Context, filename string) (r int32, err error)
}

type LoggerClient struct {
  c thrift.TClient
}

func NewLoggerClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *LoggerClient {
  return &LoggerClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewLoggerClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *LoggerClient {
  return &LoggerClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewLoggerClient(c thrift.TClient) *LoggerClient {
  return &LoggerClient{
    c: c,
  }
}

func (p *LoggerClient) Client_() thrift.TClient {
  return p.c
}
// Parameters:
//  - Filename
func (p *LoggerClient) Timestamp(ctx context.Context, filename string) (err error) {
  var _args0 LoggerTimestampArgs
  _args0.Filename = filename
  if err := p.Client_().Call(ctx, "timestamp", &_args0, nil); err != nil {
    return err
  }
  return nil
}

// Parameters:
//  - Filename
func (p *LoggerClient) GetLastLogEntry(ctx context.Context, filename string) (r string, err error) {
  var _args1 LoggerGetLastLogEntryArgs
  _args1.Filename = filename
  var _result2 LoggerGetLastLogEntryResult
  if err = p.Client_().Call(ctx, "get_last_log_entry", &_args1, &_result2); err != nil {
    return
  }
  switch {
  case _result2.Error!= nil:
    return r, _result2.Error
  }

  return _result2.GetSuccess(), nil
}

// Parameters:
//  - Filename
//  - Message
func (p *LoggerClient) WriteLog(ctx context.Context, filename string, message string) (err error) {
  var _args3 LoggerWriteLogArgs
  _args3.Filename = filename
  _args3.Message = message
  var _result4 LoggerWriteLogResult
  if err = p.Client_().Call(ctx, "write_log", &_args3, &_result4); err != nil {
    return
  }
  switch {
  case _result4.Error!= nil:
    return _result4.Error
  }

  return nil
}

// Parameters:
//  - Filename
func (p *LoggerClient) GetLogSize(ctx context.Context, filename string) (r int32, err error) {
  var _args5 LoggerGetLogSizeArgs
  _args5.Filename = filename
  var _result6 LoggerGetLogSizeResult
  if err = p.Client_().Call(ctx, "get_log_size", &_args5, &_result6); err != nil {
    return
  }
  switch {
  case _result6.Error!= nil:
    return r, _result6.Error
  }

  return _result6.GetSuccess(), nil
}

type LoggerProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler Logger
}

func (p *LoggerProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *LoggerProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *LoggerProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewLoggerProcessor(handler Logger) *LoggerProcessor {

  self7 := &LoggerProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self7.processorMap["timestamp"] = &loggerProcessorTimestamp{handler:handler}
  self7.processorMap["get_last_log_entry"] = &loggerProcessorGetLastLogEntry{handler:handler}
  self7.processorMap["write_log"] = &loggerProcessorWriteLog{handler:handler}
  self7.processorMap["get_log_size"] = &loggerProcessorGetLogSize{handler:handler}
return self7
}

func (p *LoggerProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x8 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x8.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush(ctx)
  return false, x8

}

type loggerProcessorTimestamp struct {
  handler Logger
}

func (p *loggerProcessorTimestamp) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := LoggerTimestampArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    return false, err
  }

  iprot.ReadMessageEnd()
  var err2 error
  if err2 = p.handler.Timestamp(ctx, args.Filename); err2 != nil {
    return true, err2
  }
  return true, nil
}

type loggerProcessorGetLastLogEntry struct {
  handler Logger
}

func (p *loggerProcessorGetLastLogEntry) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := LoggerGetLastLogEntryArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("get_last_log_entry", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := LoggerGetLastLogEntryResult{}
var retval string
  var err2 error
  if retval, err2 = p.handler.GetLastLogEntry(ctx, args.Filename); err2 != nil {
  switch v := err2.(type) {
    case *LoggerException:
  result.Error = v
    default:
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing get_last_log_entry: " + err2.Error())
    oprot.WriteMessageBegin("get_last_log_entry", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  }
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("get_last_log_entry", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}

type loggerProcessorWriteLog struct {
  handler Logger
}

func (p *loggerProcessorWriteLog) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := LoggerWriteLogArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("write_log", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := LoggerWriteLogResult{}
  var err2 error
  if err2 = p.handler.WriteLog(ctx, args.Filename, args.Message); err2 != nil {
  switch v := err2.(type) {
    case *LoggerException:
  result.Error = v
    default:
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing write_log: " + err2.Error())
    oprot.WriteMessageBegin("write_log", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  }
  }
  if err2 = oprot.WriteMessageBegin("write_log", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}

type loggerProcessorGetLogSize struct {
  handler Logger
}

func (p *loggerProcessorGetLogSize) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := LoggerGetLogSizeArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("get_log_size", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := LoggerGetLogSizeResult{}
var retval int32
  var err2 error
  if retval, err2 = p.handler.GetLogSize(ctx, args.Filename); err2 != nil {
  switch v := err2.(type) {
    case *LoggerException:
  result.Error = v
    default:
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing get_log_size: " + err2.Error())
    oprot.WriteMessageBegin("get_log_size", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  }
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("get_log_size", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Filename
type LoggerTimestampArgs struct {
  Filename string `thrift:"filename,1" db:"filename" json:"filename"`
}

func NewLoggerTimestampArgs() *LoggerTimestampArgs {
  return &LoggerTimestampArgs{}
}


func (p *LoggerTimestampArgs) GetFilename() string {
  return p.Filename
}
func (p *LoggerTimestampArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *LoggerTimestampArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Filename = v
}
  return nil
}

func (p *LoggerTimestampArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("timestamp_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *LoggerTimestampArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("filename", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:filename: ", p), err) }
  if err := oprot.WriteString(string(p.Filename)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.filename (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:filename: ", p), err) }
  return err
}

func (p *LoggerTimestampArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("LoggerTimestampArgs(%+v)", *p)
}

// Attributes:
//  - Filename
type LoggerGetLastLogEntryArgs struct {
  Filename string `thrift:"filename,1" db:"filename" json:"filename"`
}

func NewLoggerGetLastLogEntryArgs() *LoggerGetLastLogEntryArgs {
  return &LoggerGetLastLogEntryArgs{}
}


func (p *LoggerGetLastLogEntryArgs) GetFilename() string {
  return p.Filename
}
func (p *LoggerGetLastLogEntryArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *LoggerGetLastLogEntryArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Filename = v
}
  return nil
}

func (p *LoggerGetLastLogEntryArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("get_last_log_entry_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *LoggerGetLastLogEntryArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("filename", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:filename: ", p), err) }
  if err := oprot.WriteString(string(p.Filename)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.filename (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:filename: ", p), err) }
  return err
}

func (p *LoggerGetLastLogEntryArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("LoggerGetLastLogEntryArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Error
type LoggerGetLastLogEntryResult struct {
  Success *string `thrift:"success,0" db:"success" json:"success,omitempty"`
  Error *LoggerException `thrift:"error,1" db:"error" json:"error,omitempty"`
}

func NewLoggerGetLastLogEntryResult() *LoggerGetLastLogEntryResult {
  return &LoggerGetLastLogEntryResult{}
}

var LoggerGetLastLogEntryResult_Success_DEFAULT string
func (p *LoggerGetLastLogEntryResult) GetSuccess() string {
  if !p.IsSetSuccess() {
    return LoggerGetLastLogEntryResult_Success_DEFAULT
  }
return *p.Success
}
var LoggerGetLastLogEntryResult_Error_DEFAULT *LoggerException
func (p *LoggerGetLastLogEntryResult) GetError() *LoggerException {
  if !p.IsSetError() {
    return LoggerGetLastLogEntryResult_Error_DEFAULT
  }
return p.Error
}
func (p *LoggerGetLastLogEntryResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *LoggerGetLastLogEntryResult) IsSetError() bool {
  return p.Error != nil
}

func (p *LoggerGetLastLogEntryResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *LoggerGetLastLogEntryResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *LoggerGetLastLogEntryResult)  ReadField1(iprot thrift.TProtocol) error {
  p.Error = &LoggerException{}
  if err := p.Error.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Error), err)
  }
  return nil
}

func (p *LoggerGetLastLogEntryResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("get_last_log_entry_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *LoggerGetLastLogEntryResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteString(string(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *LoggerGetLastLogEntryResult) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetError() {
    if err := oprot.WriteFieldBegin("error", thrift.STRUCT, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:error: ", p), err) }
    if err := p.Error.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Error), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:error: ", p), err) }
  }
  return err
}

func (p *LoggerGetLastLogEntryResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("LoggerGetLastLogEntryResult(%+v)", *p)
}

// Attributes:
//  - Filename
//  - Message
type LoggerWriteLogArgs struct {
  Filename string `thrift:"filename,1" db:"filename" json:"filename"`
  Message string `thrift:"message,2" db:"message" json:"message"`
}

func NewLoggerWriteLogArgs() *LoggerWriteLogArgs {
  return &LoggerWriteLogArgs{}
}


func (p *LoggerWriteLogArgs) GetFilename() string {
  return p.Filename
}

func (p *LoggerWriteLogArgs) GetMessage() string {
  return p.Message
}
func (p *LoggerWriteLogArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *LoggerWriteLogArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Filename = v
}
  return nil
}

func (p *LoggerWriteLogArgs)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Message = v
}
  return nil
}

func (p *LoggerWriteLogArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("write_log_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *LoggerWriteLogArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("filename", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:filename: ", p), err) }
  if err := oprot.WriteString(string(p.Filename)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.filename (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:filename: ", p), err) }
  return err
}

func (p *LoggerWriteLogArgs) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("message", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:message: ", p), err) }
  if err := oprot.WriteString(string(p.Message)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.message (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:message: ", p), err) }
  return err
}

func (p *LoggerWriteLogArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("LoggerWriteLogArgs(%+v)", *p)
}

// Attributes:
//  - Error
type LoggerWriteLogResult struct {
  Error *LoggerException `thrift:"error,1" db:"error" json:"error,omitempty"`
}

func NewLoggerWriteLogResult() *LoggerWriteLogResult {
  return &LoggerWriteLogResult{}
}

var LoggerWriteLogResult_Error_DEFAULT *LoggerException
func (p *LoggerWriteLogResult) GetError() *LoggerException {
  if !p.IsSetError() {
    return LoggerWriteLogResult_Error_DEFAULT
  }
return p.Error
}
func (p *LoggerWriteLogResult) IsSetError() bool {
  return p.Error != nil
}

func (p *LoggerWriteLogResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *LoggerWriteLogResult)  ReadField1(iprot thrift.TProtocol) error {
  p.Error = &LoggerException{}
  if err := p.Error.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Error), err)
  }
  return nil
}

func (p *LoggerWriteLogResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("write_log_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *LoggerWriteLogResult) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetError() {
    if err := oprot.WriteFieldBegin("error", thrift.STRUCT, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:error: ", p), err) }
    if err := p.Error.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Error), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:error: ", p), err) }
  }
  return err
}

func (p *LoggerWriteLogResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("LoggerWriteLogResult(%+v)", *p)
}

// Attributes:
//  - Filename
type LoggerGetLogSizeArgs struct {
  Filename string `thrift:"filename,1" db:"filename" json:"filename"`
}

func NewLoggerGetLogSizeArgs() *LoggerGetLogSizeArgs {
  return &LoggerGetLogSizeArgs{}
}


func (p *LoggerGetLogSizeArgs) GetFilename() string {
  return p.Filename
}
func (p *LoggerGetLogSizeArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *LoggerGetLogSizeArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Filename = v
}
  return nil
}

func (p *LoggerGetLogSizeArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("get_log_size_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *LoggerGetLogSizeArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("filename", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:filename: ", p), err) }
  if err := oprot.WriteString(string(p.Filename)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.filename (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:filename: ", p), err) }
  return err
}

func (p *LoggerGetLogSizeArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("LoggerGetLogSizeArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Error
type LoggerGetLogSizeResult struct {
  Success *int32 `thrift:"success,0" db:"success" json:"success,omitempty"`
  Error *LoggerException `thrift:"error,1" db:"error" json:"error,omitempty"`
}

func NewLoggerGetLogSizeResult() *LoggerGetLogSizeResult {
  return &LoggerGetLogSizeResult{}
}

var LoggerGetLogSizeResult_Success_DEFAULT int32
func (p *LoggerGetLogSizeResult) GetSuccess() int32 {
  if !p.IsSetSuccess() {
    return LoggerGetLogSizeResult_Success_DEFAULT
  }
return *p.Success
}
var LoggerGetLogSizeResult_Error_DEFAULT *LoggerException
func (p *LoggerGetLogSizeResult) GetError() *LoggerException {
  if !p.IsSetError() {
    return LoggerGetLogSizeResult_Error_DEFAULT
  }
return p.Error
}
func (p *LoggerGetLogSizeResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *LoggerGetLogSizeResult) IsSetError() bool {
  return p.Error != nil
}

func (p *LoggerGetLogSizeResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *LoggerGetLogSizeResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *LoggerGetLogSizeResult)  ReadField1(iprot thrift.TProtocol) error {
  p.Error = &LoggerException{}
  if err := p.Error.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Error), err)
  }
  return nil
}

func (p *LoggerGetLogSizeResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("get_log_size_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *LoggerGetLogSizeResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.I32, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteI32(int32(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *LoggerGetLogSizeResult) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetError() {
    if err := oprot.WriteFieldBegin("error", thrift.STRUCT, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:error: ", p), err) }
    if err := p.Error.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Error), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:error: ", p), err) }
  }
  return err
}

func (p *LoggerGetLogSizeResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("LoggerGetLogSizeResult(%+v)", *p)
}


