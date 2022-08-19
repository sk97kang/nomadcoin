# Nomadcoin

make nomadcoin with Go

## 시작하기

1. Install Go
2. `go mod init <github repository>`

## Vriable

- 변수의 선언

```go
var name string = "nico"
// 변수를 생성하면서 값에서 타입을 추론하여 지정
name := "nico"
```

- 번수 타입
  - bool
  - string
  - int
  - uint

## Function

- 함수의 선언

```go
func main() {
    // go 실행 시 항상 실행되는 함수
}
```

- 함수의 사용

```go
func plus(a, b int) int {
	return a + b
}

func main() {
	result := plus(1, 2)
	fmt.Println(result)
}
```
