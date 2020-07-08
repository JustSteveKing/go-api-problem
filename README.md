# Go API Problem

A simple struct in GoLang for creating APIProblems in your API.


## Install

```bash
$ go get -u github.com/JustSteveKing/go-api-problem
```


## Usage

You use this very simply like:

```go
type server struct{}

func main() {
    s := &server{}
    http.Handle("/", func(rw http.ResponseWriter, r *http.Request) {
        if dbc := db.Where("email = ?", "user@email.com").First(&user); dbc.Error != nil {
            respondWithJSON(
                rw,
                http.StatusBadRequest,
                &APIProblem{
                    Title:  "Invalid Credentials",
                    Detail: "Unable to verify user credentials",
                    Status: strconv.Itoa(http.StatusBadRequest),
                    Code:   "IDEN-001-001",
                },
                "application/problem+json",
            )

            return
        }
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}

func respondWithJSON(rw http.ResponseWriter, code int, payload interface{}, contentType string) {
	response, _ := json.Marshal(payload)

	rw.Header().Set("Content-Type", contentType)
	rw.WriteHeader(code)
	rw.Write(response)
}
```
