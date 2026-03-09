### TCP Key-Value Store in Go

This project is a simple in-memory key-value store running on top of TCP.

Clients connect to the server over TCP and send plain text protocol commands.

#### Run the server
```bash
go run ./cmd
```

#### Protocol commands
- `PUT <key> <value>`: Store a value for a key.
- `GET <key>`: Retrieve the value for a key.
- `DEL <key>`: Delete a key-value pair.
- `LIST`: List all key-value pairs as `key=value` lines.

#### Connect and send commands
To interact with the server, use any TCP client tool such as `telnet` or `nc` (netcat), then pass your commands.

Examples:
- One-shot with `nc`: `echo "PUT name sean" | nc localhost 8080`
- One-shot with `nc`: `echo "GET name" | nc localhost 8080`
- Interactive with `telnet`: `telnet localhost 8080` and then type commands like `PUT city sf`, `GET city`, or `LIST`.
