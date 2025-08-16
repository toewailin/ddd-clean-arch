#!/bin/bash

# Step 1: Create the main project folder
echo "Creating the project directory..."
mkdir ddd-clean-arch
cd ddd-clean-arch

# Step 2: Initialize Go module for the main project
echo "Initializing Go module for the main project..."
go mod init ddd-clean-arch

# Step 3: Create the main folder structure
echo "Creating main folder structure..."
mkdir build certs cmd config deploy logs pkg script internal

# Create subdirectories under internal
mkdir -p internal/modules/{admin,bet,faq,game,order,payment,player,product,promotion,stock,user}
mkdir -p internal/router

# Step 4: Create config files
echo "Creating config files..."
mkdir -p config
touch config/config.go config/database.go config/redis.go

# Step 5: Create deploy folder with Docker and Kubernetes files
echo "Creating deploy files..."
mkdir -p deploy/docker deploy/k8s
touch deploy/docker/docker-compose.yml deploy/docker/Dockerfile deploy/docker/Dockerfile.dev

# Step 6: Create internal module files
echo "Creating internal module files..."

# FAQ Module
echo "Creating FAQ module files..."
mkdir -p internal/modules/faq
touch internal/modules/faq/{entity.go,handler.go,repository.go,router.go,service.go}

# Product Module
echo "Creating Product module files..."
mkdir -p internal/modules/product
touch internal/modules/product/{entity.go,handler.go,repository.go,router.go,service.go}

# Order Module
echo "Creating Order module files..."
mkdir -p internal/modules/order
touch internal/modules/order/{dto.go,entity.go,handler.go,repository.go,route.go,service.go}

# User Module
echo "Creating User module files..."
mkdir -p internal/modules/user
touch internal/modules/user/{entity.go,handler.go,repository.go,route.go,service.go}

# Other modules (bet, game, payment, player, promotion, stock, admin) will have a similar setup:
echo "Creating other modules..."
for module in admin bet game payment player promotion stock; do
    mkdir -p "internal/modules/$module"
    touch "internal/modules/$module/{entity.go,handler.go,repository.go,router.go,service.go}"
done

# Step 7: Create the router file
echo "Creating router file..."
touch internal/router/router.go

# Step 8: Create pkg directory files
echo "Creating pkg directory files..."

# Middleware
mkdir -p pkg/middleware
touch pkg/middleware/auth.go

# Logger
mkdir -p pkg/logger
touch pkg/logger/logger.go

# Utils
mkdir -p pkg/utils
touch pkg/utils/{i18n.go,jwt.go,success.go,utils.go,validation.go}

# Response
mkdir -p pkg/response
touch pkg/response/{error.go,pagination.go,success.go}

# Libraries
mkdir -p pkg/libs
touch pkg/libs/idgenerator.go

# Step 9: Create logs directory
echo "Creating logs directory..."
mkdir -p logs
touch logs/app.log

# Step 10: Create the main API gateway entry point
echo "Creating cmd/main.go..."
mkdir -p cmd/api
touch cmd/api/main.go

# Step 11: Create README file
echo "Creating README file..."
touch README.md

# Step 12: Create the upload.sh script file
echo "Creating upload.sh script file..."
mkdir -p script
touch script/upload.sh

echo "Project setup completed!"

# Step 13: Open VS Code (optional)
# Uncomment the following line to open the project in VS Code
# code .
