LiTimer: 一个定时触发运行配置的URL小工具
-
可能看到这个，第一反映，这不是crontab吗？

好吧，不是的，主要特性：

* 间隔最小时间分钟，基于2016-01-01 00:00:00的基准时间来推算运行时间(当然也可以自己自定义每个URL的基准时间)
* 同步运行：如果前一次URL的HTTP请求还没有结束，就不会发起第二个
* 可以设置timeout
* 详细的运行日志和配置，且均基于MySQL，可以自行开发管理界面 

----
#1.直接使用编译好的[timer]:
    1.1.该文件是在CentOS6下面编译出来
    1.2.复制timer和timer.ini,timer.sql到安装目录

#2.自己编译：
    2.1.下载golang，配置GOPATH和PATH
    2.2.测试运行[go version]看看能不能输出版本号
    2.3.安装需要使用的第三方包:
        go get github.com/go-ini/ini
        go get github.com/go-sql-driver/mysql
    2.3.切换到timer的目录，运行[go build timer.go]，即可编译成功

----
    
# 选择上面的1或2做完后：
	1.配置timer.ini中的MySQL部份
	
	2.在配置的MySQL DB中初始化表结构，表结构在timer.sql中
	
	3.运行[./timer]
	
	4.在[timer_process]表中插入要定时运行的URL(也可以修改或删除)，可以在[run_logs]查到运行的日志，可以对这二个表做一个管理后台和查询后台
	
	5.定时运行的URL必须首先输出[__run_successed__]方表示成功(后面可以接任何其它东西)，否则均表示失败

----

	有任何的其它疑问或建议可以发邮件给我：balsampears@gmail.com