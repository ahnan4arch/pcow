# TCP 透明代理

主要实现思路参考 [`cow`](https://github.com/cyfdecyf/cow) 能自动检测 IP 是否被墙, 如果被墙则使用二级代理. 二级代理目前只支持 shadowsocks

该工具实现了内网的自动[fan qiang], 同时保证了国内网站的访问速度, 也为你的 vps 节约有限的带宽/流量.

欢迎在 develop branch 进行开发并发送 pull request :)

## 快速开始

### 环境

pcow 工作在内网网关服务器, 配置 nat 端口跳转到 pcow 由它完成后续翻墙工作.

### 安装

在网关服务器(linux X86,ARM)执行(也可用于更新)
```
curl -L git.io/pcow | bash
```

### 配置
在 pcow 所在同级目录放置 `config` 文件, 内容如下(改成你的配置)

```
listen = :8888
proxy  = ss://encrypt_method:password@1.2.3.4:1985
```

配置 nat, 80, 443 交给 pcow 来自动翻墙, 当然你也可以把全部的 tcp 流量都交给 pcow 处理.

```
# 内网不用翻墙
iptables -t nat -A PREROUTING -p tcp -d 192.168.1.0/24 -j RETURN

# 80, 443 交给 pcow 处理
iptables -t nat -A PREROUTING -p tcp --dport 80 -j REDIRECT --to 8888
iptables -t nat -A PREROUTING -p tcp --dport 443 -j REDIRECT --to 8888
```

### 启动

```
./pcow
```

## 其他

* 本程序是我为公司技术部门提供的翻墙方案, 在 centos 6.x 服务器(网关)测试
* 目前还在实验阶段
* 暂不支持多个二级代理分流
* 请在小范围内使用, 避免故障导致公司大范围断网影响工作.
  > 比如在技术部门使用内部的路由器, 只有使用该路由的, 分配到翻墙网关
  > 或者由用户手工修改网关, 完成内测



