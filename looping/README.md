### Looping

#### 1. simple loops
##### formats
1. for initializer; test; incrementer {} 
--------

     eg.

     	for i := 0; i < 5; i++ { // looping by adding 1
		fmt.Println(i)
	}

--------
2. for test {}  //  DO-WHILE loop implementation
--------
   	k := 0
	for k < 5 {
		fmt.Println(k)
		k++
	}
--------
3.  for {}
--------
   

    	m := 0
	for {
		fmt.Println(m)
		m++
		if m == 5 {
			break
		}
	}
--------

#### 2. exiting early
1. break - breaks out of the immediate loop that its executing
2. continue - it doesnt break out of the current loop it just breaks out of the current iteration, goes back to the incrementer and continues execution of the for-loop and the next increment
3. labels - it breaks out of an inner loop, if you have 2 nested loops and you need to break out of the outer loop you use a label assigned to the outer loop
------
eg.

    Loop:
        for i := 1; i <= 3; i++ {
            for j := 1; j <= 3; j++ {
                fmt.Println(i * j)
                if i*j >= 3 {
                    break Loop // breaks out of the inner loop and outer loop labelled `Loop`
                }
            }
        }
------

#### 3. looping through collections (slices, arrays,maps, strings and channels)
#### SYNTAX
-----
    for k, v := range collection {}
-----
##### a. slices
------
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}
------

##### b. arrays
-------
	ar := [3]int{1, 2, 3}
	for k, v := range ar {
		fmt.Println(k, v)
	}
-------

##### c. maps
-------
    statePopulations := map[string]int{type int
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	for k, v := range statePopulations {
		fmt.Println(k, v)
	}
-------

##### d. strings
---------
	text := "hello Go!"
	for k, v := range text {
		fmt.Println(k, string(v))
	}
---------
