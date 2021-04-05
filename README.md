## mbui

#### 介绍

该项目依赖 maxwell 采集 mysql 二进制日志，元数据暂存到 redis，然后异步消费 redis 元数据，导入到 mysql，提供展示和查询功能。

#### 安装依赖

- 安装 maxwell 。配置相关账号。[参考链接](https://maxwells-daemon.io/quickstart/)

- 启动 maxwell 。
```
maxwell --user='maxwell' --password='maxwell' --host='127.0.0.1' --port=3306 --producer=redis --redis_type=lpush --output_ddl=true --output_row_query=true --output_binlog_position=true
```

#### 使用

- 编辑配置文件 config/config.yml

- go run main.go