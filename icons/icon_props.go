package icons

import "github.com/a-h/templ"

type IconProps struct {
	ID               string // specifies the <svg> id attribute.
	Class            string // add classes to the <svg>.
	Style            string // specifies the <svg> style attribute.
	templ.Attributes        // add additional attributes to the <svg>.
}
