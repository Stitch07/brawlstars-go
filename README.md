# brawlstars-go

A Go client for the unofficial [Brawl Stars API](https://docs.brawlapi.cf).

## Features:

* 100% typesafe API coverage.
* Automatic ratelimit and compression handling.
* An easy to use API with helper methods.

## Usage

```go
client := brawlstars.New(token)
player, err := client.GetPlayer("#Y2QPGG")
```

For documentation, visit the [godoc page.](https://godoc.org/github.com/Soumil07/brawlstars-go)