# jsondiff

## import

```shell
go get -u github.com/hellojqk/jsondiff
```

## example

```go
var json1 = `{"int":1,"string":"string","float":1.1,"array":[1,2,3],"object":{"int":1,"string":"string"},"objectAry":[{"int":1,"string":"string"}]}`
var json2 = `{"int":12,"string":"string2","float":1.12,"array":[12,22,32],"object":{"int":1,"string":"string2"}}`
Diff(json1, json2, false)

var json1Bytes = []byte(json1)
var json2Bytes = []byte(json2)
DiffBytes(json1Bytes, json2Bytes, false)
```
