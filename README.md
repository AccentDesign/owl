<img src="owl-logo.svg" alt="banner" width="150">

Some example components styled like [Shadcn](https://ui.shadcn.com/docs) using [Templ](https://templ.guide).
This is meant as a starting point for UI in a project and NOT as an installable package.

Things will be added, changed or deleted at will.

### What's going in here

In the past i've built several projects in various forms of web technology and have always tried to make things re-usable.
This can be tricky at times when you want to swap out different html elements such as `button`, `anchor` etc. You tend to end
up with arbitrary attributes to swap things out, or end up just duplicating the whole thing anyway.

This led me to think: Is there a better way?

Assuming the important thing is the attributes you want to define on the html tag and not the tag itself. This is where
these examples are heading to.

* Attribute structs can be defined to (in the end) return `templ.Attributes`. These I have explicitly defined in structs. See the [attrs](./attrs) for more examples.
* Separate attribute structs are created for anchor and image elements to ensure clarity.
* Elements have been defined for the like of `div`, `anchor`, `button` etc that just accept the defined attributes and render the element and it's children. See the [elem](./elem)  for more examples.
* Example css is all separate to avoid bloat. See the [css](./css) for more examples.

Simple Example:

```go
func alert() attrs.Attributes {
    return attrs.Global{
        Class: templ.CSSClasses{"owl-alert"},
    }
}

func alertTitle() attrs.Attributes {
    return attrs.Global{
        Class: templ.CSSClasses{"owl-alert-title"},
    }
}

func alertDescription() attrs.Attributes {
    return attrs.Global{
        Class: templ.CSSClasses{"owl-alert-description"},
    }
}

templ Alert() {
    @elem.Div(alert()) {
        @icons.Terminal("")
        @elem.Div(alertTitle()) { Lorem ipsum dolor sit amet! }
        @elem.Div(alertDescription()) { Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. }
    }
}
```

* Elements can be swapped out where appropriate.
* Functions that return attributes can be tailored as required.

### Usage

Copy stuff and see the examples for usage.

### Development

install air for reloading:
```bash
go install github.com/air-verse/air@latest
```

and templ for generating the html:
```bash
go install github.com/a-h/templ/cmd/templ@latest
```

then just run:
```bash
air
```

### Links

* [Golang](https://go.dev)
* [Templ](https://templ.guide)
* [Air](https://github.com/air-verse/air)
* [Icônes](https://icones.js.org/collection/lucide)
* [Tailwind](https://tailwindcss.com)
* [Shadcn](https://ui.shadcn.com/docs)
* [Htmx](https://htmx.org/)