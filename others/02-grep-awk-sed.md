# Linux文本处理利器

## grep
`grep` 是一种强大的文本搜索工具，它能使用正则表达式搜索文本，并把匹配的行打印出来。

### 语法格式
`grep [options] "pattern" [file]`

### 常用选项
- `-i`：忽略大小写(ignore case)
- `-v`：反过来(invert)，只打印不匹配的，而匹配的不打印
- `-w`：被匹配的文本只能是单词，不能是单词中的一部分
- `-c`：显示共有多少行匹配，而不是显示匹配内容
- `-cv`：显示有多少行没有匹配，注意跟`-c`的区别
- `-A n`：显示匹配的字符串所在的行及其后n行，after
- `-B n`：显示匹配的字符串所在的行及其前n行，before
- `-C n`：显示匹配的字符串所在的行及其前后各n行，context
- `-E`：开启扩展(Extend)的正则表达式
- `-r,-R`：递归查找

## awk
`awk` 命令处理一个或者多个文件时，它会依次读取文件的每一行内容，然后对其进行处理，即行处理器。

### 语法格式
`awk [options] 'pattern action' [file1 file2]`

### 常用选项-options
- `-F`：指定分隔符，默认是空格
- `-v`：给指定变量赋值
- `-f`：从脚本文件中读取awk命令

**示例**
```bash
awk -F ':' '{print "first field:"$1}' /etc/passwd  # $1不能放在双引号中

# -F ':' 等价于 -v FS=':' 如下两行命令
awk -F ':' '{print "first field:"$1}' /etc/passwd
awk -v FS=':' '{print "first field:"$1}' /etc/passwd

awk -v a="xxxx" 'BEGIN{print a}' # 自定义变量，输出xxxx
```

### 模式-pattern
可以简单理解：模式即条件  
**1.空模式**  

**2.关系运算模式**  
- `>`：大于
- `>=`：大于等于
- `<`：小于
- `<=`：小于等于
- `==`：等于
- `!=`：不等于
- `~`：正则匹配为真
- `!~`：正则不匹配为真

**3.正则模式**
　具体参看示例
**4.行范围模式**
　具体参看示例
**5.特殊模式**
- `BEGIN`：处理文本之前执行的操作
- `END`：处理所有文本之后执行的操作

**示例**
```bash
### 行范围模式
awk 'NR==1' /etc/passwd  # 打印第1行

### 正则模式
# 打印匹配`/foo/`和`/bar/`的行，二者等价
awk '/foo/ && /bar/' /etc/passwd
awk '/foo/ && /bar/{print $0}' /etc/passwd

### 特殊模式
# 第1行输出first field，第2行输出nobody
awk -F ':' 'BEGIN{print "first field"} {print $1}' /etc/passwd  

# 第1行输出nobody，第2行输出first field
awk -F ':' '{print $1} END{print "first field"} ' /etc/passwd  

# 第1行输出first line，第2行输出nobody，第3行输出last line
awk -F ':' 'BEGIN{print "first line"} {print $1} END{print "last line"} ' /etc/passwd  
```

### 动作-action
如上述示例中看到的`{print $0}` 即动作。

**1.条件控制语句**
```bash
if(条件1)
{
  语句;
}
else if(条件2)
{
  语句;
}
else
{
  语句1;
}
```

**2.循环语句**
```bash
#for循环语法格式1
for(初始化; 布尔表达式; 更新) {
  //代码语句
}

#for循环语法格式2
for(变量 in 数组) {
  //代码语句
}

#while循环语法
while(布尔表达式) {
  //代码语句
}

#do...while循环语法
do {
  //代码语句
} while(条件)
```
支持 `break`、`continue`、`exit` 等关键字。  


**示例**
```bash
### 动作示例
# 二者等价
awk -v FS=":" '{print $1;print $3}' /etc/passwd
awk -v FS=":" '{print $1}{print $3}' /etc/passwd

### 条件控制语句
# 二者等价
awk 'NR==1' /etc/passwd
awk '{if(NR==1){print $0}}' /etc/passwd

awk '{if(NR==1){print $0} else if(NR==2){print "line 2"} else {print "others"}}' /etc/passwd

### 循环语句
awk 'BEGIN{for (i=0; i<10; i++) {if (i==5) {break} else {print "hi,"i}}}'

# 数组遍历,被删除后的项输出为空
awk 'BEGIN{a[0]=0;a[1]=11;a[2]=22;a[3]=33; for (i=0; i<length(a); i++) {delete a[2]; print a[i]}}' 
```

### 其他
**1.内置变量**
- 0表示文本处理时的当前行， 1表示文本行被分隔后的第1列，$n表示文本行被分割后的第n列
- `NR`：文件中的行号，表示当前是第几行
- `NF`：文件中的当前行被分割的列数
- `$NF`：文件中的当前行被分割后最后一列
- `FS`：输入分隔符，默认分隔符为空格和制表符
- `OFS`：输出分隔符，默认分隔符为空格和制表符
- `FILENAME`：表示当前文件的名称，如同时处理多个文件，也表示当前文件名
- `RS`：行分隔符，用于分割行，默认为换行符
- `ORS`：输出记录的分隔符，默认为换行符

**2.内置函数**
- `toupper()`：用于将字符转为大写
- `tolower()`：将字符转为小写
- `length()`：长度
- `substr()`：子字符串
- `sqrt()`：平方根
- `rand()`：随机数

**示例**
```bash
awk -v FS=':' -v OFS='||' '{print $1,$3}' /etc/passwd # 输出 nobody||-2

awk -F ':' '{print toupper($1)}' /etc/passwd  # 转换成大写
```

## sed
### 语法格式
```sed [options] action [file1 file2]```

### 常用选项
- `-i`：直接编辑原文件
- `-f`：指定命令文件

**示例**
```bash
sed -i '1d' 11.txt # 删除原文件的第1行
```

### 动作
**1.替换**  
　具体参看示例  
**2.删除**
　具体参看示例  
**示例**
```bash
### 替换
echo "hi,my name is x" | sed 's/x/jordan/;s/hi/hello/' # 输出 hello,my name is jordan
echo "hi,my name is x1 and x2" | sed 's/x/jordan/' # 替换第1个匹配 hi,my name is jordan1 and x2
echo "hi,my name is x1 and x2" | sed 's/x/jordan/g' # 替换全部 hi,my name is jordan1 and jordan2
echo "hi,my name is x1 and x2" | sed 's/x/jordan/2' # 替换第2个匹配 hi,my name is x1 and jordan2
echo "hi,my name is x1 and x2" | sed 's/x/jordan/w r.txt' # 结果保存到r.txt文件
sed '2s/aa/***/' 1.log # 行寻址并替换第2行
sed '1,3s/aa/***/' 1.log # 行寻址并替换第1-3行，2,$s说明是第2至最后一行

### 删除
sed '1d' 1.log # 删除第1行
```

## 对比
- `grep` 适合文本的查找和匹配
- `awk` 适合文本的格式化
- `sed` 适合文本的编辑

具体示例可以参考 [根据日志统计API的耗时和QPS](./01-api-qps-latency.md)

## 参看文献
1. [AWK命令总结之从放弃到入门](http://www.zsythink.net/archives/tag/awk)
2. [awk命令](http://man.linuxde.net/awk)
3. [sed命令](http://man.linuxde.net/sed)
4. [grep命令](http://man.linuxde.net/grep)
