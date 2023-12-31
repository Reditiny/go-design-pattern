# 1.3 抽象工厂模式
前面介绍的工厂方法模式中考虑的是一类产品的生产，这些工厂只生产同种类产品，同种类产品称为同等级产品
本节要介绍的抽象工厂模式将考虑多等级产品的生产，将同一个具体工厂所生产的位于不同等级的一组产品称为一个产品族

## 案例
咖啡店，此时不仅提供咖啡还提供甜点，工厂可以同时提供咖啡与甜点两类产品

## 抽象工厂模式 优缺点

**优点：**

当一个产品族中的多个对象被设计成一起工作时，它能保证客户端始终只使用同一个产品族中的对象。

**缺点：**

当产品族中需要增加一个新的产品时，所有的工厂类都需要进行修改。
