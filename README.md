# gokit-example
This is a simple application using go-kit, used to demonstrate Clean Architecture in Go


# Run
```
go run ./cmd/simpleapp
```

# APIs
### Get article

```
GET 127.0.0.1/9999/article/:id
```

### Create Article
```
POST 127.0.0.1/9999/article/:id

{
    "title": "...",
    "text": "..."
}
```

The full article describing the Clean Architecture concepts can be found here:
https://orenrose.medium.com/clean-architecture-in-golang-with-go-kit-e5b716a3b881
