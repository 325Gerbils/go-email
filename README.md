# go-email

Simple package to send text emails.

### Installation:

`go get github.com/325gerbils/go-email`

### Usage:

```go
import (
    "log"
    email "github.com/325gerbils/go-email"
)

email.Auth("username@gmail.com", "myPassword", "smtp.gmail.com:587")


err := email.Send("recipient@website.com", "Subject", "This is what will appear in the email body")
if err != nil {
    log.Println("Error sending email:", err)
}
```
