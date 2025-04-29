## jammy-cloud-database

-----------------------
**jammy-cloud-database** 是一个云计算项目，主要提供RDS数据库的生命周期管理，包含数据库实例创建，删除，参数管理，备份管理等。

**组件说明：**
- example主要为代码使用示例，包含对数据库节点容器的操作以及对数据库生命周期api的调用等
- mysql-instance-management 主要为mysql实例管理
- rds-api api接口层
- mysql-monitor-agent 主要为数据库代理，安装在节点容器中，负责数据库指标监控，上报，执行各种生命周期管理的相关操作

-------------------------
**支持的数据库：**
- mysql
- postgresql
- sqlserver
-------------------------
**容器底座：**

本项目使用lxd作为RDS数据库节点容器，相比较传统的虚拟机有更小的开销，同时对比docker及containerd等容器，提供更好的隔离性

**数据库实例初始化：**

