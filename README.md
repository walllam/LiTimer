LiTimer: 一个定时触发运行配置的URL小工具
-
可能看到这个，第一反映，这不是crontab吗？ 好吧，不是

### 主要特性：

* 间隔最小时间分钟，基于2016-01-01 00:00:00的基准时间来推算运行时间(当然也可以自己自定义每个URL的基准时间)
* 同步运行：如果前一次URL的HTTP请求还没有结束，就不会发起第二个
* 可以设置timeout
* 详细的运行日志和配置，且均基于MySQL，可以自行开发管理界面 

### 开发计划：

* 增加支持HTTP API报警的功能
* 增加支持通过redis为介质异步运行一个URL(触发源点可以任意来源，介质以redis的普通key或队例实现)

----
### 编译：
    1.1.下载golang，配置GOPATH和PATH
    1.2.测试运行[go version]看看能不能输出版本号
    1.3.安装需要使用的第三方包:
        go get github.com/go-ini/ini
        go get github.com/go-sql-driver/mysql
    1.4.切换到timer/src目录，运行[go build timer.go]，即可编译成功

### 配置和运行：
    2.1.配置timer.ini中的MySQL部份
    2.2.在配置的MySQL DB中初始化表结构，表结构在timer.sql中	
    2.3.运行[./timer]
    2.4.在[timer_process]表中插入要定时运行的URL(也可以修改或删除)，可以在[run_logs]查到运行的日志，可以对这二个表做一个管理后台和查询后台
    2.5.定时运行的URL必须首先输出[__run_successed__]方表示成功(后面可以接任何其它东西)，否则均表示失败

----
有任何的其它疑问或建议可以发邮件给我：wallupmark@gmail.com
