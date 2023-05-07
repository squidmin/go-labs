# go-labs

Made with:
- **GoLand 2022.3.2**
- **go1.19.5**

### Module initialization

Run the following steps from the project root directory: 

<details>
<summary>Make a module directory</summary>

```shell
mkdir project_name
```

</details>


<details>
<summary>Initialize the Go module and give it a module path</summary>

```shell
go mod init module_path
```

The `go mod init` command creates a `go.mod` file to track your code's dependencies.
So far, the file includes only the name of your module and the Go version your code supports.
But as you add dependencies, the `go.mod` file will list the versions your code depends on.
This keeps builds reproducible and gives you direct control over which module versions to use.

If you publish a module, the module path must be a path from which your module can be downloaded by Go tools.

That would be your code's repository, e.g.: `github.com/squidmin/go-labs/example`

</details>


<details>
<summary>go mod tidy</summary>

```shell
go mod tidy
```

More info: https://go.dev/ref/mod#go-mod-tidy

</details>


---


### Build & install an executable

<details>
<summary>Expand</summary>

Run the following steps from the project root directory:

```shell
go build ./directory_name
```

```shell
export PATH=$PATH:$(echo $(go list -f '{{.Target}}'))
```

```shell
cd ./directory_name
go install
```

</details>


---


### Run an executable

<details>
<summary>Expand</summary>

```shell
executable_name -o --option -k=val --key=value
```

</details>


### Run a script without building

<details>
<summary>Expand</summary>

```shell
go run script_name.go
```

</details>


---


### go.dev "Getting started" tutorial

https://go.dev/doc/tutorial/getting-started
