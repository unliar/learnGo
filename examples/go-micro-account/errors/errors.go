package errors

import "github.com/micro/go-micro/errors"
import proto "github.com/unliar/proto/account"

var (
	UserNotFound = errors.BadRequest(proto.AccountErrors)
)
