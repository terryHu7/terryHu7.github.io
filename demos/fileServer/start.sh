
#!/bin/bash
nohup go run src/server.go  >logs/file_upload.log 2>>logs/err.log &
