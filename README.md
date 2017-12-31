# Go语言经典笔试题

在线地址  [https://goquiz.github.io/](https://goquiz.github.io/)

支持在线编辑和编译运行!

更新中...

欢迎贡献更多高质量的Go语言笔试题，欢迎提交Pull Request

# 目录

<!--CATEGORY_BEGIN-->

[语言规范](https://goquiz.github.io/#lang-spec)
- [subslice引用](https://goquiz.github.io/#subslice-grow)
- [短声明](https://goquiz.github.io/#short-declairation)
- [interface是否等于nil的判断](https://goquiz.github.io/#interface-nil)
- [map存取数据](https://goquiz.github.io/#map-ok-idiom)
- [指针基本操作](https://goquiz.github.io/#pointer)
- [interface{}万能指针](https://goquiz.github.io/#empty-interface)
- [临时变量的指针](https://goquiz.github.io/#for-pointer)
- [break外层循环](https://goquiz.github.io/#label-break)
- [全局变量](https://goquiz.github.io/#global-varible)
- [defer栈](https://goquiz.github.io/#defer-closefile)
- [panic栈](https://goquiz.github.io/#defer-panic)
- [recover](https://goquiz.github.io/#defer-recover)
- [goroutine闭包参数](https://goquiz.github.io/#goroutine-closure)
- [receiver函数的"仿重载"](https://goquiz.github.io/#type-shadowing)
- [string转换为bytes](https://goquiz.github.io/#string-bytes)
- [map并发操作的临界区](https://goquiz.github.io/#mutex-map)
- [深度比较](https://goquiz.github.io/#DeepEqual)
- [goroutine让先](https://goquiz.github.io/#Gosched)
- [map的值的内部字段可修改性](https://goquiz.github.io/#map-addressing)
- [Struct的Slice的排序](https://goquiz.github.io/#sort-Slice)

[标准库和包](https://goquiz.github.io/#lib-pack)
- [init函数](https://goquiz.github.io/#init-import)
- [json反序列化](https://goquiz.github.io/#json-unmarshal)
- [map并发存取](https://goquiz.github.io/#sync-map)
- [utf8字符串长度](https://goquiz.github.io/#utf8-len)
- [正则表达式的高级模式](https://goquiz.github.io/#regex-multiline)
- [文本文件的行遍历](https://goquiz.github.io/#readline)
- [堆容器](https://goquiz.github.io/#container-heap)
- [context超时](https://goquiz.github.io/#context-withtimeout)
- [命令行参数解析](https://goquiz.github.io/#flag)
- [文本模版](https://goquiz.github.io/#text-template)
- [html模版](https://goquiz.github.io/#html-template)
- [http服务](https://goquiz.github.io/#http-server)
- [sql查询](https://goquiz.github.io/#sql-query)

[工具链](https://goquiz.github.io/#toolchain)
- [gcflags参数](https://goquiz.github.io/#gcflags)
- [benchmark次数变量](https://goquiz.github.io/#benchmark-N)
- [govendor路径覆盖](https://goquiz.github.io/#govendor-gopath)
- [GOMAXPROCS](https://goquiz.github.io/#GOMAXPROCS)
- [共享库](https://goquiz.github.io/#shared-object)

[内部原理](https://goquiz.github.io/#internals)
- [string内部结构](https://goquiz.github.io/#unsafe-bytes-string)
- [slice内部结构](https://goquiz.github.io/#unsafe-slice-array)
- [defer的性能成本](https://goquiz.github.io/#defer-overhead)
- [map的malloc阈值容量](https://goquiz.github.io/#map-malloc)
- [内存分配](https://goquiz.github.io/#runtime-newobject)
- [SSA](https://goquiz.github.io/#SSA)
- [AST](https://goquiz.github.io/#AST)
- [Go引导过程](https://goquiz.github.io/#bootstrap)
- [channel同步异步的区别](https://goquiz.github.io/#channel-sync-async)
- [析构函数](https://goquiz.github.io/#finalizer)
- [gc工作机制](https://goquiz.github.io/#gc)
- [goroutine休眠](https://goquiz.github.io/#goroutine-sleep)
- [内存分配过程](https://goquiz.github.io/#object-memory)
- [stack和heap内存分配判定](https://goquiz.github.io/#stack-heap)
- [goroutine暂停或者终止](https://goquiz.github.io/#stop-goroutine)


<!--CATEGORY_END-->

# 如何在本地Build

Prerequisites: `"gopkg.in/yaml.v2"`

    cd _build
    make html
    make readme
