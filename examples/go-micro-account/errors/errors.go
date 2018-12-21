package errors

import (
	"fmt"
	e "github.com/micro/go-micro/errors"
)
import proto "github.com/unliar/proto/account"

func ConvertToString(e int32) string {
	return fmt.Sprintf("%v", e)
}

var (
	UserNotFound = e.BadRequest(ConvertToString(proto.AccountErrors_UserNotFound), "USER NOT FOUND")
)
