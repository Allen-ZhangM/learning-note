package main

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))
	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}
	m.EXPECT().Get(gomock.Eq("Tom2")).Return(100, nil)
	if v := GetFromDB(m, "Tom2"); v != 100 {
		t.Fatal("expected -1, but got", v)
	}
	m.EXPECT().Get(gomock.Any()).Return(630, nil)
	if v := GetFromDB(m, "Tom2"); v != 630 {
		t.Fatal("expected -1, but got", v)
	}
	m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil)
	if v := GetFromDB(m, "Tom2"); v != 0 {
		t.Fatal("expected -1, but got", v)
	}
	//m.EXPECT().Get(gomock.Nil()).Return(0, errors.New("nil"))
	//if v := GetFromDB(m, ""); v != -1 {
	//	t.Fatal("expected -1, but got", v)
	//}
	m.EXPECT().Get(gomock.Any()).Do(func(key string) {
		t.Log(key)
	})
	if v := GetFromDB(m, "key"); v != 0 {
		t.Fatal("expected -1, but got", v)
	}
	m.EXPECT().Get(gomock.Any()).DoAndReturn(func(key string) (int, error) {
		if key == "Sam" {
			return 630, nil
		}
		return 0, errors.New("not exist")
	})
	if v := GetFromDB(m, "Sam"); v != 630 {
		t.Fatal("expected -1, but got", v)
	}

	m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil).Times(2)
	GetFromDB(m, "ABC")
	GetFromDB(m, "DEF")

	//调用顺序
	//gomock.InOrder(o1, o2)

}
