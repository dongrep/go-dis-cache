# Go Distributed Cache

A distributed cache implementation in Go, featuring custom in-memory data structures, LRU and TTL mechanisms, and peer-to-peer distribution.

Repository URL: https://github.com/dongrep/go-dis-cache

## Features

- **Custom In-Memory Storage**: Utilizes Go's map data structure for efficient caching.
- **LRU (Least Recently Used)**: Keeps the cache updated by prioritizing recently accessed elements.
- **TTL (Time to Live)**: Ensures cache entries expire after a specified duration.
- **HTTP Protocol**: Enables easy integration and standardized communication.
- **Peer-to-Peer Distribution**: Supports eventual consistency across multiple instances.
- **Consistent Hashing**: Allows seamless addition of new instances for data sharding.

## Installation

1. Ensure you have Go installed on your system.

2. Clone the repository:

   ```
   git clone https://github.com/dongrep/go-dis-cache.git
   ```

3. Navigate to the project directory:

   ```
   cd go-dis-cache
   ```

4. Build the application:
   ```
   go build
   ```

## Usage

Run the application using the following command:

```
./go-dis-cache -port=<PORT> -peers=<PEER_LIST>
```

Where:

- `<PORT>` is the port number for this instance to listen on
- `<PEER_LIST>` is a comma-separated list of other instances' addresses

Example:

```
./go-dis-cache -port=8000 -peers=localhost:8001,localhost:8002
```

This starts an instance on port 8000, connected to peers on ports 8001 and 8002.

## Important Notes

- Each node must have a unique port number.
- The `-peers` flag should list all other instances, excluding the current one.
- Ensure all listed peer instances are running before starting a new node.

## Implementation Details

The distributed cache system includes:

- In-memory storage using Go's map data structure
- Custom LRU mechanism for cache entry management
- TTL-based expiration system
- HTTP server for inter-instance communication
- Consistent hashing algorithm for data sharding

## Contributing

Contributions are welcome. Please follow Go code style guidelines and include tests for new features.

## Acknowledgements

This project was inspired by the book "System Programming Essentials using Go", which provided valuable insights into building robust systems.
