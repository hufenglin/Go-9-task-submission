通常来说，应该要将这个错误抛给上层。因为DAO只执行调用者给的操作，一个操作失败，应该有调用者判断与处理。


```go
//DAO
func Dao(query string) error{
    err := mockError()

    if err == sql.ErrNoRows {
        return errors.Wrapf(NotFound, fmt.Sprintf( format:"data not found, sql: %s ", query))
    }

    if err != nil {
        return errors.Wrapf(err, fmt.Sprintf( format:"db query system error sql: %s", query))
    }
    //做一些其他的处理逻辑
    return nil
}
```

```go
//服务层
var NotFound = errors.New( message:"not found")
func Biz() error {
    err := Dao(query:"")

    if errors.Is(err, NotFound) {
        return nil
    }

    if err != nil {
        //出现了数据库查询的问题，可以继续向上传递
    }
    return nil
}
```