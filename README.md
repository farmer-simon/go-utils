# readme
golang utils

## import
```
go get github.com/farmer-simon/go-utils
```
## String
- UUID: generate unique id
```
id := utils.UUID()
```

- DefaultString: return default value when string is empty
```
var s string
s = utils.DefaultString(s, "defaultValue")
```

- RandomString return random string
```
s:=utils.RandomString(15)
```

- Convert character encoding
```
s := "GBK 与 UTF-8 编码转换测试"
gbk, err := utils.Utf8ToGbk([]byte(s))
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(string(gbk))
}

utf8, err := utils.GbkToUtf8(gbk)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(string(utf8))
}
```
## Number
- Max: compare two number and return the max number
```
max := utils.Max(1, 2) // max is 2
```

- Min: compare two number and return the min number
```
min := utils.Min(1, 2) // min is 1
```

## Json
```
s, err := utils.JsonEncode(map[string]interface{}{"name":"test"})
fmt.Println(s, err)
m := make(map[string]interface{})
decodeErr = utils.JsonDecode(s, &m)
fmt.Println(decodeErr, m)
```

## Date
- GetTimeStr: return current time, eg:2019-12-30 22:00:00
```
now := utils.GetTimeStr()
``` 

- GetDateStr(): return current date, eg: 2019-12-30
```
day := utils.GetDateStr()
```

## Encrypt
- Md5: return md5 encrypt string
```
s := utils.Md5("123456")
```

- Sha1: return hash encrypt string
```
s := utils.Sha1("123456")
```

## Convert
- Str2Byte: convert string to byte
```
b := utils.Str2Byte("123")
```

- Byte2Str: convert byte to string
```
s := utils.Byte2Str([]byte("123"))
```

## Array
- Implode: concat array with separator
```
arr := []int{1,2,3}
s := utils.Implode(arr, "-") // s is 1-2-3
```

- Explode: split string with separator
```
arr := utils.Explode("1-2-3", "-")
fmt.Println(arr.Items())
```

## File
- PathExist: check file or directory exist
```
if utils.PathExist("/home") {
  fmt.Println("exist")
}
```

- Mkdir: create directory if not exist
```
utils.Mkdir("/var/log/app", true)
```

- RuntimeCaller: return code invoke trace, always use in panic/recover
```
fmt.Println(utils.RuntimeCaller())
```
