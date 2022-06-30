<!-- 
1) Paragraph explaining that we want to create an implementation/struct for the specified interface(s).
 -->

# Issue Template
<!--
Further explanation about the issue/task.
 -->

We want to create an implementation/`struct` for the specified interface(s).

#


## Interfaces to implement
- ``


## Methods to implement in the struct
- ``


## Factories methods to implement/return
- `` 

## Unit test command

```bash
go test ./pkg/my-package/ -v
```

You can make use of the _test/main.go_ file and run it to include your implementations and do your own local tests. Commit your code to the `master` branch of your forked repo or branch from the main repo and create **Pull Request** to the `master` branch.

**NOTE:** For no reason you are _ever_ allowed to change the structure or do changes in the interfaces, these are carefully designed to interact _that_ way with other components, keeping them as cohesive as possible. Same with unit tests.

If you feel like some interface shouldn't have certain structure, you're free to open a _discussion_ and judgement about it in the **Slack channel** or thru **DM**.

Enjoy! :feelsgood: