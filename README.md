# AnimeScrapper
This API scrapes torrent magents from [AnimeTosho](https://animetosho.org/). ADD MORE!


## **Quick Start**
1. Clone the repo to a directory with:<br />
```sh 
git clone git@github.com:NebulaZer0/animeScrapper.git
```
2. move to the `cmd` directory <br />
```sh 
cd cmd
```
3. Create Executable:<br />
```go 
go build -o aniScrapper
```
4. Start Service:<br />
```sh 
./aniScrapper
```

## **Docker Setup**
1. Clone the repo to a directory with:<br />
```sh 
git clone git@github.com:NebulaZer0/animeScrapper.git
```
2. Build the docker image:<br />
```sh
docker compose build
```
3. Create container:<br />
```sh
docker compose up
```

*To remove type* `docker compose down`

## **Endpoints**
### GET `/search`
---
## **Enviroment Variables**
|Option      | Description                |
|------------|----------------------------|
|SERVER_PORT | Sets the port to listen on |

---
## **Items**
|Field     | Data Type    | Description                                        | Required | Max  | 
|----------|:------------:|----------------------------------------------------|:--------:|------|
| title    | `string`     | Title of anime to search                           | Yes      | 1    |
| filter   | `[]string`   | Returns titles that contains filter strings        | No       | 10   |
| minSeed  | `int`        | Returns titles that are greater then minSeed value | No       | None |
| maxEntry | `int`        | Returns a specfic amount of titles                 | No       | 100  |


---
## **API**
### Title Field
Gets a torrent object based on the title
```sh
curl -i -X GET localhost:8080/search \
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
### Filter Field
Shows releated torrent names based on filter strings.

```sh
curl -i -X GET localhost:8080/search \
-H "Content-Type: application/json" \
-d '{ "title": "Belle Movie", "filter": ["Episode 1", "Season 1"]}'
```
