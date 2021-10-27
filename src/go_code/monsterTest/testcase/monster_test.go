package testcase

import "testing"

func TestMonster_Store(t *testing.T) {
	//先创建一个monster
	monster := &monster{
		Name:  "海乱鬼",
		Age:   35,
		skill: "只此一刀",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("Store()错误，希望为：%v,实际为%v", true, res)
	}
	t.Logf("monster.Store()测试成功！")
}

func TestMonster_ReStore(t *testing.T) {
	//先创建monster
	var monster *monster = &monster{}
	res := monster.ReStore()

	if !res {
		t.Fatalf("ReStore()错误，希望为：%v,实际为%v", true, res)
	}
	if monster.Name != "海乱鬼" {
		t.Fatalf("ReStore()错误，希望为：%v,实际为%v", "海乱鬼", monster.Name)

	}
	t.Logf("monster.ReStore()测试成功！")
}
