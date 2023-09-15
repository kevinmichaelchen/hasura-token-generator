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

### Installing

#### With Tea

Installing with [Tea][tea] is the easiest approach.

```
tea tokesura --help
```

[tea]: https://tea.xyz/

#### With Go

```shell
go install github.com/kevinmichaelchen/tokesura@latest
```

#### With Docker

```shell
docker pull ghcr.io/kevinmichaelchen/tokesura
docker run --rm ghcr.io/kevinmichaelchen/tokesura --help
```

### Commands

#### Generate a new token

```shell
tokesura generate \
  --allowedRoles teacher,coteacher \
  --defaultRole teacher \
  --subject user-123 \
  --userID 123 \
  --secret foobar
```

#### Verify a token

```shell
tokesura verify \
  --secret foobar \
  --token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJodHRwczovL2hhc3VyYS5pby9qd3QvY2xhaW1zIjp7IngtaGFzdXJhLWFsbG93ZWQtcm9sZXMiOlsidGVhY2hlciIsImNvdGVhY2hlciJdLCJ4LWhhc3VyYS1kZWZhdWx0LXJvbGUiOiJ0ZWFjaGVyIiwieC1oYXN1cmEtdXNlci1pZCI6IjEyMyJ9LCJpYXQiOjE2OTQ4MDM3NzgsInN1YiI6InVzZXItMTIzIn0.H1S_Uqk7u0KLBRIIP2hvCP6oV4udpfQFj2803t5NFAI
```

#### Generating a secret

We occasionally have to generate secrets. For example, Hasura requires a [JWT
secret][jwt-secret] and an [admin secret][admin-secret].

[jwt-secret]:
  https://hasura.io/docs/latest/deployment/graphql-engine-flags/reference/#jwt-secret
[admin-secret]:
  https://hasura.io/docs/latest/deployment/graphql-engine-flags/reference/#admin-secret-key

```shell
# Generate an admin secret
tokesura secret generate --length 20

# Generate a JWT secret
tokesura secret generate --length 32
```
