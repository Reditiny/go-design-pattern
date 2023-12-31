# 3.3 责任链模式
通过一系列的处理器来逐级处理请求，直到请求被处理或者到达链的末端。每个处理器对象都可以处理请求，或者将请求传递给下一个处理器对象。  
责任链模式将请求的发送者和接收者解耦，使得多个对象都有机会处理请求。

## 案例
请假流程控制系统。一天以内的假小组有权同意，否则交给经理处理；1天到3天的假经理有权同意，否则交给老板处理；3天以上的假期老板最终决定是否同意。

## 责任链模式 优缺点

**优点：**

+ 耦合度低，降低了请求发送者和接收者的耦合度
+ 扩展性强，可以根据需要增加新的请求处理类，满足开闭原则
+ 灵活性高，当工作流程发生变化，可以动态地改变链内的成员或者修改它们的次序，也可动态地新增或者删除责任。
+ 复杂度低，链简化了对象之间的连接，一个对象只需保持一个指向其后继者的引用，不需保持其他所有处理者的引用，这避免了使用众多的 if 或者 if···else 语句。

**缺点：**

+ 不能保证每个请求一定被处理。由于一个请求没有明确的接收者，所以不能保证它一定会被处理，该请求可能一直传到链的末端都得不到处理。
+ 对比较长的职责链，请求的处理可能涉及多个处理对象，系统性能将受到一定影响。
+ 职责链建立的合理性要靠客户端来保证，增加了客户端的复杂性，可能会由于职责链的错误设置而导致系统出错，如可能会造成循环调用。