## json转换相关函数
```go
func Marshal(v interface{}) ([]byte, error)  
func Unmarshal(data []byte, v interface{}) error
```

在使用过程中，遇到过这种
```go
type T struct{
    Port int64 `json:"port"`
}
```
结构体成员是以int64类型存储，但是有时候想把它当作string类型使用，刚开始对json转换不熟悉，就是unmarshal之后，只能再对该成员二次转换成string。后来详细阅读golang json相关文档，发现有如下的小技巧

> 结构体的每一个成员都能够指定结构体成员的格式化字符串(存储在成员tag标签json中)来定制.这个字符串指定成员名称，还可以额外跟随一个通过逗号分割的可选项列表。如果仅仅是为了单独使用这个可选项，成员名称可以为空（使用默认的成员名称)  

### 标签可选项列表
1. "omitempty":指定如果结构体成员为空值，则该成员在encode时应该被忽略   
2. "-": 该成员转换时候均被忽略  
3. "string": 表示该字段是作为字符串形式存储在json字符串中，可以通过指定成员类型算是作为隐式转换。只能使用在字符串，folat,int,bool类型的字段中

```go
// Field appears in JSON as key "myName".
Field int `json:"myName"`

// Field appears in JSON as key "myName" and
// the field is omitted from the object if its value is empty,
// as defined above.
Field int `json:"myName,omitempty"`

// Field appears in JSON as key "Field" (the default), but
// the field is skipped if empty.
// Note the leading comma.
Field int `json:",omitempty"`

// Field is ignored by this package.
Field int `json:"-"`

// Field appears in JSON as key "-".
Field int `json:"-,"`  
//字符串转换为int类型
Int64String int64 `json:",string"`
```

### json第三方库
有时候添加了若干个结构体，可能结构体需要添加json,xorm等标签，手动一个个添加是相当的麻烦。在github上找了一番，找到了一个可以添加标签的库
[gomodifytags](https://github.com/fatih/gomodifytags),省下了相当的手工
还有有时候参考文档，文档会有示范，需要我们自己根据示范返回结果生成结构体
在线json转golang结构体 [json-to-go](https://mholt.github.io/json-to-go/) 很好用


[参考http://colobu.com/2017/06/21/json-tricks-in-Go/](http://colobu.com/2017/06/21/json-tricks-in-Go/)
[参考https://golang.org/pkg/encoding/json/](https://golang.org/pkg/encoding/json/)