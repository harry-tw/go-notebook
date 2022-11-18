# demo-go-kit

## Technique

- Architecture by [go-kit]
  - Service: `ArticleService`
  - Endpoint 
  - Transport 
- Database
  - ORM: [go-gorm]
  - CRUD Code Generation: [Code gen by go-gorm]
  - Migration: [go-gormigrate]
- Deployment
  - docker-compose
    - App: [demo-go-kit]
    - Database: [PostgreSQL]
    - DBMS: [adminer]

## How To Use
1. Launch [PostgreSQL]
2. Execute migration
3. `docker-compose up`

## Database
### Schema Change
1. Modify `model/model.go`
2. Modify `cmd/orm/genmodelquery/main.go` if there is tables added.
3. Execute `cmd/orm/genmodelquery/main.go` to auto-generate CRUD code under `model/query/`.
4. Modify `cmd/orm/migration/main.go` to prepare database migration.
5. Execute `cmd/orm/migration/main.go` with database migration. 
6. Check migration have effective by DBMS such as [adminer].

[go-kit]: https://github.com/go-kit/kit
[go-gorm]: https://gorm.io/
[Code gen by go-gorm]: https://gorm.io/gen/
[go-gormigrate]: https://github.com/go-gormigrate/gormigrate
[PostgreSQL]: https://hub.docker.com/_/postgres
[adminer]: https://hub.docker.com/_/adminer
[demo-go-kit]: ./
[Service]: ./service
