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
- **Converting to Ascii**:
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
 - 

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
