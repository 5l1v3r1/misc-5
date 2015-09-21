# Linux PAM module in go

Requires go 1.5 since prior versions do not allow creating linked libraries.


Compile:
```
go build -buildmode=c-shared -o pam_go.so -v
```
