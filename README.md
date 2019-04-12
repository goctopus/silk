# silk

又双叒叕一个golang的orm框架，取名silk，意思是使用起来如丝般顺滑。目标是在保证一定性能的情况下，尽可能做到最优雅，最大力度的提升团队开发效率。

## 特点

问：silk跟别的传统的golang orm框架比有什么不同？

答：orm存在的意义就是为了开发效率，使得代码更简洁更清晰明了，让妈妈再也不用担心我面对一堆杂乱的sql时无从下手。silk把链式调用做到极度顺滑，同步提供cli工具让你自动生成模型文件，开箱即用。结构体的调用降低代码出错率。同步提供强大的数据处理工具collection，让你处理各种数据时不需要再自己造轮子。在开发效率层面是大大的提升。

## todo

- [ ] 测试各个驱动(mysql/mssql/postgresql/sqlite)的语法
- [ ] 性能测试
- [ ] 命令行工具
- [ ] [Collection数据结构实现](https://github.com/goctopus/silk/blob/master/collection.go)
- [ ] hook支持
- [ ] 事务支持
- [ ] 模型关系支持

## 如何参与开发

[如何参与开发](https://github.com/goctopus/silk/blob/master/CONTRIBUTING.md)