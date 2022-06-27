<!-- 
We want to create an issue template to facilitate the assignment of common tasks for implementations to developers.

Such file must container:

1) Paragraph explaining that we want to create an implementation/struct for the specified interface(s).
2) A section with an empty list of interfaces to implement with the correct empty coding format for the line.
3) A section with an empty list of methods to implement in the struct with the correct empty coding format for the line.
4) A section with an empty list of Factories methods to fulfill with the correct empty coding format for the line.
 -->


## How to run specific unit tests for a module

```bash
go test ./pkg/util/ -v
```
##

You can make use of the test/main.go file and run it to include your implementations and do your own local tests. Commit your code to the master branch of your forked repo or branch from the main repo and create Pull Request to the master branch.