**Some Golang concepts needed to understand the code:**
1. *Interfaces*
- In our code we have used interfaces as regular interfaces and as function arguments too. 
- A function can accept an argument which is an interface (regular or empty).
- Such functions that take an interface argument can take any type as argumet
that implements that interface.
- A function that takes an empty interface as an argument can take any type as argument.
<br><br>

2. *...type*
- ...type when passed as argument to a functions means that we can pass any number of
parameters of that type to the function.
<br><br>


3. *maps*
- To create an empty map, use make function.
- map[key-type]value-type
<br><br>

4. *mutex*
- Mutual Exclusion ie mutex is the concept that only one goroutine can access a variable at a time
to avoid conflicts.
- We can use the mutex object to lock and unlock access to a variable (not exactly lock, we restrict
it to only one goroutine at a time).
- We can also use defer with mutex.Unlock so as to ensure that it unlocks the variable after use.
<br><br>


5. *json package*
- JSON is basically a data format for sending and receiving data on the web.
- Marshalling a go object helps us convert it to a json string.
- We can use MarshalIndent instead of Marshal to add pretty printing, by providing a prefix and an indent parameter.
- 
<br><br>


6. filepath package
- This contains many functions to use with filepaths.
- Clean function returns the shortest path name equivalent to the path by pure lexical processing.
- Dir function returns the directory (the last one) used in the path.
- Rel is a function that takes two paths as argumets and gives a string that can be added to one of them to form the other.
<br><br>


7. lumber package
- This is a package used for logging.
<br><br>


8. ioutil package
- It is a package used to work with files.
<br><br>


9. os package
<br><br>
