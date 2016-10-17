# 关于Golang你需要知道的50件事 - Part 1 入门

### Source
- [50 Shades of Go: Traps, Gotchas, and Common Mistakes for New Golang Devs](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)

---

Go是一门简单有趣的语言, 不过和其他语言一样, 这门语言也有一些所谓的'坑'...大部分这些'坑'并不完全是Go的错, 有些'坑'是你从别的语言转换到Go时必然会遇到的陷阱, 而其他的则一般是因为你在写代码进行了错误的假设或者没有注意到细节.

如果你花了时间去学习这门语言的官方spec, wiki, mailing list讨论, 以及以及Rob Pike的一些非常好的文章和源码, 那么其实这些'坑'都是显而易见的. 不过不是每个人开始学习的道路都是一样的, 如果你是一个Go新手, 那么这里的内容将能大大减少你调试代码的时间.

这篇文章涵盖了Go 1.5及以下版本.

目录:

1. [左花括号不能另起一行](#case1)
2. [未使用的变量](#case2)
3. [未使用的引入](#case3)
4. [短变量声明只能在函数内使用](#case4)
5. [使用短变量声明重复声明变量](#case5)
6. [不能使用短变量声明给字段赋值](#case6)
7. [意外的变量覆盖](#case7)
8. [不能在没确定类型的情况下使用nil初始变量](#case8)
9. [使用nil Slice/Map](#case9)
10. [Map的容量](#case10)
11. [字符串不能是nil](#case11)
12. [数组类型的函数参数](#case12)
13. [Slice和Array使用range语句时的意外值](#case13)
14. [多维Array/Slice](#case14)
15. [访问Map中不存在的key](#case15)
16. [字符串是不可变的](#case16)
17. [字符串和字节Slice的转换](#case17)
18. [字符串和索引操作](#case18)
19. [字符串并不总是UTF8编码](#case19)
20. [字符串的长度](#case20)
21. [使用多行Slice/Array/Map字面量缺少逗号](#case21)
22. [log.Fatal和log.Panic可以比log做的更多](#case22)
23. [内置数据结构的操作并不是同步的](#case23)
24. [字符串使用range语句时的迭代值](#case24)
25. [使用for range语句来遍历一个Map](#case24)
26. [switch语句中的Fallthrough行为](#case25)
27. [自增和自减](#case27)
28. ['非'位操作符](#case28)
29. [运算符优先级](#case29)
30. [未导出的字段不进行编码](#case30)
31. [在还有活动的协程时退出程序](#case31)
32. [给没有buffer的Channel发送消息](#case32)
33. [向已经关闭的Channel发送消息会导致Panic](#case33)
34. [使用nil Channel](#case34)
35. [带有接收者的方法并不能改变消息的原值](#case35)

---

<a id="case1" name="case1"/>
### 左花括号不能另起一行

在大部分使用花括号的语言里你可以自由选择放置它们的位置, 但是Go不一样, 因为Go语言的是没有分号的, 左花括号换行会导致Go编译器的分号自动插入机制(JS也有类似的机制)出错.

#### Fails:

    package main

    import "fmt"

    func main()  
    { //error, can't have the opening brace on a separate line
        fmt.Println("hello there!")
    }

#### Compile Error:

> /tmp/sandbox826898458/main.go:6: syntax error: unexpected semicolon or newline before {

#### Works:

    package main

    import "fmt"

    func main() {  
        fmt.Println("works!")
    }

---

<a id="case2" name="case2"/>
### 未使用的变量

如果你在代码中存在没有使用的变量是会导致异常进而编译失败的, 你必须使用你在函数体内声明的变量, 但是如果一个变量是全局变量就不会有这样的问题, 同时未使用的函数参数也不会报错.

当然仅仅是给一个未使用变量赋值仍然是不够的, 必须要使用这个变量才能编译通过.

#### Fails:

    package main

    var gvar int //not an error

    func main() {  
        var one int   //error, unused variable
        two := 2      //error, unused variable
        var three int //error, even though it's assigned 3 on the next line
        three = 3

        func(unused string) {
            fmt.Println("Unused arg. No compile error")
        }("what?")
    }

#### Compile Errors:

> /tmp/sandbox473116179/main.go:6: one declared and not used /tmp/sandbox473116179/main.go:7: two declared and not used /tmp/sandbox473116179/main.go:8: three declared and not used

#### Works:

    package main

    import "fmt"

    func main() {  
        var one int
        _ = one

        two := 2 
        fmt.Println(two)

        var three int 
        three = 3
        one = three

        var four int
        four = four
    }

当然还有一个方案就是删除或注释掉未使用的变量.

---

<a id="case3" name="case3"/>
### 未使用的引入

如果引入一个包但是又没有使用这个包导出的任意变量或者函数的话, 代码是会编译失败的.

如果你的确需要引入一个包又不实用它, 那么可以给它一个空的标志```_```以避免编译错误. 这个空标志用来导入一个包以获取它的副作用.

#### Fails:

    package main

    import (  
        "fmt"
        "log"
        "time"
    )

    func main() {  
    }

#### Compile Errors:

> /tmp/sandbox627475386/main.go:4: imported and not used: "fmt" /tmp/sandbox627475386/main.go:5: imported and not used: "log" /tmp/sandbox627475386/main.go:6: imported and not used: "time"

#### Works:

    package main

    import (  
        _ "fmt"
        "log"
        "time"
    )

    var _ = log.Println

    func main() {  
        _ = time.Now
    }

当然还有一个方案就是删除或注释掉未使用的引入, [goimports](https://github.com/bradfitz/goimports)这个包就是帮你完成这个任务的.

---

<a id="case4" name="case4"/>
### 短变量声明只能在函数内使用

#### Fails:

    package main

    myvar := 1 //error

    func main() {  
    }

#### Compile Error:

> /tmp/sandbox265716165/main.go:3: non-declaration statement outside function body

#### Works:

    package main

    var myvar = 1

    func main() {  
    }

---

<a id="case5" name="case5"/>
### 使用短变量声明重复声明变量

你不能在一个代码块里重复声明变量, 但是可以在```:=```左侧至少有一个新变量的情况下使用短变量声明重复声明一个已有变量.

#### Fails:

    package main

    func main() {  
        one := 0
        one := 1 //error
    }

#### Compile Error:

> /tmp/sandbox706333626/main.go:5: no new variables on left side of :=

#### Works:

    package main

    func main() {  
        one := 0
        one, two := 1,2

        one,two = two,one
    }

---

<a id="case6" name="case6"/>
### 不能使用短变量声明给字段赋值

#### Fails:

    package main

    import (  
        "fmt"
    )

    type info struct {  
        result int
    }

    func work() (int,error) {  
        return 13,nil  
    }

    func main() {  
        var data info

        data.result, err := work() //error
        fmt.Printf("info: %+v\n",data)
    }

#### Compile Error:

> prog.go:18: non-name data.result on left side of := 

就算有ticket去指出这个'坑'这个应该也不会变因为Rob Pike喜欢 :-)

你可以使用临时变量或者预先声明需要的变量然后使用标准赋值操作符.

#### Works:

    package main

    import (  
        "fmt"
    )

    type info struct {  
        result int
    }

    func work() (int,error) {  
        return 13,nil  
    }

    func main() {  
    var data info

    var err error
    data.result, err = work() //ok
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("info: %+v\n",data) //prints: info: {result:13}
    }

---

<a id="case7" name="case7" />
#### 意外的变量覆盖

短变量声明是如此方便(特别是对于从动态类型语言转过来的人)以至于容易让人以为这是赋值语句, 如果你在一个新的代码块里错误地使用了短声明语句, 其实是在一个新的闭包里声明了同名的变量, 不会对外部产生影响, 代码不一定会编译出错但是可能无法以预期的方式运行.

    package main

    import "fmt"

    func main() {  
        x := 1
        fmt.Println(x)     //prints 1
        {
            fmt.Println(x) //prints 1
            x := 2
            fmt.Println(x) //prints 2
        }
        fmt.Println(x)     //prints 1 (bad if you need 2)
    }

即使对一个经验丰富的Go开发者来说这也是一个常见的陷阱, 你可以使用vet命令找出代码中的这些问题, 但是vet默认没有执行任何被覆盖变量检测, 请确保带上了`-shadow`的flag:

    go tool vet -shadow your_file.go

注意vet命令并不会报告出所有的被覆盖变量, 使用[go-nyet](https://github.com/barakmich/go-nyet)以获得更强的被覆盖变量检测.

---

<a id="case8" name="case8" />
### 不能在没确定类型的情况下使用nil初始变量

nil可以作为接口, 函数, 指针, Map, Slice以及Channel这些类型的零值, 但是我们不能在没有指定类型的情况给变量赋nil, 编译器会因为不知道变量是否满足类型要求而报错.

#### Fails:

    package main

    func main() {  
        var x = nil //error

        _ = x
    }

#### Compile Error:

> /tmp/sandbox188239583/main.go:4: use of untyped nil

#### Works:

    package main

    func main() {  
        var x interface{} = nil

        _ = x
    }

---

<a id="case9" name="case9" />
### 使用nil Slice/Map

向nil Slice中添加项目是没问题的, 但是对Map进行同样的操作会导致运行时Panic.

#### Works:

    package main

    func main() {  
        var s []int
        s = append(s,1)
    }

#### Fails:

    package main

    func main() {  
        var m map[string]int
        m["one"] = 1 //error
    }

---

<a id="case10" name="case10" />
### Map的容量

你可以在创建Map的时候确定其容量, 但是你并不能对一个Map使用```cap()```函数.

#### Fails:

    package main

    func main() {  
        m := make(map[string]int,99)
        cap(m) //error
    }

#### Compile Error:

> /tmp/sandbox326543983/main.go:5: invalid argument m (type map[string]int) for cap

---

<a id="case11" name="case11" />
### 字符串不能是nil

这对那些曾经把nil赋给字符串变量的程序员来说是一个'坑'.

#### Fails:

    package main

    func main() {  
        var x string = nil //error

        if x == nil { //error
            x = "default"
        }
    }

#### Compile Errors:

> /tmp/sandbox630560459/main.go:4: cannot use nil as type string in assignment /tmp/sandbox630560459/main.go:6: invalid operation: x == nil (mismatched types string and nil)

#### Works:

    package main

    func main() {  
        var x string //defaults to "" (zero value)

        if x == "" {
            x = "default"
        }
    }

---

<a id="case12" name="case12" />
### 数组类型的函数参数

如果你是一个C/C++程序员, 那么对你来说数组其实是指针, 将数组作为参数传递其实是同一块内存的引用, 所以函数里对参数的操作也会影响外面的原值. 但是数组在Go里是```值类型```的, 所以将数组作为参数传递的时候会复制一份这个数组的内容.

    package main

    import "fmt"

    func main() {  
        x := [3]int{1,2,3}

        func(arr [3]int) {
            arr[0] = 7
            fmt.Println(arr) //prints [7 2 3]
        }(x)

        fmt.Println(x) //prints [1 2 3] (not ok if you need [7 2 3])
    }

如果你需要改变原始数组里的数据, 那么在传参的时候请使用数组指针.

    package main

    import "fmt"

    func main() {  
        x := [3]int{1,2,3}

        func(arr *[3]int) {
            (*arr)[0] = 7
            fmt.Println(arr) //prints &[7 2 3]
        }(&x)

        fmt.Println(x) //prints [7 2 3]
    }

另一个方案就是使用Slice, Slice作为参数传递的时候是传引用的.

    package main

    import "fmt"

    func main() {  
        x := []int{1,2,3}

        func(arr []int) {
            arr[0] = 7
            fmt.Println(arr) //prints [7 2 3]
        }(x)

        fmt.Println(x) //prints [7 2 3]
    }

---

<a id="case13" name="case13" />
### Slice和Array使用range语句时的意外值

这种情况会发生在你在其他语言中使用```for-in```或者```foreach```语句的时候, 而Go语言中的range语句是与众不同的: 它的第一个返回值是项目的索引, 第二个返回值是具体值.

#### Bad:

    package main

    import "fmt"

    func main() {  
        x := []string{"a","b","c"}

        for v := range x {
            fmt.Println(v) //prints 0, 1, 2
        }
    }

#### Good:

    package main

    import "fmt"

    func main() {  
        x := []string{"a","b","c"}

        for _, v := range x {
            fmt.Println(v) //prints a, b, c
        }
    }

---

<a id="case14" name="case14" />
### 多维Array/Slice

以下是创建多维数组或Slice的方法: 

Array:

    package main

    func main() {  
        x := 2
        y := 4

        table := make([][]int,x)
        for i:= range table {
            table[i] = make([]int,y)
        }
    }  

Slice:

    package main

    import "fmt"

    func main() {  
        h, w := 2, 4

        raw := make([]int,h*w)
        for i := range raw {
            raw[i] = i
        }
        fmt.Println(raw,&raw[4])
        //prints: [0 1 2 3 4 5 6 7] <ptr_addr_x>

        table := make([][]int,h)
        for i:= range table {
            table[i] = raw[i*w:i*w + w]
        }

        fmt.Println(table,&table[1][0])
        //prints: [[0 1 2 3] [4 5 6 7]] <ptr_addr_x>
    }

---

<a id="case15" name="case15" />
### 访问Map中不存在的key

大部分程序员都会认为这个操作会像很多其他语言一样返回一个```nil```, 其实这个操作返回的是Map中那个数据类型的```零值```, 当然, 如果零值就是nil的话的确是返回nil, 而其他类型则不一定了. 最可靠的方式是通过判断Map取值操作的第二个返回值.

#### Bad:

    package main

    import "fmt"

    func main() {  
        x := map[string]string{"one":"a","two":"","three":"c"}

        if v := x["two"]; v == "" { //incorrect
            fmt.Println("no entry")
        }
    }

#### Good:

    package main

    import "fmt"

    func main() {  
        x := map[string]string{"one":"a","two":"","three":"c"}

        if _,ok := x["two"]; !ok {
            fmt.Println("no entry")
        }
    }

---

<a id="case16" name="case16" />
### 字符串是不可变的

不要尝试去改变一个字符串中独立的字符, 结果必然是失败的, 因为字符串其实是一个只读的带有一些额外特性的字节Slice, 如果真的需要修改的话, 请先将其转换成字节Slice再进行处理.

#### Fails:

    package main

    import "fmt"

    func main() {  
        x := "text"
        x[0] = 'T'

        fmt.Println(x)
    }

#### Compile Error:

> /tmp/sandbox305565531/main.go:7: cannot assign to x[0]

#### Works:

    package main

    import "fmt"

    func main() {  
        x := "text"
        xbytes := []byte(x)
        xbytes[0] = 'T'

        fmt.Println(string(xbytes)) //prints Text
    }

注意这并不是真正意义上的修改字符串中字符的方式, 因为一个字符可能被存在多个字节中. 如果你的确需要修改一个字符串, 可以先把它转换成符号Slice, 当然即使这样一个字符也可能跨越多个符号, 比如一个带有音调的字符. 这些复杂以及可能存在歧义才使得Go语言的字符串表现为一个字节序列.

---

<a id="case17" name="case17" />
### 字符串和字节Slice的转换

当你把一个字符串转换成字节Slice或者反过来时, 一般你会得到一份原始数据的拷贝, 这个不同于别的语言里的转换操作, 并不是基于同样的底层原始数据产生新的Slice和数组.

当然Go已经对字符串和字节Slice互转提供了一些优化操作以避免额外的内存分配.

其中一个优化是在将字节Slice转换成字符串作为Map的索引的时候, 另一个则是将字符串转换成字节Slice使用```for range```的时候, 这两个情况并没有进行复制数据以避免额外的内存分配.

    package main

    import "fmt"

    func main() {
        data := make(map[string]int)
        data["test"] = 3

        str := "test"
        sbytes := []byte(str)
        fmt.Println(data[string(sbytes[0])])

        for i, v := range []byte(str) {
            fmt.Println("key: ", i, "value: ", v)
        }
    }

---

<a id="case18" name="case18" />
### 字符串和索引操作

给一个字符串使用索引取值得到的是一个字节值, 而不是像很多别的语言那样得到一个字符.

    package main

    import "fmt"

    func main() {  
        x := "text"
        fmt.Println(x[0]) //print 116
        fmt.Printf("%T",x[0]) //prints uint8
    }

如果你需要访问字符串中的特殊字符(比如Unicode符号), 可以使用for range语句, 官方的```unicode/utf8```以及```utf8string(golang.org/x/exp/utf8string)```包都是非常有用的, utf8string这个包甚至包含一个非常方便的```At()```方法, 当然另一个方案就是把字符串转换成字符Slice.

---

<a id="case19" name="case19" />
### 字符串并不总是UTF8编码

字符串的值并不是必须得是UTF8文本, 它们可以包含任意的字节, 唯一可以确定字符串是UTF8编码就是当时用字符串字面量的时候, 当然即使这样字符串里也可以通过escape的方式包含别的编码的文本.

可以通过```unicode/utf8```的```ValidString()```方法来判断一个字符串是否是UTF8格式的文本.

    package main

    import (  
        "fmt"
        "unicode/utf8"
    )

    func main() {  
        data1 := "ABC"
        fmt.Println(utf8.ValidString(data1)) //prints: true

        data2 := "A\xfeC"
        fmt.Println(utf8.ValidString(data2)) //prints: false
    }

---

<a id="case20" name="case20" />
### 字符串的长度

假设你是一个Python程序员, 那么你肯定写过像下面这样的代码:

    data = u'♥'  
    print(len(data)) #prints: 1  

当你把它转换成Go代码的时候结果可能让你惊讶:

    package main

    import "fmt"

    func main() {  
        data := "♥"
        fmt.Println(len(data)) //prints: 3
    }

这是因为内建的```len()```函数返回的是一个字符串里的字节数, 而不是像别的语言处理Unicode字符串一样返回的字符数量.

如果需要达到这样的效果请使用```unicode/utf8```包里的```RuneCountInString()```函数, 这个函数返回的字符串里Unicode符号的数量.

    package main

    import (  
        "fmt"
        "unicode/utf8"
    )

    func main() {  
        data := "♥"
        fmt.Println(utf8.RuneCountInString(data)) //prints: 1
    }

当然从技术层面来讲```RuneCountInString()```函数返回的也并不是字符的数量因为一个字符可能跨越多个Unicode符号.

    package main

    import (  
        "fmt"
        "unicode/utf8"
    )

    func main() {  
        data := "é"
        fmt.Println(len(data))                    //prints: 3
        fmt.Println(utf8.RuneCountInString(data)) //prints: 2
    }

---

<a id="case21" name="case21" />
### 使用多行Slice/Array/Map字面量缺少逗号

#### Fails:

    package main

    func main() {  
        x := []int{
        1,
        2 //error
        }
        _ = x
    }

#### Compile Errors:

> /tmp/sandbox367520156/main.go:6: syntax error: need trailing comma before newline in composite literal /tmp/sandbox367520156/main.go:8: non-declaration statement outside function body /tmp/sandbox367520156/main.go:9: syntax error: unexpected }

#### Works:

    package main

    func main() {  
        x := []int{
        1,
        2,
        }
        x = x

        y := []int{3,4,} //no error
        y = y
    }

这里注意到, 使用多行声明时, 最后一个元素也要带上逗号, 当然, 使用单行声明时这个逗号是可以省略的.

---

<a id="case22" name="case22" />
### log.Fatal和log.Panic可以比log做的更多

一般语言的log库通常会提供各个级别的log. 和其他语言log库不一样的是, Go内建```log```的```Fatal*()```和```Panic*()```方法不仅会打印log, 而且可以让程序直接终止.

    package main

    import "log"

    func main() {  
        log.Fatalln("Fatal Level: log entry") //app exits here
        log.Println("Normal Level: log entry")
    }

---

<a id="case23" name="case23" />
### 内置数据结构的操作并不是同步的

虽然Go已经有很多内建功能来原生地支持并发, 但是却并没有一个并发安全的数据结构. 所以你需要确保数据的改动是原子性的, 推荐使用协程和Channel来实现原子操作, 当然你也可以使用```sync```包如果它的确对你的应用有所裨益.

---

<a id="case24" name="case24" />
### 字符串使用range语句时的迭代值

索引值是第二个返回值中字符第一个字节的索引, 这并不是这个字符在字符串中的位置, 注意一个实际的字符可能又多个UTF8 rune组成, 当然如果你真的需要操作字符, 那么可以使用```norm(golang.org/x/text/unicode/norm)```包.

对字符串使用```for range```会尝试将字符串解释成UTF8文本, 这时所有无法被理解的内容会被转换成0xfffd rune(也就是Unicode replacement characters)而不是实际的值, 如果你有任意类型的数据存储在字符串变量里, 可以事先将其转换成字节Slice以获得真正被存储的值.

    package main

    import "fmt"

    func main() {  
        data := "A\xfe\x02\xff\x04"
        for _,v := range data {
            fmt.Printf("%#x ",v)
        }
        //prints: 0x41 0xfffd 0x2 0xfffd 0x4 (not ok)

        fmt.Println()
        for _,v := range []byte(data) {
            fmt.Printf("%#x ",v)
        }
        //prints: 0x41 0xfe 0x2 0xff 0x4 (good)
    }

---

<a id="case25" name="case25" />
### 使用for range语句来遍历一个Map

简单的说, 使用```for range```语句来遍历一个Map, 重新编译之后顺序是不确定的.

    package main

    import "fmt"

    func main() {  
        m := map[string]int{"one":1,"two":2,"three":3,"four":4}
        for k,v := range m {
            fmt.Println(k,v)
        }
    }

不过如果你使用[Go Playground](https://play.golang.org/), 你一般都会得到相同的结果, 因为除非有所改动, 不然你的代码并不会被重新编译.

---

<a id="case26" name="case26" />
### switch语句中的Fallthrough行为

```switch```语句中的```case```会默认执行完结束, 而不像其他一些语言会执行到下一个case条件.

    package main

    import "fmt"

    func main() {  
        isSpace := func(ch byte) bool {
            switch(ch) {
            case ' ': //error
            case '\t':
                return true
            }
            return false
        }

        fmt.Println(isSpace('\t')) //prints true (ok)
        fmt.Println(isSpace(' '))  //prints false (not ok)
    }

当然你可以在```case```中最后使用```fallthrough```语句来实现Fallthrough操作, 或者将若干条件写在一个case里以获得类似的效果. 

    package main

    import "fmt"

    func main() {  
        isSpace := func(ch byte) bool {
            switch(ch) {
            case ' ', '\t':
                return true
            }
            return false
        }

        fmt.Println(isSpace('\t')) //prints true (ok)
        fmt.Println(isSpace(' '))  //prints true (ok)
    }

---

<a id="case27" name="case27" />
### 自增和自减

很多语言都有自增自减操作符, 但是和别的语言不同的是, Go并没有操作符前置的版本, 而且不能把这两个操作符混用在别的语句中.

#### Fails:

    package main

    import "fmt"

    func main() {  
        data := []int{1,2,3}
        i := 0
        ++i //error
        fmt.Println(data[i++]) //error
    }

#### Compile Errors:

> /tmp/sandbox101231828/main.go:8: syntax error: unexpected ++ /tmp/sandbox101231828/main.go:9: syntax error: unexpected ++, expecting :

#### Works:

    package main

    import "fmt"

    func main() {  
        data := []int{1,2,3}
        i := 0
        i++
        fmt.Println(data[i])
    }

---

<a id="case28" name="case28" />
### '非'位操作符

很多语言使用```~```作为'非'位操作符, 但是Go重用了异或操作符```^```来达到这个目的.

#### Fails:

    package main

    import "fmt"

    func main() {  
        fmt.Println(~2) //error
    }

#### Compile Error:

> /tmp/sandbox965529189/main.go:6: the bitwise complement operator is ^

#### Works:

    package main

    import "fmt"

    func main() {  
        var d uint8 = 2
        fmt.Printf("%08b\n",^d)
    }

Go仍然使用了```^```作为异或操作符, 这可能会让一些人迷惑.

如果你愿意的话, 你可以使用一个异或操作符来实现一元的非操作符(比如0x02 XOR 0xff => NOT 0x02), 这也可以解释为什么重用异或操作符来表示取非操作.

Go还有一个特殊的AND NOT位操作符```&^```, 这让非操作符更让人困惑了, 这可以看作是为了不用括号实现```A AND (NOT B)```的一个hack.

    package main

    import "fmt"

    func main() {  
        var a uint8 = 0x82
        var b uint8 = 0x02
        fmt.Printf("%08b [A]\n",a)
        fmt.Printf("%08b [B]\n",b)

        fmt.Printf("%08b (NOT B)\n",^b)
        fmt.Printf("%08b ^ %08b = %08b [B XOR 0xff]\n",b,0xff,b ^ 0xff)

        fmt.Printf("%08b ^ %08b = %08b [A XOR B]\n",a,b,a ^ b)
        fmt.Printf("%08b & %08b = %08b [A AND B]\n",a,b,a & b)
        fmt.Printf("%08b &^%08b = %08b [A 'AND NOT' B]\n",a,b,a &^ b)
        fmt.Printf("%08b&(^%08b)= %08b [A AND (NOT B)]\n",a,b,a & (^b))
    }

---

<a id="case29" name="case29" />
### 运算符优先级

包括位擦除操作符(&^)在内, Go和别的语言一样拥有很多标准操作符, 但是操作符的优先级却不尽相同.

    package main

    import "fmt"

    func main() {  
        fmt.Printf("0x2 & 0x2 + 0x4 -> %#x\n",0x2 & 0x2 + 0x4)
        //prints: 0x2 & 0x2 + 0x4 -> 0x6
        //Go:    (0x2 & 0x2) + 0x4
        //C++:    0x2 & (0x2 + 0x4) -> 0x2

        fmt.Printf("0x2 + 0x2 << 0x1 -> %#x\n",0x2 + 0x2 << 0x1)
        //prints: 0x2 + 0x2 << 0x1 -> 0x6
        //Go:     0x2 + (0x2 << 0x1)
        //C++:   (0x2 + 0x2) << 0x1 -> 0x8

        fmt.Printf("0xf | 0x2 ^ 0x2 -> %#x\n",0xf | 0x2 ^ 0x2)
        //prints: 0xf | 0x2 ^ 0x2 -> 0xd
        //Go:    (0xf | 0x2) ^ 0x2
        //C++:    0xf | (0x2 ^ 0x2) -> 0xf
    }

---

<a id="case30" name="case30" />
### 未导出的字段不进行编码

Go中对结构体进行转码(json/xml/gob等等)时中不会包含以小写字母开头的字段, 所以重新解码的时候也会缺失这些字段的内容.

    package main

    import (
        "fmt"
        "encoding/json"
    )

    type MyData struct {  
        One int
        two string
    }

    func main() {  
        in := MyData{1,"two"}
        fmt.Printf("%#v\n",in) //prints main.MyData{One:1, two:"two"}

        encoded,_ := json.Marshal(in)
        fmt.Println(string(encoded)) //prints {"One":1}

        var out MyData
        json.Unmarshal(encoded,&out)

        fmt.Printf("%#v\n",out) //prints main.MyData{One:1, two:""}
    }

---

<a id="case31" name="case31" />
### 在还有活动的协程时退出程序

主协程并不会等待所有的协程结束, 这是新手一个常见的错误.

    package main

    import (  
        "fmt"
        "time"
    )

    func main() {  
        workerCount := 2

        for i := 0; i < workerCount; i++ {
            go doit(i)
        }
        time.Sleep(1 * time.Second)
        fmt.Println("all done!")
    }

    func doit(workerId int) {  
        fmt.Printf("[%v] is running\n",workerId)
        time.Sleep(3 * time.Second)
        fmt.Printf("[%v] is done\n",workerId)
    }

你将会看到:

    [0] is running 
    [1] is running 
    all done!

一个比较通用的解决方案是使用一个```WaitGroup```变量, 它将允许主协程等待所有工作协程完成, 如果你有一些带有信号处理机制的一些耗时很长的工作协程, 那么你最好手动给他们发送终止的信号. 另一个方案是关闭在工作协程中接收消息的Channel, 这可以把所有协程一次性全都结束. 

    package main

    import (  
        "fmt"
        "sync"
    )

    func main() {  
        var wg sync.WaitGroup
        done := make(chan struct{})
        workerCount := 2

        for i := 0; i < workerCount; i++ {
            wg.Add(1)
            go doit(i,done,wg)
        }

        close(done)
        wg.Wait()
        fmt.Println("all done!")
    }

    func doit(workerId int,done <-chan struct{},wg sync.WaitGroup) {  
        fmt.Printf("[%v] is running\n",workerId)
        defer wg.Done()
        <- done
        fmt.Printf("[%v] is done\n",workerId)
    }

运行结果将是这样:

    [0] is running 
    [0] is done 
    [1] is running 
    [1] is done

看上去主协程实在所有工作协程完成之后退出的, 然而你同时会看到这样的结果:

    fatal error: all goroutines are asleep - deadlock!

这看上去可不太好, 为什么会这样出现死锁呢? 看上去所有工作协程都退出了并且执行了```wg.Done()```, 程序应该可以工作才对.

其实这个死锁的发生是因为每个工作协程都是获得了一份原始```WaitGroup```变量的拷贝, 在工作进程中执行```wg.Done()```并没有影响到主协程中wg变量.

    package main

    import (  
        "fmt"
        "sync"
    )

    func main() {  
        var wg sync.WaitGroup
        done := make(chan struct{})
        wq := make(chan interface{})
        workerCount := 2

        for i := 0; i < workerCount; i++ {
            wg.Add(1)
            go doit(i,wq,done,&wg)
        }

        for i := 0; i < workerCount; i++ {
            wq <- i
        }

        close(done)
        wg.Wait()
        fmt.Println("all done!")
    }

    func doit(workerId int, wq <-chan interface{},done <-chan struct{},wg *sync.WaitGroup) {  
        fmt.Printf("[%v] is running\n",workerId)
        defer wg.Done()
        for {
            select {
            case m := <- wq:
                fmt.Printf("[%v] m => %v\n",workerId,m)
            case <- done:
                fmt.Printf("[%v] is done\n",workerId)
                return
            }
        }
    }

这样一来程序就能如预期的一般工作了.

---

<a id="case32" name="case32" />
### 给没有buffer的Channel发送消息

当我们声明一个Channel同时不带长度时, 也就是一个不带缓冲的Channel, 这时当消息被接收者处理时发送者并不会被阻塞住, 接收者可能并没有足够的时间来处理发送者接下来发送进来的信息, 当然这取决于你的程序里协程的具体运行环境.

    package main

    import "fmt"

    func main() {  
        ch := make(chan string)

        go func() {
            for m := range ch {
                fmt.Println("processed:",m)
            }
        }()

        ch <- "cmd.1"
        ch <- "cmd.2" //won't be processed
    }

---

<a id="case33" name="case33" />
### 向已经关闭的Channel发送消息会导致Panic

从一个已经关掉的Channel接收消息是安全的, 当从一个Channel接收的值是```false```代表已经没有数据可以接收了, 如果这个Channel带缓冲的话, 那么首先你会接收到缓冲好的数据, 知道Channel里为空才会接收到```false```.

但是向一个关掉的Channel发送消息是会导致Panic的, 这是新手常犯的一个错误, 他们可能认为发送消息和接收消息的行为应该一致.

    package main

    import (  
        "fmt"
        "time"
    )

    func main() {  
        ch := make(chan int)
        for i := 0; i < 3; i++ {
            go func(idx int) {
                ch <- (idx + 1) * 2
            }(i)
        }

        //get the first result
        fmt.Println(<-ch)
        close(ch) //not ok (you still have other senders)
        //do other work
        time.Sleep(2 * time.Second)
    }

当然, 避免这个情况出现的工作量可大可小, 取决于具体的使用场景, 不过无论如何, 你都应该避免向关掉的Channel发送消息.

上面那个有bug的示例可以通过使用一个特殊的传递结束信号的Channel来解决.

    package main

    import (  
        "fmt"
        "time"
    )

    func main() {  
        ch := make(chan int)
        done := make(chan struct{})
        for i := 0; i < 3; i++ {
            go func(idx int) {
                select {
                case ch <- (idx + 1) * 2: fmt.Println(idx,"sent result")
                case <- done: fmt.Println(idx,"exiting")
                }
            }(i)
        }

        //get first result
        fmt.Println("result:",<-ch)
        close(done)
        //do other work
        time.Sleep(3 * time.Second)
    }

---

<a id="case34" name="case34" />
### 使用nil Channel

发送消息给一个nil Channel(也就是不通过```make```声明的Channel)会导致程序死锁, 这可能让Golang新手非常疑惑, 尽管这是一个文档中明确定义的行为.

    package main

    import (  
        "fmt"
        "time"
    )

    func main() {  
        var ch chan int
        for i := 0; i < 3; i++ {
            go func(idx int) {
                ch <- (idx + 1) * 2
            }(i)
        }

        //get first result
        fmt.Println("result:",<-ch)
        //do other work
        time.Sleep(2 * time.Second)
    }

运行这个代码将会导致如下的错误:

> This behavior can be used as a way to dynamically enable and disable case blocks in a select statement.

不过这个方法的一个用处是可以动态的决定一个select里的case语句是否被执行.

    package main

    import "fmt"  
    import "time"

    func main() {  
        inch := make(chan int)
        outch := make(chan int)

        go func() {
            var in <- chan int = inch
            var out chan <- int
            var val int
            for {
                select {
                case out <- val:
                    out = nil
                    in = inch
                case val = <- in:
                    out = outch
                    in = nil
                }
            }
        }()

        go func() {
            for r := range outch {
                fmt.Println("result:",r)
            }
        }()

        time.Sleep(0)
        inch <- 1
        inch <- 2
        time.Sleep(3 * time.Second)
    }

---

<a id="case35" name="case35" />
### 带有接收者的方法并不能改变消息的原值

接收者作为函数参数和常规的函数参数一样, 如果是作为一个值声明的, 那么函数作用域中会得到一份原值的拷贝, 也就是说在函数中的操作并不会改变消息的原值除非接收者是一个Map/Slice并且你在改变其子项, 或者你所使用的接收者是指针.

    package main

    import "fmt"

    type data struct {  
        num int
        key *string
        items map[string]bool
    }

    func (this *data) pmethod() {  
        this.num = 7
    }

    func (this data) vmethod() {  
        this.num = 8
        *this.key = "v.key"
        this.items["vmethod"] = true
    }

    func main() {  
        key := "key.1"
        d := data{1,&key,make(map[string]bool)}

        fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items)
        //prints num=1 key=key.1 items=map[]

        d.pmethod()
        fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items) 
        //prints num=7 key=key.1 items=map[]

        d.vmethod()
        fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items)
        //prints num=7 key=v.key items=map[vmethod:true]
    }
