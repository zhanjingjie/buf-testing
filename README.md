# buf-testing
For testing some issues related to using buf.

Step 1: Install Buf
Brew was installed on the Mac through brew. 

Step 2: Working commands (all the Buf commands work)
* buf build protos
* buf lint protos
* buf format protos -w --exit-code
* cd protos && buf generate
* buf breaking protos --against '.git#branch=main,subdir=protos'

Step 3: Run Go
```
go run main.go
```
