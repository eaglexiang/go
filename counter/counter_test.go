/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-05-17 21:49:49
 * @LastEditTime: 2019-05-17 22:20:24
 */

package counter

import (
	"testing"
)

func Test_Up(t *testing.T) {
	c := Counter{}
	for i := 0; i < 2000; i++ {
		c.Up()
	}
	if c.Value != 2000 {
		t.Error("c.data should be 2000 but: ", c.Value)
	}
}

func Test_UpWithStep(t *testing.T) {
	c := Counter{}
	for i := 0; i < 2000; i++ {
		c.Up(2)
	}
	if c.Value != 4000 {
		t.Error("c.data should be 4000 but: ", c.Value)
	}
}

func Test_Down(t *testing.T) {
	c := Counter{
		Value: 2000,
	}
	for i := 0; i < 2000; i++ {
		c.Down()
	}
	if c.Value != 0 {
		t.Error("c.data should be 0 but: ", c.Value)
	}
}

func Test_DownWithStep(t *testing.T) {
	c := Counter{
		Value: 4000,
	}
	for i := 0; i < 2000; i++ {
		c.Down(2)
	}
	if c.Value != 0 {
		t.Error("c.data should be 0 but: ", c.Value)
	}
}

func Test_GreaterThanZero(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Error("value should be less than 0")
		}
	}()

	c := Counter{}
	c.Down()
}

func Test_OverFlow(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Error("value should be overflow")
		}
	}()

	c := Counter{}
	for i := 0; i < 10; i++ {
		c.Up(1 << 62)
	}
}
