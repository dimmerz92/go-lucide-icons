# Go Lucide Icons

A wrapper of Lucide Icons for Go (Golang) developers.

## Requirements

- None if using HTML templates
- [Templ](https://templ.guide/) if using Templ templates
```sh
go install github.com/a-h/templ/cmd/templ@latest
```

## Installation & Usage

There are two ways that `go-lucide-icons` can be used.
- As a CLI tool to generate icon templates in html/template or templ format
- A library imported in to your project for templ templates or html templates

Browse [Lucide Icons](https://lucide.dev/icons/) for a comprehensive view of all the icons available.

### The command line way

1. Install the command line utility:
```sh
go install github.com/dimmerz92/go-lucide-icons/cmd/golucide@latest
```

2. Generate icons to your project directly:
**For HTML icons**
```sh
golucide html <icon name in kebab case> [-out output-directory]
```
**For Templ icons**
```sh
golucide templ <icon name in kebab case> [-out output-directory]
```

In both cases, the `-out` flag in step 2 is optional. By default, icons will be generated to the current working directory.

### The package way

#### The HTML way

Add the package to your project.
```sh
go get github.com/dimmerz92/go-lucide-icons/html-icons
```

All componenents have the same name as the lucide icons and remain in kebab case.

Simply add the lucide icon templates to your existing templates and you're ready.
```go
tpls := template.Must(template.ParseFiles("your existing files"))
err := icons.AddLucideIcons(tpls)
```

If you are using the icon as is, there is no need to pass any data to it. However, if you wish to add any attributes to the icon, it expects data in the form of `[]template.HTMLAttr`.

```go
iconData := []template.HTMLAttr{
    `class="some-class another-class"`,
    `style="height: 2rem; width: 2rem"`,
}
err := tpls.ExecuteTemplate(w, "a-arrow-down", iconData)
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
</html>
```

#### For Templ icons

Add the package to your project.
```sh
go get github.com/dimmerz92/go-lucide-icons/templicons
```

All components have the same name as the lucide icons, except in pascal case.

E.g. a-arrow-down becomes AArrowDown.

To pass attributes to the icon, simply use the [`templ.Attributes` type](https://templ.guide/syntax-and-usage/attributes/#spread-attributes).

## License

Lucide is licensed under the ISC [LICENSE](https://lucide.dev/license).

Go Lucide Icons is provided using the MIT [LICENSE](/LICENSE).

All Icons are created by [Lucide Icons](https://github.com/lucide-icons/lucide).
