package main

import "sync"

// 1.懒汉方式
//缺点： 非线程安全

/*
type singleton struct{}

var ins *singleton

func GetIns() *singleton {
	if ins == nil {
		ins = &singleton{}
	}
	return ins
}
*/

// 2.饿汉方式
// 缺点：如果singleton创建初始化比较复杂耗时时，加载时间会延长。

/*
type singleton struct{}

var ins *singleton = &singleton{}

func GetIns() *singleton {
	return ins
}
*/

// 3.懒汉加锁
// 缺点：虽然解决并发的问题，但每次加锁是要付出代价的
/*
type singleton struct{}
var ins *singleton
var mu sync.Mutex
func GetIns()*singleton{
	mu.Lock()
	defer mu.Unlock()
	if ins == nil {
		ins = &singleton{}
	}
	return ins
}
*/

// 4.懒汉初始加锁
// 避免了每次加锁，提高代码效率
/*
type singleton struct{}

var ins *singleton
var mu sync.Mutex

func GetIns() *singleton {
	if ins == nil {
		mu.Lock()
		defer mu.Unlock()
		if ins == nil {
			ins = &singleton{}
		}
	}
	return ins
}
*/

// 5.sync.Once实现

type singleton struct{}

var ins *singleton
var once sync.Once

func GetIns() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}
