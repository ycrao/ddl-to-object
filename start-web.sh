#!/bin/bash

# DDL to Object Web Service Startup Script

set -e

# Color definitions
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${BLUE}========================================"
echo -e "    DDL to Object Web Service"
echo -e "========================================${NC}"
echo

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${YELLOW}Warning: Go is not installed, please install Go first${NC}"
    exit 1
fi

# Enter web directory
cd web

echo -e "${GREEN}🚀 Starting Web Server...${NC}"
echo -e "${BLUE}📱 Open in browser: http://localhost:8080${NC}"
echo -e "${BLUE}🔗 API endpoint: http://localhost:8080/api/convert${NC}"
echo
echo -e "${YELLOW}Press Ctrl+C to stop the server${NC}"
echo

# Start server
go run server.go