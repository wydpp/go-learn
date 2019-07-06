package _map

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1 =  %d", len(m1))

	m2 := map[int]int{}
	m2[4] = 6
	t.Logf("len m2 =%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	//不存在0
	t.Log(m1[1])
	m1[2] = 0
	//2也是0？
	t.Log(m1[2])
	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("key 3 = %d", v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 3, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
