# go-example
## 项目说明

```
├── api 暴露给前端的api文档
├── build 构建参数
├── cmd 启动命令
├── configs 配置文件
├── docs 使用文档
├── examples 示例
├── internal 内部使用的包
│   ├── pkg  内部使用的封装
│   ├── smodel 模型模块
│   │   └── controller 控制层
│   │   └── errors 错误
│   │   └── service 业务层
│   │   └── testing 测试案例
│   │   └── router 内部路由
│   ├── router 路由
├── pkg 允许外部引用的封装
│   ├── ginwrap gin框架包装
│   ├── tools 工具
├── scripts
│   ├── install 安装脚本
│   ├── lib shell脚本库封装
├── web 前端单页应用, 资源等
├── Dockerfile 编译docker容器
├── go.mod  依赖管理
├── main.go 入口
└── README.md
```
## 代码规范
参考 https://github.com/yann66666/coding-standards/blob/main/README.md