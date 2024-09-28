# Golang Prisma Postgresql Railway

## Models

- post

`go run main.go`

## Generate the prisma lib

We do this step every time we edit the schema

```bash
go run github.com/steebchen/prisma-client-go generate dev
```

## Migrate and deploy prima in the database

This command will create two file, one for preview other will implement that migration

```bash
go run github.com/steebchen/prisma-client-go migrate dev --preview-feature --create-only
```

This command will direct implement the migration without preview the sql

```bash
go run github.com/steebchen/prisma-client-go migrate dev
```

## Project tree example

`tree -I '.github|.trunk|prisma'`

```tree
.
├── README.md
├── TODO.md
├── api
│   ├── v1
│   │   ├── account_routes.go
│   │   ├── address_routes.go
│   │   ├── category_routes.go
│   │   ├── listing_routes.go
│   │   ├── media_routes.go
│   │   ├── phone_routes.go
│   │   ├── post_routes.go
│   │   ├── publication_media_routes.go
│   │   ├── publication_routes.go
│   │   ├── role_routes.go
│   │   ├── user_media_routes.go
│   │   └── user_routes.go
│   └── v2
├── common
│   └── web_reponse.go
├── config
│   └── database.go
├── domain
│   ├── account
│   │   ├── controller
│   │   │   └── account_controller.go
│   │   ├── dto
│   │   │   ├── account_create_dto.go
│   │   │   ├── account_response_dto.go
│   │   │   └── account_update_dto.go
│   │   ├── model
│   │   │   └── account.go
│   │   ├── repository
│   │   │   ├── account_repository.go
│   │   │   └── account_repository_impl.go
│   │   └── service
│   │       ├── account_service.go
│   │       └── account_service_impl.go
│   ├── post
│   │   ├── controller
│   │   │   └── post_controller.go
│   │   ├── dto
│   │   │   ├── post_create.go
│   │   │   ├── post_respose.go
│   │   │   └── post_update_request.go
│   │   ├── model
│   │   │   └── post.go
│   │   ├── repository
│   │   │   ├── post_repository.go
│   │   │   └── post_repository_impl.go
│   │   └── service
│   │       ├── post_service.go
│   │       └── post_service_impl.go
│   ├── user
│   │   ├── controller
│   │   │   └── user_controller.go
│   │   ├── dto
│   │   │   ├── user_create.go
│   │   │   ├── user_response.go
│   │   │   └── user_update.go
│   │   ├── model
│   │   │   └── user.go
│   │   ├── repository
│   │   │   ├── user_repository.go
│   │   │   └── user_repository_impl.go
│   │   └── service
│   │       ├── user_service.go
│   │       └── user_service_impl.go
│   └── user_media
│       ├── controller
│       │   └── user_media_controller.go
│       ├── dto
│       │   ├── user_media_create.go
│       │   ├── user_media_response.go
│       │   └── user_media_update.go
│       ├── model
│       │   └── user_media.go
│       ├── repository
│       │   ├── user_media_repository.go
│       │   └── user_media_repository_impl.go
│       └── service
│           ├── user_media_service.go
│           └── user_media_service_impl.go
├── go-arepas
├── go.mod
├── go.sum
├── helper
│   ├── error.go
│   └── json.go
├── main.go
├── router
│   └── router.go
├── thunder-collection_go-prisma.json
└── trunk
```


## CORS

In your .env file or environment settings, define the allowed origins for different environments.

In dev

```bash
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:4322,http://localhost:8080
CORS_ALLOWED_METHODS=GET,POST,PUT,DELETE,OPTIONS
CORS_ALLOWED_HEADERS=Content-Type,Authorization
```

In production, you would set the CORS_ALLOWED_ORIGINS to your actual domain(s):

```bash
CORS_ALLOWED_ORIGINS=https://your-astro-client.com,https://your-android-client.com,https://your-ios-client.com
```
