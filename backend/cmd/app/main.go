package main

import (
    "log"
    "net/http"
    "backend/internal/data"
    "backend/internal/graphql"
)

func main() {
    data.LoadMockData()

    http.Handle("/graphql", graphql.NewGraphQLHandler())

    log.Println("Server running at http://localhost:8080/graphql")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
