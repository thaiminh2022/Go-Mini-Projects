# URL shortener

A url shortener web app with frontend + backend

## Running the app

```bash
$ go run cmd/url_shortener.go
```

## How to use

1. Run the webapp, open your website and go to localhost:3000
2. Fill in the form, make sure it's a valid url (https://example.com)
3. Press submit, the server will redirect you to localhost:300/s

   - This is an api end point for shorten url
   - Make a post request with query path POST localhost:3000/s?url=https://example.com
   - This will response with a html body with the shorten url

4. Copy the url and use

### The /r/{id} endpoint

- This endpoint is for the backend to read the id and match with a local url in the database
- If your shorten url is localhost:3000/r/abcdxyznyt, your id will abcdxyznyt
- Manually enter this id will still make the server redirect if the database have a key with that id
