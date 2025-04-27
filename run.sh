#!/bin/bash

echo "            ------------ Protobuf Generated Code ------------"
echo "1) For auth service"
echo "2) For email service"
echo "3) For notification service"
echo "4) For orchestrator service"
echo "5) For subscription service"
echo "                       ---------------------------"

generate_proto() {
    local service=$1
    local file_name=$2
    local grpc_path="$service/pkg/grpc"
    local proto_path="$grpc_path/proto/$file_name"
    local output_path="$grpc_path/pb/"

    protoc --go_out="$output_path" --go-grpc_out="$output_path" "$proto_path"
}

read -p "Type: " cmd
if [[ $cmd == 1 ]]; then
    generate_proto "auth" "auth-manager.proto"
elif [[ $cmd == 2 ]]; then
    generate_proto "email" "email-service.proto"
elif [[ $cmd == 3 ]]; then
    generate_proto "notification" "subscription-manager.proto"
elif [[ $cmd == 4 ]]; then
    generate_proto "orchestrator" "auth-manager.proto"
    generate_proto "orchestrator" "subscription-manager.proto"
elif [[ $cmd == 5 ]]; then
    generate_proto "subscription" "publication-authorization.proto"
    generate_proto "subscription" "subscription-manager.proto"
fi
