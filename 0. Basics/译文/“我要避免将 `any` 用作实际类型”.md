原文：[“我要避免将  用作实际类型”](https://utcc.utoronto.ca/~cks/space/blog/programming/GoAvoidingAnyAsAType)

---


 
 Chris's Wiki :: blog/programming/GoAvoidingAnyAsAType 






[Chris Siebenmann](/~cks/) ::
[CSpace](/~cks/space/) »
 [blog](/~cks/space/blog/) »
 [programming](/~cks/space/blog/programming/) »
 GoAvoidingAnyAsAType
Welcome, guest.




## In Go, I'm going to avoid using '`any`' as an actual type


January 25, 2024

As modern Go programmers know, when Go introduced generics it also
introduced a new '`any`' type. This is [officially documented](https://go.dev/ref/spec#Interface_types) as:



> For convenience, the predeclared type `any` is an alias for the empty
> interface.
> 
> 
> 


The 'any' type (alias) exists because it's extremely common in code
that's specifying generic types to want to be able to say 'any
type', and the way this is done in generics is 'interface{}', the
empty interface. This makes generic code clearly easier to read and
follow. Consider these two versions of the signature of [reflect.TypeFor](/~cks/space/blog/programming/Go122ReflectTypeFor)



> 
> ```
> 
> func TypeFor[T any]() Type
> func TypeFor[T interface{}]() Type
> 
> ```
> 
> 


These are semantically equivalent but the first is clearer, because
you don't have to remember this special case of what 'interface{}'
means. Instead, it's right in the name 'any' (and there's less
syntactic noise).


But after Go generics became a thing, there's been a trend of
using this new '`any`' alias outside of generic types, instead of
writing out 'interface{}'. I don't think this is a good idea.
To show why, consider the following two function signatures,
both of which use 'any':



> 
> ```
> 
> func One[T any](v T) bool
> 
> ```
> 
> 
> ```
> 
> func Two(v any) bool
> 
> ```
> 
> 


These two function signatures look almost the same, but they have
wildly different meanings, even if (or when) they're invoked with
the same argument. The effects of '`One(10)`' are rather different
from '`Two(10)`', since 'One' is a generic function while 'Two' is
a regular one. Now consider them written this way:



> 
> ```
> 
> func One[T any](v T) bool
> 
> ```
> 
> 
> ```
> 
> func Two(v interface{}) bool
> 
> ```
> 
> 


Now we see clearly what `Two()` is doing differently than `One()`;
it's obvious that it isn't taking 'any type' as such, but instead
it's taking a generic interface as the argument type. This makes
it obvious that a non-interface value will be converted to an
interface value (and will tell some people that [an interface value
will lose its interface type](/~cks/space/blog/programming/GoNilIsTypedSortOf)).


This increased immediate clarity without needing to remember what
'`any`' is why I'm planning to use 'interface{}' in my code in
the future, and why I think you should too. Yes, '`any`' is shorter
and it has a well defined meaning in the specification and we can
probably remember the special meaning all of the time. But why give
ourselves that extra cognitive burden when we can be explicit?


(In generics, the argument goes the other way; 'any' really does
mean 'any type', and the 'any' name is clearer than writing
'interface{}' and then needing to remember that that's how generics
do it.)


In a sense the 'any' name is a misnomer when used as a type. It's true
that 'interface{}' will accept any type, but used as a type, it doesn't
mean 'any type'; it means specifically the type 'an empty interface',
which is to say an interface that has no methods, which implies
interface type conversion (unless you already have an 'interface{}'
value). Since 'any' does mean 'any type' in the context of generics,
I think it's better to use a different name for each thing, even if
Go formally makes the names equivalent. The names of things are
fundamentally for people.



([6 comments](/~cks/space/blog/programming/GoAvoidingAnyAsAType?showcomments#comments).)
Written on [25](/~cks/space/blog/2024/01/25/) [January](/~cks/space/blog/2024/01/) [2024](/~cks/space/blog/2024/).   



|  |  |  |  |  |  |
| --- | --- | --- | --- | --- | --- |
| 

|  |  |
| --- | --- |
|  «  | [The cooling advantage that CPU integrated graphics has](/~cks/space/blog/tech/CPUIGPCoolingAdvantage) |

 | 

|  |  |
| --- | --- |
| [Histogram data is most useful when they also provide true totals](/~cks/space/blog/tech/HistogramsNeedTotalsToo) |  »  |

 |




 These are my [WanderingThoughts](/~cks/space/blog/)   

([About the blog](/~cks/space/AboutBlog))


[Full index of entries](/~cks/space/blog/__Index)   

[Recent comments](/~cks/space/blog/__RecentComments)


This is part of [CSpace](/~cks/space/FrontPage), and is written by [ChrisSiebenmann](/~cks/space/People/ChrisSiebenmann).   

Mastodon: [@cks](https://mastodon.social/@cks)   

Twitter @thatcks


\* \* \*


Categories: [links](/~cks/space/blog/links/), [linux](/~cks/space/blog/linux/), [programming](/~cks/space/blog/programming/), [python](/~cks/space/blog/python/), [snark](/~cks/space/blog/snark/), [solaris](/~cks/space/blog/solaris/), [spam](/~cks/space/blog/spam/), [sysadmin](/~cks/space/blog/sysadmin/), [tech](/~cks/space/blog/tech/), [unix](/~cks/space/blog/unix/), [web](/~cks/space/blog/web/)   

Also: [(Sub)topics](/~cks/space/blog/__Topics)


This is a [DWiki](/~cks/space/dwiki/DWiki).   

[GettingAround](/~cks/space/help/GettingAround)   

([Help](/~cks/space/help/Help))


 
 Search:  





---

 Page tools: [View Source](/~cks/space/blog/programming/GoAvoidingAnyAsAType?source), [Add Comment](/~cks/space/blog/programming/GoAvoidingAnyAsAType?writecomment). 

Search: 

Login: 
Password: 


 

Atom Syndication: [Recent Comments](/~cks/space/blog/programming/GoAvoidingAnyAsAType?atomcomments).


---

 Last modified: Thu Jan 25 23:03:50 2024   

This dinky wiki is brought to you by the Insane Hackers
Guild, Python sub-branch.



