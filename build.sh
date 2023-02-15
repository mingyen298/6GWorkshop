# sudo docker rmi 6g-workshop:model_storage 6g-workshop:aiml_mitlab 6g-workshop:inference_node

sudo docker rm -f ws-model_storage
sudo docker rmi 6g-workshop:model_storage
sudo docker build -t 6g-workshop:model_storage -f Dockerfile.model_storage .

sudo docker rm -f ws-aiml_mitlab
sudo docker rmi 6g-workshop:aiml_mitlab
sudo docker build -t 6g-workshop:aiml_mitlab -f Dockerfile.aiml_mitlab .

sudo docker rm -f ws-inference_node
sudo docker rmi 6g-workshop:inference_node
rm -rf role/inference_node/main
go build -o role/inference_node/main role/inference_node/main.go
sudo docker build -t 6g-workshop:inference_node -f Dockerfile.inference_node .