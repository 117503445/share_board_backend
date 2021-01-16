# share_board_backend

## 介绍

基于 Websocket 的 同步 Web 手写板

## 地址

前端 <https://github.com/117503445/share_board_frontend>

后端 <https://github.com/117503445/share_board_backend>

## 接口

### 客户端发送状态更改

Websocket /api/v1/ws

客户端发送

```json
{
    "boardid":"fdc39841-618e-4bd4-9799-2153a6fca002", // boardid 指示画板
    "pagenumber":1,
}
```

服务器返回:正常

```json
{
    "code": 1,
    "msg": "ok",
    "data":{} // data 为全部笔迹数据
}
```

服务器返回:白板不存在 / 页数不合法

```json
{
    "code": 2,
    "msg": "board not exists or page number is illegal.",
    "data":{} // data 为全部笔迹数据
}
```

### 客户端发送新增笔迹

Websocket /api/v1/ws

客户端发送

```json
{
    "type":"add",
    "data":{} // data 为新添加的笔迹数据
}
```

服务器不返回,向其他在此页上的客户端广播此请求,并在数据库中添加此笔迹

### 客户端接收新增笔迹

Websocket /api/v1/ws

服务器发送

```json
{
    "type":"add",
    "data":{} // data 为新添加的笔迹数据
}
```

客户端不返回,在本地画板添加此笔迹

### 客户端发送全部笔迹

Websocket /api/v1/ws

客户端发送

```json
{
    "type":"replace",
    "data":{} // data 为全部笔迹数据
}
```

服务器不返回,向其他在此页上的客户端广播此请求,并在数据库中替换数据为data

### 客户端接收全部笔迹

Websocket /api/v1/ws

服务器发送

```json
{
    "type":"replace",
    "data":{} // data 为全部笔迹数据
}
```

客户端不返回,本地画板替换为此笔记

## 返回状态码

- 1 正常
- 2 资源不存在

## Todo

- 验证密码

- 部署文档编写
