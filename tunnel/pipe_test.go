/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-08-02 19:32:39
 * @LastEditTime: 2019-09-21 14:02:53
 */

package tunnel

import (
	"fmt"
	"testing"
)

func Test_pipe(t *testing.T) {
	var in = NewVirtualConn()
	var out = NewVirtualConn()

	p := newPipe()
	p.In = in
	p.Out = out

	go p.Flow()

	go func() {
		for i := 0; i < 1000; i++ {
			msg := fmt.Sprint(i)
			in.Write([]byte(msg))
		}
	}()

	for c := 0; c < 1000; c++ {
		valid := fmt.Sprint(c)

		var b = make([]byte, 1024)
		l, err := out.Read(b)
		if err != nil {
			break
		}

		var r = string(b[:l])
		if r != valid {
			t.Error("r should be: ", valid, " but: ", r)
		}
	}
}
