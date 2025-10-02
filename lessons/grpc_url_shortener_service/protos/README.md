


## команды


* для протока

```bash
$env:PATH += ";C:\Users\glkru\OneDrive\Desktop\moilubimiy\prrotoc\bin"
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc -I proto proto/sso/sso.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative 
```

* или для крутых перцев

```bash
$env:PATH += ";C:\Users\glkru\OneDrive\Desktop\moilubimiy\prrotoc\bin"
 $env:PATH += C:\Users\glkru\AppData\Local\Microsoft\WinGet\Links\task.exe
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
 task generate
```