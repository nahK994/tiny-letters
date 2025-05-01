# TinyLetters

TinyLetter is a microservices-based scalable newsletter infrastructure designed to streamline the process of creating, managing, and delivering newsletters to subscribers. The project leverages a distributed architecture to ensure modularity, scalability, and ease of maintenance.

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
   - Sends tasks to the Email Service for email dispatch.
    - **Why is it needed?**  
     If a user wants to receive notifications via WhatsApp or other channels, the Notification Service will handle it.  

5. **Email Service**
   - Handles rate-limiting.
   - Sends emails to subscribers.

6. **Saga Orchestration Service**
   - Implements the Orchestration-based Saga pattern.
   - Manages distributed transactions across microservices.
   - Ensures data consistency through a series of coordinated steps.
   - Rolls back changes if any service fails during the transaction.

---

### **Diagram**

![Architecture Diagram](./architectural_diagram.jpg "Architecture Diagram")

---

## **Tech Stack**

- **Programming Language**: Go (Golang)
- **Database**: PostgreSQL
- **Message Queue**: Kafka
- **Authentication**: JWT
- **Containerization**: Docker
- **Orchestration**: Kubernetes (Minikube for local testing)
- **API Gateway**: Traefik or Kong (optional)

---

## **Database Design**

### **Tables**

1. **Users**
   - `id`, `email`, `password`, `role`

2. **Publications**
   - `id`, `publisher_id`, `created_at`, `updated_at`

3. **Posts**
   - `id`, `publication_id`, `title`, `content`, `is_premium`, `created_at`, `updated_at`, `is_published`

4. **Publisher Subscriptions**
   - `id`, `user_id`, `plan_id`

5. **Subscriber Subscriptions**
   - `user_id`, `is_premium`, `publication_id`

6. **Publisher Subscription Plans**
   - `id`, `name`, `order`

7. **Audience Limits**
   - `plan_id`, `size`

8. **Permissions**
    - `name`, `plan_id`

---

## **Getting Started**

### **Prerequisites**

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)
- [Minikube](https://minikube.sigs.k8s.io/docs/)
- [Kafka](https://kafka.apache.org/)

### **Setup Instructions**

1. Clone the repository:
   ```bash
   git clone https://github.com/nahK994/tiny-letter.git
   cd tiny-letter
   ```

2. Start DB using Docker Compose:
   ```bash
   docker compose -f db.yml up -d
   ```

3. Run services locally:
   ```bash
   // Will be updated soon
   ```

---

## **Contributing**

Contributions are welcome! Please open an issue or submit a pull request on the [GitHub repository](https://github.com/nahK994/tiny-letters).

---

## **Contact**

For any inquiries, please contact **Shomi Khan** at [shomikhan043@gmail.com](mailto:shomikhan043@gmail.com).
