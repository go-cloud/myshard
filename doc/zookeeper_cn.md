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

route用来存放路由表信息


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
