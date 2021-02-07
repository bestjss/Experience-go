package hello

import (
	"errors"
	"fmt"
)

/**
* 初始化用 生命周期函数
**/
func init()  {
	fmt.Println("world init")
}

/**
* 非生命周期函数 所以未输出
**/
func a()  {
	fmt.Println("a")
}

func World(key string) (string, error)  {
	if key == "" {
		return "", errors.New("key is empty")
	}
	return " world "+ key, nil
}
