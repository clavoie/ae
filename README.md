# ae [![GoDoc](https://godoc.org/github.com/clavoie/ae?status.svg)](http://godoc.org/github.com/clavoie/ae) [![Build Status](https://travis-ci.org/clavoie/ae.svg?branch=master)](https://travis-ci.org/clavoie/ae) [![Go Report Card](https://goreportcard.com/badge/github.com/clavoie/ae)](https://goreportcard.com/report/github.com/clavoie/ae)

Wrappers around Google's AppEngine Go packages, meant for use with a dependency injection system and unit tests. Currently only the `appengine`, `datastore`, and `taskqueue` pakcages are wrapped.

## Usage

The wrapper interfaces aim to keep the names, parameter types, and parameter orders the same as the functions of the packages they wrap. Using the wrappers should consist of setting up your dependency injection system and then using the wrappers as you would the package functions directly.

```go
fnuc MyHandler(app ae.AppEngine, db ae.Db, log logu.Logger) {
  // wraps the appengine package
  if app.IsDevAppServer() {
    return;
  }
  
  // wraps the datastore package
  q := db.NewQuery("Address")
  q = q.Filter("Street =", someValue).KeysOnly()
  addresses := make([]Address, 0)
  keys, err := q.GetAll(addresses)
  
  if err != nil {
    log.Errorf("Could not get addresses: %v", err)
    return
  }
  
  // etc
} 
```
## Setup

Below is a more complete example of how main.go file would be set up for an App Engine service:

```go
package main

import (
  "myPackage"
	"net/http"

	"github.com/clavoie/di"
	"github.com/clavoie/erru"
	"github.com/clavoie/logu"
)

var httpResolver di.IHttpResolver

func onResolveErr(err *di.ErrResolve, w http.ResponseWriter, r *http.Request) {
	logger := logu.NewAppEngineLogger(r)
	logger.Errorf("err encountered while resolving dependencies: %v", err.String())

	httpErr, isHttpErr := err.Err.(erru.HttpErr)
	if isHttpErr {
		w.WriteHeader(httpErr.StatusCode())
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

var httpDefs = []*di.HttpDef{
  {myPackage.HomeHandler, "/"},
  {myPackage.AddressHandler, "/address"},
  // etc...
}

func init() {
	var err error

	httpResolver, err = di.NewResolver(onResolveErr, 
    ae.NewDiDefs(), 
    logu.NewAppEngineDiDefs(),
    // etc, your service deps here
    )

	if err != nil {
		panic(err)
	}

	err = httpResolver.SetDefaultServeMux(httpDefs)
}
```
