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

A simple properties struct is made available where `id`, `class`, and `style` attributes can be set directly, and any other attribute can be set using the [`templ.Attributes` property](https://templ.guide/syntax-and-usage/attributes/#spread-attributes).

```go
type IconProps struct {
	ID               string // specifies the <svg> id attribute.
	Class            string // add classes to the <svg>.
	Style            string // specifies the <svg> style attribute.
	templ.Attributes        // add additional attributes to the <svg>.
}
```

**No Classes**
```templ
@icons.AArrowDown(icons.IconProps{})
```

```templ
@icons.AArrowDown(icons.IconProps{Class: "h-4 w-4 text-red-600"})
```

### Not building with templ?

Templ will still remain a dependency, however, you can still use these icons with Go's html/template package.

With this method, you can inject the icon as HTML into your templates from your program.

```templ
icon, err := templ.ToGoHTML(context.Background(), icons.AArrowDown(icons.IconProps{Class: "h-4 w-4 text-red-600"}))
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
