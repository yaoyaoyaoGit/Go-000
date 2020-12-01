 # 问题

数据库操作的时候，比如DAO层中遇到一个sql.ErrNoRows的时候，是否应该Wrap这个error，抛给上层？为什么？写出代码

# 答案
不应该，大部分情况下应该直接return空值。
例如：
```go
func dao() (string, error)
	id := 43
	var username string
	err = stmt.QueryRowContext(ctx, id).Scan(&username)
	switch {
	case err == sql.ErrNoRows:
		return "", nil
	case err != nil:
		return "", err
	default:
		return username, nil
	}
```

1. 如果DAO层返回了sql.ErrNoRows错误，Service层需要去判断这个Error是不是sql.ErrNoRows，这个Dependency是多余的。并且再判断完这个Error后，仍然需要去判断
数据结构是否为空，防止panic。
例如：
```go
func biz() error{
    user, err := dao()
    if err != nil{
        if err.Is(sql.ErrNoRows){
            return fmt.Errof("404 error")
        }
        else{
            return err
        }
    }
    if user == nil{
        return fmt.Errof("404")
    }
    return nil
}
```
2. 主要原因是sql.ErrNoRows这个错误，在大部分情况下和Service Unavailable也就是500的错误是不在同一优先级的，大部分应用不在意这个错误。

3. https://github.com/xo/xo/issues/60 这里的讨论也挺不错有道理的的，只有Go语言把No results当成一个error，并且ErrNoRows和数据体是nil是完全等价的。