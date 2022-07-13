# 模糊测试
## 1.模糊测试是什么？
模糊测试是一种向程序提供随机意外的输入以测试可能的崩溃或者边缘情况的方法。通过模糊测试可以揭示一些逻辑错误或者性能问题，因此使用模糊测试可以让程序的稳定性和性能都更有保证。

我们进行单元测试，一般都是表格驱动测试：自己思考一些边界条件和可能的情况，以表格的形式传递给测试程序。模糊测试可以解决边界条件不够完善的情况。

## 2.怎么进行模糊测试？
模糊测试是go1.18正式加入的功能，需要先将你的go版本升级到1.18及以后。

### 2.1 升级go1.18
 https://go.dev/dl/

### 2.2 开始模糊测试
待测试函数：
```go
func Equal(a []byte, b []byte) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
```

模糊测试写法：
```go
func FuzzEqual(f *testing.F) {
	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		Equal(a, b)
	})
}
```

执行模糊测试：

>  `go test -fuzz .`

### notice

1. 模糊测试函数申明以`Fuzz`开头，形参为`*testing.F`
2. 指定模糊测试时长：`go test -fuzz -fuzztime=10s`



## refenrence

- [官方文档](https://go.dev/doc/fuzz/)
- [Go 1.18 让写测试的代码量骤减，你会开始写测试吗？](https://mp.weixin.qq.com/s/7I0zB_AsltzDLmc9ew48Bg)

