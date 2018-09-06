# ae [![GoDoc](https://godoc.org/github.com/clavoie/ae?status.svg)](http://godoc.org/github.com/clavoie/ae) [![Build Status](https://travis-ci.org/clavoie/ae.svg?branch=master)](https://travis-ci.org/clavoie/ae) [![Go Report Card](https://goreportcard.com/badge/github.com/clavoie/ae)](https://goreportcard.com/report/github.com/clavoie/ae)

Wrappers around Google's AppEngine Go packages, meant for use with a dependency injection system.

```go
fnuc MyHandler(app ae.AppEngine, db ae.Db, log logu.Logger) {
  if app.IsDevAppServer() {
    return;
  }
  
  q := db.NewQuery("Address")
  q = q.Filter("Street =", someValue).KeysOnly()
  addresses := make([]Address, 0)
  keys, err := q.GetAll(addresses)
  
  if err != nil {
    logu.Errorf("Could not get addresses: %v", err)
    return
  }
} 
```
