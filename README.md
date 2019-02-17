# Instagram-parcer

### Instagram Parcer for Golang.

Parcing Account and Media information from public instagram accounts

#### Instalation
##### Install:
```
$ go get -u github.com/sparrowganz/instagram-parcer/instagram
```

##### Import:
```go
import "github.com/sparrowganz/instagram-parcer/instagram"
```

#### Usages

##### Get Intagram Account

```go
//By Username
account , err := GetAccountByUsername("username")
```

```go
//By Url
account , err := GetAccountByUrl("https://instagram.com/username/")
```

##### Get Medias
_only main page from account_
```go
//By Username
medias , err := GetLastMediasByUsername("username")
```

```go
//By Url
medias , err := GetLastMediasByUrl("https://instagram.com/username/")
```

##### Get One Media 

```go
//By Url
media , err := GetMediaByUrl("https://instagram.com/p/shortcode/")
```

```go
//By ShortCode
media , err := GetMediaByShortCode("shortcode")
```

#### Exceptions

if `err != nil` then it may be one of the mistakes:
* Media not found
* URL domain is not instagram.com
* Invalid JSON
* Information not found
* Don`t read HTML document
* Inability to read request body 
* HTTP status is not equal 200



