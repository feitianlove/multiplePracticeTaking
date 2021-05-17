package server

import (
	"github.com/smallnest/rpcx/server"
	"testing"
)

func TestArith_Mul(t *testing.T) {
	s := server.NewServer()
	err := s.RegisterName("Arith", new(Arith), "")
	if err != nil {
		panic(err)
	}
	func() {
		_ = s.Serve("tcp", "127.0.0.1:8972")
	}()
}
