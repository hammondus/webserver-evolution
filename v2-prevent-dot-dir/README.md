# v2. Prevent showing directory listings and .dot files

This fixes 2 potential cons from v1.
- Prevents a directory listing being returned if a directory is requested and it doesn't contain index.html
- Optionally prevents access to .dot files

## Cons
- Uses the DefaultServeMux, so things like timeouts can't be set, and this default server runs in a global scope, so any part of your code and touch it.
- The "root" directory isn't protected against files that link to outside the directory structure.
- The default html file served is always `index.html`
- no logging
- everything hard coded. 

`go run .` will run this, serving files from the current directory.

if you browse to localhost:8080, it will give you the default index.html
Same if you go to localhost:8080/static
If you browse to locahost:8080/noindex, a directory with no index.html, it should no longer display a directory listing
`dotFileAccess: true,`  can be set to `false` to prevent access to any file or directory that starts with a .dot

`escape` is a link to `../` created with `ln -s ../ escape`
This server will allow this link to be used to escape out of the directoy this server considers root

If you go to localhost:8080/noindex  there is no index.html, but now you just get "Not Found" with a http 404 response code.

## Changes
As v1 was basically 3 lines of code, the changes are pretty much everything.