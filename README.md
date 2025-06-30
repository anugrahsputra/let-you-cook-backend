# Let You Cook

Let You Cook is a RESTful API service for a simple to-do list application. It allows users to manage their tasks, categorize them, and track their work sessions using the Pomodoro technique.

## Features

- User authentication (Register and Login)
- Manage tasks (Create, Read, Update, Delete)
- Manage categories for tasks
- Track Pomodoro sessions
- User profile management

## Getting Started

### Prerequisites

- [Go](https://golang.org/) (version 1.23.4 or later)
- [MongoDB](https://www.mongodb.com/)
- [MinIO](https://min.io/)

### Installation

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/let-you-cook.git
    cd let-you-cook
    ```

2.  **Install dependencies:**

    ```bash
    go mod tidy
    ```

3.  **Set up environment variables:**

    Create a `.env` file in the root directory by copying the example file:

    ```bash
    cp .env.example .env
    ```

    Update the `.env` file with your database credentials and other configurations:

    ```
    # .env
    MONGO_URI=mongodb://localhost:27017
    DB_NAME=let-you-cook
    MINIO_ENDPOINT=localhost:9000
    MINIO_ACCESS_KEY_ID=minioadmin
    MINIO_SECRET_ACCESS_KEY=minioadmin
    MINIO_BUCKET_NAME=let-you-cook
    JWT_SECRET=your-secret-key
    ```

### Running the Application

To run the application, execute the following command:

```bash
go run main.go
```

The application will start on port `42069`.

## API Endpoints

All endpoints are prefixed with `/api/v1`.

### Authentication

- `POST /auth/register`: Register a new user.
- `POST /auth/login`: Login a user.

### Users

- `GET /users`: Get all users.
- `GET /users/:id`: Get a user by their ID.

### Profile

- `POST /profile`: Create a user profile.
- `GET /profile`: Get the profile of the logged-in user.
- `PATCH /profile`: Update the profile of the logged-in user.

### Tasks

- `POST /tasks`: Create a new task.
- `GET /tasks`: Get all tasks for the logged-in user.
- `GET /tasks/category`: Get all tasks grouped by category.
- `PATCH /tasks/:id`: Update a task.
- `DELETE /tasks/:id`: Delete a task.

### Categories

- `POST /category`: Create a new category.
- `GET /category`: Get all categories for the logged-in user.
- `GET /category/:id`: Get a category by its ID.
- `PATCH /category/:id`: Update a category.
- `DELETE /category/:id`: Delete a category.

### Pomodoro Sessions

- `POST /session/create`: Create a new Pomodoro session.
- `PATCH /session/start/:id`: Start a Pomodoro session.
- `PATCH /session/end/:id`: End a Pomodoro session.
- `GET /session`: Get all Pomodoro sessions for the logged-in user.
- `PATCH /session/:id`: Update a Pomodoro session.

## Contribution

Contributions are welcome! Please feel free to submit a pull request.

By contributing to this project, you agree that your contributions will be licensed under the MIT License.
