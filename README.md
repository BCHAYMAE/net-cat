# Net-Cat

A TCP group chat server written in Go, inspired by the NetCat utility.

Net-Cat allows multiple clients to connect to a shared chat room over TCP, exchange messages in real time, and receive previous messages from the conversation.

## Features

- TCP server with a default port (`8989`).
- Custom port configuration.
- Multiple simultaneous client connections.
- Maximum of 10 connected clients.
- Username registration and validation.
- Unique usernames.
- Username length limited to 20 characters.
- Real-time message broadcasting.
- Message timestamps and usernames.
- Message history for newly connected clients.
- Join and leave notifications.
- Empty messages are ignored.
- Concurrent client handling using goroutines.
- Shared data protection using mutexes.

## Requirements

- Go installed on your system.
- NetCat (`nc`) installed to connect to the server.

## Installation

Clone the repository:

```bash
git clone https://github.com/BCHAYMAE/net-cat.git
```

Navigate to the project directory:

```bash
cd net-cat
```

Build the project:

```bash
go build -o TCPChat .
```

## Usage

Start the server using the default port:

```bash
./TCPChat
```

Start the server using a custom port:

```bash
./TCPChat 2525
```

The server listens on port `8989` by default.

## Connecting to the Server

Open another terminal and connect using NetCat:

```bash
nc localhost 8989
```

For a custom port:

```bash
nc localhost 2525
```

Enter a username when prompted:

```text
[ENTER YOUR NAME]: User1
```

You can now exchange messages with other connected clients.

## Message Format

Messages are displayed using the following format:

```text
[2026-11-04 20:30:15][User1]: Hello everyone!
```

The server also displays notifications when clients join or leave:

```text
[User1] has joined our chat...
[User1] has left our chat...
```

## Testing

Run the automated tests:

```bash
go test -v
```

Run the server with Go's race detector:

```bash
go run -race .
```

Check the code using Go Vet:

```bash
go vet ./...
```

## Project Structure

```text
net-cat/
├── main.go
├── main_test.go
├── go.mod
└── README.md
```

- `main.go`: TCP server, client handling, message broadcasting, and validation.
- `main_test.go`: Unit tests for port and username validation.
- `go.mod`: Go module configuration.
- `README.md`: Project documentation.

## Concepts Used

This project demonstrates several Go programming concepts:

- TCP networking with the `net` package.
- Goroutines for concurrent client handling.
- Mutexes for protecting shared data.
- Slices for managing connected clients and message history.
- Buffered input with `bufio`.
- Error handling.
- Unit testing with the `testing` package.

## License

This project was created for educational purposes.