## Microservice Essentials

The repository contains a **TODO microservice** written in Go, using modern approaches, domain driven designs(DDD), idiomatic coding styles and integration with Hashicorp Vault, ElasticSeach, Kafka and many more modern tools.

### Note

The application uses some features in supported in `Go 1.16+`

- usage of embed to bundle static files
- `os.WriteFile` to write files

## Tools and packages required

```bash
# runs and manages sql migration both from CLI and programatically.
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.14.1

# converts sql queries into type-safe golang code.
go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.6.0

# creates mock of any interface, basic usgage for unit testing.
go install github.com/maxbrunsfeld/counterfeiter/v6

# generate open api 3 client and types from spec.
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen
```

## Open API and Swagger

OpenAPI is a specification to document REST API's in standard format, while Swagger offer tools to implement these specifications. Read more anout OpenAPI and swagger [here](https://swagger.io/blog/api-strategy/difference-between-swagger-and-openapi/).

<br>

**Some Blog post covering implementation of swagger in go**

- [For OpenAPI 2 diff b/w go-swagger and swag](https://medium.com/@pedram.esmaeeli/generate-swagger-specification-from-go-source-code-648615f7b9d9) - covers pros and cons of both packages.
- [serving swagger ui in go](https://ribice.medium.com/serve-swaggerui-within-your-golang-application-5486748a5ed4) - how to setup swagger ui for different frameworks in go.
- [detailed impl with go-swagger](https://www.ribice.ba/swagger-golang/) - have a detailed impl with go-swagger
- [openapi](https://mariocarrion.com/2021/05/02/golang-microservices-rest-api-openapi3-swagger-ui.html) - covers about open api.

**Some tools and packages availiable**

- [oapi-codegen](https://github.com/deepmap/oapi-codegen) - generate client and types for openAPI 3 spec
- [swag](https://github.com/swaggo/swag#getting-started) - easy tool to get openAPI 2 specs
- [go-swagger](https://github.com/go-swagger/go-swagger) - for openAPI 2
- [kin-openapi](https://github.com/getkin/kin-openapi) - create and generate source files
- [swagger-ui](https://github.com/swagger-api/swagger-ui) - download ui files from here

To embed the static-ui, `embed` is required. Detail explanantion covered [here](https://harsimranmaan.medium.com/embedding-static-files-in-a-go-binary-using-go-embed-bac505f3cb9a)

<br>

For better yaml parsing, used this [pkg](https://github.com/ghodss/yaml) instead of native marshaler.

## Author
**Akshit Sadana <akshitsadana@gmail.com>**

- Github: [@Akshit8](https://github.com/Akshit8)
- LinkedIn: [@akshitsadana](https://www.linkedin.com/in/akshit-sadana-b051ab121/)
