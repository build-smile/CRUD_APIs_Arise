## Getting Started

### Prerequisites

- Go 1.18+
- Docker & Docker Compose

### Setup

1. Clone the repository:
    ```
    git clone https://github.com/yourusername/CRUD_APIs_Arise.git
    cd CRUD_APIs_Arise
    ```

2. Build and run with Docker Compose:
    ```
    docker-compose up --build
    ```

3. The API will be available at `http://localhost:3000`

### API Endpoints

- `POST /members` - Create a new member
- `GET /members` - List all members
- `GET /members/{id}` - Get a member by ID
- `PUT /members/{id}` - Update a member
- `DELETE /members/{id}` - Delete a member

## Configuration

Edit `config/config.yaml` to change application settings.

## Export file postman
You can import the Postman collection from `Arise.postman_collection.json` to test the API endpoints.
