## GO 工程化实践

#### 工程项目结构

![](F:\陈晓会\GO\极客大学\Go-000\Week04\项目工程结构.png)

* go中所有的包都是一等公民. first class. 都是相同的.

* 只不过包放在了不同的目录.

* 当更多的人共同参与一个POC时候,需要更多的目录结构.

  >  鼓励开发一个小工具toolkit. 快速创建工程模板.

* kit project 工具包,基础库;  一个公司只有一个; 

* service application 服务端应用程序. 微服务.

  ```go
  /api  API协议
  /configs  YAML文件
  /test  	测试数据;   SQL文件;   GO编译忽略"." "_"开头的文件或目录.
  ```

  

* 不要用/src

* 服务数 service tree    业务.服务.子服务

* model:   放一些数据结构;  数据层;

* DAO面向表/redis的key;   数据访问层;   

* 

#### API设计

* gRPC: IDL
* PB: protocol buffers 定义/文档/代码
* google api design guide (谷歌API设计指南)中文版

#### 配置管理

* 环境配置

* 静态配置

* 动态配置

* 全局配置

  > I believe that we, as Go grogrammers, should work hard to ensure that nil is never a parameter that needs to be passed to any public function.

#### 包管理

#### 测试

* <<google 软件测试之道>>
* 微软测试之道

