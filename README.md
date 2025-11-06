# Fybrk Relay Server

Ultra-lightweight relay server for Fybrk file synchronization when direct P2P connections fail.

## Quick Start

### Docker (Recommended)
```bash
docker run -p 8080:8080 fybrk/relay
```

### Go Binary
```bash
go run main.go
```

## Deployment

### Hetzner Cloud (€3.29/month)
```bash
# Create server
hcloud server create --type cx11 --image ubuntu-22.04 --name fybrk-relay

# Deploy with Docker
ssh root@your-server
docker run -d -p 80:8080 --restart unless-stopped fybrk/relay
```

### Fly.io ($1.94/month)
```bash
fly launch
fly deploy
```

## Health Check
```bash
curl http://localhost:8080/health
```

## Protocol

WebSocket endpoint: `/relay`

Messages:
- `{"type": "register", "device_id": "abc123"}` - Register device
- `{"type": "relay", "target": "def456", "data": {...}}` - Relay message

The server simply forwards messages between registered devices.
