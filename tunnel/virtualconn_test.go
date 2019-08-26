/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-08-03 10:39:12
 * @LastEditTime: 2019-08-03 10:46:14
 */

package tunnel

import (
	"fmt"
	"testing"
)

func Test_virtualConn(t *testing.T) {
	var c = NewVirtualConn()

	go func() {
		for i := 0; i < 1000; i++ {
			msg := fmt.Sprint(i)
			c.Write([]byte(msg))
		}
		c.Close()
	}()

	for i := 0; i < 1000; i++ {
		b := make([]byte, 1024)
		valid := fmt.Sprint(i)

		l, err := c.Read(b)
		if err != nil {
			break
		}

		msg := string(b[:l])

		if msg != valid {
			t.Error("valid is: ", valid, " but result is: ", msg)
		}
	}
}
