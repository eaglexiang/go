/*
 * @Author: EagleXiang
 * @LastEditors: EagleXiang
 * @Email: eagle.xiang@outlook.com
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-02-19 15:38:19
 * @LastEditTime: 2019-02-19 15:50:44
 */

package version

import "testing"

func Test_Version(t *testing.T) {
	v0 := createVersion(t, "1")
	v1 := createVersion(t, "2")
	if v0.compareWith(v1) != Less {
		t.Error("v0 should be less than v1")
	}
	v1 = createVersion(t, "1")
	if v0.compareWith(v1) != Equal {
		t.Error("v0 should equal v1")
	}
	v1 = createVersion(t, "0")
	if v0.compareWith(v1) != Greater {
		t.Error("v0 should be greater than v1")
	}

	v1 = createVersion(t, "0.1")
	if v0.compareWith(v1) != Greater {
		t.Error("v0 should be greater than v1")
	}
	v1 = createVersion(t, "1.1")
	if v0.compareWith(v1) != Less {
		t.Error("v0 should be less than v1")
	}

	v0 = createVersion(t, "1.0.5")
	if v0.compareWith(v1) != Less {
		t.Error("v0 should be less than v1")
	}
}

func createVersion(t *testing.T, v string) Version {
	version, err := CreateVersion(v)
	if err != nil {
		t.Error(err)
	}
	return version
}
