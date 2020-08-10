# OSRM Routes API
======

First step is working.  Able to retrieve data and return it back to the user

command used:
`curl "127.0.0.1:80/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219"`

JSON is unfortunately not sorted yet. I will add this soon as one wil receive this:
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

It is working, so just need to make it look present it in a specific order (sorted by driving time and distance) and not by alphabetical order