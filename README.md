# Go Templ Lucide

A Templ wrapper of Lucide Icons for Go developers.

## Requirements

* [Templ](https://templ.guide/) `go install github.com/a-h/templ/cmd/templ@latest`

## Installation

`go get github.com/dimmerz92/go-templ-lucide/icons`

## Usage

Browse [Lucide Icons](https://lucide.dev/icons/) for a comprehensive view of all the icons available.

All components have the same name as the lucid icons, except in pascal case.

E.g. a-arrow-down becomes AArrowDown

All components accept an optional variadic argument of strings which are all given to the svg class attribute.

Passing an argument is not mandatory.

**No Classes**
```templ
@icons.AArrowDown()
```

**With Classes (Single Arg)**
```templ
@icons.AArrowDown("h-4 w-4 text-red-600")
```

**With Classes (Multi Arg)**
```templ
@icons.AArrowDown(
    "h-4",
    "w-4",
    "text-red-600",
)
```

### Not building with templ?

Templ will still remain a dependency, however, you can still use these icons with Go's html/template package.

With this method, you can inject the icon as HTML into your templates from your program.

```templ
icon, err := templ.ToGoHTML(context.Background(), icons.AArrowDown("h-4 w-4 text-red-600"))
if err != nil {
    log.Printf("Failed to render Templ to Go HTML: %v", err)
}

if err = existingTemplates.Execute(w, icon); err != nil {
    log.Printf("Failed to execute template: %v", err)
}
```

See the [Templ docs](https://templ.guide/syntax-and-usage/using-with-go-templates#using-a-templ-component-withhtmltemplate).

## License

Go Templ Lucide is provided using the MIT License.

All Icons are created by [Lucide Icons](https://github.com/lucide-icons/lucide).
