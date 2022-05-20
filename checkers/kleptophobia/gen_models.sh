#!/bin/bash

python3 -m grpc_tools.protoc -I models --python_out=models --grpc_python_out=models models/models.proto
sed -i 's/import models_pb2 as models__pb2/import models.models_pb2 as models__pb2/' models/models_pb2_grpc.py