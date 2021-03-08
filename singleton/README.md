# 单例模式
单例模式（Singleton Pattern）指确保一个类在任何情况下都绝对只有一个实例，并提供一个全局访问点，属于创建型设计模式

# 常见单例模式
1. 饿汉式单例:
    - 类/模块初始化时立即创建的静态全局单例

2. 双重检查单例:
    - 提供获取单例的方法
    - 在该方法中通过双重检查静态实例是否为null,从而延迟初始化全局单例
    - 首次检查不加同步锁，二次检查加同步锁，以保证单例只创建一次

3. 容器式单例:
    - 存在全局唯一的Bean容器
    - 通过Bean的名称读取或设置与该名称绑定的bean单例
    - 通过读写锁来控制并发安全
    - 通常用于需要持有大量单例的场景，如IOC容器

# 说明
单例模式的优点:
1. 单例模式可以保证内存里只有一个实例，减少了内存的开销。 
2. 可以避免对资源的多重占用。 
3. 单例模式设置全局访问点，可以优化和共享资源的访问。 

单例模式的缺点: 
1. 单例模式一般没有接口，扩展困难。如果要扩展，则除了修改原来的代码，没有第二种途径，违背开闭原则。 
2. 在并发测试中，单例模式不利于代码调试。在调试过程中，如果单例中的代码没有执行完，也不能模拟生成一个新的对象。 （
3. 单例模式的功能代码通常写在一个类中，如果功能设计不合理，则很容易违背单一职责原则。

