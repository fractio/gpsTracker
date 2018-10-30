# GPS tracker app

The purpose of this application is to test some tech

- Golang
- Boltload (bbolt)
- API sever in Golang
- websocket in golang (via channels)
- Flow types
- Auto generate front end types from golang structs

## Tools

### Make types for front end from Location struct

typewriter -file ./Location.go -lang flow -v -out ./client/src/models.js

### run frontend in dev mode

```bash
cd client
yarn start
```

### run backend in dev mode

```bash
go run ./...
```
