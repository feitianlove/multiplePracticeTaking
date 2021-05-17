package server

import (
	"context"
)

/*
	1、必须是可导出类型的方法
	2、接受3个参数，第一个是 context.Context类型，其他2个都是可导出（或内置）的类型。
	3、第3个参数是一个指针
	4、有一个 error 类型的返回值
*/
type Calculate struct {
	A, B int
}
type Result struct {
	R int
}

type Arith int

func (a Arith) Mul(ctx context.Context, args *Calculate, r *Result) error {
	r.R = args.A * args.B
	return nil
}
