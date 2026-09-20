`Go`语言学习:

## 1.简单理解:

**`Go`是`Google`开发的一门工程型、面向开发的语言，简单、原生并发，限制自由度，把常见正确方案内置进语法，减少程序员出错的语言。**

设计目标就是高效、并发、适合网络服务、适合工程项目、减少出错，强调少一些语言机制，多一些工程上的直接性。

`Go`没有头文件，没有`#include`，没有命名空间，没有类，不用手动`free、delete`。

## 2.`Go`程序的基本结构

```go
package main//当前程序属于哪个包
如果一个
import "fmt"//导入fmt包
func main(){//程序入口
fmt.Println("hello")
}
```

`package main`——包声明：

​	`Go`的代码是按包组织的，包是代码模块化最小单元，相当C++里的命名空间+源码文件分组。

​	`package main`：只有包名为`main`的包才能编译成可执行程序。其他包比如`package raft`，只是库包，只能被别人`import`导入，不能单独运行。可以类比C++靠有没有`int main()`来判断是不是可执行程序。要满足两个条件才能运行:1.写`package main`，2.包里要有无参无返回的函数`func main`，`func main`是程序的入口。

`import "fmt"`——导入包

​	`import`作用：引入别的包里的代码，等价于`#include <stdio.h>`

​	包的访问规则：导入包的程序使用包内函数或变量要`包名.函数名`，如果包内函数大写开头比如`Println()`就可以导出，外部包可以访问，但如果是小写`println()`，只能在自己的包内部使用，外部访问不了，类似`c++`的`public、private`。

​	其他写法:`import _ "net/http/prof"`，匿名导入，只执行包的`init`函数，不用里面符号，如果导入了包但是完全不用会报错，所以有了这种导入写法。

## 3.变量

```go
var a int = 10
var a int//Go会给它默认值0
var name string = "alice"
a := 10
name := "alice"
```

`:=`叫做**短变量声明**，`Go`自动推断类型。

变量基本类型：

```go
int
int32
int64

uint
uint32
uint64

float32
float64

bool
string
```

`c++`的`int x`可能是未定义值，但是`go`会默认零值，`Go`规定变量没有显式初始化时，也有一个确定的默认值。

`Go`没有隐式类型转换，比如`int a、int64 a`之间的转换，必须写`a=int(b)`。

## 4.`fmt.Printf`

`%v`任意类型，`%T`类型

```c++
x := 10
fmt.Printf("value=%v\n", x)
fmt.Printf("type=%T\n", x)
```

`fmt.Println()`和`printf`的不同是它直接把东西打印出来：

```go
fmt.Println("name=", name, "age=", age, "height=", height)
```

## 5.`if`

`Go`的`if`不带括号，`Go`没有`while`，`for`充当`while`的作用：

```go
if ok {
}
for {
//无限循环
    break;
}
```

注意`Go`的一个语法细节：`else`必须和前面的`}`在同一行。

`Go`的`if`可以在判断之前声明变量！

```go
if age:=20;age>=18 {
fmt.Println("成年人")
}
```

## 6.`switch`

`Go`的`switch`和`C++`的有一个明显区别：

`Go`默认执行完一个`case`就自动退出`switch`，但是`C++`会继续执行下去。

如果`Go`想要继续执行下去可以明确写**`fallthrough`**:

```go
switch x {
case 1:
    fmt.Println("one")
    fallthrough
case 2:
    fmt.Println("two")
}
```

## 7.`range`

用来遍历容器：`slice/map/string/channel`，就相当于`for(auto t:times)`，但是`range`返回两个值：`i`索引，`v`对应元素拷贝

(1)遍历`slice`:比如现在有一个`slice`：

```go
nums:=[]int{10,20,30}
for i,value :=range nums {
fmt.Println(i,value)
}
```

这里的`i`是下标，`range`是元素，如果只要元素，不要下标可以:

```go
for _,value :=range nums {
fmt.Println(value)
}
```

`_`叫做**空白标识符**，意思就是这个值我不要

只要下标，不要元素可以:

```go
for i:= range nums {
fmt.Println(i)
}
```

`Go`规定不能定义了变量但是不用，否则会编译报错

！！！注意这里的`value`是拷贝过来的值，修改`value`不会修改原数组，要:

```go
for i:= range nums {
nums[i]+=100
}
```

(2)遍历`map`：

```go
m:=map[string]int{"a":1,"b":2}
for k,v := range m {
fmt.Println(k,v)
}
```

`map`的`range`返回：第一个返回`key`，第二个返回`value`拷贝，`map`遍历顺序不固定

(3)`range channel`

```go
ch :=make(chan int,3)
ch <- 1
close(ch)
for v:=range ch {
fmt.Println(v)
}
```

`range channel`：持续从`channel`接收数据，直到`channel`被`close`，循环自动退出。如果`channel`没有`close`，`range`会永久阻塞在这里，`goroutine`泄漏。

(4)`Go range`没有引用版本，`range`第二个返回永远是拷贝，想要原地修改只能用下标。

(5)`range`在循环开始时一次性拿到容器的长度，遍历`slice`时如果循环内`append`扩容，不会自动迭代新增元素。

## 8.函数

(1)最基本的函数

```go
func add(a int,b int) int {
return a+b
}
result:=add(10,20)
fmt.Println(result)
```

`func` 函数名 参数 返回值类型

多个参数也可以简写成:

```go
func add (a,b int)int{
}
```

(2)`Go`函数最重要的特性：多返回值

```go
func get()(int,int){
return 1,2
}
a,b := get()
```

为什么要设计多返回值？

`Go`非常喜欢用**结果+错误**的形式，通常最后一个返回值用来放错误`error`，6.824`RPC`全部遵守这个规范，可以用`_`忽略错误。`Go`的错误处理风格时显式检查错误，而不是大量依赖异常。

(3)`Go`的`error`和`C++`的`errno`

`C`没有多返回值，系统调用失败时函数返回`-1`，只能用全局变量`errno`记录错误码，但是这个`errno`全局共享，多线程下很容易被覆盖，这是`C`的错误码的缺点：不和本次调用绑定，多线程并发时很容易被覆盖。

`Go`的`error`不是全局变量!`error`是一个接口类型:

```go
type error interface {
Error()string
}
```

每次函数出错返回一个独立的`error`对象，`nil`代表没有错误，通常先判断`err`是否非`nil`，再使用返回结果。

```go
func div(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("除数不能为0")
    }

    return a / b, nil
}
```

有错误就`fmt.Errorf()`返回错误码，没有错误就返回`nil`.

(4)一个函数多个返回值，我不想要怎么办？

直接用`_`前面说过的空白标识符

(5)命名返回值

```c++
func add(a, b int) (result int) {
    result = a + b
    return
}
```

给返回值起了名字`result`，`return`之后自动返回`result`

(6)函数也是一种值

`Go`可以把函数赋值给变量：

```go
func add(a,b int)int{
return a+b
}
f:=add
result:=f(1,2)
```

(7)匿名函数

还可以直接写一个没有名字的函数

```go
f:=func(a,b int)int{
return a+b
}
```

## 9.`value,ok`

比如`map`:

```go
users:=map[string]int{
"alice":20,
"bob":21,
}
```

我们去`age:=users["alice"]`，如果想知道`alice`到底存在不存在，避免拿到个空值，可以：

```go
age,ok :=users["alice"]
```

存在`ok=true`，不存在`ok=false`。

## 10.数组

(1)也可以让`Go`自动计算长度

```go
a:=[5]int{1,2,3,4,5}
a:=[...]int{1,2,3,4,5}
```

`...`表示你放几个元素我就自动推断数组长度

(2)数组长度是类型的一部分，2个`int`数组长度不同类型就不同

## 11.`slice`切片

类似`vector`，定义：

```go
var s []int
```

`[5]int`固定长度数组，`[]int`长度可以变化`slice`

创建`slice`:

```go
s:=[]int{1,2,3}
```

添加元素用`append`:

```go
s:=[]int{1,2,3}
s=append(s,4)
s=append(s,5)
s==append(s,6)
//变成{1,2,3,4,5,6}
```

`slice`有两个非常重要的概念：`len(s) cap(s)`

`len`计算当前有多少个元素，`cap`计算容量

切片操作：

```go
s:=[]int{1,2,3,4,5}
```

取`s[1:4]`，得到2,3,4，`s[1:4]`是左闭右开，也就是`1<=index<4`，注意这里有一个非常重要的点，如果我们`b:=a[1:3]`，此时`b`和`a`共享底层数组：

```go
a
┌────┬────┬────┬────┐
│ 10 │ 20 │ 30 │ 40 │
└────┴────┴────┴────┘
       ↑    ↑
       └────┘
         b
```

如果此时修改`b[0]=999`，`a`也会被修改

## 12.`make`

经常会看到`s:=make([]int,3,10)`，三个参数分别是`[]int类型`，`3 len`，`10 cap`，也可以`s:=make([]int,3)`，此时`len=cap=3`。

## 13.`map`

可以类比成`C++`的`unordered_map<string,int>`，`Go`就是：

```go
ages:=make(map[string]int)
ages["alice"]=20
ages["bob"]=10
value,ok:=ahes["bob"]
if ok{
    fmt.Println(value)
}else{
    fmt.Println("不存在")
}
```

删除元素：

```go
delete(ages,"alice")
```

遍历`map`:

```go
for key,value:=range ages{
fmt.Println(key,value)
}
```

## 14.指针

和`C++`的指针差不多

声明：`var p *int`

和`C++`指针的最大区别就是`Go`没有指针运算，不允许`p++`，`Go`的指针主要用于间接访问数据，修改数据，避免不必要的数据复制。

`nil`指针：

```go
var p*int//这时p=nil
```

## 15.`struct`

`Go`并没有类这个思想，基本靠结构体`struct`来完成。

```go
s:=student{
Name:"alice",
age:20,
}
fmt.Println(s.name)
fmt.Println(s.age)
```

`Go`分布式系统代码的核心就是`struct+method+interface+goroutine+mutex`.

`Go`所有结构体字段自动零值初始化，`C`结构体变量不会自动清零，`Go`统一`.`访问结构体成员变量，不论指针还是结构体对象，编译器会自动判断是不是指针。

结构体是值类型，赋值的时候会完整拷贝整个结构体，如果结构体很大，传值开销大，而且是副本，所以`Raft`代码几乎全部用`*Raft`指针。

(1)结构体绑定方法

`Go`的方法不属于结构体内部，`struct`的函数方法`method`就是一个绑定了额外的第一个参数(接收者)的普通函数，这个接收器其实就相当于`C++`里类成员函数里的`this`，接收器有两种：值接收器/指针接收器。

```go
// 指针接收器：(rf *Raft)，操作原对象，Raft绝大多数方法都是这种
func (rf *Raft) GetTerm() int {
    rf.mu.Lock()
    defer rf.mu.Unlock()
    return rf.currentTerm
}

// 值接收器：(rf Raft)，传入结构体拷贝，修改不会影响原对象
func (rf Raft) GetTermCopy() int {
    return rf.currentTerm
}
```

```go
type Raft struct {
age int
}
func (rf Raft) add1(){
rf.age+=1
}
func (rf*Raft) add2(){
rf.age+=1
}
func main(){
rf:=Raft(0)
rf.add1()
fmt,Println(rf.age)//这里结果还是0,因为调用的方法函数是值接收，修改的是拷贝的那一份数据
rf.add2()
fmt.Priintln(rf.age)//这里结果是1,因为调用的方法函数是指针接收，通过指针直接修改的是原数据
}
```

和普通函数的区别:

```go
//普通函数
func add(rf*Raft){
rf.age++
}
//方法
func (rf*Raft)add(){
rf.age++
}
```

其实就是一个是传参数，去计算，一个是把这个函数绑定到这个类型，唯一的区别就是：调用语法+绑定到类型上，可以满足接口`interface`。

(2)大小写导出规则

结构体字段首字母大写：可以导出到别的包，别的包可以访问；

小写：包内私有，外部包完全看不见

```go
type Raft struct {
    CurrentTerm int // 大写！导出，labrpc序列化RPC参数必须读到
    votedFor    int // 小写！只有raft包内部可以访问，RPC无法序列化
}
```

(3)结构体和锁

`sync.Mutex`是结构体，不能拷贝：

`Mutex`里面保存锁状态，如果拷贝`Mutex`会复制一份无效锁，程序出现`bug`，所以`Raft`结构体里放`sync.Mutex`，`Raft`必须全程传指针，绝对不能传值拷贝`Raft`。

只要结构体包含`sync.Mutex、sync.Cond`，永远传指针，禁止值拷贝

## 16.`interface`

(1)`Go`的`interface`是方法的集合，只要某个类型拥有接口全部方法，就自动实现这个接口，不用显示声明`implements`。接口只规定方法，不实现方法，交给类型实现。

自动实现`Handler`:编译器自动检查，只要你的类型拥有接口里的规定的全部方法(签名完全一致)，就认为这个类型实现了这个接口，不需要你写任何代码声明"我实现了`Handler`"，只有类型满足接口的所有方法时，这个类型才可以赋值给接口变量。

作用大白话：接口定义一套方法约定，不提供实现，作用就是让不同的具体类型可以被同一套代码统一处理，来实现多态。

```go
        Animal
          ↑
       接口类型
       /      \
     Dog      Cat
```

```go
//接口
type Speaker interface {
Speak()string
}
//类型一
type Dog struct{}
func(d*Dog)Speak()string{
return "汪汪"
}
//类型二
type Cat struct{}
func (c*Cat)Speak()string {
return "喵喵"
}
//这个函数接收Speaker接口变量，外部主要是调用这个函数
func Make(s Speaker){
fmt.Println(s.Speak())
}
func main(){
    //把*Dog赋值给接口变量s1
var s1 Speaker=&Dog{}
    //ba1*Cat赋值给接口变量s2
var s2 Speaker=&Cat{}
Make(s1)
Make(s2)
}
```

函数根本不关心:

```go
你是不是Gog 
你是不是Cat
```

它只关心：**你有没有`Speak()`这个方法**

调用者只需要知道`Handler`，不需要关心底层是什么，这就是**接口的核心思想**：面向能力，而不是面向具体实现。

(2)值接收者和指针接收者在`interface`上的区别

```go
type Handler interface {
add()int
}
```

如果是值接收者，`func(a A)add()int`：

类型`A`：拥有`add`这个方法，`A`实现`Handler`

类型`*A`:`Go`有语法糖，指针可以自动接引用调用值接收者方法，`*A`也实现`Handler`

不管传`A`还是传`&A`，都可以找到`add()`。

如果是指针接收者：`func (b*B)add()int`：

类型`*B`:拥有`add`，`*B`实现`Hnalder`

类型`B`:没有`add`方法，`B`不满足`Handler`

(3)空接口`interface`

空接口定义：里面没有任何方法

```go
var x interface{
}
```

契约：不需要实现任何方法，所以`Go`里面所有类型，全部自动实现空接口。

空接口能干什么？

空接口变量可以存任意类型的值，举个例子，存进去是万能盒子，这个盒子内部存了类型信息+值，拿出来需要类型断言，编译器会告诉你是什么类型

```go
var x interface{}=10
v,ok := x.(int)
if ok {
fmt.Println(v)
}
```

`type switch`多类型判断:

```go
switch val := x.(type){
case int:
case string:
defalut:
}
```

(4)和`C++`的一个核心区别

`C++`子类要继承可能必须显式写`implements`，但是`Go`只要这个类型满足接口的所有规则所有函数类型定义，就满足接口并且自动实现这个接口。

(5)`nil interface`

接口变量在`Go`内部存了两个东西：动态类型、动态值，只有动态类型==`nil`并且动态值==`nil`，整个接口变量才等于`nil`

```go
var i interface{}
fmt.Println(i==nil)//true
```

```go
type Raft struct{
age int
var rf*Raft=nil//rf是*Raft类型的指针，值是nil
var i interface{}=rf
    //i内部：type=*Raft,data=nil
fmt.Println(i==nil)//false
}
```

## 17.`Package`

可以理解成一个代码模块，比如聊天室当中我们的`client`文件夹，一个目录通常对应一个`Package`，同一个目录下的多个`.go`文件共同组成这个`Package`包。

这里有一个非常重要的规则：

`func Add()`和`func add()`一个大写一个小写的区别就是`Add`可以被其他`package`使用，`add`只能在当前`package`里被使用，遇到大写开头的函数时要意识到这个函数可能就是专门提供给其他`package`使用的。`struct`也遵循这个规则。

## 18.`Go Module`

运行一个简单的`Go`程序我们就是`go run main.go`，但是一个稍微正式一点的项目通常会有`go.mod`，我们创建一个`Go Module`的时候就会生成这个文件：

```go
mkdir hello-go
cd hello-go
go mod init hello-go
```

会产生:

```go
hello-go/
└── go.mod
```

`go.mod`是干什么的？

可以暂时把它理解成：这个`Go`项目的身份证+依赖管理文件，它告诉`Go`:这个项目叫什么，使用哪个`Go`版本，依赖哪些其他模块，可能你的项目需要使用一个第三方库：`github.com/xxx/redis`，那么`go.mod`中可能出现:

```go
module chatroom
go 1.24
require (
    github.com/redis/go-redis/v9 v9.x.x
)
```

如果项目里:

```go
import "github.com/xxx/xxx"
```

但是`go.mod`里还没有记录这个依赖，执行:

```go
go mod tidy
```

`Go`会帮助分析代码，找出依赖，下载需要的模块，更新`go.mod`，更新`go.sum`。

## 19.必须认识的命令

```go
go run .	//运行当前项目
go build	//生成可执行文件
go mod init xxx	//初始化modle
go mod tidy	//整理依赖
go test	//测试
```

## 20.`Goroutine`

(1)`Goroutine`是`Go`中用来执行并发任务的轻量级执行单元。`Go`轻量协程。

```go
go task()
```

的意思就是启动一个`Goroutine`，让它去执行`task()`。`main`函数结束，所有`Goroutine`结束，程序退出，不会去等待`Goroutine`，所以我们以后需要用到`sync.WaitGroup`来等待`Goroutine`来完成。

(2)`Goroutine`和`std::thread`有什么区别？

`std::thread`是操作系统内核管理的线程；`goroutine`是`Go runtime`管理的轻量任务，复用`OD`线程。

网络编程`goroutine`写法：

```go
func handleClient(conn net.Conn) {
    // 处理客户端
}

func main() {
    listener, _ := net.Listen("tcp", ":8888")

    for {
        conn, _ := listener.Accept()

        go handleClient(conn)
    }
}
```

分维度对比：

**1.栈**

`std::thread`:创建时分配固定大小栈，通常8`MB`

`goroutime`:初始栈很小，2`KB`，按需扩容缩容，可以轻松开出成千上万个`goroutine`

传统线程通常需要比较大的栈空间和操作系统资源，而`goroutine`的设计目标就是：**让你可以很低成本地创建大量并发任务。

**2.调度主体**

`std::thread`:由操作系统内核调度，切换线程的时候需要陷入内核，上下文切换开销大

`goroutine`:由`Go`运行时`runtime`在用户态调度，`goroutine`之间切换大多都在用户态，不进内核，开销低，只有`goroutine`发生系统调用的时候才会进入内核。

**3.任务和操作系统线程绑定**

`std::thread`:1:1绑定，一个任务一个`OS`线程，任务阻塞，`OS`线程跟着阻塞，资源被占用

`goroutine`:`M:N`映射，多个`goroutine`复用少量`OS`线程，当某个`goroutine`阻塞在系统调用的时候，`runtime`把这个`goroutine`剥离，`runtime`找另一个空闲`goroutime`，继续跑其他就绪`goroutime`，这样一个`goroutine`卡住不会阻塞其他`goroutine`。

来举个例子：

`G1`这个`goroutine`当前跑在`M1`这个`OS`线程上，`M1`绑定`P`这个调度器，当`G1`系统调用或者其他原因阻塞的时候，`M1`跟着`G1`一起进入内核，被`OS`挂起阻塞，`P`不能卡在这等`M1`，`P`和`M1`解绑，然后去找另一个空闲的`OS`线程`M2`，去绑定这个`M2`，然后从就绪队列拿出`G2`，让`M2`执行`G2`的代码。

所以不是原来阻塞的`M1`继续跑`G2`，原来`M1`已经卡斯在内核里了，`P`换了一个新`M`来跑别的`G`。

**4.创建开销**

`std::thread`:很重，内核要创建`TCB`，分配栈，创建成本高，不适合高频大量创建销毁

`goroutine`:极轻，`runtime`层面创建，不是内核创讲

**5.并发竞争**

这里二者的共享内存并发问题完全一样，都是多个执行单元读写同一个共享变量的时候都要加锁保护：

`std::thread`:`std::mutex`、`std::condition_variable`

`goroutinr`:`sync.Mutex`7、`sync.Cond`

**6.生命周期**

`std::thread`:主线程退出，进程还可以继续跑子线程，这取决于`detach/join`

`goroutine`:`main`函数`return`之后整个进程直接退出，所有`goroutine`都退出，不等待。

**7.并发理念**

`std::thread`:通过共享内存进行通信

`goroutine`:通过通信共享内存

(3)`goroutine`的写法

```go
go task()
go task(10)
go func(){//匿名函数，()是立即调用
fmt.Println("12")
}()
```

## 21.`channel`

(1)**定义**:`goroutine`负责干活，`channel`负责让`goroutine`之间传递数据协作。

`channel`就是`goroutine`之间传递数据的管道。

```go
ch := make(chan int)
```

这里创建了一个 `chane int`，意思是这个`channel`里面传递的是`int`。

(2)向`channel`发送数据

```go
ch <- 10
```

`<-`代表数据流动方向，这句意思就是把10放进`ch`。

从`channel`接收数据：

```go
value := <-ch
```

意思从`ch`中取出一个数据，保存到`value`。

所以完整流程就是`channel->value`。

(3)`channel`默认是无缓冲`channel`

```go
func main() {
    ch := make(chan int)

    ch <- 100

    fmt.Println("hello")
}
```

如果发送数据时没有其他`goroutine`接收那么发送方阻塞，上面的例子，发送方阻塞，`hello`不会正常打印。

但是`channel`也有缓冲:

```go
ch := make(chan int,3)
```

这就是容量为3的缓冲`channel`。当已经发了3个数据，要继续发第4个数据的时候就阻塞了。

所以：**无缓冲**更强调同步

​	   **有缓冲**可以让发送方暂时不用等待接收方

(4)`channel`不只是传数据，它还能做同步：

```go
func worker(done chan bool) {
    fmt.Println("worker工作中")

    // 模拟工作
    time.Sleep(time.Second)

    fmt.Println("worker完成")

    done <- true
}
func main() {
    done := make(chan bool)

    go worker(done)

    <-done

    fmt.Println("main继续")
}
```

(6)`channel`可以传什么？

`int`、`string`、`struct`、`channel`、指针

(7)`channel`可以传递多个数据

```go
func worker(ch chan int) {
    for i := 0; i < 5; i++ {
        ch <- i
    }
}
func main() {
    ch := make(chan int)

    go worker(ch)

    for i := 0; i < 5; i++ {
        value := <-ch
        fmt.Println(value)
    }
}
```

这就是非常经典的生产者-消费者模型：

```go
生产者                消费者

worker                main
  │                     │
  │ 0 ────────────────→ │
  │ 1 ────────────────→ │
  │ 2 ────────────────→ │
  │ 3 ────────────────→ │
  │ 4 ────────────────→ │
```

(8)`channel`关闭

用`close(ch)`，表示不会再往这个`channel`发送数据了，并不是说销毁了`channel`，通常由发送方来`close`。

```go
for value := range ch {
    fmt.Println(value)
}
```

会一直接收，直到`channel`被关闭。

## 22.`waitgroup`

`waitgroup`用来等待多个`goroutine`完成，可以把`waitgroup`理解成一个`goroutine`的计数器。

```go
var wg sync.WaitGroup

wg.Add(1)

go func() {
    defer wg.Done()

    fmt.Println("worker")
}()

wg.Wait()
```

`wg.wait()`就是等`goroutine`完成，等待计数器归零。

核心函数:

1.`wg.Add(int)`：修改计数器，任务开始前，增加任务数量，告诉`waitgroup`我要等待几个任务。

2.`wg.Done()`:计数器减1,一般`defer wg.Done()`放在协程函数开头，保证函数退出一定计数减1，一个任务完成，计数器-1

3.`wg.wait()`:阻塞当前`goroutine`，直到内部计数器为0，如果还有任务没有完成就阻塞等待

```go
package main
import (
    "fmt"
    "sync"
)
func main() {
    var wg sync.WaitGroup

    for i := 0; i < 3; i++ {
        wg.Add(1) // 每开一个协程，计数+1
        go func(peerID int) {
            defer wg.Done() // 协程退出，计数-1
            fmt.Printf("向peer %d 发送RPC\n", peerID)
        }(i) // 传入i拷贝，避开闭包循环变量陷阱！Raft里发投票RPC必踩坑
    }

    wg.Wait() // 卡住main，等3个协程全部Done，计数器=0才继续
    fmt.Println("所有RPC请求全部发起完成")
}

```

!!!`Add`必须在启动`goroutine`之前执行，`wait()`可以多次调用，计数器为0时直接返回。

!!!`waitgroup`不能拷贝，如果把`wg`当作函数参数传值，会复制一份`waitgroup`对象，原`wg`和副本完全独立。必须传指针`func f(wg *sync.waitgroup)`。

```go
// 错误示范：拷贝wg
func bad(wg sync.WaitGroup) {
	defer wg.Done()
}
// 正确：传指针
func good(wg *sync.WaitGroup) {
	defer wg.Done()
}

```

```go
WaitGroup
    ↓
主要解决“等几个 Goroutine 完成”

Channel
    ↓
主要解决“Goroutine 之间传递数据/同步”
```

## 23.`mutex`

既然有`goroutine`并发就肯定会有竞态条件，之前学习`c++`我们都知道`a++`这个操作不是原子的，可能两个`goroutine`同时执行`a++`最终结果并不是2,是1。

`go`里的锁：`var mu sync.Mutex`

```go
mu.Lock()
a++
mu.Unlock()
```

同一时间只有`goroutine`可以进入临界区。临界区就是加锁和解锁中间这段代码。

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var counter int
    var mu sync.Mutex
    var wg sync.WaitGroup

    wg.Add(100)

    for i := 0; i < 100; i++ {
        go func() {
            defer wg.Done()

            mu.Lock()
            counter++
            mu.Unlock()
        }()
    }

    wg.Wait()

    fmt.Println("counter =", counter)
}
```

推荐`defer mu.Unlock()`函数退出自动解锁，避免忘记解锁。

`go`还有一个`sync.RWMutex`读写锁，它允许多个`goroutine`同时读，但是写的时候必须独占。比如:

```go
mu.RLock()
value:=data["name"]
mu.RUnlock()

mu.Lock()
data["name"]="alice"
mu.Unlock()
```

## 24.`Context`

(1)**概念**：`context`主要用来控制一个任务的生命周期：取消、超时、截止时间，以及在调用链中传递请求级信息。可以理解成`context`来控制这个任务还要不要继续。

用处:比如客户端给服务端发消息，服务端去数据库查东西然后返回，但是这个时候客户端断开了也就是现在其实服务器没有必要去查数据了，这时就需要用到`context`，通过`context`去取消，结束相关函数。如果是`c++`就只能自己写一个原子`boo`l，每个线程循环轮询这个`bool`。

(2)`context`长什么样

`go`标准库:

```go
import "context"
```

**核心接口**：

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool) // 返回截止时间
    Done() <-chan struct{}  // 只读channel：取消时这个chan会被close
    Err() error             // 为什么取消：超时 / 手动cancel
    Value(key any) any      // 传附加数据（尽量少用）
}

```

`Done()`:最核心，`ctx.Done()`返回一个`<-chan struct{}`，也就是说`ctx.Done()`本质上是一个`channel`，这个`channel`可以理解成`context`的取消通知通道，只要`context`取消，这个`channel`就关闭，`goroutine`里`select {case <-ctx.Done():...}`就能捕获信号。

一旦取消，`ctx.Err()`返回错误，`Done()`永久关闭，`context`只能取消一次，取消后不可逆。

(3)**4个创建`context`的函数**:

`context.Bckground()`:根`context`，永远不会取消，没有超时，一般`main`、初始化用，作为父顶层`context`.

`context.TODO()`:不知道该用啥`context`的时候临时占位，代码重构标记，和`Background`功能一样

上面这两个是空上下文，不能取消，只能当父节点派生子`ctx`。

`context.WithCancel(parent Context) (ctx Context,cancel CancelFunc)`:派生一个子`ctx`，返回`cancel()`函数和`ctx`，调用这个`cancel()`就会把这个`ctx`取消，并且所有子子孙孙`context`全部跟着取消。

```go
ctx, cancel := context.WithCancel(context.Background())
go func() {
    select {
    case <-ctx.Done():
        fmt.Println("收到取消信号，退出")
    }
}()
// ...业务...
cancel() // 手动触发取消，所有继承这个ctx的goroutine都会收到Done信号
```

一般`defer cancel()`避免忘记导致资源泄漏。

`context.WithTimeout(parent Context,duration time.Duration) (ctx Context,cancel CancelFunc)`:超时自动取消，也可以手动提前`cancel`，通常用来设置超时时间然后取消。

```go
// 3秒超时
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()
select {
case <-time.After(5 * time.Second):
    fmt.Println("任务完成")
case <-ctx.Done():
    fmt.Println("超时！", ctx.Err()) // context deadline exceeded
}

```

`context.WitDeadline(parent Context,deadline time.Time)`:指定一个未来的时间点，到点自动取消，本质和`Timeout`一样，知识参数形式不同。

`withcancel`就是自己取消，`withtimeout`设置超时时间，`withdeadline`设置超时时间点。

举一个简单的`context`例子：

```go
package main

import (
    "context"
    "fmt"
)

func worker(ctx context.Context) {
    fmt.Println("worker开始")

    select {
    case <-ctx.Done():
        fmt.Println("worker被取消")
        return

    default:
        fmt.Println("worker继续执行")
    }
}

func main() {
    ctx := context.Background()

    worker(ctx)
}
```

(5)父子树机制：

`context`是一棵树：父`ctx`取消，所有子`ctx`全部取消。子`ctx`取消不会影响父`ctx`。

`context`一层一层传，所有函数共享同一个`context`，上层调用`cancel`下面所有`cancel`，这叫取消信号向下传播。

如果`RPC`超时，`context`取消，服务器相关任务停止。

(6)`ctx`只传递取消/超时信号，不传递锁、业务变量。

`cancel`一定要`defer`。

`Done channel`是只读`<-chan struct{}`，只能读，不能往里面`send`。

## 25.`Go`常用高级特性

(1)`defer`

之前讲过，用途就是资源释放，和`c++`的`RAII`很类似。

`defer`是后进先出：

```go
func test() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")

    fmt.Println("开始")
}
```

输出：

```go
开始
3
2
1
```

因为`defer`是一个栈。

(2)`panic()`

```go
func main(){
panic("出错了")
}
```

正常错误就用`error`进行处理，严重异常用`panic`进行处理。

`recover()`可以捕获`panic`，从`panic`中恢复，但是必须在`defer`中使用:

```go
func safe() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("捕获到 panic:", r)
        }
    }()

    panic("出问题了")
}

func main() {
    safe()

    fmt.Println("程序继续运行")
}
```

流程:

```go
safe()
  │
  ├── defer 注册恢复函数
  │
  ├── panic()
  │
  ↓
开始执行 defer
  │
  ├── recover()
  │
  ↓
恢复
  │
  ↓
safe() 返回
  │
  ↓
main继续执行
```

(3)闭包`Closure`

闭包就是一个匿名函数，它捕获了函数外面作用域里的变量，匿名函数可以记住和访问外面的变量。`go`和`c++`的捕获规则完全不一样。

```go
func main(){
x := 10
f := func(){
fmt.Println(x)
}
f()
}
```

`func(){}`就是匿名函数，这个匿名函数引用了外部的`x`，它就是闭包，闭包捕获的是变量引用，不是变量当时的值。`c++ lambda`默认是值捕获，复制一份快照，但是`go`闭包永远捕获变量本身，不是拷贝值。

```go
package main
import (
	"fmt"
	"time"
)
func main() {
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println(i) // 闭包捕获i的引用！所有协程共享同一个i变量
		}()
	}
	time.Sleep(100 * time.Millisecond)
}

```

这个程序输出333,因为闭包访问的是`i`的地址，等`goroutine`跑起来的时候，`for`循环早跑完了。修复方案：

```go
for i := 0; i < 3; i++ {
	go func(peer int) { // peer是形参，每次调用复制一份i的值
		fmt.Println(peer)
	}(i) // 这里立刻把当前i拷贝传进去
}
```

```go
for i := 0; i < 3; i++ {
	peer := i // 每次循环创建新变量peer
	go func() {
		fmt.Println(peer)
	}()
}
```

因为捕获的是引用，闭包内部可以修改外层变量。

闭包可以延长变量生命周期，如果闭包捕获了一个外部局部变量，编译器会把这个变量逃到堆上，函数退出变量依然存活。

(4)类型断言`type assertion`

```go
var x any = "hello"
s := x.(string)
fmt.Println(s)
var x any = 100
s := x.(string)//这里这样写程序会panic,因为x不是string类型，所以需要用到ok
s, ok := x.(string)
if ok {
    fmt.Println("是字符串:", s)
} else {
    fmt.Println("不是字符串")
}
```

不知道一个`interface`是什么类型可以:

```go
switch v := x.(type) {
case int:
    fmt.Println("int:", v)

case string:
    fmt.Println("string:", v)

case bool:
    fmt.Println("bool:", v)

default:
    fmt.Println("unknown")
}
```

## 26.`Go`并发

来实现一个`Worker Pool`工作池

第一版：`goroutine+waitgroup`

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()

    fmt.Println("worker", id, "开始工作")

    time.Sleep(time.Second)

    fmt.Println("worker", id, "工作完成")
}

func main() {
    var wg sync.WaitGroup

    wg.Add(3)

    go worker(1, &wg)
    go worker(2, &wg)
    go worker(3, &wg)

    wg.Wait()

    fmt.Println("所有 worker 完成")
}
```

第二版：`channel+goroutine`

```go
tasks <-chan int//表示这个channel只能接收任务
```

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(
    id int,
    tasks <-chan int,
    wg *sync.WaitGroup,
) {
    defer wg.Done()

    for task := range tasks {
        fmt.Println(
            "worker", id,
            "处理任务", task,
        )

        time.Sleep(500 * time.Millisecond)
    }

    fmt.Println("worker", id, "退出")
}

func main() {
    tasks := make(chan int)

    var wg sync.WaitGroup

    // 创建 3 个 worker
    wg.Add(3)

    go worker(1, tasks, &wg)
    go worker(2, tasks, &wg)
    go worker(3, tasks, &wg)

    // 发送任务
    for i := 1; i <= 10; i++ {
        tasks <- i
    }

    // 告诉 worker：没有更多任务了
    close(tasks)

    // 等待 worker 全部退出
    wg.Wait()

    fmt.Println("所有任务处理完成")
}
```

第三版：`mutex`

我们现在想统计总共处理了多少个任务，`count++`

```go
var mu sync.Mutex
var count int
mu.Lock()
defer mu.Unlock()
count++
```

第4版:加入`context`

```go
func worker(
    ctx context.Context,
    id int,
    tasks <-chan int,
    wg *sync.WaitGroup,
) {
    defer wg.Done()

    for {
        select {
        case <-ctx.Done():
            fmt.Println("worker", id, "取消")
            return

        case task, ok := <-tasks:
            if !ok {
                fmt.Println("worker", id, "退出")
                return
            }

            fmt.Println(
                "worker", id,
                "处理任务", task,
            )
        }
    }
}
```

## 26.`TCP`网络编程

(1)`Go`创建`TCP Server`

```go
package main

import (
    "fmt"
    "net"
)

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("listen error:", err)
        return
    }

    defer listener.Close()

    fmt.Println("server listening on :8080")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("accept error:", err)
            continue
        }

        fmt.Println("new client:", conn.RemoteAddr())

        go handleClient(conn)
    }
}
```

`net.Listen`相当于`c++`的`socket+bind+listen`

`listener.Accept()`相当于`c++`的`accept()`

`net.conn`就是客户端连接本身，相当于`c++`里的`connfd`，可以对他进行`conn.Read()、conn.Write()、conn.close()`。

(2)处理客户端

```go
func handleClient(conn net.Conn) {
    defer conn.Close()

    buf := make([]byte, 1024)

    for {
        n, err := conn.Read(buf)

        if err != nil {
            fmt.Println("client disconnected")
            return
        }

        fmt.Println("收到:", string(buf[:n]))

        conn.Write([]byte("server received\n"))
    }
}
```

`go handleClient(conn)`这句话意味着每有一个客户端连接成功就给这个客户端开个`goroutine`。

(3)`Go Client`

```go
package main

import (
    "fmt"
    "net"
)

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:8080")

    if err != nil {
        fmt.Println("dial error:", err)
        return
    }

    defer conn.Close()

    conn.Write([]byte("hello server"))

    buf := make([]byte, 1024)

    n, err := conn.Read(buf)

    if err != nil {
        return
    }

    fmt.Println("收到:", string(buf[:n]))
}
```

`net.Dial()`相当于`c++`的`socket()、connect()`

```go
C++ Client                Go Client

socket()                  net.Dial()
connect()                 net.Dial()
send()/write()            conn.Write()
recv()/read()             conn.Read()
close()                   conn.Close()
```

(4)`TCP`是字节流，我们之前写聊天室是自己写了一个协议，按照`\n`进行分割消息，`Go`里面用`bufio.Scanner`帮我们按照换行读取：

````go
func handleClient(conn net.Conn) {
    defer conn.Close()

    scanner := bufio.NewScanner(conn)

    for scanner.Scan() {
        message := scanner.Text()

        fmt.Println("收到:", message)

        conn.Write(
            []byte("server: " + message + "\n"),
        )
    }
}
````

(5)`Go`写聊天室不需要`epoll`

```go
你的 Go 代码

conn.Read()
     │
     ▼
Go net 包
     │
     ▼
Go runtime
     │
     ▼
Linux 网络 I/O 机制
     │
     ▼
epoll 等
```

这里不需要自己手写`epoll`，`Go runtime`帮我们完成了这份工作。

`Go runtime`在`linux`底层依然在用`epoll`，只是`runtime`封装了一层`netpoller`网络轮询器，把`epoll`的回调事件循环包装成了同步阻塞的`goroutine`写法。

**流程：**

`Go的socket`全部默认设置成非阻塞

当`goroutine`调用`conn.Read()`的时候缓冲区此时如果没有数据，系统调用会立刻返回`EAGAIN`，`runtime`捕获这个`EAGAIN`，把这个`sockfd`注册到全局唯一`epoll`实例也就是`netpoller`，`gopark()`把当前`G`挂起，状态变成`_Gwaiting`，把`G`绑定到这个`fd`的结构体，等`fd`就绪之后找这个`G`，把`M`还给`P`，`M`不卡住，`P`立刻去调度另一个就绪`G`运行，此时`G`等待网络事件，`M`去跑别的`goroutinr`，没有线程阻塞在`epoll_wait`，等到收到数据内核通知`epoll`这个`fd`就绪了，`runtime`的某个`M`调用`netpoll()`，内部调用`epoll_wait`拿到就绪`fd`，根据`fd`找到之前挂起的`G`，把`G`放回`P`的运行队列，`P`调度到这个`G`，恢复执行`conn.Read()`，这次读到数据，函数返回。

**总结**：

`epoll`负责内核通知哪个`fd`就绪，`netpoller`负责`fd`就绪之后唤醒对应的`goroutine`，`c++ reactor`和`go netpoller`最大的区别就是`go netpoller`不会阻塞`OS`线程。

## 27.`RPC`

(1)`RPC`就是让你像调用本地函数一样去调用另一台机器上的函数，分布式系统最核心的事件之一就是`RPC`。

(2)`RPC`都要完成什么？

```go
① 客户端调用
       ↓
② 把函数名 + 参数序列化
       ↓
③ 通过网络发送
       ↓
④ Server 接收
       ↓
⑤ 反序列化
       ↓
⑥ 找到对应函数
       ↓
⑦ 执行
       ↓
⑧ 序列化返回值
       ↓
⑨ 网络发送
       ↓
⑩ Client 反序列化
```

`RPC`本质上还是网络通信，只不过把网络通信封装成了调用函数的形式。

```go
应用程序
   │
   ▼
RPC
   │
   ▼
TCP
   │
   ▼
IP
   │
   ▼
网卡
```

`RPC`就是负责把函数调用变成网络请求

(3)`RPC`为什么需要序列化？

客户端内存里：

```go
Args{
	A: 10,
	B: 20,
}
```

`server`不能直接访问`client`的内存，两台机器是隔离的，所以必须:

````go
Go Struct
   ↓
序列化
   ↓
字节
   ↓
网络
   ↓
字节
   ↓
反序列化
   ↓
Go Struct
````

注意`Go`的`JSON`序列化有一个规则：结构体字段必须是大写开头，才能被`json`包访问并序列化

不同`RPC`框架可能使用不同序列化方式：

```go
JSON
Protobuf
gob
自定义二进制协议
```

(4)`RPC`最大的问题：网络不可靠，还有`server`和`client`一些函数执行状态是否一致，都是分布式系统的难点。

(5)为什么`Raft`需要`RPC`?

因为`Raft`本身就是多个`server`，`Leader`需要告诉其他节点：

```go
“我是 Leader”
“投我一票”
“这是新的日志”
“提交到这里”
```

这些都通过`RPC`，例如:

```go
RequestVote RPC

Leader/候选人
     │
     │ RequestVote
     ▼
Follower
     │
     │ Reply
     ▼
Candidate
```

(6)客户端用来发起远程调用，只有一个核心函数

```go
func (end *ClientEnd) Call(serviceMethod string, args interface{}, reply interface{}) bool
ok := client.Call("Server.Add", &args, &reply)
```

`serviceMethod`:字符串，结构体名.方法名，例如`Raft.RequestVote`

`args`:请求参数结构体指针，要发给服务器的数据

`reply`:返回值结构体指针，服务器返回的数据会写到这里面

返回值`bool`:`true`=成功收到恢复，`false`=失败，丢包、超时

所有`RPC`请求全部都是调用`Call`。

(7)服务端3个关键函数，框架内部，不用自己调用但是要注册

**`labrpc.MakeServer()`**创建`RPC`服务对象，用来接收别人发来的`RPC`请求。

```go
server := labrpc.MakeServer()
```

**`server.Register(rcvr interface{})`**注册你的对象

```go
server.Register( &MyServer{} )
```

告诉`RPC`服务器，这个结构体上面的所有方法，可以被远程调用，`Register`内部会扫描结构体所有符合`RPC`格式的方法，建立一张映射表：`"server.Add"->Add`这个方法。

只有注册过的结构体，它的方法才能被远程调用，没有注册`call`调用失败。

**自己写的那个`RPC`处理函数，就是被远程调用的那个函数，`labrpc`强制规定固定模板：**

```go
func (s *Server) Add(args *AddArgs, reply *AddReply) error {
    reply.Sum = args.A + args.B
    return nil
}
```

第一个参数：请求参数`args`

第二个参数：输出结果写到`reply`

返回`error`，成功返回`nil`

(8)整体流程：

你写客户端代码 → 调用 `Call()`，发起远程请求

`RPC `服务端收到包 → 框架内部自动反射调用你的 `Add / RequestVote `函数

你永远不会手动直接调用 `Add`；`Add `是被 `RPC` 框架调用。

(9)代码简单举例：

```go
package main

import (
	"fmt"
	"labrpc"
)

// ========== 1. 服务端：定义结构体（RPC服务对象） ==========
type MathServer struct{}

// 定义RPC请求参数结构体：客户端发给服务端的数据
type MultiplyArgs struct {
	A int
	B int
}

// 定义RPC返回结构体：服务端计算后，返回给客户端的数据
type MultiplyReply struct {
	Result int
}

// ========== 2. 被远程调用的RPC方法【固定模板！】 ==========
// 规则：(接收器指针)，args指针，reply指针，返回error
func (s *MathServer) Multiply(args *MultiplyArgs, reply *MultiplyReply) error {
	// 服务端本地执行计算
	reply.Result = args.A * args.B
	return nil
}

func main() {
	// ========== 3. 服务端初始化 ==========
	// 创建RPC服务器实例
	server := labrpc.MakeServer()
	// 注册MathServer对象
	// Register会扫描MathServer所有符合RPC格式的方法，存入映射表
	srv := &MathServer{}
	server.Register(srv)

	// ========== 4. 创建客户端，绑定这个服务端 ==========
	client := labrpc.MakeClient(server)

	// ========== 5. 客户端发起远程调用 Call ==========
	args := MultiplyArgs{
		A: 3,
		B: 5,
	}
	var reply MultiplyReply // 用来接收返回结果

	// Call(服务名.方法名, 参数指针, 返回指针)
	ok := client.Call("MathServer.Multiply", &args, &reply)

	if ok {
		fmt.Printf("RPC调用成功，结果 = %d\n", reply.Result)
	} else {
		fmt.Println("RPC调用失败")
	}
}

```

## 28.创建项目

```go
go mod init myproject//生成go.mod

# 格式化整个项目
go fmt ./...

# 运行 Server
go run ./server

# 运行 Client
go run ./client

# 编译 Server
go build ./server

# 编译 Client
go build ./client

# 测试整个项目
go test ./...  //要有server_test.go client_test.go文件，必须要以_test.go结尾，也可以go test ./server  go test ./client

# 查看详细测试
go test -v ./...

# 整理依赖
go mod tidy
```

**`go run` 用来实际运行程序，`go test` 用来自动验证程序，检查你的实现是否正确**
