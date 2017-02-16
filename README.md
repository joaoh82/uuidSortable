# uuidSortable
[WIP] Package to generate a Sortable UUID based on RFC 4122 with a timestamp.

The main.go file is an example of a usage generating 1000 UUID in less then one second.

This code generates a UUID with a sortable timestamp.
```
u, err := NewIDSortable()
if err != nil {
    panic(err)
}
```

In the example there is also profiling enabled to check for memory usage.

## Maintainer

Created in 2017 by João Henrique Machado Silva joaoh82@gmail.com @joaoh82. Under the MIT License.
