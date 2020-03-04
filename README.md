# go-email

Simple package to send text emails. THIS IS NOT THREAD-SAFE. To hack in thread-safety, lock a mutex before calling the functions and unlock it afterwards. See usage examples

### Installation:

`go get github.com/325gerbils/go-email`

### Usage:

```go
import (
    "log"
    email "github.com/325gerbils/go-email"
)

// Add authentication data
email.Auth("username@gmail.com", "myPassword", "smtp.gmail.com:587")

// Send email
err := email.Send("recipient@website.com", "Subject", "This is what will appear in the email body")
if err != nil {
    log.Println("Error sending email:", err)
}
```

Thread-safe hack:

```go
// Thread-safe sending:
import "sync"

mut := sync.Mutex{}

mut.Lock()
email.Auth("username@gmail.com", "myPassword", "smtp.gmail.com:587")
mut.Unlock()

mut.Lock()
err := email.Send("recipient@website.com", "Subject", "This is what will appear in the email body")
if err != nil {
    log.Println("Error sending email:", err)
}
mut.Unlock()
```
