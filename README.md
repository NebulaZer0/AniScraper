# animeScrapper

TODO: description

## Endpoints
### GET `/`

Gets a torrent object based on the title

#### Request
```sh
curl -i -X GET localhost:8080/ \
-H "Content-Type: application/json" \
-d '{ "title": "Belle Movie" }'
```

#### Response
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sat, 22 Oct 2022 20:33:23 GMT
Content-Length: 694

{
        "Results": [
                {
                        "name": "[EMBER] Belle ...",
                        "size": "4.036 GB",
                        "seed": "53",
                        "magnet": "magnet:?..."
                }
        ]
}%
```
