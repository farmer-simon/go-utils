# egu
easy golang utils

## import
```
go get github.com/ebar-go/egu
```
## String
- UUID: generate unique id
```
id := egu.UUID()
```

- DefaultString: return default value when string is empty
```
var s string
s = egu.DefaultString(s, "defaultValue")
```

## Number
- Max: compare two number and return the max number
```
max := egu.Max(1, 2) // max is 2
```

- Min: compare two number and return the min number
```
min := egu.Min(1, 2) // min is 1
```

## Json
```
s, err := egu.JsonEncode(map[string]interface{}{"name":"test"})
fmt.Println(s, err)
m := make(map[string]interface{})
decodeErr = egu.JsonDecode(s, &m)
fmt.Println(decodeErr, m)
```

## Date
- GetTimeStr: return current time, eg:2019-12-30 22:00:00
```
now := egu.GetTimeStr()
``` 

- GetDateStr(): return current date, eg: 2019-12-30
```
day := egu.GetDateStr()
```

## Encrypt
- Md5: return md5 encrypt string
```
s := egu.Md5("123456")
```

- Sha1: return hash encrypt string
```
s := egu.Sha1("123456")
```

## Convert
- Str2Byte: convert string to byte
```
b := egu.Str2Byte("123")
```

- Byte2Str: convert byte to string
```
s := egu.Byte2Str([]byte("123"))
```

## Array
- Implode: concat array with separator
```
arr := []int{1,2,3}
s := egu.Implode(arr, "-") // s is 1-2-3
```

- Explode: split string with separator
```
arr := egu.Explode("1-2-3", "-")
fmt.Println(arr.Items())
```

## File
- PathExist: check file or directory exist
```
if egu.PathExist("/home") {
  fmt.Println("exist")
}
```

- Mkdir: create directory if not exist
```
egu.Mkdir("/var/log/app", true)
```

- RuntimeCaller: return code invoke trace, always use in panic/recover
```
fmt.Println(egu.RuntimeCaller())
```