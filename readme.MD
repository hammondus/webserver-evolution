# Evolution


# Stuff to add in later versions.

## Headers
`w.Header().Set("X-Content-Type-Options", "nosniff")`
This prevents browsers from mime type sniffing.
Browsers sometimes ingore the content-type header and try to guyess the file type by examining the content. For a web server that can host user provided content, this can be exploited by hackers.

eg, Someone could upload a file claiming to be a jpeg, but it actually contains javascript.
When someone views that file, it might execute it.