# User Management gRPC Application

This is a simple user management gRPC application for managing user details and includes a search capability.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Build and Run](#build-and-run)
- [gRPC Service Endpoints](#grpc-service-endpoints)
- [gRPC Service Models](#grpc-service-models)

## Prerequisites

Before you begin, ensure you have the following installed on your system:

- Go 1.16 or higher
- Protocol Buffers compiler (`protoc`)
- gRPC Go plugin for the Protocol Buffers compiler

## Build and Run

### Step 1: Clone the Repository

```bash
git clone <repository-url>
cd <repository-directory>
```

### Step 2: Generate gRPC Code from Protocol Buffers
Ensure you have protoc and the protoc-gen-go plugin installed. If not, install them using the following commands:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Generate the Go code from the .proto file:

```bash
protoc --go_out=. --go-grpc_out=. proto/users.proto
```

### Step 3: Build the Application
Build the application using the following command:

```bash
go build -o <object_file>
```

### Step 4: Run the Application
Run the application using the following command:

```bash
./<object_file>
```

The server will start and listen on port 50051.

## gRPC Service Endpoints

### GetUserById
Retrieves a user by ID.

**Request**
```protobuf
message UserRequest {
    int32 id = 1;
}
```

**Response**

```protobuf
message UserResponse {
    User user = 1;
    string message = 2;
}
```

### GetUsersByIds
Retrieves multiple users by their IDs.

**Request**

```protobuf
message UsersRequest {
    repeated int32 ids = 1;
}
```

**Response (stream)**

```protobuf
message UserResponse {
    User user = 1;
    string message = 2;
}
```

### SearchUsers
Searches for users based on criteria.

**Request**

```protobuf
message SearchRequest {
    repeated SearchCriteria criteria = 1;
}
```

**Response (stream)**

```protobuf
message UserResponse {
    User user = 1;
    string message = 2;
}
```

## gRPC Service Models

### User
Represents a user entity.

```protobuf
message User {
    int32 id = 1;
    string fname = 2;
    string city = 3;
    uint64 phone = 4;
    float height = 5;
    bool married = 6;
}
```

### SearchCriteria
Defines a search criterion.

```protobuf
message SearchCriteria {
    string field = 1;  // e.g., "fname", "city", "phone", "height"
    string keyword = 2;
}
```

