# Stringer

## Goals

Feed aggregator to work with Reeder app.

## Testing

Start app - runs on port 8888
Start mitmproxy --mode reverse:http://localhost:8888 -p 4444
Add account to Reeder (Reader, self-hosted Google Reader API) Server: http://localhost:8888, username: hello, password: world


## Current Problem

Item IDs

- https://www.inoreader.com/developers/stream-ids
- https://www.inoreader.com/developers/item-ids
  - user/-/state/com.google/read
  - user/-/state/com.google/starred
  - user/-/state/com.google/like

## Links

- https://www.inoreader.com/developers/
- https://github.com/FreshRSS/FreshRSS/blob/master/p/api/greader.php
- https://freshrss.github.io/FreshRSS/en/users/06_Mobile_access.html
- https://developers.google.com/identity/protocols/oauth2/native-app?csw=1
- https://code.google.com/archive/p/pyrfeed/wikis/GoogleReaderAPI.wiki
- https://feedhq.readthedocs.io/en/latest/api/reference.html
- https://github.com/devongovett/reader
- https://github.com/mrusme/journalist/blob/master/api/greaderAPI.go
