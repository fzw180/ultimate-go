package main

import (
	"fmt"
	"github.com/grpc/grpc-go/status"
	"google.golang.org/grpc/codes"
)

// panic, 立刻stops execution the current function, then unwinding the stack, then run deferred functions
// panic,
// recover, 恢复control of the goroutine and execution
// recover, 只在defer function有用, 因为unwinding the stack时只会运行defer function里面的code

func panic1() {
	panic("error in panic1")
}

func panic2() {
	panic("error in panic2")
}

func main() {
	defer func() {
		if errRecover := recover(); errRecover != nil {
			var err error
			isError, ok := errRecover.(error)
			if ok {
				err = status.Error(codes.Aborted, isError.Error())
			} else {
				err = status.Error(codes.Aborted, fmt.Sprintf("%v", errRecover))
			}
			fmt.Println(err)
		}
	}()
	panic("error in main")
	recover() // unreachable code
	panic1()  // unreachable code
	panic2()  // unreachable code

}
