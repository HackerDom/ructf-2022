#!/bin/bash

cp ../../services/kleptophobia/proto/grpc.proto proto/grpc.proto
cp ../../services/kleptophobia/proto/models.proto proto/models.proto
sed -i 's/proto\/models.proto/models.proto/' proto/grpc.proto

python3 -m grpc_tools.protoc -Iproto --python_out=models proto/models.proto
python3 -m grpc_tools.protoc -Iproto --grpc_python_out=models proto/grpc.proto

sed -i 's/import models_pb2 as models__pb2/import models.models_pb2 as models__pb2/' models/grpc_pb2_grpc.py