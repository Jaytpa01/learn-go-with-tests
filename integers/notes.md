## Testable Examples

- <https://go.dev/blog/examples>

Example functions / testable examples are snippets of Go code that are displayed in the generated godoc. This ensures all examples are up to date, as they are actually compiled and verified by running them as tests.

They can also be run by a user visiting the godoc web page for the package and clicking the associated “Run” button.

They begin with Example (much like test functions begin with Test), and reside in a package's `_test.go` files.

i.e

```go
func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}
```

> [!NOTE]
> The `// Output` comment is actually important. It will be considered the `expected` value written to stdout.

Some examples:

- <https://pkg.go.dev/testing#hdr-Examples>
- <https://pkg.go.dev/crypto/sha256#pkg-examples>
- <https://pkg.go.dev/crypto/sha256#example-New>
