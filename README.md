# tokesura

A CLI and library to generate [Hasura JWTs][hasura-jwt].

[hasura-jwt]: https://hasura.io/docs/latest/auth/authentication/jwt

## Motivation

Working with JWTs in general -- let alone Hasura JWTs -- can be tricky.

- How do you pick a secure secret?
  - Should you really be going to [jwt.io][jwt-io]?
- How do you generate JWTs?
- How do you verify them?
- How do you generate a JWT for one specific user?
  - (For example, if you're testing permission rules).

This project offers a CLI and library that can answer all those questions.

[jwt-io]: https://jwt.io

## Usage

### Generate a new token

```shell
go run main.go generate \
  --secrets SECRET \
  --subject user \
  --allowedRoles user \
  --defaultRole user
```

### Verify a token

```shell
go run main.go verify \
  --secrets SECRET \
  --token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJodHRwczovL2hhc3VyYS5pby9qd3QvY2xhaW1zIjp7IngtaGFzdXJhLWFsbG93ZWQtcm9sZXMiOlsiZXh0LWxpY2Vuc2Utd3JpdGVyIl0sIngtaGFzdXJhLWRlZmF1bHQtcm9sZSI6ImV4dC1saWNlbnNlLXdyaXRlciJ9LCJpYXQiOjE2NzM2NDAwMDQsInN1YiI6ImJvb21pIn0.iiopKLqb61hta2WgU7DCsdoei4vs2rMpIIQLLsycTgk
```

### Generating a secret

We occasionally have to generate secrets. For example, Hasura requires a
[JWT secret](https://hasura.io/docs/latest/deployment/graphql-engine-flags/reference/#jwt-secret)
and an
[admin secret](https://hasura.io/docs/latest/deployment/graphql-engine-flags/reference/#admin-secret-key).

```shell
# Generate an admin secret
go run main.go secrets generate --length 20

# Generate a JWT secret
go run main.go secrets generate --length 32
```
