# Quality Trace
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![React](https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB)
![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?&style=for-the-badge&logo=redis&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/jdasilvalima/qualityTrace?style=for-the-badge)
[![GitHub last commit](https://img.shields.io/github/last-commit/jdasilvalima/qualityTrace?style=for-the-badge)](https://github.com/jdasilvalima/qualityTrace/commits)

## I. PROJECT DESCRIPTION
### I.1 Introduction
QualityTrace is a real-time inventory and quality traceability system designed for the food industry. Built with Go and GraphQL, it empowers companies to track product batches from receiving to distribution, ensure compliance with safety standards, and manage stock levels efficiently. With features like expiration alerts, quality control logs, and supplier traceability, QualityTrace helps streamline operations and reduce waste.

### I.2 Goals
- Deepen my knowledge of **Go** and its concurrency model to build efficient, scalable systems.
- Gain hands-on experience with **GraphQL**, mastering query design, schema development, and optimization for complex data relationships.
- Explore integration of in-memory databases like **Redis** for caching, session management, and real-time notifications.

### I.3 Web Application Overview
WIP

## II. PROJECT SETUP
### II.1 Requirements
- [Docker](https://www.docker.com/) neeeds to be installed

### II.2 BACKEND
```bash
  cd .\backend\
  docker build -t qualitytrace-backend .
  docker run -p 8080:8080 --rm -v ${PWD}:/app -v /app/tmp --name qualitytrace-api-go-air qualitytrace-backend
```

### II.3 FRONTEND
WIP

## III. REFERENCES
### III.1 LINKS
- Air - [Live reload for Go apps](https://github.com/air-verse/air)
- Standard Go Project Layout - [Github](https://github.com/golang-standards/project-layout)
- graphql-go [Tutorial](https://www.howtographql.com/graphql-go/0-introduction/)
- [gqlgen](https://gqlgen.com/) is a library for creating GraphQL applications in Go.
- Example of go / graphql project on [graphql-golang](https://github.com/howtographql/graphql-golang) repository
- [Database migrations](https://github.com/golang-migrate/migrate)  written in Go

### III.2 COMMANDS
```bash
  go run github.com/99designs/gqlgen generate
  go run server.go
```