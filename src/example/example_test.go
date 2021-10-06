package example

import "testing"

func TestAdd(t *testing.T) {
	r := sub(12, 6)
	if r != 6 {
		t.Error("计算错误！")
	}
	t.Logf("这是一个测试！")
}
