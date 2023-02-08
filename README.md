# buf-testing
For testing some issues related to using buf.

Step 1: Install Buf
Brew was installed on the Mac through brew. 

Step 2: Working commands
* buf build protos
* buf lint protos
* buf format protos -w --exit-code
* cd protos && buf generate

Step 3: Non-working command
* buf breaking protos --against '.git#branch=main'

Error: protos/zhanjingjie/buftesting/buftesting/v1/buftesting_api.proto:10:8:zhanjingjie/buftesting/buftesting/v1/buftesting.proto: does not exist
