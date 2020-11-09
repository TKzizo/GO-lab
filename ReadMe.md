# GO-lab 

>## Variales 
 - **variable must always be used else u get an error**.
 ---
 - **variable declaration**:

    1.  *implicit type declaration*:

     ```go
      var {name} {type} = [value] 
      ```

    2.  *let the compiler decide the type*:

    ```go
    {name} := {value}
    ```
---
- **Can't redeclare a varible but can shadow it be declaring inside another block.**
---
- **variable scoope**: 
    1. ***first*** *lower case variables*: 
        
        are only visible in their package .

    2. ***first*** *upper case variables*: 

        are visible to other packages.
_______________________

- **Converting variables**:

    ```GO
    var i int = 42 
    var k float32

    k = float32(i)
    
     ```
---
- **Converting to utf-8**:
    ```go
     k := 42 
     sign = string(k)

     fmt.Print(sign)
    ```
    *output*:
    ```go
    * //utf-8 value of 42
    ```
---
>## Packages 
 - ***fmt*** contains all main functions
 ---
 - **strconv** string convertion package :        
    - **Itoa** converts integer to string
 ---
 - ***reflect*** to use tags
 ---
 - ***math*** math functions:
    - **Abs()** absolute value. 
    - **Sqrt()** square root.
    - **Pow()** power.
---
 -  ***log*** show loggings:
    - **Println()** formated error
    - **Fatal** send -1 return  
---
 - ***Net/http*** net tools:
    - **Get()** return response and err
---
 - ***io/ioutil*** input/output utilities:
    - **ReadAll** read bytes
---
 - ***sync*** Go routines:
    - **WaitGroup{}** 
    - **add** add go routine to stack to wait for before terminating the execution.
    - **Wait()** wait for go routines to finish.
    - **Done()** send signal that the go routine finished.
---
 - ***time*** :
    - **Sleep( )**
    - **time.Millesecond**
---
 - ***runtime***:
    - **GOMAXPROCS()** can set how many theads program can run on, if given negative value it returns how many threads are available. 

>## Primitive types
- **Boolean**
    ```go
     var right_answer bool = true
     var wrong_answer bool = false    
    ```
---
- **Numerical**
    - ***Integers***
        ```go
         int    //default min 32 can go up 64
         int8   // -128 _ 127
         int16  // -32 768 _ 32 767
         int32  // -2 147 483 648 _ 2 147 483 647
         int64  // -9 223 372 036 854 775 808 _ 9 223 372 036 854 775 807
         ```
    - ***Unsigned integers***
        ```go
         uint8  //0 255
         uint16
         uint32
        ```
    - ***Floating numbers***
        ```go
        var f1 float32 
        var f2 float64  
        ```
    - ***complex numbers**
        ```go
        var i1 complex64 = 12 + 3i //both are float32
        var i2 complex128 = 5 + 2i //both are float64
        real(i1) // real part of complex
        imag(i2) // imaginary part of complex

        var i3 complex128 = complex(real(i1),real(i2)) //convert float to conplex

        ```
---
- **TEXT**
    - ***String***
        - string are immutable.
        ```go
          s := "helllo world"
          b := s[3] // is a byte encoding of letter "l" Unsigned integer 
          c := string(s[3]) // is letter l  

          b := []byte(s) // [values of each char in utf-8]
        ```
    - ***Rune***
        ```go 
        ///used when working with utf-32    
         var c rune = 'a' // c is int32 value 
        ```
>## Operations 
- ***Numerical operations***
    ```go
     +   //addition
     -   //substraction
     *   //multiplication
     /   //division
     %   //modulo     
     ```
- ***bitwise     operations***
     ```go
     &   // logic AND
     |   // logic OR
     ^   // XOR where either is set to 1 
     &^  // NOR where neither is set to 1
     value >> num // bit shift right num-times
     value << num // bit shift left num-time
    ```
>## Constants
 - **Constant declaration**
    ```go
        const myConst int = 42 //if it starts with upper case i will be exported as gloabl const

    ```
---
 - Constant can be ***Shadowed*** just like variables.

---
 - the type can be determined at comopilation:
    ```go
        const myConst = 44 
        var num int16 = 20

        var num2 = num + myConst // now myConst will be int16 as will be num
    ```
---
 - ***Enumerating constants***:
    ```go
        //declare a cosntant bloc 
        const (
            const1 = iota //iota is function used to keep track of const created
            const2 = iota 
            const3 // since we are in const bloc it uses same pattern and assign iota()
        )

        //const1 has value 0
        //cosnt2 has value 1
        //const3 has value 2
    ```
 - ***iota*** usage:
    ```go
        //enumeration of values
        const (
            _     = iota //if we don't need 0
            user1 = iota
            user2
        )
    ```
    ```go
        // since const must be defined at compilation and the power func is package
        // we can use bit shift with iota to replace that
        const (
            _   = iota 
            KB  = 1 << (10 * iota) 
            MB // this will be 1 << (10 * 2)
            GB
            TB
        )
    ```
    ```go
        const (
            isAdmin = 1 << iota //00000001
            isManager //00000010
            isSalesman //00000100

            canUseMetro //00001000
            canUseCar //00010000
            canUseTrain //00100000
        )
        var employee byte = isManager | canUseTrain //00100010  
    ```
>## Data Structures

>>### Arrays and slices:
   - **Array declaration:**
     ```go
        /// implicit
         name := [size]type {values, .. , ..}  
         /// same as:
         name := [...]type {values , .. , ..}
        /// explicit 
         var name [size]type  // empty array

         //Examples 
         name[2] = "hello"

         var length int = len(name) // length of the array

         var capaciy int = cap(name) // capacity of array 

         // multidiementional array 
         var 2d_table [3][3]int = [3][3]int {
                                                [3]int {3,3,3}, 
                                                [3]int {2,2,2},
                                                [3]int {1,1,1}
                                            }
      ```
   - arrays are ***passed as data***:
      ```go
        a := [...]int {1,2,3}
        b := a // now b has a copy of array a so if we change b nothing happens to a
        c := &a // c points to the same table a  

      ```
   - Slice declaration 
     ```go
      a := []int {1,3,4,5.10,22}

      b = a // point to same slice a 

      length := len(b) // legth of slice b
      capacityy := cap(b) // capacity of slcie b
      
      c := a[2:] // elemtst from 3rd index to last
      d := a[:6] // elemnts before 6th index
      e := a[1:5] //figure it out xDDD

     // create slice from array
      arr := [3]int {3,2,4}
      slic := arr[:]

      // add element to slice 
     a = append(a, values , ..)
    
     // concatinate slice to slace 
     a = append(a , []b{1,2,3}...) // the ... is called spreading 

     ```
   - ***Make*** function
     ```go
        a := make( []string , length )
        a := make( []int , length , capacity)
     ```
>>### Maps:
   -  **Note about Maps**:
        1. map is pair of key string.
        
        1. the key can be any primitive type as well as arrays but not slice.
        1. the order of the elements in maps isn't guaranteed.
        1. Maps are passed by reference.
   - **Maps declaration**:
     ```go
        a_map := map[key]type_of_values {
            "key1" : value1,
            "key2" : value2
        }

        //using the make function
        2nd_map := make(map[string]int) //gives an empty map

     ```
   - **Maps manipulation**:
     ```go
        // to get an element from the map we  specify the key:
        a_map := map[key]type_of_values {
            "key1" : value1,
            "key2" : value2
        }
        var1 := a_map["key1"]

        // to check if the key actualy exits we can also do this:
        _ , exist? := a_map["key"]
        fmt.Printf(exist?) // return true if key exist else false
                            // _ will contain always 0 if the key doesn't exist 

        // to delete an element from the map we can use "delete " function:
        delete(name_of_map , the_key)


        // to add an element to the map:
        a_map["new key"] = new_value
     ```  
>>### Structs:
   - **Notes about *Structs***:
        1. structs follow same capital rule as variables which means if struct name is catpital letter then it will be available to other packages however if its elemetns are lower case then the struct is visible but they are not.

        1. structs are passed by value.
        1. to pass struct pas reference we use **&** same as arrays.

   - **structs declaration**:
     ```go
     // first we define the struct 
     type Name_of_str struct {
        first_element type1
        second_elemetn type2
     }
            //now
            var1 := Name_of_str {
                first_element : value,
                second_element : value2
            }  
        
        // Anomymous struct:
            str := struct{name string}{name : "kat"}
        }
     ```
   - **Structs manipulation**:
        ```go
            // access element of struct:
            struct_name.name_element
        ```
     - similar to **inheretance** in other languages in Go we have **composition**:
     ```go
            type Shape struct{
                surface float32
            }
            type triangle struct{
                Shape
                _type string
            }
            // now the type Shape is embedded in triangle and we can access Surface same as any other elment of triangle

            //now to initialize a var we have two methods
            //1st method is:
            b := triangle{}

                b.surface = 12
                b._type = "straight"
            //2nd method 
            b := triangle{
                Shape : Shape{surface : 64},
                _type : "straight"
            }
            
     ```
     - **Tags** we use them to provide a text relative to a certain value, we need **reflect** package :
     ```go 
            type Animal struct {
                origin string 
                Name string `required Max "25"` 
            } 

            //to use the tag
                t := reflect.TypeOf(Animal{})
                field , _ := t.FieldByName("Name")
            // field == required Max: "25"

     ```

>## Conditionals:

>>### if statements:

   - **must always have "{}"**.

   - if we are testing multiple statement using  (OR " || " ) then the compliler stops at the first true statement , this is called **short circuting**.
    
   - format: "**else if**".
   -  

   ___
   - **useful usage**:
        ```go
            maps := map[string]int {
                "key1" : 123,
                "key2" : 456
            }

        if value,ok := maps["key3"] /* this is called initialiser*/ ; ok {
            fmt.Print(value) //value is only available inside the scope
        }

        //return booleen test 
            var1 := 50
            var2 := 44

            fmt.Println(var1 > var2 , var1 == var2) 
            //output 
                true false
        
        ```
>>### Switch statement:
   - the **break** is implicit so no need to write it ,but we can use it in some situations.
   - so if we want two cases to execute we use **fallthrough**.
---
   ```go
    switch value /* can be initialiser same as if statement*/ {
        case value1:
            //do this
        case value2: 
            //do that
        case value3, value4:
            //do somthing
        default:
            //do shit
    }
```
---
   - ***Type switch***
  ```go
    var i interface{} = 1 // or "hello" or [3]int{} or any primitive type
    
    switch i.(type) {

        case int:
            //do this
                if bla_bla { break} //when we don't want the rest of the bloc to execute
            //else do another thing
        case string:
            //do that:
        default:
            //something
    }
  ```
>## Loops: "**For**"
   - In Go there is only the **For** loop.
---
   - **simple loops**
     ```go
        //single index
        for i := 0; i < 5 ; i++ {
            //do somthing
        }

        //multiple indexes
        for i , j := 0,0 ; i < 5 ; i , j = i+1 , j+1{
            //do something crazy
        }

        //while looop
         var i int = 0
         for i < 5 {
             //do somethig
             i++
         }

        //do_while loop
        for {
            //do something
            if bla_bla {break}
        }
     ```
   - **Lables**:
     ```go
        // if we want to break off nested loops
        i := 0
        Loop: 
            for i < 10 {
                for j:=0;j<5;j++{
                    //do bla bla
                    if bla_bla {break Loop}
                }
                i++
            }
     ```

   - **looping collections**
     ```go
        //collection can be of any type array, slice, maps,String
        for key, value := range collection {
            fmt.Println(key , value)
        }     
        
        //if we just want the key like in maps we can
        for key := range maps{
            //print key
        }
     ```

>## Control Flow

>>### Defer:
   - multiples defer are executed in LIFO.
--- 
 ```go
    //suppose we made an http request and we need to work with the response before closing the response,
    //inorder not to forget that we can use "defer" to close it at any point but it will be closed just before the current function return.
     import (
         "fmt"
         "log"
         "net/http"
         "io/ioutil"
     )
     func main(){
        res , err := http.Get("http://www.google.com/robots.txt")

        if err != nil {
            log.Fatal(err) //couldn't get the document
        }

        defer res.Body.Close() //this should be last thing to do but since we putted defer it will be last to execute

        robots , err := ioutil.ReadAll(res.Body) //robots contains bytes
        if err != nil {
            log.Fatal(err) //couldn't read
        }
        fmt.Printf("%s",robots) // %v prints bytes

     }
 ```
 ```go
  // here is notice:
   a := "string"
   defer fmt.Print(a)    
   a = "hello"

   //output: 
    "string"

 ```
>>### Panic:
   - panic is to be used only in must situation and not just anywhere.

   - in Go we don't have exeptions but we have errors and panic throw an error.

   - all defer executes before terminating the function that panicked.
 ```go
    // we can stop the execution of program by panic at some erreur 

    fmt.Println("hello beautful world")
    panic("shit we are in the underẃorld")
    fmt.Println("i love this place") // this will not printed

    //output:
        hello beautful world
        panic: shit we are in the underẃorld
 ```
>>### Recover:
   - not to stop all the program from executing if a function panics we use recover on that function so the rest of the program can execute and we can get info about the err.
 ```go
    func main() {
        fmt.Println("hello there !!!")
        panicker()
        fmt.Println("goodbye my darkest old friend")
    }

    func panicker() {
        fmt.Println("it's about to go down")
        
        defer func() {
            if err := recover(); err != nil {
                log.Println("Error:",err)
            }
        }() // anonymous function

        panic("this is hell")
        fmt.("i'm sure this not gonna be written")
    }

    //output:
    """
        hello there !!!
        it s about to go down
        yyyy/mm/dd hh:mm:ss Error: this is hell
        goodbye my darkest old friend
    """
    // in case we can't handle the error we just panic(err) after the log.
 ``` 

>## Pointers:
   - we can't do arithmetic with pointers.
   ---  
 ```go
        //declaring pointers:
        a := 42
        var point *int 
        point = &a // hexadecimal number

        // structs and pointers:
        type doc struct {
            name int
        }
        
        var point *doc // <nil> pointer to structer 
        //or 
        point := new(doc) // &{0} pointer to empty structure
        //or 
        point := &doc{}

        (*point).name = 42
        //same as
        point.name = 42

 ```
 >## Functions:
   - **Notes**:
     1. functions are case sensitive same as variables.

   - **declaring functions**:
     ```go
        func functions_name(parameteres,..) return_type {
            //body 
        }

        //example:
        func hello(msg1,msg2 string,) {
            fmt.Println(msg1,msg2)
        }

        //working with poiters:
        func hello(name,msg *string) {
            // do stuff
            //...
            //...
        }
        //calling the functions:
         hello (&variable,&variable2)
     ```
   - **Go sugar functions xDD** aka **variadic parameters**
     ```go
      //we can send multiple data to a function without knowing how many until runtime
      //there is only one condition it has to be the last parameter:

      func koko(msg string, values ...int) int {

          fmt.Println(msg)
          result := 0
          for _ , val := range(values){
              result += val
          }
        return result
      }
      //calling the function
        sum := koko("hello world",1,2,3,5,9,8)

      //another way to declare the function is like this
      func koko(msg string, values ...int) (result int) { //result is instanced here
          
          for _,val := range values {
              result += val
          }

          return  // we dont have to precise because we already did in the begining
      }  
     ```  
   - **Returning pointers to local variables**:
     ```go
      // this may seem weid espicialy if you know how the local stack and the heap works
      // but in Go we have the abiliy to reference a local variable:
      func main() {
          variable := nameless() // this is pointer to varible name in func nameless()
          fmt.Println(*variable)
      }      

      func nameless() *string {
        name := "My name is Jeff"
        return &name
      }
     ```  
   - **Multiple parameters and errors**
     ```go
        func main() {

            d , err := divide(6.2 , 6.2)
            if err != nil {
                fmt.Println(err)
                return // to end the main function
            }

            fmt.Println(d)
        }

        func divide(val1 float64, val2 float64) (float64,error) {
            if b == 0 {
                return 0.0, fmt.Errorf("cannot divide by 0")
            }
            return a/b , nil
        }
     ``` 
   - **anonymous function**:
     ```go
        // we can create nameless functions and execute them instantly by adding () after {}

        func() {
            //do stuff
            //do stuff
        }()     

        // we can also pass function as variables 
        var f func(float64,int ) (int , error)

        f = func(a ,b float64) (int , error) {
            //do stuff
            //do stuff
        }

        val , err := f(6.0, 5.2)
     ```  
     
   - ***Methods in Go*** 
     ```go
        func main() {
            g := DOC {
                name : "john",
                spec: "surgery"
            }

            g.operate()
        } 
        
        type DOC struct {
            name string
            spec string
        }

        func (k DOC) operate() {  
            // know that the struct DOC is passed by value
            // we can pass it by refence by addig *, by doing that we can make changes to k

            fmt.Print("Dr.",k.name," is doing a ",k.spec)
        }
        
     ```
> ## Interfaces:
 - ***Notes*** :
    1. unlike struct, interfaces don't store data but behavior instead.
    1. In Go interfaces are **implemented implicitly** by creating a method that replicate their behavior.
    1. by convention the name of an interface should end with a "er"
    1. declarations inside interfaces must start with capitas to be used outside of package.
    1.

 - ***declaring an Interface*** :
    ```go
        type name(er) interface {
            // we declare a method without defining it:
            Create([]byte) (int)
        }

        type anything struct {
            // elemetns
        }

        func (obj (*)anything ) Create([]byte) int {
            // we must define
        }
    ```
 - ***Some use cases***:
     - suppose we want to create a new type of **int** that has a certain method:
        ```go
            func main() {
                it := Inter(0) // the brackets because it empty int 
                var ic IntCounter = &it
                for i := 0; i < 10 ; i++ {
                    fmt.Println( ic.incrementer() )
                }
            }
            type IntCounter interface {
                incrementer() int
            }

            type Inter int

            func (obj *Inter) incrementer() int {
                    *obj++
                    return int(*obj)
            }
        ```
     - composing multiple interfaces:
        ```go
            import "bytes" // for the sac of the example 

            func main() {
                var bwc WriterCloser = NewBufferedWriterCloser() 

                bwc.Write([]byte("hello world from my first real usage interface"))
                // so this will print 8 characters by line 
                bwc.Close()
                // if there is a less than 8 characters it will print them
            }

            type Writer interface {
                Write([]byte) (int,error)
            }
            type Closer interface {
                Close() error
            }

            type WriterCloser interface {
                Writer
                Closer
            }

            type BufferedWriterCloser struct {
                buffer *byte.Buffer
            } 

            //now BufferedWriterCloser must implement both Write and Close methods
            // the writer method
            func (obj *BufferdWriterCloser) Write(data []byte) (int,error) {
                n,err := obj.buffer.Write(data) // write date to buffer
                if err != nil {
                    return 0,err
                }
                
                v := make([]byte , 8) // makes slice with array capacity of 8 bytes 
                for obj.buffer.len() > 8 {

                    _ ,err := obj.buffer.Read(v) // read from buffer to v 
                    if err != nil {
                        return 0,err
                    }

                    _, err := fmt.Println(string(v))
                    if err != nil {
                        return 0,err
                    }
                }
            return n,nil
            }
            
            // the Closer method
            func (obj *BufferredWriterClose) Close() err {
                for obj.buffer.len() > 0 {
                    data := obj.buffer.Next(8) // read 8 bytes or until end of buffer and flush it
                    _, err := fmt.Println(string(data)) 
                    if err != nil {
                        return 0, err
                    }
                }
                return nil
            }

            // this function is only used to initialise the buffer inside the struct
            func NewBufferedWriterCloser() *BufferedWriterCloser {
                return &BufferedWriterclosed {
                    buffer : bytes.NewBuffer([]byte{})
                }
            }
        ```
     - now in order to access directly to the struct attributs we need to typecat our interface type variable into the type of the instance and to do that:

        ```go
            // we will continue use the same beforehand defined struct and interface
            func main() {
                var wc WriterCloser = NewBufferedWriterCloser()
                wc.Write([]byte("hello there again"))
                wc.close()
                
                // since NewBufferredWriterCloser() return a *BufferedWriterCloser we should type cast our wc variable to that, and we can assign it to another vaiable 

                bwc := wc.(*BufferedWriterCloser) 
                // Note:
                bwc := wc.(BufferedWriterCloser) //this will raise an error or panic we will know why later on

                // in order to be sure that there is no error we use same syntax as maps
                bwc,ok := wc.(*BufferedWriterCloser)
                if ok {
                    fmt.Println(bwc)
                }else {
                    fmt.Println("conversion failed")
                }
            }   

         ``` 
     - empty interfaces:

        ```go
            // in case we need to work with multiple non-compatible interfaces, the empty interface comes in handy because it allows us to use them all, by typecasting without getting any error like we did in the example before, because it doesn't require us to implement any methor:

            var x interface{}/*this is how we declare it*/ = NewBufferedWriterCloser()

            if wc, ok := x.(WriterCloser); ok {
                wc.Write([]byte("hello for the third time"))
                wc.close()
            }
            // and now x is WriterClosr so is wc
            // and if we do this:
            y, ok := x.(io.Reader) // just an interface from io package
            if ok {
                fmt.Println(y)
            }else {
                fmt.Println("conversion failed)
            }
        ``` 
     - since the start we have been using * in our method definition instead of the type directly, and we said in example above that not using * will raise an error:
        - that's because when we implemented the methods Write and Close:
            ```go
            func (obj *BufferedWriterClose) Close() error 
            func (obj *BufferdWriterCloser) Write(data []byte) (int,error)
            ```
            we had pointer receiver (*Buf..,) and not value receiver (Buff...); we did that because we needed to access internel data the byte[] slice and to avoid creating a copy by passing it by value; now if implemented them using value reciver we still be able to access them using pointer.
            ```Go
                var x WriterReceiver = &BuffferedWriterReader{}
            ```
     - ***BEST PRACTICES***:
        1. use small interfaces, that will avoid lot of conflict and debug,since strongest interfaces are always the smallest: 
            **io.Reader**, **io.Writer**, **interface**.

        1. Don't export interfaces for types that will be consumed, unlike db package, by that u can is without implementing every method.

        1. Do export interfaces for types that will be used by the package, something you can't do in java because of its explicit implementation of interfaces.
        
        1. design function and method to accept interfaces whenever possible, when you don't need acces to internel attributs.
        
>## Go-Routines:
   - ***GOROUTINES*** is not paralellism.
   - ***GOROUTINES*** are high level abstraction of thread which makes manipulating them more easy and with less risk.
   - the **main** function is a goroutine, the main one.


     ```go

     import ("fmt","time", "runtime","sync")

     var wg = sync.WaitGroup{}

     func main() {

         var msg = "Routines are awesome"
         wg.Add(1)
         go msger() // output: routines sucks

         msg = "routines sucks"
         //we can use time.sleep(100 * Millisecond) to replace the following but is really bad practine
         wg.Wait()
         
     }

     func msger() {
         fmt.Println(msg)
         wg.Done()
     }

     ```

>#### Reference 
   - all appreciation goes to FreeCodeCamp and Micheal Van Sickle for the inspiration and help with the creation of this document.

     **[FreecodeCamp Golang crash course  \*-*  ](https://www.youtube.com/watch?v=YS4e4q9oBaU)**
