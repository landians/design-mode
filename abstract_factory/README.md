# 抽象工厂
抽象工厂模式（Abstract Factory Pattern）指提供一个创建一系列相关或相互依赖对象的接口，无须指定它们具体的类。意思是客户端不必指定产品的具体类型，创建多个产品族中的产品对象。 

# 说明
抽象工厂模式的优点: 
1. 当需要产品族时，抽象工厂可以保证客户端始终只使用同一个产品的产品族。 
2. 抽象工厂增强了程序的可扩展性，对于新产品族的增加，只需实现一个新的具体工厂即可，不需要对已有代码进行修改，符合开闭原则。 

抽象工厂模式的缺点:
1. 规定了所有可能被创建的产品集合，产品族中扩展新的产品困难，需要修改抽象工厂的接口.
2. 增加了系统的抽象性和理解难度。

# 对比工厂方法说明:
抽象工厂是创建一系列对象, 且这些对象之间可能存在相关性. 工厂方法通常是创建单一类别的对象