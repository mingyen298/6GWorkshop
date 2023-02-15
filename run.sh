
sudo docker rm -f ws-model_storage
sudo docker run --network workshop --name ws-model_storage -itd -p 4504:4504 6g-workshop:model_storage

sudo docker rm -f ws-aiml_mitlab
sudo docker run --network workshop --name ws-aiml_mitlab -itd -p 4503:4503 6g-workshop:aiml_mitlab

sudo docker rm -f ws-inference_node
sudo docker run --network workshop --name ws-inference_node -itd -p 4501:4501 -p 4502:4502 6g-workshop:inference_node
curl -X POST 192.168.2.129:4502/inference_node/model/predict -H "Content-Type:application/json" -d '{"input":[4,5,6]}'
curl -X POST 192.168.2.129:4503/aiml_mitlab/model/update/1
# curl -X POST 192.168.2.129:4503/aiml_mitlab/data/upload -H "Content-Type:application/json" -d '{"input":[4,5,6]}'