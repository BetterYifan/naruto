![discovery](3.png)

1. 实例定位：通过appid(唯一，服务名)+hostname定位实例

2. Node: discovery的服务节点

3. Provider: 服务提供者，服务部署在容器后，发起register请求给Discovery Server，后定期30s心跳一次

4. Consumer: 服务消费者，启动后拉取node内存中的APPID对应的容器节点信息