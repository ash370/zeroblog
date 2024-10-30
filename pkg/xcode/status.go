package xcode

/*import (
	"github.com/emicklei/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func toXCode(grpcStatus *status.Status) Code {
	grpcCode := grpcStatus.Code()
	switch grpcCode {
	case codes.OK:
		return OK
	case codes.InvalidArgument:
		return RequestErr
	case codes.NotFound:
		return NotFound
	case codes.PermissionDenied:
		return AccessDenied
	case codes.Unauthenticated:
		return Unauthorized
	case codes.ResourceExhausted:
		return LimitExceed
	case codes.Unimplemented:
		return MethodNotAllowed
	case codes.DeadlineExceeded:
		return Deadline
	case codes.Unavailable:
		return ServiceUnavailable
	case codes.Unknown:
		return String(grpcStatus.Message())
	}

	return ServerErr
}

func GrpcStatusToXCode(gstatus *status.Status) XCode {
	details := gstatus.Details()
	for i := len(details) - 1; i >= 0; i-- {
		detail := details[i]
		if pb, ok := detail.(proto.Message); ok {
			return FromProto(pb)
		}
	}

	return toXCode(gstatus)
}*/
