# API Gateway Documentation

## Overview

This project is an API Gateway built using the Gin framework in Go. The gateway provides centralized routing, rate limiting, authentication, and service discovery via Consul. It is designed for scalability, efficiency, and security, following modern microservices architectural principles.

---

## Features

### **1. Service Routing Middleware**
- **Purpose**: Dynamically forwards incoming requests to appropriate backend services based on predefined routes or via service discovery using Consul.
- **How it works**:
  - Checks the requested path against predefined routes or dynamically queries Consul for the backend service.
  - Uses a reverse proxy to forward the request to the backend.
  - Handles errors such as invalid paths or unreachable services gracefully.
- **Use Case**: Enables seamless request forwarding in a microservices architecture.

---

### **2. JWT Authentication Middleware**
- **Purpose**: Ensures that only authenticated users can access protected routes by validating JSON Web Tokens (JWTs).
- **How it works**:
  - Checks for the `Authorization` header in incoming requests.
  - Verifies the token's signature and ensures it hasn't expired.
  - Extracts user-specific claims, such as `userID`, and makes them available in the request context.
- **Use Case**: Secures sensitive APIs by enforcing authentication.

---

### **3. Rate Limiting Middleware**
- **Purpose**: Controls the rate of incoming requests to prevent abuse and ensure fair usage of resources.
- **How it works**:
  - Implements a token bucket algorithm to manage request limits.
  - Refills tokens at a predefined rate, allowing for controlled bursts.
  - Responds with `429 Too Many Requests` when the limit is exceeded.
- **Use Case**: Prevents denial-of-service attacks and ensures resource availability.

---

### **4. Health Check Endpoint**
- **Purpose**: Provides a simple endpoint to check if the API Gateway is running.
- **How it works**:
  - Responds to a specific path (e.g., `/ping`) with a success message.
- **Use Case**: Useful for monitoring tools and uptime checks.

---

### **5. Service Routing with Consul**
- **Purpose**: Automatically discovers and routes requests to microservices registered in Consul.
- **How it works**:
  - Queries Consul to find the service address and port based on the incoming request path.
  - Uses a reverse proxy to forward requests to the discovered service.
  - Handles cases where the service is not found or unreachable.
- **Use Case**: Simplifies microservices communication by eliminating hardcoded URLs.

---

## How to Run the API Gateway

1. **Set Up Consul**:
   - Ensure a Consul server is running and accessible from the API Gateway.
   - Register all backend services with Consul, specifying service names, addresses, and ports.

2. **Run the API Gateway**:
   - Compile and start the API Gateway.
   - Configure routes or rely on Consul for service discovery.

3. **Access the Health Check Endpoint**:
   - Use `/ping` to verify that the gateway is up and running.

---

## Testing the API Gateway

### **Authentication**
- Test authenticated routes by generating a JWT token and attaching it to the `Authorization` header.
- Validate that requests without a token or with an invalid/expired token are rejected.

### **Rate Limiting**
- Send multiple requests to a rate-limited endpoint.
- Verify that the gateway returns `429 Too Many Requests` after exceeding the configured limit.

### **Service Routing**
- Test requests to various backend services registered in Consul.
- Confirm that requests are forwarded to the correct services or rejected if the service is unavailable.

### **Error Handling**
- Test invalid requests (e.g., unregistered routes, unreachable services).
- Confirm that the gateway returns appropriate error messages and status codes.

---

## Project Structure

- **`middlewares/`**: Contains all middleware implementations.
- **`main.go`**: Initializes the API Gateway, configures middlewares, and starts the server.
- **`internal/common/`**: Contains shared configurations and utility functions.

---

## Future Enhancements

1. **Circuit Breaker Middleware**:
   - Automatically disable requests to failing services to improve system reliability.

2. **Load Balancing**:
   - Distribute requests across multiple instances of a service.

3. **Improved Logging and Monitoring**:
   - Add structured logging and integrate with monitoring tools like Prometheus and Grafana.

4. **Enhanced Security**:
   - Support for roles-based access control (RBAC) and encrypted communication between services.

---

## Conclusion

This API Gateway provides a robust foundation for managing microservices in a secure, scalable, and efficient manner. It demonstrates expertise in middleware design and modern API gateway principles. This project is well-suited for real-world applications and showcases strong engineering practices.

