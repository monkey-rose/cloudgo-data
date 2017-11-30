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
## 程序结构
## 服务性能
# Answer
## Q:orm 是否就是实现了 dao 的自动化？
orm（object relational mapping）对于数据库操作，不再需要自己写SQL代码，而是用go对象mapping，自动生成SQL代码，,实现了数据存取的自动化

## Q:使用 ab 测试性能
#### orm
![post](http://ww1.sinaimg.cn/large/0060lm7Tly1fm000ptmu8j31ke198wpj.jpg)
![get](http://ww3.sinaimg.cn/large/0060lm7Tly1fm000l7630j313q162gux.jpg)
#### database/sql
![post](http://ww2.sinaimg.cn/large/0060lm7Tly1fm000ae0uqj31kw17jtjx.jpg)
![get](http://ww3.sinaimg.cn/large/0060lm7Tly1fm000hij19j316m17kqcx.jpg)
