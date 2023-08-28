# 3.4 命令模式
命令模式将请求（命令）封装成对象，从而允许将请求参数化、延迟执行或者将请求放入队列中，以及支持撤销操作。  
命令模式将请求的发送者和接收者解耦，使得不同的请求可以灵活地被处理和执行。

## 案例
餐馆点餐，服务员是调用者，负责发起命令（命令中包含订单）；大厨是接收者，真正命令执行的对象

## 命令模式 优缺点

**优点：**

+ 降低系统的耦合度。命令模式能将调用操作的对象与实现该操作的对象解耦。
+ 增加或删除命令非常方便。采用命令模式增加与删除命令不会影响其他类，它满足“开闭原则”，对扩展比较灵活。

**缺点：**

+ 使用命令模式可能会导致某些系统有过多的具体命令类。
+ 系统结构更加复杂。