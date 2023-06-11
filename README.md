# go-examples

This repo includes some examples about Go.

---
### References:

- [interface](https://200lab.io/blog/interface-trong-golang-cach-dung-chinh-xac/)
- [embedded struct](https://gobyexample.com/struct-embedding)
- [generics](https://gobyexample.com/generics)
- [errors](https://gobyexample.com/errors)
---
### Document (Vietnamese)

#### 1. Interface

Interface trong Golang là một kiểu được định nghĩa bởi tập hợp của các method (hàm trong Golang). Interface có thể chứa bất kỳ giá trị gì miễn là nó có implement các method này.

```
type Speaker interface {
    Speak() string
}
```

Cách sử dụng Interface:

- Implemented ngầm định:

    Với interface Speaker chỉ cần khai báo một struct và viết định nghĩa cho function Speak() cho struct đó:
    ```
    type Foo struct{}

    func (Foo) Speak() string {
        return "Hello, I am Foo"
    }
    ```
    Implement đầy đủ tại: [gointerface/simple.go](./gointerface/example/simple.go)

-  Empty Interface thay thế cho mọi kiểu dữ liệu:

    Dùng cú pháp `interface{}` để định nghĩa dạng dữ liệu tương tự như `any` trong TypeScript: [gointerface/empty.go](./gointerface/example/empty.go)
    ```
    func main() {
        var i interface{}
        describe(i)

        i = 42
        describe(i)

        i = "hello"
        describe(i)
    }

    func describe(i interface{}) {
        fmt.Printf("(%v, %T)\n", i, i)
    }
    ```
    Tại đây ta có thể thấy kiểu dữ liệu và giá trị được thay đổi không gây lỗi như khi ta xác định sẵn kiểu dữ liệu ban đầu.

    Một ví dụ phổ biến của với `interface{}` là sử dụng làm `value` cho `map`:
    ```
    product := make(map[string]interface{}, 0)

    product["name"] = "Iphone 13 Pro Max"
    product["price"] = 31000000
    product["quantity"] = 40
    ```

    Ép kiểu dữ liệu Interface thông qua cơ chế `Type assertions`:
    ```
    var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
    ```
    
Một số vấn đề thường gặp khác nên tham khảo tại [reference go interface](#references)

---

#### 2. Embedded struct

Đầu tiên ta định nghĩa một struct `base` như sau:

```
type base struct {
    num int
}
func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}
```

Sau đó tạo một struct `container` nhúng (`embed`) struct `base` bằng cách:

```
type container struct {
    base
    str string
}
```

Để khởi tạo struct `container` thì cần truyền cả dữ liệu của `base`:

```
co := container{
    base: base{
        num: 1,
    },
    str: "some name",
}
```

Khi này ta có thể truy cập trực tiếp field hay function của `base` như một field hay function của container hoặc truyền full path vào:

- `fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)`

    `fmt.Println("describe:", co.describe())` 

- `fmt.Println("also num:", co.base.num)`

Đồng thời function của `embedded struct` có thể coi như một cách để implement interface cho struct cha:

```
type describer interface {
    describe() string
}
var d describer = co
```

Full implement: [embedded struct](./embedded-struct/main.go)

---

#### 3. Generics

Chỉ support với Go version từ 1.18, có cách gọi khác là `type parameters`

Hiểu đơn giản thì đây là một phương pháp lập trình cho phép tham số hóa kiểu dữ liệu, tạo ra các function, struct có thể chấp nhận nhiều kiểu dữ liệu khác nhau.

Ví dụ cơ bản như tạo ra một hàm tính tổng cho mảng, khi chưa sử dụng `generics` ta cần định nghĩa function với mỗi kiểu dữ liệu khác nhau:
```
// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}
```

Với `generics` ta không định nghĩa rõ ràng kiểu dữ liệu của tham số đầu vào:

```
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```

Ngắn gọn hơn:

```
type Number interface {
    int64 | float64
}

// SumNumbers sums the values of map m. It supports both integers and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```

Mở rộng cho struct:

```
type List[T any] struct {
    head, tail *element[T]
}

type element[T any] struct {
    next *element[T]
    val  T
}

func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

func (lst *List[T]) GetAll() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}
```

Lưu ý:
- Có thể định nghĩa function cho generics type nhưng mà tham số truyền vào là generic type đã được định nghĩa ví dụ như trong trường hợp ở trên là List[T] chứ không phải List.
- comparable là cú pháp cho phép sử dụng == hoặc !=. Đây là bắt buộc với key trong map.

---

#### 4. Errors

Để throw error thì cần import `"errors"`, đây là một bộ thư viện đã được built-in sẵn trong Go.

Khởi tạo error:

```
errors.New("This is error message")
```

`nil` là một kiểu error nhưng có nghĩa là không có lỗi để có thể kiểm soát kết quả một function có lỗi hay không thì ta sử dụng như sau:

```
func f1(arg int) (int, error) {
    if arg == 42 {
         return -1, errors.New("can't work with 42")
    }

    return arg + 3, nil
}
```

Kiểm soát lỗi khi sử dụng function:

```
var value = 42

result, err := f1(value)

if err != nil {
    fmt.Println("f1 failed:", err)
} else {
    fmt.Println("f1 worked:", result)
}
```

Custom error:

```
type argError struct {
    arg  int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
```

Để khởi tạo custom error cần sử dụng `&`:

```
&argError{arg, "can't work with it"}
```

Nâng cao: [error blog post](https://go.dev/blog/error-handling-and-go)