# CONTRIBUTING

## Programs / Dependencies

1. Go 1.18
2. Make 4.2.1

## Development

1. Fork the repo.
2. Clone it into your local machine.
3. Install dependencies:

```bash
go get ./...
```

4. Make sure to have a good extension for handling Go in your Text Editor, in case of VSCode: https://marketplace.visualstudio.com/items?itemName=golang.Go.
5. (Optional) In case of not counting on a Go Editor Extension, make sure to run:

```bash
go fmt
```

Before any commit you make to keep consistency along the codebase.

6. Create your `structs` with your implementations from your assigned `interfaces` and return your instances through the corresponding factories.
7. Run:

```bash
go test ./...
```

And make sure that your assigned tests are running correctly, otherwise, apply necessary corrections.

8. Commit your changes to the `master` branch.
9. Create new **Pull Request** from your `master` branch to the main repo's `master` branch and await for Code Reviews.

## Accessing repo's documentation

You can access the Go's repository documentation by running the following command:

```bash
godoc -http=:6060
```

And then following: http://localhost:6060/pkg/github.com/edgaralexanderfr/cyber-kill-server/