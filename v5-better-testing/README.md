# v3. Prevent escape from the root filesystem

This fixes one to the biggest security flaws. The ability to use a file system link to access outside the directoy the webserver is started from.

## Cons
- Uses the DefaultServeMux, so things like timeouts can't be set, and this default server runs in a global scope, so any part of your code and touch it.
- The default html file served is always `index.html`
- no logging
- everything hard coded. 

`go run .` will run this, serving files from the current directory.

if you browse to localhost:8080, it will give you the default index.html
Same if you go to localhost:8080/static
If you browse to locahost:8080/noindex, a directory with no index.html, it should no longer display a directory listing
`dotFileAccess: true,`  can be set to `false` to prevent access to any file or directory that starts with a .dot

`escape` is a link to `../` created with `ln -s ../ escape`
This no longer works

If you go to localhost:8080/noindex  there is no index.html, but now you just get "Not Found" with a http 404 response code.

## Changes
Only a few additional lines added to main.go