# AnimeScrapper
This API scrapes torrent magents from [AnimeTosho](https://animetosho.org/). ADD MORE

## **Endpoints**
### GET `/search`

---
## **Items**
|Field     | Data Type    | Description                                        | Required | Max  | 
|----------|:------------:|----------------------------------------------------|:--------:|------|
| title    | `string`     | Title of anime to search                           | Yes      | 1    |
| filter   | `[]string`   | Returns titles that contains filter strings        | No       | 10   |
| minSeed  | `int`        | Returns titles that are greater then minSeed value | No       | none |
| maxEntry | `int`        | Returns a specfic amount of titles                 | No       | 100  |


---
## **API**
### Title
Gets a torrent object based on the title
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
