/*
	Package enry implements multiple strategies for programming language identification.

	Identification is made based on file name and file content using a service
	of strategies to narrow down possible option.
	Each strategy is available as a separate API call, as well as a main enty point

		GetLanguage(filename string, content []byte) (language string)

	It is a port of the https://github.com/github/linguist from Ruby.
	Upstream Linguist YAML files are used to generate datastructures for data
	package.
*/
package enry // import "github.com/go-enry/go-enry/v2"

//go:generate make code-generate

import "github.com/go-enry/go-enry/v2/data"

// Type represent language's type. Either data, programming, markup, prose, or unknown.
type Type int

// Type's values.
const (
	Unknown     Type = Type(data.TypeUnknown)
	Data             = Type(data.TypeData)
	Programming      = Type(data.TypeProgramming)
	Markup           = Type(data.TypeMarkup)
	Prose            = Type(data.TypeProse)
)
