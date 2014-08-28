# go-instagram

[![Build Status](https://travis-ci.org/nowk/go-instagram.svg?branch=master)](https://travis-ci.org/nowk/go-instagram)
[![GoDoc](https://godoc.org/gopkg.in/nowk/go-instagram.v0?status.svg)](http://godoc.org/gopkg.in/nowk/go-instagram.v0)

Instagram client in Go

## Install

    go get gopkg.in/nowk/go-instagram.v0

## Usage

    import "log"
    import "gopkg.in/nowk/go-instagram.v0"

    func main() {
      api := instagram.NewClient(<...access_token...>)

      data, err := api.Users.SelfFeed()
      if err != nil {
        log.Fatalf("fatal: %s", err)
      }

      if data.Meta.code != 200 {
        log.Fatalf("error: %s: %s", data.Meta.ErrorType, data.Meta.ErrorMessage)
      }

      for _, media := range data.Data {
        log.Print(media.User.FullName)
      }
    }

## API

Return a new client

    api := instagram.NewClient(<...access_token...>)

##### Endpoints

Endpoints are organized and mapped based on the documentation http://instagram.com/developer/endpoints/

    api.Users         // => Users Endpoints
    api.Relationships // => Relationships Endpoints
    api.Media         // => Media Endpoints
    api.Comments      // => Comments Endpoints
    api.Likes         // => Likes Endpoints
    api.Tags          // => Tags Endpoints
    api.Locations     // => Locations Endpoints
    api.Geographies   // => Geographies Endpoints

---

`go-instagram` comes with an included `jsons` package which contains all the structures for the returned JSON payloads. This package is available to allow you to create your own `go channels`.

---

##### Enforce Signed Header

    api.SetSignedHeader("secret", "192.168.0.1", "192.168.0.2")

## License

MIT
