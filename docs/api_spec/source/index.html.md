# Introduction

This document is describe [project](https://github.com/koungkub/wongnai-challenge) for [WeChallengeProgram](https://careers.wongnai.com/development/wechallenge1).

# Docs

## Home Page

> **Success** Response header

```http
HTTP/1.1 200 OK
Content-Type: text/plain
```

Request

| Method | Path | Body |
| ------ | ---- | ---- |
| GET    | /    | false |

Home page for testing API that it worked or health check.

## Get Review

> **Success** Response header

```http
HTTP/1.1 200 OK
Content-Type: text/html
```

> **Failure** Response header

```http
HTTP/1.1 422 Unprocessable Entity
Content-Type: text/html
```

Request

| Method | Path | Body |
| ------ | ---- | ---- |
| GET    | /reviews/{id} | false |

Parameters

| Parameter | Type | Default | Description |
| --------- | ---- | ------- | ----------- |
| id        | integer | false | review_id for get correct review |

get review by specific review_id and correct review must be shown.

## Search Review

> **Success** Response header

```http
HTTP/1.1 200 OK
Content-Type: text/html
```

> **Failure** Response header

```http
HTTP/1.1 422 Unprocessable Entity
Content-Type: text/html
```

Request

| Method | Path | Body |
| ------ | ---- | ---- |
| GET    | /reviews?query={text} | false |

Parameters

| Parameter | Type | Default | Description |
| --------- | ---- | ------- | ----------- |
| text        | string | false | keyword for search review |

get review by specific keyword and correct review must be shown.

## Edit Review

> **Request body**

```json
{
	"data": {
		"comment": "very aroi"
	}
}
```

> **Success** Response header

```http
HTTP/1.1 200 OK
Content-Type: text/html
```

> **Failure** Response header

```http
HTTP/1.1 422 Unprocessable Entity
Content-Type: text/html
```

Request

| Method | Path | Body |
| ------ | ---- | ---- |
| PUT    | /reviews/{id} | true |

Parameters

| Parameter | Type | Default | Description |
| --------- | ---- | ------- | ----------- |
| id        | id | false | review_id for edit review |

edit review by specific review_id. If the request failure, it will be caused by 2 things.

1. review not changed.
2. review_id not found.
