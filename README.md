#### Fibermongo
基于fiber + mongodb 构建的 REST API 

```
#插入数据
curl http://127.0.0.1:2022/user -X POST --header "Content-Type:application/json" -d '{"name": "cc","location": "Guangzhou","title": "Phper"}'

#通过ID查询一条数据
curl  http://127.0.0.1:2022/user/62a01b1a050e209164550288

#修改数据
curl http://127.0.0.1:2022/user/62a01b1a050e209164550288 -X PUT --header "Content-Type:application/json" -d '{"name": "dd","location": "Beijing","title": "Gopher"}'

#删除数据
curl http://127.0.0.1:2022/user/62a01b1a050e209164550288 -X DELETE

#查询所有数据
curl  http://127.0.0.1:2022/users
```