
![Logo](https://raw.githubusercontent.com/angelsebamc/GoShort/main/static/images/goshort_logo.png)


# GoShort

Working on a basic URL Shortener API made with GO

## External Libraries

- [gin](https://pkg.go.dev/github.com/gin-gonic/gin)
- [mongo-driver](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo)
- [jwt](https://pkg.go.dev/github.com/golang-jwt/jwt)
- [crypto](https://pkg.go.dev/golang.org/x/crypto)
- [validator](https://pkg.go.dev/github.com/go-playground/validator)

## API Reference

### User

#### User Sign in

```http
  POST /api/user/signin
```

| Parameter | Type     | Description | Required  |
| :-------- | :------- | :---------- | :-------- |
| `username` | `string` | User's username | Yes  |
| `email` | `string` | User's email       | Yes  |
| `password` | `string` | User's password | Yes  |

#### User Authorization

```http
  POST /api/user/auth
```

| Parameter | Type     | Description     | Required |
| :-------- | :------- | :---------------| :------- |
| `email`   | `string` | User's email    | Yes      |
| `password`| `string` | User's password | Yes      |


### Link

#### Create a Link

```http
  POST /api/link/create
```

> 🔒 Requires a JWT token in the Authorization headers.


| Parameter      | Type     | Description              | Required |
| :------------- | :------- | :----------------------- | :--------|
| `original_url` | `string` | The URL to be shortened  | Yes      |


#### Delete a Link

```http
  POST /api/link/delete
```

> 🔒 Requires a JWT token in the Authorization headers.


| Parameter | Type     | Description              | Required |
| :-------- | :------- | :----------------------- | :------- |
| `id`      | `string` | The ID of the link       | Yes      |

#### Get user links

```http
  GET /api/link/user_links
```

> 🔒 Requires a JWT token in the Authorization headers.