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
