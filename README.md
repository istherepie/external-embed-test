# External Embed Test

Testing the new embedding feature in go 1.16.
Is it possible to import an external module with embedded files ?

## Usage

Download dependecy:

```
	go get github.com/istherepie/external-embed-test

```

And import the `ui` package:

```
	import (
		"github.com/istherepie/external-embed-test/pkg/ui"
	)

	embeddedFiles := ui.StaticAssets()

```


## Examples

An example has been included showcasing how to use an embedded HTML/CSS/JS webapp.

Please see: `examples/webserver.go`

Or run it locally:

```
	# git clone https://github.com/istherepie/external-embed-test

	# cd external-embed-test

	(external-embed-test)# go run examples/webserver.go

```


## Author

Feel free to contact me @ <hello@istherepie.com>!
