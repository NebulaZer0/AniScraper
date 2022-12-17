# AniScraper
This API scrapes anime torrent magnets from [AnimeTosho](https://animetosho.org/).

---

## **Quick Start**
1. Clone the repo to a directory with:<br />
```sh 
git clone git@github.com:NebulaZer0/AniScraper.git
```
2. move to the `cmd` directory <br />
```sh 
cd cmd
```
3. Create Executable:<br />
```go 
go build -o aniscraper
```
4. Start Service:<br />
```sh 
./aniscraper
```
---
## **Docker Setup**
1. Clone the repo to a directory with:<br />
```sh 
git clone git@github.com:NebulaZer0/AniScraper.git
```
2. Build the docker image:<br />
```sh
docker compose build
```
3. Create container:<br />
```sh
docker compose up
```

*To remove type:* `docker compose down`

---

## **Enviroment Variables**
|Option      | Description                           | Default |
|------------|---------------------------------------|:-------:|
|SERVER_PORT | Sets the port to listen on            | 8080    |
|MAX_PAGE    | Sets how many pages to scrape through | 15      |
|ACAO| Set Access-Control-Allow-Origin header address | * |

---
# **API**

## GET `/search` 

---
### **Items**
|Field     | Data Type    | Description                                        | Required | Max  | 
|----------|:------------:|----------------------------------------------------|:--------:|------|
| title    | `string`     | Title of anime to search                           | Yes      | 1    |
| filter   | `[]string`   | Returns titles that contains filter strings        | No       | 10   |
| minSeed  | `int`        | Returns titles that are greater then minSeed value | No       | None |
| maxEntry | `int`        | Returns a specfic amount of titles                 | No       | 100  |

## Examples

### Title Field
Gets a torrent object based on the title
```sh
curl -i -X GET localhost:8080/search?title="animeName" \
-H "Content-Type: application/json"
```
### Filter Field
Shows releated torrent names based on filter strings.

```sh
curl -i -X GET localhost:8080/search \
-H "Content-Type: application/json"
```

### Minimum Seed Field
Get seeds that are greater then the minimum seed value.

```sh
curl -i -X GET localhost:8080/search \
-H "Content-Type: application/json"
```

### Max Entry Field
Gets a specific amount of entrys based on max entry value.

```sh
curl -i -X GET localhost:8080/search \
-H "Content-Type: application/json"
```
---
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
}
```

## Libraries Used
### [Go Figure](https://github.com/common-nighthawk/go-figure)<br />
### [Gorilla/Mux](https://github.com/gorilla/mux)<br />
### [GoDotEnv](https://github.com/joho/godotenv)<br />
### [Colly](https://github.com/gocolly/colly)

