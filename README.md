# Golang

It's fast, statically typed complied language that feels like a dynamically typed, interpreted language

Go has a concept of types that is either explicity declared by a programmer or is inferred by the complier

## Modules & CLI

A module is a group of packages
It contains a go.mod file with configuration & metadata

CLI manipulates the module

-- go mod init
-- go build
-- go run filename.go (to run single)
-- go run . (to run module)
-- go test
-- go get

## Defining variables

```go
// Data types goes after identifier
// variables have nil by default
// constants can be only bool, string or numbers

| Verb         |  Description  |
| -----------  | -----------   |
| int          | 0             |
| float64      | 0.0           |
| string       | ""            |
| bool         | false         |
| pointers     | nil           |
| functions    | plain string  |
| interfaces   | true or false |
| maps         | floating numbers |

var x int // x is 0
var name string // name is 0
var isTest bool // isTest is false
const y = 2 // y is 2

// const variable cannot be re-assigned or it is immutable

// varible initialization shortcut (only for var)
// this shortcut cannot be used outside function 
otherText := "Bye!"

// Go implicitly get the type based on the value u have assign
var text = "string";
text1 := "string";


// we can declare variables outside of the main func
var url = "string"
func main(){
  print(" url ", url)
}

func multipleValues(x int, y int) int {
  return x + y
}

// or

func multipleValues(x, y int) int {
  return x + y
}

// func can return muliple values
func returnMultipleValues(x, y int) (int, int) {
  sum := x + y
  diff := x - y
  return sum, diff
}

sum, difference = returnMultipleValues(20, 10);

// named return parameters
func operation(x, y int) (sum int, diff int) {
  // we just initalizing the value, bz it is already declared in the function return
  sum = x + y
  diff = x -y
  return // return keyword should be end of the function
}

// variadic functions
func sumOfNumbers(numbers ...int) int {
  sum := 0
  for _, value := range numbers {
    sum += value
  }
  return sum
}

print(sumOfNumbers(10))
print(sumOfNumbers(10, 30, 40))
print(sumOfNumbers(10, 29, 89))

func printDetails(name string, subjects ...int) {
  println(" hey ", student, ", here are your subjects -")
  for _, sub := range subjects {
    printf("%s ", sub)
  }
}

// anonymous func
 x := func(x , y int) int {
  return x + y
 }

fmt.Printf("%T \n", x) // func(int, int) int
fmt.Println(x(20, 30)) // 50

// anonymous func without arguments
 x := func(x , y int) int {
  return x + y
 }(20, 30)


fmt.Printf("%T \n", x) // func(int, int) int
fmt.Println(x) // 50

// higher order function is a func that receives a func as an argument
func printResult(radius float64, calcFunc func(r float64) float64){
  result := calcFunc(radius)
  fmt.Println("Result: ", result)
  fmt.Println("Thank you! ")
}

func getFunc(query int) func(r float64) float64{
  query_to_func := map[int]func(r float64) float64{
    1: calcArea,
    2: calcPerimeter,
    3: calcDiameter
  }

  return query_to_func(query)
}

// Arrays & slices
[1, 2, 4, 5] 

["foo", "bar"]

[7.0, 9.43]

// array declaration
var grades [5]int
print(grades) // [0,0,0,0,0]
var fruits [2]strng
print(fruits) // [ ]

// array initialization
var grades [3]int = [3]int{10, 20, 30}
// or
var grades := [3]int{10, 20, 30}
// or
var grades := [...]int{10, 20, 30}

// array methos
print(len(grades))

// looping through array

for i:=0; i < len(grades); i++ {
  print(grades[i])
}

//or
for index, ele := range grades {
  print(index, "=>" , ele)
}

// multi dimensional arrays
arr := [3][2]int{{2,4}, {4,5}}
print(arr[0][0])


// slice
// continous segment of an underlying array
// variable typed (elements can be added or removed)
// more flexible

arr := [4]int{10,20,40,50}
slice := arr[0:1] // 10

// another way declaring slice, using make func
slice := make([]int , 5, 8) // make([]<data_type>, length, capacity)

print(slice) // [0,0,0,0,0]
print(len(slice)) // 5
print(cap(slice)) // 8

// maps

"x" -> 30 // key : value 
1 -> 100 // integer to integer map key : value
"key" -> "value" 


// unordered collection of key/value pairs
// implemented by hash tables
// provide efficient add, get & delete operations
var map map[<key_data_type>]<value_data_type>
var my_map map[string]int

var codes map[string]string
map := map[<key_data_type>]<value_data_type>{<key-value-pairs>}

codes := map[string]string{"en" : "English", "fr": "French"}
print(codes) // map[en:English fr:French]

print(codes["en"]) //English


// getting a key
value, found := map_name[key]

value, found := codes["en"] // true English
value, found := codes["ed"] // false ""

// adding new k-v
codes["it"] = "Italian"
// updating k-v
codes["en"] = "English updated"
// deleting k-v
delete(codes, "en")

// iterate

for key, value := range codes {
  print(key , "=>" , value)
}

for key, value := range codes {
  print(key , "=>" , value)
}

// truncate a map

for key, value := range codes {
  delete(codes, key)
}
// map[]
// or

codes = make(map[string]string) // map[]
```

## Built-in Data Types

-- string (16 bytes)
-- Integer values: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64
int - 4 bytes for 32-bit machines, 8 bytes for 64-bit machines
int8 - 8 bits or 1byte
int16 - 16bits or 2bytes
-- Floating point values: float32, float64
-- Boolean: bool (1 byte)

Boolean operators: ==, !=, <, >, <=, >=, && , ||, !

## Printf - format specifiers

| Verb      |  Description  |
| ----------- | ----------- |
| %v          | default format |
| %T          | type of the value |
| %d          | integers |
| %c          | char |
| %q          | quoted characters/string |
| %s          | plain string |
| %t          | true or false |
| %f          | floating numbers |
| %.2f          | floating numbers upto 2 decimal places |

```go

// Find the type of the variable
var grades int = 21
print("grade = %v is of type %T \n", grades, grades)

// using the reflect.TypeOf()
import ("fmt", "reflect")

print(" Type: \n", reflect.TypeOf(1000))



// converting btw types

// int to float
var i int = 90
var f float64 = float64(i)

// int to float
var f float64 = 45.89
var i int = int(f) // 45, but it end up losing the precision of float

// strconv package (can convert string to int & vice versa)
Itoa()
  -- converts integer to string
  -- returns one value - string formed with the given integer

Atoi()
  -- converts string to integer
  -- returns two values - the corresponding integer, error (if any)

import "strconv"

var i int = 42
var s string = strconv.Itoa(i) // convert int to string
print(" %q ", s) // "42"

var s string = "200"
i, err := strconv.Atoi(s) 
print(" %v, %T ", i, i) // 200, int
print(" %v, %T ", err, err) // <nil>, <nil> (if there is not error occured)
```

## user input & output

```go
var name string
var age int
print(" What is your name : ")
fmt.Scanf("%s ", &name)
print(" Hello ", name)

// reat multiple inputs
print(" What is your name : & your age")
fmt.Scanf("%s %d", &name, &age)
print(" Hello ", name, age)

// scanf return values

Count: the no of argms that the fn writes to
err: any error thrown during the execution of the fn

count, err := fmt.Scanf("%s, %d", &a, &b)

```

## condition statements

```go
var a string = "happy"
if a == "happy" {
  print(a)
} else if a == "not bad" {
  print(" not bad")
} else {
  print(" bad")
}

// switch

var i int = 10
switch i {
  case -5:
    print("-5")
  case 10:
    print("10")
  case 20:
    print("20")
  default:
    print("default")
}


// fallthrough keyword
// is used in switch-case to force the execution flow to fall through the successive case block
switch i {
  case -5:
    print("-5")
  case 10:
    print("10")
    fallthrough
  case 20:
    print("20")
    fallthrough
  default:
    print("default")
}

output:
10
20
default


// for loop

for i := 1; i <=5; i++ {
  print(i*i)
}
// or
i :=1
for i <=5 {
  print(i*i)
  i += 1
}

```

## Pointer

When we talk about memory management in programming we generally deal with RAM, whenever variable is declare a certain amount of memory in the RAM is allocated for it & it is based on the data type of the variable being used in the program.

This memory allocation is done while the program is running, & hence the variables might get a different address every time a program is run.

Pointer is a varible that holds the memory address of another variable.

They point where the memory is allocated & provide ways to find or even change the value located at the memory address

```go
 x := 1

 | memory address | memory |
 | 0x0301         | 1      |
 | 0x0302         |        |
 | 0x0303         |        |
 | 0x0304         | 0x0301 |
 | 0x0305         |        |

 var ptr *int := &x // 0x0301
```

### address & dereference operators

-- & operator (address-of) : The address of a variable can be obtained by preceding the name of a variable with an ampersand sign (&), known as address-of operator.
-- * operator (dereference operator): It is known as the dereference operator. When placed before an address, it returns the value that address.

```go
x := 1

 | memory address | memory |
 | 0x0301         | 1      |
 | 0x0302         |        |

 &x = 0x0301 (it will return address of x)
 *0x0301 = 77 (it would give us the value at that address)
```

```go
i := 10
printf("%T %v \n" , &i, &i) // *int 0xc000018c008
printf("%T %v \n" , *(&i), *(&i)) // int 10

```

### declaring a pointer

var <pointer_name> *<data_type>

var ptr_i *int

var ptr_i *string

var <pointer_name> *<data_type> = &<variable_name>

```go
var i *int // <nil> zero of a pointer is nil
var s *string // <nil>

i := 10
var ptr_i *int = &i

// another method to initialize the pointer
var <pointer_name> = &<variable_name> // datatype internally determined by the complier
s := hello
var ptr_s = &s

// shorthand operator
<pointer_name> := &<variable_name>
s := "hello"
ptr_s := &s


s := "hello"
var a *string = &s
println(a)
var b = &s
println(b)
c := &s
println(c)

```

### dereferencing a pointer

By using dereferencing pointer, we can change the value of that particular address

```go
s := "hello"
printf("%T %v \n", s, s)
ps := &s
*ps = "world"
printf("%T %v \n", ps, ps)

// output
string hello
string world
```

### passing by value

There are two ways of passing arguments to a function, Passing by value & passing by reference.

Passing by value in functions

-- Func is called by directly passing the value of the variable as an argument
-- the Paramter is copied into another location of your memory
-- So, when accessing or modifying the variable within your function, only copy is accessed or modified & original value is never modified
-- All basic types(int, float, bool, string, array) are passed by value

```go
  func modify(s string){
    s = "world"
  }

  func main(){
    a := "hello"
    println(a) // hello
    modify(a)
    println(a) // hello
  }
```

### Passing by reference function

-- Go supports pointers, allowing you to pass reference to values within your program.
-- In call by reference/pointer, the address of the varible is passed into the function call as the actual parameter
-- All the operations in the func's are performed on the value stored at the address of the actual parameters.

Slices & Maps are basically passed by reference by default

```go

  func modify(s *string){
    *s = "world"
  }

  func main(){
    a := "hello"
    println(a) // hello
    modify(&a)
    println(a) // world
  }

```

## Structs, Methods & Interfaces

Struct is a user-defined data type
-- a structure that groups together data elements.
-- provide a way to reference a series of grouped values through a single variable name.
-- used when it makes sense to group or associate two or more data variables.

```go

type <struct_name> struct {
  // list of fields
}

type Circle struct {
  x float64
  y float64
  r float64
}

type Student struct {
  name string
  rollNo int
  marks []int
  grades map[string]int
}

// intialization


type Student struct {
  name string
  rollNo int
  marks []int
  grades map[string]int
}

var s Student
// +v special debug 
print("%+v", s) // {name: rollNo:0 marks:[] grades: map[]}


// An instance of a struct can also be created using the new keyword

<variable_name> := new(<struct_name>)

st := new(Student)


// using new keyword
type Student struct {
  name string
  rollNo int
  marks []int
  grades map[string]int
} 

st := new(Student)
printf("%+v", st)

// &{ name: rollNo:0 marks:[] grades:map[] }


// another way of initalization of struct
<variable_name> := <struct_name> {
  <field_name> : <value>,
  <field_name> : <value>,
}

//
st := Student {
  name: "Suresh",
  rollNo : 490
}

// 
type Student struct {
  name string
  rollNo int
}

st := Student {
  name : "Suresh",
  rollNo: 490
}

fmt.Printf("%v+", st) // {name: Joe rollNo: 12}

// omit the field names
st := Student{"Joe", 12}

fmt.Printf("%v+", st) // {name: Joe rollNo: 12}

```

### accessing the struct fields

<variable_name>.<field_name>

```go
 type Circle struct {
  x int
  y int
  radius int
 }

 var c Circle
 c.x = 5
 c.y = 5
 c.radius = 5
 fmt.Printf("%+v \n", c) // {x:5 y:5 radius:5}
 fmt.Printf("%+v \n", c.area) // c.area undefined (type Cirlce has no field or method area)

```

### compare struct

```go
type c1 struct {
  x int
}

type c2 struct {
  x int
}

c := s1{x:5}
c1 := s2{x:5}

if c == c1 {
  print("yes")
}

// invaild operation: c == c1 (mismatched types s1 & s2)

type c1 struct {
  x int
}

c := s1{ x: 5 }
c1 := s1{ x: 6 }
c2 := s1{ x: 5 }
 
if c != c1 {
  fmt.Println(" c & c1 have different values ") // c & c1 have different values
}

if c == c2 {
  fmt.Println(" c is same as c2 ") // c is same as c2 
}

```

### Methods sets

-- set of methods that are available to a data type.
-- useful way to encapsulate functionality

-----------------------------

## Advanched Go

### Concurrency

Sequential Programming: The commands or instructions are executed one after another in a linear fashion. This means command or instructions are not executed until the previous command is entirely completed.
Multitasking is the ability of the operating system to execute more than one task simultaneously on CPU. This basically happens in CPU by switching btw jobs using a small time quantum.
Concurrency: is the notion of multiple things happening at the same time. It is the potential for multiple processes to be in progress at the same time. We can also say that concurrency is achieved through multitasking.

### Go routines

Considered as a lightweight thread that has a separate independent execution

Can execute concurrently with other go-routines

Managed entirely by Go runtime

Go-routines are non-deterministic

```go
// syntax
// golang uses a special keyword called go for starting a go-routine

go func() or method()

// the func or method executed in the go-routine
```

### main go-routines

The main function in the main package is the main go routine. All go-routine are started from the main go-routine. These go-routines can then start multiple other go-routines & so on. But the main go-routine represents the main program.

The go-routines do not have parents or children. when you start the go-routine, it just executes alongside all other go-routines. each go-routine exists only when its function returns. The only execution to that is that all go-routines exit when the main go-routine, the one that runs function main exists.

Go-routines are independent

### Anonymous go-routine

-- In golang anonymous func are those funs that don't have any name. Simply put, anonymous funcs don't use any variables as a name when they are declared.
-- Anonymous funcs in golang can also be called using go-routine.

```go
go func(){

}(arg...)

import ("time")

func main(){
  go func(){
    println("In anonymous method")
  }()

  time.sleep(1 * time.Second)
}

```

### go runtime scheduler

What happens when you start go program?

### Wait groups

1) The problem with go-routines was the main go-routine terminating before the go-routines completed or even began their execution.

2) To wait for multiple go-routines to finish, we can use a wait group.

3) A wait group is a synchronization primitive that allows multiple go-routines to wait for each other.

```go
import "sync"

var wg sync.WaitGroup


```
