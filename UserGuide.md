daoliet使用手册
=========

在开始使用daolinet之前，请确保系统已经完成[daolinet安装](../../../daolinet/blob/master/InstallGuide.md)过程，同时确保各个服务都正常启动。daolinet提供了一套命令行工具daolictl，并结合docker命令共同完成对docker网络的管理。

#### daolictl命令说明

daolictl工具提供了不同网段间的连通和隔离的安全组(group/member)功能，防火墙入站(firewall)功能以及虚拟机单独的隔离功能(cut/uncut)。所有子命令都可以通过--help选项查看具体说明。

#### daolinet使用说明

以下-H参数为启动api server指定的地址和端口，默认端口为3380，如果在其它非swarm manager节点上执行命令，请指定完整的ip地址和port端口。

1.创建网络

使用docker network命令创建网络，同时指定--driver为daolinet。

    docker -H :3380 network create --subnet=10.1.0.0/24 --gateway=10.1.0.1 --driver=daolinet dnet1
    docker -H :3380 network create --subnet=192.168.0.0/24 --gateway=192.168.0.1 --driver=daolinet dnet2

2.启动容器

使用docker run命令启动容器，同时指定--net为第一步中创建的网段名称。

    # 启动容器指定10.1.0.0/24的网络
    docker -H :3380 run -ti -d --net=dnet1 --name test1 centos # 10.1.0.2
    docker -H :3380 run -ti -d --net=dnet1 --name test2 centos # 10.1.0.3

    # 启动容器指定192.168.0.0/24的网络
    docker -H :3380 run -ti -d --net=dnet2 --name test3 centos # 192.168.0.2
    docker -H :3380 run -ti -d --net=dnet2 --name test4 centos # 192.168.0.3

3.容器网络测试

daolinet默认容器网络规则为：***同一网络能够通信，不同网络不通直接通信***

    # 进入容器test1
    docker -H :3380 attach test1

    # 从容器test1中ping容器test2，*正常通信*
    >> ping 10.1.0.3
    # 从容器test1中ping容器test3和test4，*不能通信*
    >> ping 192.168.0.2
    >> ping 192.168.0.3

4.安全组

缺省情况下不同网段dnet1网络和dnet2网络是不能直接通信，接下来通过daolictl命令创建安全组及添加成员规则连通不同网络。

    # 创建一个安全组
    daolictl group create G1

    # 向安全组中添加网络
    daolictl member add --group G1 dnet1
    daolictl member add --group G1 dnet2
    daolictl group show G1
    # 再从容器test1中ping容器test3和test4，*正常通信*
    >> ping 192.168.0.2
    >> ping 192.168.0.3

    # 从安全组中删除网络
    daolictl member rm --group G1 dnet2
    # 再从容器test1中ping容器test3和test4，*不能通信*
    >> ping 192.168.0.2
    >> ping 192.168.0.3

5.细粒度控制

安全组可以控制不同网络之间的连通和隔离，充当路由器的角色，对于容器与容器之间的细粒度控制，我们提供如下操作完成。

    # 创建隔离规则
    daolictl cut test1:test2
    # 再从test1中ping容器test2，*不能通信*
    >> ping 10.1.0.3

    # 恢复通信
    daolictl uncut test1:test2
    # 再从test1中ping容器test2，*正常通信*
    >> ping 10.1.0.3

6.防火墙端口映射

如果容器中启动服务，可以通防火墙端口映射将服务port映射到服务器port以对外提供服务。

> **注意，请先登录agent节点添加服务镜像**
>
> 例如，进入agent-node节点下载ssh服务和apache服务镜像：
>
>       ssh agent-node
>       docker pull daolicloud/centos6.6-ssh
>       docker pull daolicloud/centos6.6-apache

    # 添加一条名为fw-ssh的规则，将容器testssh中ssh服务22端口映射到服务器20022端口
    daolictl firewall create --container testssh --rule 20022:22 fw-ssh
    # 访问容器ssh服务，<GATEWAY IP>为容器所在服务器ip地址
    daolictl firewall show testssh
    ssh <GATEWAY IP> -p 20022

    ＃ 添加一条名为fw-web的规则，将容器testweb中apache服务80端口映射到服务器20080端口
    daolictl firewall create --container testweb --rule 20080:80 fw-web
    # 访问容器apache服务，<GATEWAY IP>为容器所在服务器ip地址
    daolictl firewall show testweb
    curl -L http://<GATEWAY IP>:20080

    # 取消防火墙端口规则
    daolictl firewall delete fw-ssh fw-web
