package HelloService

import (
	"context"
	"learnGo/examles/go-grpc-blog/proto/Message"
)

type Server struct {
}

func (s *Server) GetId(ctx context.Context, req *Hello.Req) (res *Hello.Res, err error) {
	return &Hello.Res{Id: 11}, nil
}
