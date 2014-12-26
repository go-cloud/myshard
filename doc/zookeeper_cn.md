在myshard里面，我们使用zookeeper（后续以zk代替）来进行全局配置管理，以及协调多个服务。

## MySQL

我们使用group来表明一组MySQL，一组MySQL至少需要有一个master，可以有0个或者多个slave。也就是最常用的master/slave模式，我们不建议使用master/master架构模式，因为这样可能引入更多的复制问题，如果真要支持，myshard也可能只会选择一个进行写入操作，而将另一个当做slave。另外，还需要考虑gelera的架构模式，应该可以退化到单master情况，（gelera虽然有三台MySQL，但外部通过一个入口访问。）

group是全局唯一的，使用数字1，2，3来表示，不同group不允许出现相同的MySQL实例。


group:

```
/zk/myshard/groups/group_1
/zk/myshard/groups/group_2
/zk/myshard/groups/group_3
```

group node

```
type GroupNode struct {
    ID int 
    Status string
}
```

group MySQL

```
/zk/myshard/groups/group_1/10.20.151.151:3306
/zk/myshard/groups/group_1/10.20.151.152:3306
```

MySQL 

```
/zk/myshard/mysqls/10.20.151.151:3306
/zk/myshard/mysqls/10.20.151.152:3306
```

MySQL node

```
type MySQLNode struct {
    ID int      //MySQL server id, must be unique in the global topology
    Addr string 
    Type string //master or slave
    Status string
    GroupID int 
}
```

## Agent

agent用来监控MySQL是否存活，通常跟MySQL在一台机器上面，agent在zk上面注册一个ephemeral节点，当MySQL当掉，或者所在机器当掉，agent与zk的连接断开，zk就能通知其他监控程序后续处理，譬如邮件通知，或者自动failover处理。

但我们还可能面临一种情况就是agent挂掉了，但MySQL仍然存活，所以我们外部的程序仍然需要额外判断MySQL是不是真正当掉了。

agent

```
/zk/myshard/agents/10.20.151.151:3306
```

## Proxy

proxy是一个无状态的服务，所以我们使用sequence flag来在zk上面创建相应的node

```
/zk/myshard/proxys/proxy_seq
```

proxy node

```
type ProxyNode struct {
    ID int 
    Addr string
    HttpAddr string
    Status string
}
```

## Route

route用来存放路由表信息，myshard会根据首先解析sql，然后根据响应的路由规则将sql发到不同的group里面的一台MySQL去执行。

对于客户端来说，它只知道sql的db和table，但是在myshard内部，我们可能将该数据拆分到不同的子table上面，也就是说，对外可能是user表，但在
myshard内部，我们使用的是user_1, user_2, ... user_1024。

而对于一个db，它可能存放到不同的group上面

```
/zk/myshard/dbs/db_1
/zk/myshard/dbs/db_2
```

db node

```
type DBNode struct {
    Name string
    GroupIDs []int //group ids
    Status string

    //如果table没有指定group，则默认在该group上面
    DefaultGroupID int
}
```

table

```
/zk/myshard/dbs/db_1/user
```

table node

```
type TableNode struct {
    //如果改table没有切分，则使用groupid，否则，各个子表需要使用各自对应的groupid
    GroupID int

    //因为可能会有多个子表，为了简单，key的格式可能如下
    // 1      表示table_1
    // 1,2,3  表示table_1, table_2, table_3
    // 1-3    表示table_1, table_2, table_3
    SubTables map[string]int
}
```

```
/zk/myshard/routes/db_1/user_route
```

route node

```
type RouteNode struct {
    DB string
    Table string

    Type string    //hash or range
    Column string  //table column for route

    RangeMap map[string]int  // key like "1-100" -> table id 
}
```

## Action

sequence flag

```
/zk/myshard/actions/action_seq
```

action node

```
type ActionNode struct {
    ID int
    CreateTime int
    Command string
    Args []interface{}
}
```

## Lock

任何对全局配置的更改都必须获取lock

```
/zk/myshard/lock
```

## Migrate

如果预估到一个表数据量会很大，我们需要在开始的时候就对其进行分表处理，譬如1024张表

myshard对于数据迁移是以子表为单位的，假设现在结构如下

    group1 -> user_1, user_2, user_3
    group2 -> user_4

因为group1的压力，我们需要将user_2移动到group2

+ 我们先将当前user_2表数据拷贝到group2上面
+ 因为这时候可能会有新的写入操作发生，我们只需要进行binlog同步就可以
+ 当binlog同步完成或者两个group之间的相差不大，我们通过config像zk发起迁移action
+ proxy收到改action之后，会拒绝所有对user表的写操作，但读仍然可行，同时等待先前所有对user表的写操作完成
+ proxy通过zk告知config可以迁移
+ config等待binlog同步完成，然后更新路由表
+ proxy收到路由表更新事件，在本地更新路由表，同时告知config更新完成
+ config发起迁移完成action，proxy收到之后开始新的处理。因为这时候仍然有老的连接在使用就得数据查询，所以我们会在一段时间之后才从group1删除user_2

迁移过程中的错误处理，后续完善

## Up and Down MySQL

通过config对某个group内的MySQL集群进行操作，大概流程类似migrate，只是proxy需要处理的是group，而不是特定的表

config在操作的时候一定要首先获取lock，也就是同一个时间只允许一个全局更新操作