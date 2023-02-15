#!/bin/bash
 
# 啟動 Apache
./main &

# 啟動 top
python main.py
 
# 無窮迴圈
while [[ true ]]; do
    sleep 1
done