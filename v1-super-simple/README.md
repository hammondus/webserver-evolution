# v1. Super Simple Web Server

This is a webserver that is about as simple as it gets.

[http @ go.dev](https://pkg.go.dev/net/http) still gives the example of using http.Dir, but there a multiple ways of getting similar results.

The below work in a similar way, but there are some advantages to useing os.DirFS over http.Dir

# Pros of this simple web server
- Super simple. Can have a webserver running in only a few lines of code.
- sets various http headers automatically for you such as content-type
- The "root" directory where the file server is started, is protected from directory traversal by user putting in ../../../otherdir etc in the URL

# Cons
- Many....
- Uses the DefaultServeMux, so things like timeouts can't be set, and this default server runs in a global scope, so any part of your code and touch it.
- The "root" directory isn't protected against files that link to outside the directory structure.
- The default html file served is always `index.html`
- if there is no `index.html` in the directory being served, it shows a complete directory listing
- dot files, which you often want hidden or restricted from can be accessed.
- no logging
- everything hard coded. 

`go run .` will run this, serving files from the current directory.
if you browse to localhost:8080, it will give you the default index.html
Same if you go to localhost:8080/static
If you go to localhost:8080/more  there is no index.html, so you will get a directory listing

`escape` is a link to `../` created with `ln -s ../ escape`
This server will allow this link to be used to escape out of the directoy this server considers root

There are a number of ways of getting a similr result.
```go
http.FileServer(http.Dir("."))          // old way since go 1.0 and used in my go code examples
http.FileServer(http.FS(os.DirFS("."))) // introduced in go 1.16
http.FileServerFS(os.DirFS("."))        // introduced in go 1.22

// the go 1.16 and 1.22 versions are functionaly the same. v1.22 is just some syntatic sugar to make the code nicer.
```

There are advantages to using os.DirFS, so that's what this webserver will use.
http.Dir is limited to serving files from the native OS filesystem.

os.DirFS uses the newer fs.FS interface, so it can be used for:
- memfs - in memory file system
- embed.FS - embedded files
- httpfs - remote http server
- cloud storage (s3, azure blob etc)
- gitfs - git repo
- vaultfs - hashicorp vault
