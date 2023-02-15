# 6G Workshop


### role

#### model storage(4504)
用來放存放及取用模型

### inference node(4501、4502)
process1(4501): 主要工作:收xApp資料並即時傳回datalake及接收更新模型的通知，收到通知後將下載新模型並發請求讓process2 reload model
process2(4502): 主要工作:接收xApp的資料作即時的推論

#### aiml_mitlab
將xApp傳來的資料寫入datalake、通知inference node有新模型



### build images
```bash=
bash build.sh
```

### run
```bash=
bash run.sh
```


### test api

#### inference 

```bash=
curl -X POST {IP}:4502/inference_node/model/predict -H "Content-Type:application/json" -d '{"input":[4,5,6]}'
```

#### upload data 

```bash=
curl -X POST {IP}:4501/inference_node/data/upload -H "Content-Type:application/json" -d '{"data":[4,5,6]}'
```

#### update model

```bash=
curl -X POST {IP}:4503/aiml_mitlab/model/update/1
```