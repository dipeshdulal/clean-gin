## Clean Gin 

Trying to implement clean architecture with gin framework. 

#### Environment Variables
|  Key | Value | Desc |
|------|-------|------|
| `ServerPort`  | `:5000` | Port at which app runs | 
| `Environment` | `development|production` | App running Environment |
| `LogOutput`   | `./server.log` | Output Directory to save logs |

#### Checklist
- [x] Implement Dependency Injection 
- [x] Routing
- [x] Environment Files
- [x] Logging (file saving on `production`) [zap](https://github.com/uber-go/zap)
- [ ] Middlewares 
- [ ] ORM