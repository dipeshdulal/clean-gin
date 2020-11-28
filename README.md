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
- [x] Dockerize Application with Debugging Support Enabled. Debugger runs at `5002`. Vs code configuration is at `.vscode/launch.json` which will attach debugger to remote application. 
