### jq
> 官方文档: https://jqlang.github.io/jq/
> 

### 测试json
```shell
cat << EOF > param.json
[
  {
    "id": "5AA1FD09-268B-4A32-815A-137047E34FB6",
    "name": "兴趣点 1",
    "sub_items": [
      {
        "id": "C5ABC66D-EC9D-4212-8E29-A929EFF83353",
        "name": "加速度",
        "type": 3,
        "result_id": "816D1233-7DFB-478B-848D-06827496E83C"
      }
    ]
  },
  {
    "id": "6FBC7C91-FBAE-4AC2-A3F2-E2988D3F9FBA",
    "name": "兴趣点 2",
    "sub_items": [
      {
        "id": "4A4CD0D6-5986-4C97-8CB2-55A0FCC29FFF",
        "name": "位移",
        "type": 3,
        "result_id": "57CC591B-9B0C-4C72-B76F-43470DA3044E"
      }
    ]
  }
]
EOF
```

1. json格式化
```shell
cat param.json | jq
```

2. 查看字段
```shell
cat param.json | jq -j '.[0].name'
```

3. 遍历数组
```shell
cat param.json | jq -j '.[] | .name'
```

4. 过滤
```shell
cat param.json | jq 'map(select(.sub_items[0].name == "加速度"))' | jq '.[] | .name'
```

5. 迭代(遍历数组)
```shell
cat param.json|jq -j '.[] | .name'
```

6. 更新值
```shell
cat param.json| jq -j '.[0].name = "test01"' 
```