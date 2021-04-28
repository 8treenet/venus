# extjson
###### 更灵活和定制的json。

## Overview
- 兼容官方
- 可配置风格
- 可使用Tag 指定格式化时间


## 安装
```go
$ go get github.com/8treenet/venus
```


#### JSON命名使用小写驼峰风格
```go
import (
	"testing"
	"time"

	"github.com/8treenet/venus/extjson"
)

type Style struct {
    StyleText    string
    StyleNumber  int
    StyleBoolean bool
    StyleTag     string `json:"tagtagtag"` //如果指定name，会使用指定name
}

func TestLowerCamelCase(t *testing.T) {
    //设置配置 NamedStyle: extjson.NamedStyleLowerCamelCase 小写驼峰风格
    extjson.SetDefaultOption(extjson.ExtOption{NamedStyle: extjson.NamedStyleLowerCamelCase})

    //extjson.Marshal 序列化
    styleBytes, _ := extjson.Marshal(Style{StyleText: "extjson", StyleNumber: 100, StyleBoolean: true, StyleTag: "Tag"})
    fmt.Println(string(styleBytes))
    //输出: {"styleText":"extjson","styleNumber":100,"styleBoolean":true,"tagtagtag":"Tag"}

    var read Style
    extjson.Unmarshal([]byte(`{"styleText":"extjson","styleNumber":100,"styleBoolean":true,"tagtagtag":"Tag"}`), &read)
    fmt.Println(read)
    //输出: {extjson 100 true Tag}
}
```

#### JSON命名使用下划线
```go
import (
	"testing"
	"time"

	"github.com/8treenet/venus/extjson"
)

func TestUnderScoreCase(t *testing.T) {
    //extjson.NamedStyleUnderScoreCase : 下划线风格
    extjson.SetDefaultOption(extjson.ExtOption{NamedStyle: extjson.NamedStyleUnderScoreCase})
    styleBytes, _ := extjson.Marshal(Style{StyleText: "extjson", StyleNumber: 100, StyleBoolean: true, StyleTag: "Tag"})
    fmt.Println(string(styleBytes))
    //输出: {"style_text":"extjson","style_number":100,"style_boolean":true,"tagtagtag":"Tag"}

    var read Style
    extjson.Unmarshal([]byte(`{"style_text":"extjson","style_number":100,"style_boolean":true,"tagtagtag":"Tag"}`), &read)
    fmt.Println(read)
    //输出: {extjson 100 true Tag}
}
```

#### JSON空数据风格
```go
import (
	"testing"
	"time"

	"github.com/8treenet/venus/extjson"
)

func TestNull(t *testing.T) {
    extjson.SetDefaultOption(extjson.ExtOption{
        NamedStyle:       extjson.NamedStyleLowerCamelCase, //小写驼峰
        SliceNotNull:     true, //空数组不返回null, 返回[]
        StructPtrNotNull: true, //nil结构体指针不返回null, 返回{}
    })
    var out struct {
        Slice  []*Style
        Slice2 []Style
        Struct *Style
    }
    outBytes, _ := extjson.Marshal(out)
    fmt.Println(string(outBytes))
    //输出: {"slice":[],"slice2":[],"struct":{}}
}

```


#### JSON时间格式化
```go
import (
	"testing"
	"time"

	"github.com/8treenet/venus/extjson"
)

type Timeformat struct {
    Time1 time.Time `json:",timeformat=ms"`
    Time2 time.Time `json:",timeformat=sec"`
    Time3 time.Time `json:",timeformat=2006-01-02 15:04:05"`
    Time4 time.Time
    Time5 *time.Time `json:",timeformat=sec"`
}

func TestTimeformat(t *testing.T) {
    now := time.Now()
    out := Timeformat{
        Time1: now,
        Time2: now,
        Time3: now,
        Time4: now,
        Time5: &now,
    }
    outBytes, _ := extjson.Marshal(out)
    t.Log(string(outBytes))
    //out : {"Time1":"1597917504308","Time2":"1597917504","Time3":"2020-08-20 17:58:24","Time4":"2020-08-20T17:58:24.308275+08:00","Time5":"1597917504"}

    var in Timeformat
    extjson.Unmarshal([]byte(`{"Time1":"1597917504308","Time2":"1597917504","Time3":"2020-08-20 17:58:24","Time4":"2020-08-20T17:58:24.308275+08:00","Time5":"1597917504"}`), &in)
    t.Log(in)
    //out : {2020-01-12 11:20:00.468 +0800 CST 2020-01-12 11:20:00 +0800 CST 2020-01-12 11:20:00 +0000 UTC 2020-01-12 11:20:00.468543 +0800 CST}
}
```

#### Gin集成
```go
import (
    "github.com/8treenet/venus/extjson"
    "net/http"
)
func init() {
    extjson.SetDefaultOption(extjson.ExtOption{NamedStyle: extjson.NamedStyleUnderScoreCase})
}

func handle(c *gin.Context) {
    var data struct {
        UserName string
        UserID   int
    }
    data.UserName = "zhangsan"
    data.UserID = 1001
    c.Render(http.StatusOK, extjson.GinRender(data)) 
}
```

#### iris集成
```go
type MyContext struct {
	iris.Context    //继承iris的上下文
	session *sessions.Session
}
// JSON 继承重写
func (ctx *MyContext) JSON(v interface{}, options ...JSON) (int, error) {
    //实现extjson.Marshal
}

// ReadJSON 继承重写
func (ctx *MyContext) ReadJSON(jsonObjectPtr interface{}) error{
    //实现extjson.Unmarshal
}

func Handler(h func(*MyContext)) iris.Handler {
	return func(original iris.Context) {
        myCtx := &MyContext{Context : original}
		h(myCtx)
		release(ctx)
	}
}

func main() {
    app := iris.New()
    app.Use(Handler)
}

//by implementing a new `context.Context` (see https://github.com/kataras/iris/blob/master/_examples/routing/custom-context/new-implementation/main.go)
````