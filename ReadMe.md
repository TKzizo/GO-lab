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
- ***byte operations***
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
                field , _ := FieldByName("Name")
            // field == required Max: "25"

        ```
            


>#### Reference 
    all appreciation goes to FreeCodeCamp and Micheal Van Sickle for the inspiration and help with the creation of this document.