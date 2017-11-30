# xorm程序运行测试
## INSERT
### 插入2个item
![result](http://ww2.sinaimg.cn/large/0060lm7Tly1flzyud9n44j31bo0budir.jpg)
### 自动建表且插入成功(struct UserInfo, 数据库不含大写字母，表名将大写转为_)
![db](http://ww3.sinaimg.cn/large/0060lm7Tly1flzyuaczfdj30m80csta9.jpg)
### username为空，报错
![error](http://ww3.sinaimg.cn/large/0060lm7Tly1flzyuihw01j31a00343z6.jpg)
## QUERY(存在及不存在)
![query](http://ww2.sinaimg.cn/large/0060lm7Tly1flzyuljl67j311u0aywgl.jpg)
## server log
![log](http://ww2.sinaimg.cn/large/0060lm7Tly1flzyuobp3xj314c06ewhg.jpg)
# 对比 database/sql 与 orm 实现
## 编程效率
一般来说，orm可以让你获得更高的编程效率，因为封装了SQL语句，剪掉了dao的实现，不过对于很熟悉SQL语句的人来说，直接使用SQL语句进行数据库的操作，可以实现比较复杂的语句，或许更高效一些。
## 程序结构
database/sql实现：dao中用SQL语句对数据库操作，提供数据库操作接口，handler处理请求函数中调用数据库接口。
orm实现：创建engine，engine调用特定函数
## 服务性能
orm反射过程会牺牲一部分性能，与database/sql相比，对请求服务的时间会增加。不过orm的服务还有其他优势，一些orm库，比如go语言的xorm支持session事务和回滚以及乐观锁，诸如此类的数据库操作提供了更方便的服务。
# Answer
## Q:orm 是否就是实现了 dao 的自动化？
orm（object relational mapping）对于数据库操作，不再需要自己写SQL代码，而是用go对象mapping，自动生成SQL代码，,实现了数据存取的自动化

## Q:使用 ab 测试性能
#### orm
post
![post](http://ww1.sinaimg.cn/large/0060lm7Tly1fm000ptmu8j31ke198wpj.jpg)
get
![get](http://ww3.sinaimg.cn/large/0060lm7Tly1fm00pfz6ocj313q162gux.jpg)
#### database/sql
post
![post](http://ww2.sinaimg.cn/large/0060lm7Tly1fm000ae0uqj31kw17jtjx.jpg)
get
![get](http://ww3.sinaimg.cn/large/0060lm7Tly1fm000hij19j316m17kqcx.jpg)
