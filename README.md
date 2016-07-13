# Goa UUID Example

## Depends on
- [Flyway](https://flywaydb.org/)
- [Goa](https://github.com/goadesign/goa)
- [Gorma](https://github.com/goadesign/goa)
- [Postgres](https://www.postgresql.org/)
- [Glide](https://github.com/Masterminds/glide)

## Usage

I passed up all of the dependencies with the sample project to minimize setup
```bash
$ go install
$ goa-uuid-ex
```

## Endpoint
```js
GET    /users/:id
POST   /users
PUT    /users/:id
DELETE /users/:id
```
