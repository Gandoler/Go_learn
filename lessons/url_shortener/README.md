# МЕГА Сокращатор по [видео](https://www.youtube.com/watch?v=rCJvW2xgnk0&list=LL&index=3&t=205s)

> оригинальный [репозиторий](https://github.com/GolangLessons/url-shortener/tree/main)




## Библиотеки ......


```bash
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
go get github.com/go-chi/render

go get github.com/fatih/color
go get github.com/go-playground/validator/v10
go get github.com/stretchr/testify/assert
go get github.com/stretchr/testify/require
go get github.com/stretchr/testify/mock

go install github.com/vektra/mockery/v2@latest

```
go:generate go run github.com/vektra/mockery/v2@latest --name=URLSaver

>не уверен что это все


а это уже в тестах

```bash
go get  github.com/brianvoe/gofakeit/v6
go get 	github.com/gavv/httpexpect/v2

```
