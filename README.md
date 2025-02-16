# Go Lucide Icons

A wrapper of Lucide Icons for Go (Golang) developers.

## Requirements

* None if using HTML templates
* [Templ](https://templ.guide/) if using Templ templates
```bash
go install github.com/a-h/templ/cmd/templ@latest
```

## Installation & Usage

There are two ways that `go-lucide-icons` can be used. That is, as a command line tool to generate only the icons you need, or imported as a module in to your project for easy access to all icons.

Browse [Lucide Icons](https://lucide.dev/icons/) for a comprehensive view of all the icons available.

### The command line way

1. Install the command line utility:
```bash
go install github.com/dimmerz92/go-lucide-icons/cmd/golucide@latest
```

2. Generate icons to your project directly:
* For Templ icons:
```bash
golucide add [-o <output directory>] templ <icon-name (kebab-case)>
```
* For HTML icons:
```bash
golucide add [-o <output directory>] html <icon-name (kebab-case)>
```

3. Add the `IconProps` struct to your project to be used with the generated icons.

```go
type IconProps struct {
    ID               string // specifies the <svg> id attribute.
    Class            string // add classes to the <svg>.
    Style            string // specifies the <svg> style attribute.
    templ.Attributes        // add additional attributes to the <svg>.
}
```

In both cases, the `-o` flag in step 2 is optional. By default, icons will be generated to the current working directory.

### The package way

#### For Templ icons

Add the package to your project.
* For Templ icons `go get github.com/dimmerz92/go-lucide-icons/pkg/templ/icons`

All components have the same name as the lucide icons, except in pascal case.

E.g. a-arrow-down becomes AArrowDown.

A simple properties struct is made available where `id`, `class`, and `style` attributes can be set directly, and any other attribute can be set using the [`templ.Attributes` property](https://templ.guide/syntax-and-usage/attributes/#spread-attributes).

NOTE: If you are using the package way, you do not need to add this struct, it can be found at `icons.IconProps{}`.

```go
type IconProps struct {
	ID               string // specifies the <svg> id attribute.
	Class            string // add classes to the <svg>.
	Style            string // specifies the <svg> style attribute.
	templ.Attributes        // add additional attributes to the <svg>.
}
```

**With no properties**
```templ
@icons.AArrowDown(icons.IconProps{})
```

**With properties**
```templ
@icons.AArrowDown(icons.IconProps{
    ID: "my-icon",
    Class: "my-class another-class",
    Style: "--my-variable: red",
    Attributes: templ.Attributes{"attribute": "value"},
})
```

#### For HTML icons

Add the package to your project.
* For HTML icons
```bash
go get github.com/dimmerz92/go-lucide-icons/pkg/html/icons
```

All componenents have the same name as the lucide icons and remain in kebab case.

Simply add the lucide icon templates to your existing templates and you're ready.

```go
tpls := template.Must(template.ParseFiles("your existing files"))
icons.AddLucideIcons(tpls)
```

If you are using the icon as is, there is no need to pass any data to it. However, if you wish to add any attributes to the icon, it expects data in the form of `[]template.HTMLAttr`.

```go
iconData := []template.HTMLAttr{
    `class="some-class another-class"`,
    `style="height: 2rem; width: 2rem"`,
}
```

To use the icons within your own templates, simply add them as named templates.

```html
<!DOCTYPE html>
<html>
    <head>
	<title>My Page</title>
    </head>
    <body>
	<p>some text</p>
	<!-- expecting data -->
	{{ template "worm" . }}
	<!-- not expecting data -->
	{{ template "fish" }}
    </body>
    </body>
</html>
```

## License

Lucide is licensed under the ISC [LICENSE](https://lucide.dev/license).

Go Lucide Icons is provided using the MIT [LICENSE](/LICENSE).

All Icons are created by [Lucide Icons](https://github.com/lucide-icons/lucide).
