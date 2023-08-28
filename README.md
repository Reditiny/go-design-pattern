# go-design-pattern

## 设计模式
设计模式是在软件设计中反复出现的问题的解决方案。它们是经过多次实践和验证的，旨在提供可重用、可维护和可扩展的代码结构。以下是一些常见的设计模式

### 创建型模式
这些模式关注对象的创建机制，确保对象被正确地创建和初始化。
+ [简单工厂模式（Simple Factory）](./1_1_simple_facoty/READMED.md)
+ [工厂方法模式（Factory Method Pattern）](./1_2_facotry_method/READMED.md)
+ [抽象工厂模式（Abstract Factory Pattern）](./1_3_abstract_factory/READMED.md)
+ [单例模式（Singleton Pattern）](./1_4_singleton/READMED.md)
+ [建造者模式（Builder Pattern）](./1_5_builder/READMED.md)
+ [原型模式（Prototype Pattern）](./1_6_prototype)

### 结构型模式
这些模式处理对象之间的组合，以形成更大的结构，以及如何简化对象之间的交互。
+ [适配器模式（Adapter Pattern）](./2_1_adapter/READMED.md)
+ [装饰器模式（Decorator Pattern）](./2_2_decorator/READMED.md)
+ [代理模式（Proxy Pattern）](./2_3_proxy/READMED.md)
+ [外观模式（Facade Pattern）](./2_4_facade/READMED.md)
+ [桥接模式（Bridge Pattern）](./2_5_bridge/READMED.md)
+ [组合模式（Composite Pattern）](./2_6_composite/READMED.md)
+ [享元模式（Flyweight Pattern）](./2_7_flyweight/READMED.md)

### 行为型模式
这些模式处理对象之间的通信和协作，定义了对象如何相互合作以实现更大的功能
+ [策略模式（Strategy Pattern）](./3_1_strategy/READMED.md)
+ [观察者模式（Observer Pattern）](./3_2_observer/READMED.md)
+ [责任链模式（Chain of Responsibility Pattern）](./3_3_responsibility_chan/READMED.md)
+ [命令模式（Command Pattern）](./3_4_command/READMED.md)
+ [解释器模式（Interpreter Pattern）](./3_5_interpreter/READMED.md)
+ [迭代器模式（Iterator Pattern）](./3_6_iterator/READMED.md)
+ [中介者模式（Mediator Pattern）](./3_7_mediator/READMED.md)
+ [备忘录模式（Memento Pattern）](./3_8_memento/READMED.md)
+ [状态模式（State Pattern）](./3_9_state/READMED.md)
+ [访问者模式（Visitor Pattern）](./3_10_visitor/READMED.md)

## 设计原则
软件设计原则是在软件开发过程中指导设计决策的基本指导方针。这些原则旨在帮助开发人员创建更加可维护、可扩展和可重用的软件，以及降低代码的复杂性。以下是一些常见的软件设计原则：

### 单一职责原则（Single Responsibility Principle，SRP）
一个类应该只有一个引起它变化的原因。换句话说，一个类应该只负责一个单一的功能或职责

### 开放封闭原则（Open-Closed Principle，OCP）：
软件实体（类、模块、函数等）应该是可扩展的，但不可修改的。新功能的添加不应该修改现有的代码，而是通过扩展来实现。

### 里氏替换原则（Liskov Substitution Principle，LSP）：
子类应该能够替代父类，并且不影响程序的正确性。换句话说，派生类应该保持父类的行为。

### 依赖倒置原则（Dependency Inversion Principle，DIP）：
高层模块不应该依赖于低层模块，两者都应该依赖于抽象。同时，抽象不应该依赖于具体细节，具体细节应该依赖于抽象。

### 接口隔离原则（Interface Segregation Principle，ISP）：
不应该强迫客户端实现它们不使用的接口。换句话说，应该保持接口的小而专一，而不是大而笨重。

### 迪米特法则（Law of Demeter，LoD）：
一个对象应该对其它对象有最少的了解，只与直接的朋友进行交互。避免暴露对象内部的细节。

### 组合/聚合复用原则（Composition/Aggregation Reuse Principle，CARP）：
优先使用组合或聚合关系来实现代码的复用，而不是继承。继承关系容易导致紧耦合和继承链的问题。