# Week01 (2022-04-11 - 2022-04-17)

## 循环

## 变量和常量
1. 常量 const identifier type
2. 变量 var identifier type

### 变量定义

### 类型转换
```go
var i int = 42
var f float64 = float64(i)
```
或者：
```go
i := 42
f := float64(i)
```

### 类型推导
```go
var i int
j := i
```


# 数组

## 定义
var identifier [len]type
示例：myArray := [3]int{1,2,3}

## 切片
数组定义中不指定长度，即为切片，未初始化之前，默认为nil，长度为0
var identifier []type


### 定义
New返回指针地址
Make 返回第一个元素，可预设内存空间

# Map
var map1 map[keytype]keytype
```go
myMap := make(map[string]string, 10)
myMap["a"] = "b"
```

# 函数

## main函数

## init函数
会在包初始化时运行



# 领教教学

- [Apple Silicon M1虚拟机现状-小果冻之家](https://www.pimspeak.com/apple-m1-virtualization-situation.html)

## Go切片扩容机制：

切片扩容机制：当原切片长度小于1024时，新切片的容量会直接翻倍。而当原切片的容量大于等于1024时，会反复地增加25%，直到新容量超过所需要的容量

注意：1.18之后扩容机制变了！https://github.com/golang/go/commit/2dda92ff6f9f07eeb110ecbf0fc2d7a0ddd27f9d

```
starting cap    growth factor
256             2.0
512             1.63
1024            1.44
2048            1.35
4096            1.30
```

## 函数：多返回值



[GOPROXY](https://goproxy.io/zh/)
