# A CGI in Go

This small project demonstrates building a fully self-contained web app (including html and Javascript) as a single Go binary.

## Why?

A long, long time ago, [CGIs](https://en.wikipedia.org/wiki/Common_Gateway_Interface) were the defacto way of building web apps. You basically wrote a C program that would generate dynamic pages (in the most  "advanced" cases, even serve static assets such as external html files!). Deploying webapps has grown exponentially more complex since (although argably easier these days - unless you're stuck in deploying Java WARs and EARs inside heavyweight containers), this experiment proposes deploying a webapp by simply building a binary you can run and point web traffic to. No frills, no folders full of resources, nothing - a single executable "website" that actually bundles all the web resources as in-memory byte arrays.

## Dependencies

Install everything by running:

```
go get -u github.com/jteeuwen/go-bindata/...
```

## Running this

Build the assets and the "cgi" by running:

```
./build.sh
```

Run it simply with:

```
./cgi-go
```

You can also serve assets from disk (useful when editing html/js/etc) by running `./cgi-go dev`

## Wait that's it?

Yes - that's it. Deploying your "website", now, is just a matter of [cross-compiling](http://golangcookbook.com/chapters/running/cross-compiling/) your Go binary to your platform of choice, then running it in your server. No dependencies, no resources to copy over, no configurations, nada (you may want to put this behind a web proxy such as Nginx if you want to do stuff such as handling SSL certificates, serving multiple sites from the same port, etc).

You can obviously expand the code to do anything else - serve a full client-side JS app, add API endpoints, use a more fully-featured web framework, etc.

## Improvement points (aka TODO)

- Chunked transfers (for bigger assets)
- Parameters support?
- Caching headers
- Performance testing: is this *much slower* than serving files from disk?
- Godep?