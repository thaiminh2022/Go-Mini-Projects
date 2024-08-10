# Currency Exchange

A simple app for converting currency in the terminal

## Running the app

### .env file

Create a .env file in the root directory with your open exchange app key:
(Replace {your app key} with your app key)

```text
APP_ID={your app key}

```

**Offline mode will be support later if i have time**

### Start the program

```bash
$ go run cmd/currency_exchange.go
```

### Reliability

Support real time update if you change 60 to 0

```go
if since > 60 || len(db.ExchangeRates) == 0 {
		rates := RequestMoreData()
		db.ExchangeRates = rates
	}

```

The code above means that we will fetch new exchange rates if it has been 60 minutes since we last fetched the exchange rate. This should be enough for a hobby project
