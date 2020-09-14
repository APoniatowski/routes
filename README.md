# OSRM Routes API
======

An API to get distances and driving times between your source and multiple destinations.

command used:
`curl "127.0.0.1:80/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219"`

```json
{
   "routes" : [
      {
         "destination" : "13.397634,52.529407",
         "distance" : 1884.9,
         "duration" : 251.6
      },
      {
         "destination" : "13.428555,52.523219",
         "distance" : 3841.7,
         "duration" : 394.1
      }
   ],
   "source" : "13.388860,52.517037"
}
```

Response times will vary from 0.50ms (baremetal kubernetes) to 200ms (google app engine)
