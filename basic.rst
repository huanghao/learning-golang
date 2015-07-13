various types::

  var v1 int
  var v2 string
  var v3 [10]int    // array?
  var v4 []int      // slice
  var v5 struct {
    f int
  }
  var v6 *int
  var v7 map[string]int
  var v8 func(a int) int    // function? or point to a function?

multiple vars::

  var (
    v1 int
    v2 string
  )

initialization::

  var v1 int = 10
  var v2 = 10
  v3 := 10

assignment::

  var v1 = int
  v1 = 123

  i, j = j, i

  _, _, nickName := GetNames()  //returns firstName, lastName and nickName

literal::

  -12
  3.14
  3.2+12i
  true
  "foo"

constant::

  const Pi float64 = 3.1415926
  const zero = 0.0
  const (
    size int64 = 1024
    eof = -1
  )
  const u, v float32 = 0, 3
  const a, b, c = 3, 4, "foo"
  const mask = 1 << 3 // compile time init

iota: the 9th letter of Greek alphabet::

  const (
    c0 = iota   // c0 = 0
    c1 = iota
    c2 = iota   // c2 = 2
  )

  const (
    c0 = iota
    c1
    c2
  )

  const (
    c0 = 1 << iota
    c1 = 1 << iota
    c2 = 1 << iota
  )

  const (
    c0 = 1 << iota
    c1
    c2
  )

  const x = iota  // 0  
  const y = iota  // 0

  const (
    Sun = iota
    Mon
    Tue
    Wed
    Thr
    Fri
    Sat
    numberOfDays  // private
  )

types
=====

- bool: true, false
- numeric: int8, byte, int16, int, uint, uintptr, float32, float64, complex64, complex128
- string, rune
rune: a letter of an ancient Germanic alphabet, related to Roman alphabet
- error

- pointer
- array
- slice
- map
- chan
- struct
- interface

Bool
----

::

  var b bool
  b = 1         // wrong
  b = bool(1)   // wrong
  b = (1 != 0)  // right

Integers
--------

- int8, unit8(byte)
- int16, uint16
- int32, uint32
- int64, uint64
- int, unit: recommended
- uintptr

int and int32 are different types::

  var v2 int32
  v1 := 3
  v2 = v1         // compile error
  v2 = int32(v1)  // compile pass

numeric operations: +, -, \*, /, %
compare operations: >, <, ==, >=, <=, !=

different types of integers can'be compared, for example int8 and int.

bitwise operations: <<, >>, ^, &, \|, ^

Float
-----

float32, float64(double)

comprasion between floats::

  import "math"

  // p is user defined percise, such as 0.00001
  func IsEqual(f1, f2, p float64) bool {
    return math.Fdim(f1, f2) < p
  }

Complex
-------

complex64 consists of two float32 for its real and image parts

complex64 + complex64 -> complex128

::

  z = complex(x, y)
  real(z) == x
  imag(z) == y

std libs: math, complx

String
------

It's immutable::

  var str string
  str = "Hello world"
  ch := str[0]

  str[0] = 'X'    // wrong
  // cannot assign to str[0]

  fmt.Printf("Length of str is %d\n", len(str))

golang only support unicode and UTF-8.
See go-iconv: https://github.com/xushiwei/go-iconv

string operations::

  "Hello" + "123"
  len("Hello")
  "Hello"[1]

stdlib: strings

iteration::

  // byte array
  for i := 0; i < len(str); i++ {
    ch := str[i]  // type "byte"
    fmt.Println(i, ch)
  }

  // unicode
  for i, ch := range str {
    fmt.Println(i, ch)  // type "rune"
  }

stdlib: unicode

Array
-----

Fixed size of the same type of elements::

  [32]byte
  [2*N] struct { x, y int32 }
  [1000]*float64
  [3][5]int
  [2][2][2]float64

  for i := 0; i < len(array); i++ {
    fmt.Println("Element", i, "of array is", array[i])
  }

  for i, v := range array {
    fmt.Println("Array Element[", i, "]=", v)
  }

array is treated as value type when passing to function as argument::

  package main
  import "fmt"

  func modify(array [10]int) {
    array[0] = 10
    fmt.Println("In modify():", array)
  }

  func main() {
    array := [5]int{1, 2, 3, 4, 5}
    modify(array)
    fmt.Println("In main():", array)
  }

Slice
-----

::

  s = a[5:]
  s := make([]int, 5)
  s := make([]int, 5, 10) // reverse capacity of 10
  s := []int{1, 2, 3, 4}

  for i := 0; i < len(s); i++ {
    s[i]
  }
  for i, v := range s {
    i, v
  }

  len(s) == 5
  cap(s) == 10

  s = append(s, 1, 2, 4)

  s2 := int{8, 9, 10}
  s = append(s, s2...)

  var s = make([]int, 5, 10)
  fmt.Println(s)  // [0 0 0 0 0]

  var s = make([]int, 3, 100)
  fmt.Println(s)
  var s2 = make([]int, 0, 10)
  s2 = append(s2, 1, 2, 3)
  s = append(s, s2[2:5]...)
  fmt.Println(s)

Map
---

make::

  m := make(map[string] int)
  m["kilo"] = 1000
  m["million"] = 1000 * 1000

  m = map[string] int {
    "byte": 8,
    "word": 16,
  }

delete::

  delete m["byte"]
  delete m["not exists"]

  delete m[nil] // panic

get::

  var m = make(map[int]int)
  m[1] = 0
  fmt.Println("value is: ", m[1], "\nnot exist:", m[2])
