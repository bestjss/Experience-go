# Env
- [Go mode](https://golang.org/ref/mod)
- [Learn](https://learn.go.dev/)
- [Go Packages](https://golang.org/pkg/)
- [tutorial](https://golang.org/doc/tutorial/getting-started)
- [Tour](https://pkg.go.dev/golang.org/x/tour)
- [External Packages](https://pkg.go.dev/)
- [Support guyan0319/golang_development_notes 写的挺好的](https://github.com/guyan0319/golang_development_notes/)

```sh
# 安装Go
$ wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
$ tar -C /usr/local -xzf go1.13.4.linux-amd64.tar.gz
$ mkdir -p /root/go
$ echo 'export GOROOT=/usr/local/go export PATH=$PATH:$GOROOT/bin export GOPATH=/root/go' >> /etc/profile
$ source /etc/profile

# 七牛源
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
------
# 阿里源
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

# 13版本以下
$ export GOPROXY=https://mirrors.aliyun.com/goproxy/
```

```sh
# 初始化包管理器
$ go mod init demo
# 编译
$ go build
# 下载依赖
$ go get
# 查看项目依赖列表
$ go list -m all

$ go mod download 下载 go.mod 文件中指明的所有依赖

$ go mod tidy 整理现有的依赖，删除未使用的依赖。

$ go mod graph 查看现有的依赖结构

$ go mod init [module] module为项目目录名，生成 go.mod 文件 (Go 1.13 中唯一一个可以生成 go.mod 文件的子命令)

$ go mod edit 编辑 go.mod 文件

$ go mod vendor 导出现有的所有依赖 (事实上 Go modules 正在淡化 Vendor 的概念)

$ go mod verify 校验一个模块是否被篡改过

$ go clean -modcache 清理所有已缓存的模块版本数据。

$ go mod 查看所有 go mod的使用命令。
```

## Info

```go
// 方法定义
<-方法-><-方法名-><-参数名-><-参数类型-><-返回类型->
  func   Hello   (  name     string  )   string
// 说明
// 方法名 Hello 大写开头标识public 方法
// := 标识初始化并赋值
```

### Testing
```sh
# Add a file endwith [_test.go] in the same package which needs to test
## Sample : use hello_test.go to testing hello.go
$ go test
```

### Install
```shell
$ go build .
$ go install
# Run
$ study
```

## Basic knowledge

### 数组
```sh
# 数组
## var identifier [len]type
> var arr1 [5]int

## 初始化
var a = [10]int{1,2,3,4,5,6,7,8,9,0}
// 不足自动补0
var a = [10]int{1,2,3,4}
// 不定长数组
a := [...]int{1,2,3,4,5}  

## 遍历数组
//第一种
for i := 0; i < len(a); i++ {
   fmt.Println(a[i])
}
//第二种
for _, e := range a {
   fmt.Println(e)
}

## 统计数组长度
len(a)

## 多维数组 三行四列
var a = [3][4]int{[4]int{1,2,3},[4]int{1,2,3},[4]int{1,2}} //不足补0
[
  [1 2 3 0] 
  [1 2 3 0] 
  [1 2 0 0]
]
```

### 切片
```sh
## var identifier []type （不需要说明长度）
var s []string
或
s:=[]string{}

s = append(s, "fsd") 

## 取截取数组到切片
s := []string{"dd", "fsdf"}
slice1 := s[1:2]   //[开始：结束]  开始和结束不能大于数组实际长度
fmt.Println(slice1)

## 注意切片不能再重新切片
slice1 := slice1[:1]  这样会报错：no new variables on left side of :=

## 但可以对切片添加内容
slice1 = append(slice1, "aa", "bb")

## 合并两个切片
s1 := []int{1, 2}
s2 := []int{1, 4}
s1 = append(s1, s2...)
fmt.Println(s1)
注意： append切片时候要记着+"...",这是因为append()一个可变参数函数，...允许您从切片将多个参数传递给可变参数函数。

## 插入指定位置
s := []string{"aa", "bb", "cc", "dd"}
//插入指定位置
index := 2
temp := append([]string{}, s[index:]...)
s = append(s[:index], "ee")
s = append(s, temp...)
fmt.Println(s) //[aa bb ee cc dd]

## 删除指定索引元素
s := []string{"aa", "bb","ee", "cc", "dd"}
//删除指定索引元素
index = 2
s = append(s[:index], s[index+1:]...)
fmt.Println(s) //[aa bb cc dd]

## 删除指定值元素
s := []string{"aa", "bb", "cc", "dd"}
value := "dd"
//删除指定值
for k, v := range s {
  if v == value {
    s = append(s[:k], s[k+1:]...)
  }
}
fmt.Println(s) //[aa bb cc]

## make() 创建一个切片
当相关数组还没有定义时，我们可以使用 make() 函数来创建一个切片 同时创建好相关数组：
var slice1 []type = make([]type, len)。
也可以简写为 slice1 := make([]type, len)，这里 len 是数组的长度并且也是 slice 的初始长度。
示例：
var slice1 []string = make([]string , 10)

## new() 和 make() 的区别
看起来二者没有什么区别，都在堆上分配内存，但是它们的行为不同，适用于不同的类型。
- new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体；它相当于 &T{}。
- make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 和 channel。
换言之，new 函数分配内存，make 函数初始化；
----
p := new([2]int) //p指针，也可以使用[]进行赋值
p[0] = 22
b := make([]int, 10, 50) //第一个参数是类型，第二个参数是分配的空间，第三个参数是预留分配空间
b[0] = 13
fmt.Println(p, b)
打印结果：
&[22 0] [13 0 0 0 0 0 0 0 0 0]
----
如果想用预留空间

package main
import "fmt"
func main() {
   p := new([2]int)
   p[0] = 22
   b := make([]int, 10, 50) //第一个参数是类型，第二个参数是分配的空间，第三个参数是预留分配空间
   a := b[:cap(b)]
   fmt.Println(a)
   // [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
}

## 多维切片（slice）**
也是同数组一样，slice 可以组合为多维的 slice：
slice := [][]int{{10}, {20, 30}}
slice[0] = append(slice[0], 20)
```

### Map

```sh
# Map 是一种无序的键值对的集合,一个key=>value的hash结构
var map1 map[keytype]valuetype
# E.g.
var  mapLit map[string]int   // 未初始化的 map 的值是 nil
# 初始化
m = map[string]int{}
m = make(map[string]int)
m ：= map[string]int{"one": 1, "two": 2}

# 赋值
var m map[string]int
m = map[string]int{"one": 1, "two": 2}
m["d"] = 2
fmt.Println(m)

# 多维map
m := map[string]map[string]int{}
d := map[string]int{}
d["b"] = 23
m["a"] = d
fmt.Println(m)

# 删除单个元素
var m map[string]int
m = map[string]int{"one": 1, "two": 2}	
fmt.Println(m)
delete(m, "one") //删除key为 one
fmt.Println(m, len(m))

# 删除整个map元素
var m map[string]int
m = map[string]int{"one": 1, "two": 2}
m = nil//清空map 元素
fmt.Println(m, len(m))

# 判断key存不存在
var m map[string]int
m = map[string]int{"one": 1, "two": 2}
if _,ok:=m["one"];!ok{
	fmt.Println("The key does not exist")
}

# 比较两个map
package main

import (
   "fmt"
   "reflect"
)

func main() {
   m1 := map[string]int{"a": 1, "b": 2, "c": 3}
   m2 := map[string]int{"a": 1, "c": 3, "b": 2}
   //方法一
   fmt.Println(reflect.DeepEqual(m1, m2)) // 其中方法一用到了反射，效率相对比较低，相差大约10倍
   //方法二
   fmt.Println(cmpMap(m1, m2))
}

func cmpMap(m1, m2 map[string]int) bool {
   if len(m1) == len(m2) {
      for k1, v1 := range m1 {
         if v2, ok := m2[k1]; ok {
            if v1 != v2 {
               return false
            }
         } else {
            return false
         }
      }
      return true
   }
   return false
}

```

### 结构体 struct

- Go的struct是字段的类型集合。
- struct 主要应用两个方面：
  - 可以将不同类型的数据存放到struct，我们都知道数组只能存放单一的数据类型.
  - 由于在GO中没有class的关键字，也就是其它语言经常在面向对象中使用的方面，但GO是通过struct结构与method方法组合来实现的面向对象

- 例子
```go
package main
import (
  "fmt"
  )
type Person struct { //结构也是一种类型
  Name string //定义struct的属性
  Age  int
}
func main() {
  a := Person{}
  a.Name = "jerry" //对struct的属性进行操作，类型与class的使用方法
  a.Age = 19
  fmt.Println(a)
}
```

- struct 作为一种类型和其他类型结合用

```go
package main
import (
	"fmt"
)
type Person struct {
	//结构也是一种类型
	Name string //定义struct的属性
	Age  int
}
func main() {
	m := map[string]Person{}
	p := Person{Name: "jerry", Age: 12}
	m["ONE"] = p
	fmt.Println(m)
}
```
- struct 面向对象示例

```go
package main
import "fmt"
type Person struct {
   //结构也是一种类型
   Name string //定义struct的属性
   Age  int
}

// 对象方法
// 结构体中添加 list方法
func (p *Person) list() {
   fmt.Println(p.Name)
}

func main() {   
   p := Person{Name: "jerry", Age: 12}
   p.list()
}
```

- 匿名字段

```go
package main

import "fmt"

type YellowPerson struct {
   // 类型  Person 名字 Person
   Person //anonymous field（匿名自字段，名字就是Person）
   // 类型  string 名字 string
   string //anonymous field（匿名自字段，名字就是string）
}
type Person struct {
   //结构也是一种类型
   Name string //定义struct的属性
   Age  int
}

func main() {
   p := YellowPerson{Person: Person{Name: "jerry", Age: 12}, string: "jerry"}
   fmt.Println(p.Person)
   fmt.Println(p.Name) // 重点  Name属于 Person 结构体中
   fmt.Println(p.Age)  // 重点  Age属于 Person 结构体中
   fmt.Println(p.string)
}

// ---------------------------
// 下面的例子是错的，不可以像上面Main中 p.Name , p.Age 使用
type YellowPerson struct {
   Person Person 
}

func main() {
   p := YellowPerson{Person: Person{Name: "jerry", Age: 12}, string: "jerry"}
   fmt.Println(p.Name)
   fmt.Println(p.Age)
}
// >p.Name undefined (type YellowPerson has no field or method Name)，
// >p.Age undefined (type YellowPerson has no field or method Age)
```

- 匿名结构体

```go
// 在函数外部定义匿名结构体并赋值给 config
var config struct {
  APIKey string
  OAuthConfig oauth.Config // 匿名结构体
}
// ---------- 分割 ------------
// 定义并初始化并赋值给 data
data := struct {
  Title string
  Users []*User
}{
  title,
  users
}
```

### 常量
- 常量使用关键字 const 定义，用于存储不会改变的数据。
- 常量的定义格式：const identifier [type] = value
```go
const name = "ok"  // 隐式类型定义
const name1 string= "ok"  // 显式类型定义
fmt.Println(name)
fmt.Println(name1)
```

### 变量
- 声明变量的一般形式是使用 var 关键字：var identifier type
```go
// 相同类型一起定义
var a, b *int // 指针
var a, b int
// 不同类型
var a int
var b bool
var str string 
// 赋值
a = 15
b = false

// 变量声明后必须使用，全局变量除外
// 省略var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误。
var b = 10
var c //错误
c : = 10

// 必须声明类型 ，以下会报错
var c // 定义
c=11  // 赋值
// 以上错误，可以改成
var c int
c=11
```

### 流程控制

#### if

```go
if condition {
    // do something 
}

if condition {
    // do something 
} else {
    // do something 
}

if condition1 {
    // do something 
} else if condition2 {
    // do something else    
}else {
    // catch-all or default
}
```

#### switch
- 相比较 C 和 Java 等其它语言而言，Go 语言中的 switch 结构使用上更加灵活。它接受任意形式的表达式：
```go
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
```

#### for
- for 初始化语句; 条件语句; 修饰语句 {}

```go
for i := 0; i < 5; i++ {
  fmt.Printf("This is the %d iteration\n", i)
}
```

#### for-range
- 这是 Go 特有的一种的迭代结构，您会发现它在许多情况下都非常有用。它可以迭代任何一个集合，语法上很类似其它语言中 foreach 语句。

```go
for pos, char := range str {
...
}
// --- Sample ---
var a = [10]int{1,2,3,4,5,6,7,8,9,0}

for _, e := range a { 
  fmt.Println(e) 
}
```

#### Break 与 continue
- break退出循环， continue调到下一个循环

```go
// 循环三次 每次打印0-5
for i:=0; i<3; i++ {
  for j:=0; j<10; j++ {
      if j>5 {
          break   
      }
      print(j)
  }
  print("  ")
}

for i := 0; i < 10; i++ {
  if i == 5 {
      continue
  }
  print(i)
  print(" ")
}
```

#### label
- Go语言也支持label(标签)语法：分别是break labe、 goto label和continuelabel。

```go
package main

import "fmt"

func main() {
  Loop:
    for i := 0; i < 10; i++ {
        fmt.Println("label i is ", i)
        for j := 0; j < 10; j++ {
          if j > 5 {
              // 跳到外面去啦，但是不会再进来这个for循环了
              break Loop
          }
        }
    }
    //跳转语句 goto语句可以跳转到本函数内的某个标签
    gotoCount := 0
  GotoLabel:
    gotoCount++
    if gotoCount < 10 {
        goto GotoLabel //如果小于10的话就跳转到GotoLabel
    }
    fmt.Println(gotoCount)
    
  Continuelabel:
    for a < 20 {
      if a == 15 {
        a++
        continue Continuelabel
      }
      fmt.Println(a)
      a++
    }
  }
}
// label i is 0 10 10 11 12 13 14 16 17 18 19
```

#### 反射reflect