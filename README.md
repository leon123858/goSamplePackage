# goSamplePackage
goSamplePackage

A golang package sample

## How to create

1. 利用 go.mod 機制在 local 直接撰寫一個專案, 像這個 github repo 一樣, 一個 repo 一個 mod
2. 第一層 package 名稱要跟 mod 名稱一樣, 但往下層可以放很多其他 package
3. 需利用 github 的 Release 功能設置 Release, tag 即為該次 release 的版本號
4. (optional)版本號從 v0.0.1 開始, 並且要符合 Semantic Versioning
5. (optional)也可以用指令操作 `git tag v0.0.1"` 並且 `git push origin v0.0.1` 來達成
6. 別的專案使用 `go get github.com/leon123858/goSamplePackage` 即可使用
7. 編輯本專案: 改完 code 再 release 下一版本即可

## use the package

1. `go get github.com/leon123858/goSamplePackage`
2. `import "github.com/leon123858/goSamplePackage"`

```go
goSamplePackage.Fibonacci(15)
```