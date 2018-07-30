
# GoDemo <img src="http://i.imgur.com/hTeVwmJ.png" width="40" height="40" alt=":walrus:" class="emoji" title=":walrus:"/>
##学习go语言的Demo
#### 2018.7.26 学习Deployment部署和Maintenance维护，
* 1、日记应用 尼玛逼啊！还要一个依赖：https://github.com/golang/sys 
* 2、seelog的Demo完成
* 3、错误的处理：panic和recover是针对自己开发package里面实现的逻辑，针对一些特殊情况来设计。
* 5、错误的处理
*  6、网站错误处理：数据库错误（连接错误、查询错误、数据错误）；应用运行时错误（文件系统和权限、第三方应用和接口错误）；HTTP错误；操作系统出错；网络出错
*  7、错误处理的目标：通知访问用户出现错误了；记录错误；回滚当前的请求操作；保证现有程序可运行可服务
*  8、如何处理错误
*  9、应用部署： daemon：Go程序还不能实现daemon，详细的见这个Go语言的bug：<http://code.google.com/p/go/issues/detail?id=227>，大概的意思说很难从现有的使用的线程中fork一个出来，因为没有一种简单的方法来确保所有已经使用的线程的状态一致性问题
* 10、Supervisord可惜啊，不支持window系统啊日了狗 
* 11、备份和恢复