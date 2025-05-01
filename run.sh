#!/bin/bash

echo "       -- Protobuf Generated Code --"
echo "1) For auth service"
echo "2) For email service"
echo "3) For notification service"
echo "4) For orchestrator service"
echo "5) For subscription service"
echo "       ----------------"
echo ""
echo "       ---- Auth -----"
echo "6) Run"
echo "7) Kill"
echo "       ---------------"
echo ""
echo "       --- Content ---"
echo "8) Run"
echo "9) Kill"
echo "       ---------------"
echo ""
echo "       ---- Email ----"
echo "10) Run"
echo "11) Kill"
echo "       ----------------"
echo ""
echo "       - Notification -"
echo "12) Run"
echo "13) Kill"
echo "       ----------------"
echo ""
echo "       - Orchestrator -"
echo "14) Run"
echo "15) Kill"
echo "       ------------------"
echo ""
echo "       -- Subscription --"
echo "16) Run"
echo "17) Kill"
echo "       ------------------"
echo ""
echo "       --- Containers ---"
echo "18) Up"
echo "19) Down"
echo "       ------------------"
echo ""

generate_proto() {
    local service=$1
    local file_name=$2
    local grpc_path="$service/pkg/grpc"
    local proto_path="$grpc_path/proto/$file_name"
    local output_path="$grpc_path/pb/"

    protoc --go_out="$output_path" --go-grpc_out="$output_path" "$proto_path"
}

kill_process() {
    local rest_port=$1
    local grpc_port=$2

    if [[ "$rest_port" != -1 ]]; then
        echo "🛑 Killing process on port $rest_port..."
        sudo kill -9 $(lsof -t -i:$rest_port 2>/dev/null) 2>/dev/null || true
    fi

    if [[ "$grpc_port" != -1 ]]; then
        echo "🛑 Killing process on port $grpc_port..."
        sudo kill -9 $(lsof -t -i:$grpc_port 2>/dev/null) 2>/dev/null || true
    fi
}

start_service() {
    local service=$1

    cd "$service/cmd" || exit
    echo "Downloading dependencies for $service service..."
    go mod download
    echo "Starting $service service..."
    go run main.go
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
elif [[ $cmd == 6 ]]; then
    start_service "auth"
elif [[ $cmd == 7 ]]; then
    kill_process 8000 50000
elif [[ $cmd == 8 ]]; then
    start_service "content"
elif [[ $cmd == 9 ]]; then
    kill_process 8001 50001
elif [[ $cmd == 10 ]]; then
    start_service "email"
elif [[ $cmd == 11 ]]; then
    kill_process -1 50003
elif [[ $cmd == 12 ]]; then
    start_service "notification"
elif [[ $cmd == 13 ]]; then
    kill_process -1 -1
elif [[ $cmd == 14 ]]; then
    start_service "orchestrator"
elif [[ $cmd == 15 ]]; then
    kill_process 8080 -1
elif [[ $cmd == 16 ]]; then
    start_service "subscription"
elif [[ $cmd == 17 ]]; then
    kill_process -1 50002
elif [[ $cmd == 18 ]]; then
    docker compose -f docker-compose.yml up -d
elif [[ $cmd == 19 ]]; then
    docker compose -f docker-compose.yml down
fi
