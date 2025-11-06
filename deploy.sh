#!/bin/bash

echo "🚀 Deploying Fybrk Relay Server"

# Build and start
docker-compose up -d --build

echo "✅ Relay server deployed!"
echo "📍 Health check: curl http://localhost/health"
echo "🔗 WebSocket endpoint: ws://localhost/relay"
echo ""
echo "💡 To use custom relay in fybrk:"
echo "   Edit ~/.fybrk/config.json and change relay_servers to:"
echo '   ["ws://your-server.com/relay"]'
