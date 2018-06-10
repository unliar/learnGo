# 命名规则

1. 名字必须是一个字母或者下划线开头

2. 区分大小写
> 基本上所有语言都区分,mysql语句好像不区分

3. 保留关键字不能用于命名,避免使用预定义的名字用于命名
> 基本上所有语言都有保留字

4. 包的名字一般总是用小写
> 包和Node.js的模块的定义很像

5. 包里的顶级定义名字的大小写决定了在导出的时候外部的可见性。
> 意思是如果你的变量是首字母是大写,相当于Node.js模块里主动exports导出,Go直接用大小写来决定是否导出

6. 推荐使用**驼峰命名**,不推荐使用下划线线,缩略词请用大写
> Python语法里用下划线的比较多,HTML类似的缩略词记得用大写

```apple js
保留关键词

break      default       func     interface   select
case       defer         go       map         struct
chan       else          goto     package     switch
const      fallthrough   if       range       type
continue   for           import   return      var

内建常量

true false iota nil

内建类型
          
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
float32 float64 complex128 complex64
bool byte rune string error

内建函数
 
make len cap new append copy close delete
complex real imag
panic recover

```