# Calculator api

A simple headless api for simple arithmetic (+, -, \*, /)

## Running the api

```bash
$ go run cmd/calculator_api.go
```

## Body

a json object with a, b number property

```json
{
  "a": 69,
  "b": 420
}
```

## Endpoints

Return a json object with a result property

```json
{
  "result": 489
}
```

- "/add": returns a + b
- "/minus": returns a - b
- "/multiply": returns a \* b
- "/divide": returns a / b

## Errors

Returns an error json object with the reason why the error happened

```json
{
  "message": "division by 0 (b must be non zero)"
}
```
