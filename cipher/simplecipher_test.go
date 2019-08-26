/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-08-20 21:18:28
 * @LastEditTime: 2019-08-26 21:23:00
 */

package cipher

import (
	"testing"

	"github.com/eaglexiang/go/bytebuffer"
)

func Test_SimpleCipher(t *testing.T) {
	var encryptor StreamEncryptor = SimpleCipher{key: 13}
	var decryptor StreamDecryptor = SimpleCipher{key: 13}

	b := bytebuffer.GetBuffer()
	b.SetString("test")

	c := encryptor.StreamEncrypt(b)
	if c.String() == "test" {
		t.Error("fail 2 encrypt")
	}

	p := decryptor.StreamDecrypt(c)
	if p.String() != "test" {
		t.Error("p.String() should be test but: ", p.String())
	}
}
