## Clean Gin

Trying to implement clean architecture with gin framework.

#### Environment Variables

| Key            | Value                    | Desc                             |
| -------------- | ------------------------ | -------------------------------- |
| `SERVER_PORT`  | `5000`                   | Port at which app runs           |
| `ENV`          | `development,production` | App running Environment          |
| `LOG_OUTPUT`   | `./server.log`           | Output Directory to save logs    |
| `DB_USER`      | `username`               | Database Username                |
| `DB_PASS`      | `password`               | Database Password                |
| `DB_HOST`      | `0.0.0.0`                | Database Host                    |
| `DB_PORT`      | `3306`                   | Database Port                    |
| `DB_NAME`      | `test`                   | Database Name                    |
| `JWT_SECRET`   | `secret`                 | JWT Token Secret key             |
| `ADMINER_PORT` | `5001`                   | Adminer DB Port                  |
| `DEBUG_PORT`   | `5002`                   | Port that delve debugger runs in |

#### Migration Commands

| Command             | Desc                                           |
| ------------------- | ---------------------------------------------- |
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
