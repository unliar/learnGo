package utils

import (
	"context"
	"github.com/micro/go-micro/metadata"
)

func AddMetaData(ctx context.Context, m map[string]string) context.Context {

	md, ok := metadata.FromContext(ctx)
	if ok {
		for k, v := range md {
			m[k] = v
		}
	}

	ctx = metadata.NewContext(ctx, m)
	return ctx

}

func ContainItem(arr []string, item string) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}
