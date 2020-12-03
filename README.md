## Clean Gin

Trying to implement clean architecture with gin framework.

#### Environment Variables

| Key           | Value                    | Desc                          |
| ------------- | ------------------------ | ----------------------------- |
| `ServerPort`  | `:5000`                  | Port at which app runs        |
| `Environment` | `development,production` | App running Environment       |
| `LogOutput`   | `./server.log`           | Output Directory to save logs |
| `DBUsername`  | `username`               | Database Username             |
| `DBPassword`  | `password`               | Database Password             |
| `DBHost`      | `0.0.0.0`                | Database Host                 |
| `DBPort`      | `3306`                   | Database Port                 |
| `DBName`      | `test`                   | Database Name                 |
| `JWTSecret`   | `secret`                 | JWT Token Secret key          |

#### Migration Commands

| Command            | Desc                                           |
| -------------- | ---------------------------------------------- |
| `make migrate-up`   | runs migration up command                      |
| `make migrate-down` | runs migration down command                    |
| `make force`        | Set particular version but don't run migration |
| `make goto`         | Migrate to particular version                  |
| `make drop`         | Drop everything inside database                |
| `make create`       | Create new migration file(up & down)           |

#### Checklist

- [x] Implement Dependency Injection (go-fx)
- [x] Routing (gin web framework)
- [x] Environment Files
- [x] Logging (file saving on `production`) [zap](https://github.com/uber-go/zap)
- [x] Middlewares (cors)
- [x] Database Setup (mysql)
- [x] Models Setup and Automigrate (gorm)
- [x] Repositories
- [x] Implementing Basic CRUD Operation
- [x] Authentication (JWT)
- [x] Migration
- [x] Dockerize Application with Debugging Support Enabled. Debugger runs at `5002`. Vs code configuration is at `.vscode/launch.json` which will attach debugger to remote application.
