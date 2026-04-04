# 📝 Task Manager CRUD API

A simple full-stack task management application built with **Go**, **PostgreSQL**, and **Docker**. Users can securely manage personal to-do lists with authentication and a clean web interface.

---

## 🚀 Features

* 🔐 User authentication (login/logout)
* ✅ Create, read, update, delete (CRUD) tasks
* 📌 Mark tasks as complete/incomplete
* ⚡ Fast backend using Go
* 🐘 PostgreSQL database with indexing
* 🐳 Dockerized setup for easy deployment
* 🌐 Simple frontend with HTML + Tailwind CSS
* 🛠 Infrastructure automation using Terraform & Ansible (optional)

---

## 🏗️ Project Structure

```
task-manager/
│
├── cmd/server/          # Application entrypoint
├── internal/            # Core application logic
│   ├── handlers/        # HTTP handlers
│   ├── models/          # Data models
│   ├── services/        # Business logic
│   ├── repository/      # Database queries
│   ├── middleware/      # Auth & session middleware
│   └── config/          # Configuration
│
├── web/                 # Frontend
│   ├── templates/       # HTML templates
│   └── static/          # CSS/JS assets
│
├── migrations/          # SQL schema
├── docker/              # Docker configs
├── terraform/           # Infrastructure as code
├── ansible/             # Server configuration & deployment
├── .devcontainer/       # VS Code dev environment
└── README.md
```

---

## ⚙️ Technologies Used

* **Backend:** Go 1.25
* **Database:** PostgreSQL 17
* **Frontend:** HTML + Tailwind CSS
* **Authentication:** bcrypt password hashing
* **Containerization:** Docker & Docker Compose
* **Infrastructure (Optional):** Terraform + Ansible

---

## 🔌 API Endpoints

| Method | Endpoint       | Description       |
| ------ | -------------- | ----------------- |
| GET    | `/`            | Login page        |
| POST   | `/login`       | Authenticate user |
| GET    | `/logout`      | Logout user       |
| GET    | `/dashboard`   | View tasks        |
| POST   | `/task/create` | Create task       |
| POST   | `/task/update` | Update task       |
| GET    | `/task/toggle` | Toggle completion |
| GET    | `/task/delete` | Delete task       |

---

## 🗄️ Database Schema

### Users Table

* `id`
* `username`
* `password_hash`
* `created_at`

### Tasks Table

* `id`
* `user_id`
* `title`
* `description`
* `is_completed`
* `created_at`
* `updated_at`

> Indexed on `user_id` for performance.

---

## 🐳 Getting Started (Docker)

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/task-manager.git
cd task-manager
```

### 2. Set Environment Variables

Edit `docker/docker-compose.yml` or create a `.env` file:

```
DEFAULT_USER=admin
DEFAULT_PASS=admin123
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=taskdb
```

### 3. Run the Application

```bash
cd docker && docker-compose up --build -d
```

### 4. Access the App

Open your browser:

```
http://localhost:8080
```

Login with:

```
Username: admin
Password: admin123
```

---

## 💻 Development (Dev Container)

This project includes a `.devcontainer` setup for seamless development in VS Code.

### Steps:

1. Install **Dev Containers** extension in VS Code
2. Press:

```
Ctrl + Shift + P → Reopen in Container
```

### Benefits:

* Preconfigured Go environment
* PostgreSQL runs alongside app
* No local setup required
* Consistent dev environment

---

## 🚀 Deployment

### 1. Provision Infrastructure (Terraform)

```bash
cd terraform
terraform init
terraform apply
```

---

### 2. Configure Server (Ansible)

```bash
cd ansible
ansible-playbook -i inventory playbook.yml
```

---

### 🔄 Deployment Flow

```
Terraform → Provision server
        ↓
Ansible → Install dependencies
        ↓
Ansible → Deploy Docker app
        ↓
Application runs
```

---

## 🔧 Cross-Platform Support

Supports multiple Linux distributions:

* Debian / Ubuntu (APT)
* RedHat / CentOS (YUM)

---

## 🌟 Future Improvements

* JWT-based authentication
* RESTful API redesign
* Pagination & filtering
* Task deadlines & reminders
* Real-time updates (WebSockets)
* CI/CD pipeline (GitHub Actions)

---

## 🤝 Contributing

Contributions are welcome!

1. Fork the repository
2. Create a new branch
3. Submit a pull request

---

## 📄 License

This project is licensed under the MIT License.

---

## 👨‍💻 Author

Built with Go for learning and practical DevOps integration.

---
