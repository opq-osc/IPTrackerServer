很精简的用于窥屏探测插件的服务端

## 运行参数

`./tracker` 默认参数运行

`-h` 打印帮助
`-port` 运行端口号
`-verbose` 是否打印日志

## 接口

均是 GET 请求, 返回均为 200

### 添加访问记录

GET `/key`

添加此次记录，并绑定在 key, 可选参数(Query)`r`为跳转链接，不提供则是默认图片

如: `/key?r=https://github.com`

不管服务端发不发生错误都返回自定义的重定向链接或默认图片

### 获取记录

GET `/key.info`

获取标识为 key 的所有访问记录

```jsonc
{
  "code": 0, // 为1表示错误
  "msg": "ok", // code 为1时, 表示错误信息
  "result": [] // code 为0时存在，返回记录列表，没有记录则为`null`
}
```

### 检查 key

GET `/key.check`

检查 key 是否已经使用

```jsonc
{
  "code": 0, // 为1表示错误
  "msg": "ok", // code 为1时, 表示错误信息
  "result": true // code 为0时存在，true表示该key尚未使用，false表示已经使用
}
```

### 举例

按顺序进行

1. `GET /abc?r=https://github.com`
   跳转到对应地址

2. `GET /abc.info`

   ```json
   {
     "code": 0,
     "msg": "ok",
     "result": [
       {
         "ip": "127.0.0.1",
         "user_agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36 Edg/94.0.992.47",
         "timestamp": 1634190615
       }
     ]
   }
   ```

3. `GET /abc.check`

   ```json
   {
     "code": 0,
     "msg": "ok",
     "result": false
   }
   ```

4. `GET /abcd.info`

   ```json
   {
     "code": 0,
     "msg": "ok",
     "result": null
   }
   ```

5. `GET /abcd.check`
   ```json
   {
     "code": 0,
     "msg": "ok",
     "result": true
   }
   ```
