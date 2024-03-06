原文：[类型断言与类型开关](https://rednafi.com/go/type_assertion_vs_type_switches/)

---

Type assertion vs type switches in Go | Redowan's Reflections



Despite moonlighting as a gopher for a while, the syntax for type assertion and type
switches still trips me up every time I need to go for one of them.

So, to avoid digging through the docs or crafting stodgy LLM prompts multiple times, I
decided to jot this down in a gobyexample[1](#fn:1) style for the next run.

## Type assertion[#](#type-assertion)

Type assertion in Go lets you access an interface’s underlying concrete type. It’s a way to
get the dynamic type of an interface. The syntax is `x.(T)`, where `x` is an interface and
`T` is the type you’re asserting.

### Basic usage[#](#basic-usage)


```
var i interface{} = "Hello" // or use `any` as an alias for `interface{}`

s := i.(string)
fmt.Println(s)

```
Here, `s` gets the type `string`, and the program outputs `Hello`.

### Checking types and values[#](#checking-types-and-values)


```
var i interface{} = 42

if v, ok := i.(int); ok {
 fmt.Println("integer:", v)
}

```
This code checks if `i` is an `int` and prints its value if so. The value of `ok` will be
`false` if `i` isn’t an integer and nothing will be printed to the console.

## Checking composite types and values[#](#checking-composite-types-and-values)


```
var i interface{} = []string{"apple", "banana", "cherry"}

if v, ok := i.([]string); ok {
 fmt.Println("slice of strings:", v)
}

```
This will print `slice of strings: [apple banana cherry]` to the console.

Similar to primitive types, you can also perform type assertions with composite types. In
the example above, we check whether the variable `i`, which is of an interface type, holds a
value of the type ‘slice of strings’.

### Handling failures[#](#handling-failures)


```
var i interface{} = "Hello"

f := i.(float64) // This triggers a panic

```
Wrong assertions, like attempting to convert a string to a float64, cause runtime panics.

## Type switches[#](#type-switches)

Type switches let you compare an interface’s type against several types. It’s similar to a
regular switch statement, but focuses on types.

### Basic usage[#](#basic-usage-1)


```
var i interface{} = 7

switch i.(type) {
case int:
 fmt.Println("i is an int")
case string:
 fmt.Println("i is a string")
default:
 fmt.Println("unknown type")
}

```
This outputs `i is an int`.

### Using a variable in a type switch case[#](#using-a-variable-in-a-type-switch-case)


```
var i interface{} = []byte("hello")

switch v := i.(type) {
case []byte:
 fmt.Println(string(v))
case string:
 fmt.Println(v)
}

```
Notice how we’re assinging variable `v` to `i.(type)` and then reusing the extracted value
in the case statements. The snippet converts `[]byte` to a string and prints `hello`.

### Handling multiple types[#](#handling-multiple-types)


```
var i interface{} = 2.5

switch i.(type) {
case int, float64:
 fmt.Println("i is a number")
case string:
 fmt.Println("i is a string")
}

```
The `case T1, T2` syntax works like an OR relationship, outputting `i is a number`.

### Addressing composite types[#](#addressing-composite-types)


```
var i interface{} = map[string]bool{"hello": true, "world": false}

switch i.(type) {
case map[string]bool:
 fmt.Println("i is a map")
case []string:
 fmt.Println("i is a slice")
default:
 fmt.Println("unknown type")
}

```
Similar to primitive types, you can check for composite types in the case statement of a
type switch. Here, we’re checking whether `i` is a `map[string]bool` or not. Running this
will output `i is a map`.

## Similarities and differences[#](#similarities-and-differences)

### Similarities[#](#similarities)

* Both handle interfaces and extract their concrete types.
* They evaluate an interface’s dynamic type.

### Differences[#](#differences)

* Type assertions check a single type, while type switches handle multiple types.
* Type assertion uses `x.(T)`, type switch uses a switch statement with `i.(type)`.
* Type assertions can panic or return a success boolean, type switches handle mismatches
more gracefully.
* Type assertions are good when you’re sure of the type. Type switches are more versatile
for handling various types.
* Type assertion can get the value and success boolean. Type switches let you access the
value in each case block.
* Type switches can handle multiple types, including a default case, offering more
flexibility for various types.

Fin!



---

1. [Go by example](https://gobyexample.com/) [↩︎](#fnref:1)
## Recent posts

- [Retry function in Go](https://rednafi.com/go/retry_function/)
- [Patching pydantic settings in pytest](https://rednafi.com/python/patch_pydantic_settings_in_pytest/)
- [Omitting dev dependencies in Go binaries](https://rednafi.com/go/omit_dev_dependencies_in_binaries/)
- [Eschewing black box API calls](https://rednafi.com/misc/eschewing_black_box_api_calls/)
- [Annotating args and kwargs in Python](https://rednafi.com/python/annotate_args_and_kwargs/)
- [Rate limiting via Nginx](https://rednafi.com/go/rate_limiting_via_nginx/)
- [Statically enforcing frozen data classes in Python](https://rednafi.com/python/statically_enforcing_frozen_dataclasses/)
- [Planning palooza](https://rednafi.com/zephyr/planning_palooza/)
- [Reminiscing CGI scripts](https://rednafi.com/go/reminiscing_cgi_scripts/)
- [Debugging dockerized Python apps in VSCode](https://rednafi.com/python/debug_dockerized_apps_in_vscode/)
* [Go](https://rednafi.com/tags/go/)
* [TIL](https://rednafi.com/tags/til/)

[« Prev  
Retry function in Go](https://rednafi.com/go/retry_function/)[Next »  
Patching pydantic settings in pytest](https://rednafi.com/python/patch_pydantic_settings_in_pytest/)[blogroll](/blogroll) • [reads](/reads) • [uses](/uses)  
  
© 2020-2024 [Redowan Delowar](https://rednafi.com/)let menu=document.getElementById("menu");menu&&(menu.scrollLeft=localStorage.getItem("menu-scroll-position"),menu.onscroll=function(){localStorage.setItem("menu-scroll-position",menu.scrollLeft)}),document.querySelectorAll('a[href^="#"]').forEach(e=>{e.addEventListener("click",function(e){e.preventDefault();var t=this.getAttribute("href").substr(1);window.matchMedia("(prefers-reduced-motion: reduce)").matches?document.querySelector(`[id='${decodeURIComponent(t)}']`).scrollIntoView():document.querySelector(`[id='${decodeURIComponent(t)}']`).scrollIntoView({behavior:"smooth"}),t==="top"?history.replaceState(null,null," "):history.pushState(null,null,`#${t}`)})})var mybutton=document.getElementById("top-link");window.onscroll=function(){document.body.scrollTop>800||document.documentElement.scrollTop>800?(mybutton.style.visibility="visible",mybutton.style.opacity="1"):(mybutton.style.visibility="hidden",mybutton.style.opacity="0")}document.querySelectorAll("pre > code").forEach(e=>{const n=e.parentNode.parentNode,t=document.createElement("button");t.classList.add("copy-code"),t.innerHTML="copy";function s(){t.innerHTML="copied!",setTimeout(()=>{t.innerHTML="copy"},2e3)}t.addEventListener("click",t=>{if("clipboard"in navigator){navigator.clipboard.writeText(e.textContent),s();return}const n=document.createRange();n.selectNodeContents(e);const o=window.getSelection();o.removeAllRanges(),o.addRange(n);try{document.execCommand("copy"),s()}catch{}o.removeRange(n)}),n.classList.contains("highlight")?n.appendChild(t):n.parentNode.firstChild==n||(e.parentNode.parentNode.parentNode.parentNode.parentNode.nodeName=="TABLE"?e.parentNode.parentNode.parentNode.parentNode.parentNode.appendChild(t):e.parentNode.appendChild(t))})
