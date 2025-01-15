# TinyLetter

TinyLetter is a microservices-based scalable newsletter infrastructure designed to streamline the process of creating, managing, and delivering newsletters to subscribers. The project leverages a distributed architecture to ensure modularity, scalability, and ease of maintenance.

---

## **Features**

- **User Management**: Registration, login, and role-based access control for publishers, subscribers, and admins.
- **Content Management**: Create and manage publications and posts with support for premium content.
- **Subscription Management**: Manage subscriptions for both publishers and subscribers, with support for multiple subscription plans.
- **Notification System**: Automatically notify subscribers about new posts via email.
- **Broadcast Service**: Rate-limited email dispatch to ensure efficient and reliable delivery.
- **Message Queue**: Decoupled communication between services using RabbitMQ for scalability.

---

## **Architecture Overview**

### **Microservices**

1. **User Service**
   - Handles user registration and login.
   - Manages user roles and permissions.
   - Communicates with Subscription Service via gRPC for publisher onboarding.

2. **Content Service**
   - Manages publications and posts.
   - Publishes messages to RabbitMQ when new posts are created.

3. **Subscription Service**
   - Manages subscriptions for publishers and subscribers.
   - Provides APIs for retrieving subscribers for a given publication.

4. **Notification Service**
   - Consumes messages from RabbitMQ.
   - Identifies the target audience by querying Subscription Service.
   - Sends tasks to the Broadcast Service for email dispatch.

5. **Broadcast Service**
   - Handles rate-limiting.
   - Sends emails to subscribers.

---

## **Tech Stack**

- **Programming Language**: Go (Golang)
- **Database**: PostgreSQL
- **Message Queue**: RabbitMQ
- **Authentication**: JWT
- **Containerization**: Docker
- **Orchestration**: Kubernetes (Minikube for local testing)
- **API Gateway**: Traefik or Kong (optional)

---

## **Database Design**

### **Tables**

1. **Role**
   - `id`: Unique identifier
   - `name`: Name of the role

2. **User**
   - `id`, `name`, `email`, `password`

3. **User Role**
   - `user_id`, `role_id`

4. **Publication**
   - `id`, `publisher_id`, `created_at`, `updated_at`

5. **Post**
   - `id`, `publication_id`, `content`, `is_premium`, `created_at`, `updated_at`

6. **Publisher Subscription Plan**
   - `id`, `name`, `applicable_for`

7. **Publisher Subscription**
   - `user_id`, `plan_id`, `is_deleted`

8. **Subscriber Subscription**
   - `user_id`, `is_premium`, `publication_id`, `is_deleted`

9. **Audience**
   - `plan_id`, `size`

10. **Permission**
    - `name`, `plan_id`

---

## **Getting Started**

### **Prerequisites**

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)
- [Minikube](https://minikube.sigs.k8s.io/docs/)
- [RabbitMQ](https://www.rabbitmq.com/)

### **Setup Instructions**

1. Clone the repository:
   ```bash
   git clone https://github.com/nahK994/TinyLetter.git
   cd TinyLetter
   ```

2. Start DB using Docker Compose:
   ```bash
   docker compose -f db.yml up -d
   ```

3. Run services locally:
   ```bash
   go run ./cmd/user
   go run ./cmd/content
   go run ./cmd/subscription
   go run ./cmd/notification
   go run ./cmd/broadcast
   ```

4. Access RabbitMQ dashboard:
   ```
   http://localhost:15672
   ```

---

<!-- ## **Testing Locally**

1. Start Minikube:
   ```bash
   minikube start
   ```

2. Deploy services to Minikube:
   ```bash
   kubectl apply -f k8s/
   ```

3. Access the application via Minikube's IP address:
   ```bash
   minikube service list
   ```

--- -->

---

## **Contributing**

Contributions are welcome! Please open an issue or submit a pull request on the [GitHub repository](https://github.com/nahK994/TinyLetter).

---

## **License**

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## **Contact**

For any inquiries, please contact **Shomi Khan** at [nkskl6@gmail.com](mailto:nkskl6@gmail.com).
