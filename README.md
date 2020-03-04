# go-email

Simple package to send text emails. It's a wrapper for go's smtp library to reduce the lines of code required to send an extremely simple email.

It consists of two functions Auth() and Send(). Auth creates a private smtp.PlainAuth and stores it for later use, Send encodes the arguments into an email body, then uses the PlainAuth to send the email. View the source code for a better understanding of what's going on. The source code is shorter than this documentation.

THIS IS NOT THREAD-SAFE. To hack in thread-safety, lock a mutex before calling the functions and unlock it afterwards. See usage examples

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

import (
    "log"
    "sync"
    
    email "github.com/325gerbils/go-email"
)

mut := sync.Mutex{}

// Lock to this thread and auth
mut.Lock()
email.Auth("username@gmail.com", "myPassword", "smtp.gmail.com:587")
mut.Unlock()

// Lock to this thread and send
mut.Lock()
err := email.Send("recipient@website.com", "Subject", "This is what will appear in the email body")
if err != nil {
    log.Println("Error sending email:", err)
}
mut.Unlock()
```
