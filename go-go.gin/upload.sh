#!/bin/bash
curl -X POST http://127.0.0.1:8080/upload -F "file=@./1.txt" -H "Content-Type: multipart/form-data"

curl -X POST http://127.0.0.1:8080/multi-upload -F "upload[]=@./2.txt" -F "upload[]=@./3.txt" -H "Content-Type: multipart/form-data"