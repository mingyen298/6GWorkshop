import os
import tensorflow as tf
from tensorflow import keras
from keras.models import load_model
from flask import Flask
from flask import render_template
from flask import request
from flask import Flask, request, jsonify
from tensorflow import keras
from keras.models import load_model
import numpy as np

flag = False
class ModelProcess():
  def __init__(self):
    self.save_model_path = 'models/model'

  def load_model(self):
    self.model = load_model(self.save_model_path)

  def predict(self, input_nplist):
    score = self.model.predict(input_nplist)
    return score

model = ModelProcess()
model.load_model()
################## Flask ######################
app = Flask(__name__)

@app.route('/inference_node/model/predict', methods=['POST']) 
def model_predict():
    input_list = request.json['input']
    input_nplist = np.array(input_list).astype(float).reshape(1,3,1)
    score = model.predict(input_nplist)[0][0]
    return {'score': f'''{score}'''}

@app.route('/inference_node/model/reload', methods=['POST']) 
def model_reload():
    model.load_model()
    return {'message': 'success'}

@app.route('/inference_node/model/test', methods=['GET']) 
def model_test():
    return {'message': 'hello world'}


if __name__ == '__main__':
    # port = os.environ.get('MODEL_PORT')
    # if port == "" :
    port = "4502"
    
    app.run(host='0.0.0.0', port=int(port))