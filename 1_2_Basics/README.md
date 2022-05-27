# Chapter 1. Basics

* [For Loop](#for-loop)
* [~~While Loop~~ (`for` is Go's `while`)](#while-loop-for-is-gos-"while")
* [If and Else](#if-and-else)
* [Switch](#switch)
* [Defer](#defer)

> Exercise:
> 1. [Exercise 1](/1_2_Basics/exercise_1.go)
> 2. [Exercise 2](/1_2_Basics/exercise_2.go)

### For Loop

* `break`
* `continue`

* Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the `for` statement and the braces `{ }` are always required.
* The init and post statements are optional.

```go
sum := 0

for i := 0; i < 10; i++ {
	sum += i
}

for ; sum < 1000; {
	sum += sum
}
```

### ~~While Loop~~ (For is Go's "while")

```go
for sum < 1000 {
	sum += sum
}

for { // forever
}
```

### If and Else

```go
if x < 0 {
	return "less than 0"
}
```

* The if statement can start with a short statement to execute before the condition.

    ```go
    func pow(x, n, lim float64) float64 {
        if v := math.Pow(x, n); v < lim {
            return v
        }
        return lim
    }
    ```

* with `else`

```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
```

### Switch

* Go only runs the selected case, not all the cases that follow.
* The break statement that is needed at the end of each case in those languages is provided automatically in Go.
* Go's switch cases need not be constants, and the values involved need not be integers.

```go
switch os := runtime.GOOS; os {
case "darwin":
	fmt.Println("OS X.")
case "linux":
	fmt.Println("Linux.")
default:
	// freebsd, openbsd,
	// plan9, windows...
	fmt.Printf("%s.\n", os)
}
```

* Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

```go
switch i {
    case 0:
        ...
    case f(): // does not call f if i==0.
        ...
}

fmt.Println("When's Saturday?")
today := time.Now().Weekday()
	
switch time.Saturday {
case today + 0:
	fmt.Println("Today.")
case today + 1:
	fmt.Println("Tomorrow.")
case today + 2:
	fmt.Println("In two days.")
default:
	fmt.Println("Too far away.")
}
```

* Switch without a condition is the same as switch true. (This construct can be a clean way to write long if-then-else chains.)

```go
t := time.Now()
switch {
case t.Hour() < 12:
	fmt.Println("Good morning!")
case t.Hour() < 17:
	fmt.Println("Good afternoon.")
default:
	fmt.Println("Good evening.")
}
```

### Defer

A defer statement defers the execution of a function until the surrounding function returns. \
The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

* Stacking defers

Deferred function calls are pushed onto a stack.
> last-in-first-out

```go
fmt.Println("counting")

for i := 0; i < 10; i++ {
	defer fmt.Println(i)
}

fmt.Println("done")
```

```plain text
counting
done
9
8
7
6
5
4
3
2
1
0
```