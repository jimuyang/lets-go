package main

import (
	"context"
	"fmt"
)

func entry(ctx context.Context) {
	logList := make([]string, 0)
	ctx = context.WithValue(ctx, "DEBUG_LOG", &logList)
	invoke1(ctx)
	ptr := ctx.Value("DEBUG_LOG").(*[]string)
	fmt.Println(*ptr)
}

func invoke1(ctx context.Context) {

	ptr := ctx.Value("DEBUG_LOG").(*[]string)
	*ptr = append(*ptr, "22222")

	invoke2(ctx)
}

func invoke2(ctx context.Context) {
	ptr := ctx.Value("DEBUG_LOG").(*[]string)
	*ptr = append(*ptr, "11111")
}
