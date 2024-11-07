# Fulove - Bath Tracking Application

A bath tracking application that helps users maintain regular bathing habits and earn points.

## Features

- User authentication (signup/signin)
- Bath time tracking
- Point system
- Preferred bath time settings
- User profile with custom icons

## Tech Stack

### Frontend
- Flutter
- Provider for state management
- HTTP package for API communication

### Backend
- Go with Gin framework
- PostgreSQL for user data
- GORM for database operations
- JWT for authentication

## Project Structure

```
fulove/
├── lib/                    # Flutter frontend code
│   ├── screens/           # UI screens
│   ├── services/          # API services
│   └── main.dart          # Entry point
├── services/              # Backend services
│   └── user/              # User service
│       ├── internal/      # Internal packages
│       │   ├── database/  # Database connection
│       │   ├── handlers/  # HTTP handlers
│       │   └── models/    # Data models
│       ├── main.go        # Service entry point
│       └── Dockerfile     # Service container
└── docker-compose.yml     # Container orchestration
```

## Getting Started

### Prerequisites
- Flutter SDK
- Go 1.21+
- Docker and Docker Compose
- PostgreSQL

### Setup and Running

1. Clone the repository:
```bash
git clone https://github.com/yourusername/fulove.git
cd fulove
```

2. Start the backend services:
```bash
docker-compose up --build
```

3. Run the Flutter application:
```bash
flutter pub get
flutter run
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)