## 评论系统架构设计

> PDF: 如何写一个架构设计的文档

#### 功能模块

> 先搞清功能

#### 架构设计

![](F:\陈晓会\GO\极客大学\Go-000\Week05\评论架构全局设计.png)

> 大的架构框框, 理解整个产品体系在系统中的定位. 
>
> 不要做需求翻译机,code machine.   先理解业务背后的本质, 事物的初衷.
>
> 系统背后是什么? 
>
> 好的架构可以在一个信封上画出来.

API层: API网关

BFF层(backend for frontend 服务于前端的后端):服务编排,**业务逻辑收敛**; 根据**业务逻辑**组装一些不同的服务; 类似FS中的LUA脚本功能. 比如读取评论时候,需要获取用户等级信息. 写评论时候要判断用户是否够等级了,是否被关小黑屋等等逻辑.还有写评论要调用敏感词过滤系统进行过滤.**这层不要太重**,比如把不同等级的业务逻辑都在这里进行区别化处理. **这是一个平台功能.**

Service层:去掉业务逻辑,专注在 评论功能的API实现上,比如发布/读取/删除上.以及稳定性_降级,可用性.

job层: 消息队列的最大用途消峰处理. job处理写请求. 然后job来更新redis和mysql. 

admin:comment-admin  管理平台;运营体系管理. 划分运营平台.按角色/安全等级来划分运营服务.  这层跟Service层共享存储层. 

> mysql是OLTP级联事物处理型数据库. 面向线上. 不适合进行分析计算. 也不适合进行后台运营搜索数据库,常常需要建立很多索引,搞的非常复杂,但是还是比不上ES.总得来说,就是一个慢查询.  状态层.
>
> redis: 只是一个cache,不是storage.  是可以被丢失的. 易失的.
>
> cannel将binlog的日志进行订阅,然后同步到ES.

dependency: account-service,  filter-service. 

> 架构设计==数据设计.  梳理清楚**数据走向和逻辑**. 尽量避免环形依赖,数据双向请求.

> 边缘缓存模式(cached-aside pattern,旁路缓存): 应用程序先从cache取数据,没有得到,则从数据库中取数据,成功后,放到缓存中.

![](F:\陈晓会\GO\极客大学\Go-000\Week05\kafka模型.png)

| 名称          | 解释                                                         |
| ------------- | ------------------------------------------------------------ |
| producer      | 消息的生产者                                                 |
| consumer      | 消息的消费者                                                 |
| consumerGroup | 消费者组,可以并行消费topic中的partition的消息.               |
| broker        | 缓存代理,kafka集群中的一台或多台服务器统称broker.            |
| topic         | kafka处理资源的消息源(feeds of messages)的不同分类           |
| partition     | topic物理上的分组,一个topic可以分为多个partition,每个partition是一个有序的队列.partition中每条消息都会被分一个有序的ID(offset). |
| message       | 消息,是通信的基本单位,每个producer可以向一个topic(主题)发布一些消息. |
| producers     | 消息和数据生产者,向kafka的一个topic发布消息的过程叫做producers |
| consumers     | 消息和数据的消费者,订阅topic并处理其发布的消息的过程叫做consumers |



> KAFKA: 全局并行,局部串行,生产者-消费者模型.  hash(comment_subject) % N(partitions) 分发消息. 导致形同主题的数据被分在一起了,方便消费.
>
> topic:  多个partition
>
> partition: 一个小队列.
>
> 

#### 存储设计 

> 核心:3张表

* comment_subject_[0-49]
  * id int64 主键
  * obj_id int64 对象ID  文章/视频/漫画
  * obj_type int8 对象类型  文章/视频/漫画
  * member_id int64 作者用户ID
  * count int32 评论总数
  * root_count int32 根评论总数
  * all_count int32 评论+回复总数
  * state int8 状态(0-正常,1-隐藏)
  * attrs int32 属性(bit 0-运营置顶, 1-up置顶, 2-大数据过滤)
  * create_time timestamp 创建时间
  * update_time timestamp 修改时间
* comment_index_[0-199]
  * id int64 主键
  * obj_id int64 对象ID
  * obj_type int8 对象类型
  * member_id int64 发表者用户ID
  * root int64 根评论ID,不为0,是回复评论
  * parent int64 父评论ID,为0是root评论
  * floor int32 评论楼层
  * count int32 评论总数
  * root_count int32 跟评论总数
  * like int32 点赞数
  * hate int32 点踩数
  * state int8 状态(0-正常,1-隐藏)
  * attrs int32 属性
  * create_time timestamp 创建时间
  * update_time timestamp 修改时间
* comment_content_[0-199]
  * comment_id int64 主键
  * at_member_ids string 对象ID
  * ip int64 对象类型
  * platform int8 发表者用户ID
  * device string 设备信息
  * message string 评论内容
  * meta string 评论元数据;背景/字体
  * create_time timestamp 创建时间
  * update_time timestamp 修改时间

> Graph存储.  dgraph/hugegraph

### 存储设计_缓存设计

* comment_subject_cache[string]
  * key string old_type   #object's id+type
  * value int64 subject marshal string
  * expire duration 24h  过期时间
* comment_index_cache[sorted set]
  * key string cache key:oid_type_sort, 其中sort为排序方式,0:楼层, 1:回复数量
  * member int64 comment_id:评论ID
  * score double 楼层号,回复数量,排序得分
  * expire duration 8h
* comment_content_cache[string]
  * key string comment id
  * value int64 content marshal string
  * expire duration 24h

#### 可用性设计

* singleflight 单飞:单个进程中只交给一个人来回源
*  热点:
  * hashmap:key/count
  * 读热点: 
  * overlord: 缓存代理
  * aster: mesh 网格
* Scaling Memcache At Facebook
* 环形数组



