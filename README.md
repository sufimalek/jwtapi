### `README.md`

### Directory Structure

```
jwtapi/
├── cmd/
│   └── jwtapi/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   │   └── auth.go
│   │   │   └── user.go
│   │   ├── middleware/
│   │   │   └── jwt.go
│   │   └── routes.go
│   ├── config/
│   │   └── config.go
│   ├── models/
│   │   └── user.go
│   ├── repository/
│   │   └── user_repository.go
│   ├── service/
│   │   └── auth_service.go
│   │   └── user_service.go
│   └── utils/
│       └── erros.go
│       └── jwt.go
│       └── logger.go
├── migrations/
│   └── 001_create_users_table.sql
├── go.mod
├── go.sum
└── README.md
```


```markdown
# jwtapi - Go Application with JWT Authentication and MySQL

This is a production-ready Go application with JWT authentication, user management, and MySQL integration. It uses `gorilla/mux` for routing and follows a well-structured directory layout.

---

## Features

- **JWT Authentication**: Secure login and token-based authentication.
- **User Management**: Create, update, delete, and list users.
- **Logging**: Basic logging using the standard `log` package.
- **Error Handling**: Custom error handling with meaningful error messages.

---

## Setup

### Prerequisites

- Go 1.20 or higher
- MySQL 8.0 or higher
- Environment variables set up (see `.env` example below)

### Environment Variables

Create a `.env` file in the root directory with the following variables:

```bash
DB_USERNAME=root
DB_PASSWORD=password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=jwtapi
SERVER_PORT=8080
JWT_SECRET=my_secret_key
```

### Database Setup

1. Create a MySQL database:

   ```sql
   CREATE DATABASE jwtapi;
   ```

2. Run the migration to create the `users` table:

   ```sql
   CREATE TABLE users (
       id INT AUTO_INCREMENT PRIMARY KEY,
       username VARCHAR(255) NOT NULL UNIQUE,
       password VARCHAR(255) NOT NULL,
       email VARCHAR(255) NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
   );
   ```

### Install Dependencies

Run the following command to install the required dependencies:

```bash
go mod tidy
```

---

## Running the Application

To start the application, run:

```bash
go run cmd/jwtapi/main.go
```

The server will start on port `8080` by default.

---

## API Endpoints

### API Endpoint for User Registration

#### Register User
- **URL**: `/register`
- **Method**: `POST`
- **Request Body**:
  ```json
    {
        "username": "newuser",
        "password": "newpassword",
        "email": "newuser@example.com"
    }
  ```
- **Response**:
  ```json
    {
        "id": 1,
        "username": "newuser",
        "email": "newuser@example.com",
        "created_at": "2023-10-01T12:00:00Z",
        "updated_at": "2023-10-01T12:00:00Z"
    }
  ```


- **Example curl request:**:
    ```json
        curl -X POST http://localhost:8080/register \
        -H "Content-Type: application/json" \
        -d '{
            "username": "newuser",
            "password": "newpassword",
            "email": "newuser@example.com"
        }'
    ```

### Authentication

#### Login
- **URL**: `/login`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
      "username": "testuser",
      "password": "testpassword"
  }
  ```
- **Response**:
  ```json
  {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
  ```

### User Management (Protected Routes)

#### Create User
- **URL**: `/api/users`
- **Method**: `POST`
- **Headers**: `Authorization: Bearer <token>`
- **Request Body**:
  ```json
  {
      "username": "newuser",
      "password": "newpassword",
      "email": "newuser@example.com"
  }
  ```
- **Response**:
  ```json
  {
      "id": 1,
      "username": "newuser",
      "email": "newuser@example.com",
      "created_at": "2023-10-01T12:00:00Z",
      "updated_at": "2023-10-01T12:00:00Z"
  }
  ```

#### Update User
- **URL**: `/api/users/{id}`
- **Method**: `PUT`
- **Headers**: `Authorization: Bearer <token>`
- **Request Body**:
  ```json
  {
      "username": "updateduser",
      "email": "updateduser@example.com"
  }
  ```
- **Response**:
  ```json
  {
      "id": 1,
      "username": "updateduser",
      "email": "updateduser@example.com",
      "created_at": "2023-10-01T12:00:00Z",
      "updated_at": "2023-10-01T12:30:00Z"
  }
  ```

#### Delete User
- **URL**: `/api/users/{id}`
- **Method**: `DELETE`
- **Headers**: `Authorization: Bearer <token>`
- **Response**: `204 No Content`

#### List Users
- **URL**: `/api/users`
- **Method**: `GET`
- **Headers**: `Authorization: Bearer <token>`
- **Response**:
  ```json
  [
      {
          "id": 1,
          "username": "user1",
          "email": "user1@example.com",
          "created_at": "2023-10-01T12:00:00Z",
          "updated_at": "2023-10-01T12:00:00Z"
      },
      {
          "id": 2,
          "username": "user2",
          "email": "user2@example.com",
          "created_at": "2023-10-01T12:30:00Z",
          "updated_at": "2023-10-01T12:30:00Z"
      }
  ]
  ```

---

## Logging

Logs are printed to the console with the prefix `jwtapi:`. You can replace this with a more advanced logging library like `zap` or `logrus` if needed.

---

## Error Handling

Errors are returned with a status code and a meaningful message. For example:

```json
{
    "error": "Failed to create user",
    "message": "username already exists"
}
```

---

## Testing

To test the application, you can use tools like [Postman](https://www.postman.com/) or [curl](https://curl.se/).

---


# **Logging Setup with Grafana, Loki, and Promtail**

## **Table of Contents**
1. [Overview](#overview)
2. [Prerequisites](#prerequisites)
3. [Setup](#setup)
   - [Docker Compose Configuration](#docker-compose-configuration)
   - [Promtail Configuration](#promtail-configuration)
   - [Loki Configuration](#loki-configuration)
   - [Grafana Configuration](#grafana-configuration)
4. [Running the Stack](#running-the-stack)
5. [Viewing Logs in Grafana](#viewing-logs-in-grafana)
6. [Troubleshooting](#troubleshooting)
7. [References](#references)

---

## **Overview**
- **Promtail**: Collects logs from your `jwtapi` API and sends them to Loki.
- **Loki**: Stores logs and allows querying them using LogQL.
- **Grafana**: Visualizes logs stored in Loki.

---

## **Prerequisites**
- Docker and Docker Compose installed.
- Basic understanding of Docker and YAML configuration.
- Your `jwtapi` API configured to write logs to a file (e.g., `/var/log/jwtapi/jwtapi.log`).

---

## **Setup**

### **Docker Compose Configuration**
The `docker-compose.yml` file defines the services for `jwtapi`, Promtail, Loki, and Grafana. Ensure the following configuration is present:

```yaml
version: '3.7'

services:
  jwtapi:
    image: jwtapi
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    volumes:
      - /var/log/jwtapi:/var/log/jwtapi
    env_file:
      - .env
    depends_on:
      - loki

  promtail:
    image: grafana/promtail:latest
    volumes:
      - ./promtail-config.yaml:/etc/promtail/config.yml
      - /var/log/jwtapi:/var/log/jwtapi
    command: -config.file=/etc/promtail/config.yml
    depends_on:
      - loki

  loki:
    image: grafana/loki:latest
    ports:
      - "3100:3100"
    volumes:
      - ./loki-config.yaml:/etc/loki/local-config.yaml
      - loki_data:/tmp/loki
    command: -config.file=/etc/loki/local-config.yaml

  grafana:
    image: grafana/grafana:latest
    ports:
      - "3000:3000"
    environment:
      - GF_SECURITY_ADMIN_PASSWORD=admin
    volumes:
      - grafana_data:/var/lib/grafana
      - ./provisioning:/etc/grafana/provisioning
    depends_on:
      - loki

volumes:
  loki_data:
  grafana_data:
```

---

### **Promtail Configuration**
Promtail is responsible for collecting logs from your `jwtapi` API and sending them to Loki. Create a `promtail-config.yaml` file with the following content:

```yaml
server:
  http_listen_port: 9080
  grpc_listen_port: 0

positions:
  filename: /tmp/positions.yaml

clients:
  - url: http://loki:3100/loki/api/v1/push

scrape_configs:
  - job_name: jwtapi
    static_configs:
      - targets:
          - localhost
        labels:
          job: jwtapi
          __path__: /var/log/jwtapi/*.log
```

---

### **Loki Configuration**
Loki stores the logs collected by Promtail. Create a `loki-config.yaml` file with the following content:

```yaml
auth_enabled: false

server:
  http_listen_port: 3100

common:
  path_prefix: /tmp/loki
  storage:
    filesystem:
      chunks_directory: /tmp/loki/chunks
      rules_directory: /tmp/loki/rules
  replication_factor: 1
  ring:
    kvstore:
      store: inmemory

schema_config:
  configs:
    - from: 2020-10-24
      store: boltdb-shipper
      object_store: filesystem
      schema: v11
      index:
        prefix: index_
        period: 168h

filesystem:
  directory: /tmp/loki/chunks

limits_config:
  reject_old_samples: true
  reject_old_samples_max_age: 168h
  ingestion_rate_mb: 16
  ingestion_burst_size_mb: 32
  allow_structured_metadata: false
```

---

### **Grafana Configuration**
Grafana is used to visualize logs stored in Loki. To automate the setup of the Loki data source, create a `provisioning/datasources/loki.yaml` file:

```yaml
apiVersion: 1

datasources:
  - name: Loki
    type: loki
    access: proxy
    url: http://loki:3100
    isDefault: true
    version: 1
    editable: true
```

---

## **Running the Stack**
1. Build and start the services:
   ```bash
   docker-compose up --build
   ```

2. Verify that all services are running:
   ```bash
   docker-compose ps
   ```

---

## **Viewing Logs in Grafana**
1. Open Grafana in your browser: `http://localhost:3000`.
2. Log in with the username `admin` and password `admin`.
3. Go to **Explore**.
4. Select the **Loki** data source.
5. Query logs using LogQL. For example:
   ```logql
   {job="jwtapi"}
   ```
   This will display logs from your `jwtapi` API.

---

## **Troubleshooting**
### **1. Logs Not Appearing in Grafana**
- Ensure Promtail is running and collecting logs:
  ```bash
  docker-compose logs promtail
  ```
- Verify that Loki is receiving logs:
  ```bash
  docker-compose logs loki
  ```

### **2. Grafana Data Source Not Found**
- Check the Grafana logs for errors:
  ```bash
  docker-compose logs grafana
  ```
- Ensure the `provisioning/datasources/loki-datasource.yaml` file is correctly mounted.

### **3. Log File Not Found**
- Verify that your `jwtapi` API is writing logs to `/var/log/jwtapi/jwtapi.log`.
- Check the file permissions:
  ```bash
  docker-compose exec jwtapi ls -l /var/log/jwtapi
  ```

---

## **References**
- [Grafana Documentation](https://grafana.com/docs/)
- [Loki Documentation](https://grafana.com/docs/loki/latest/)
- [Promtail Documentation](https://grafana.com/docs/loki/latest/clients/promtail/)

---

This setup will allow you to collect, store, and visualize logs from your `jwtapi` API using Grafana, Loki, and Promtail. Let me know if you need further assistance!



---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
```
