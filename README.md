# go-labs

A comprehensive repository for learning and experimenting with Go, set up and tested using:

- **IDE**: GoLand 2022.3.2
- **Go Version**: go1.19.5

## Getting Started

Below are the initial setup steps to prepare your environment and start experimenting with Go modules.
Ensure your terminal is pointed to the root folder of this project before proceeding.

### Module initialization

#### Creating a Module Directory

<details>
<summary>Click to expand</summary>

#### Use Case

Creating a module directory is your first step in organizing your Go project.
Modules in Go serve as containers for packages, and having a dedicated directory for each module helps in managing dependencies, versioning, and package distribution efficiently.
This step is crucial when you're starting a new project or adding a new module to an existing project to ensure that your codebase remains organized and scalable.

```shell
mkdir module_name
```

</details>

#### Initializing a Go Module

<details>
<summary>Click to expand</summary>

#### Use Case

Initializing a Go module with `go mod init` sets the foundation for managing your project's dependencies.
This command creates a `go.mod` file, marking the current directory as the root of a module.
It enables Go to track and manage the versions of external packages your project depends on, facilitating reproducible builds and simplifying dependency management.
This step is essential at the beginning of project development or when adding new dependencies.

```shell
go mod init module_path
```

The `go mod init` command creates a `go.mod` file to track your code's dependencies.
Initially, this file will only include your module's name and the Go version your code supports.
As you add dependencies, `go.mod` will list the versions your code depends on, ensuring reproducible builds and direct control over module versions.

>**Note**: For published modules, the module path must be a downloadable path by Go tools, typically your code's repository URL, e.g., `github.com/squidmin/go-labs/example`.

</details>

#### Cleaning Up Dependencies

<details>
<summary>Click to expand</summary>

### Use Case

Running `go mod tidy` ensures that the `go.mod` and `go.sum` files accurately reflect the dependencies actually used in your project.
This command adds missing dependencies, removes unused ones, and updates `go.mod` and `go.sum` to match the source code.
It's particularly useful after adding or removing imports in your code or when preparing to commit your code to version control, ensuring a clean, up-to-date record of dependencies.

```shell
go mod tidy
```

Use the `go mod tidy` command to clean up your `go.mod` and `go.sum` files.
This step adds any missing module dependencies and removes unused ones, ensuring that your module dependencies are accurate and up-to-date.

More info: [`go mod tidy`](https://go.dev/ref/mod#go-mod-tidy)

</details>

---

## Building and Running Your Code

### Building and Installing an Executable

<details>
<summary>Click to expand</summary>

#### Use Case

Compiling your Go code into an executable with `go build` and then installing it with `go install` makes it easy to distribute and run your application.
This process is crucial for creating standalone applications that can be executed without the Go runtime.
It's particularly relevant when you're ready to deploy your application or share it with users who may not have Go installed.

#### Steps

1. **Navigate to Your Project Directory**: Ensure you're in the directory containing your `main` package.

   ```shell
   cd ./directory_name
   ```
   
2. **Build Your Code (Optional)**: This step compiles your code and generates an executable in the current directory.
   It's useful for testing the build process or when you need an executable in a specific location.

   ```shell
   go build
   ```

   > **Note**: You can also build code from outside the module by specifying a path:
   >
   > ```shell
   > go build ./directory_name
   > ```
   
3. **Install Your Application**:

   ```shell
   cd ./directory_name # Navigate to the module
   go install
   ```

4. **Adding the Executable to Your `PATH` (Optional)**:

   After building or installing your Go application, you might want to make the executable globally accessible from the command line, regardless of your current working directory.
   This is particularly useful for tools or applications you plan to use frequently across various projects or directories.

   To add the executable to your `PATH`, you can use the `export` command in Unix-like operating systems, including Linux and macOS.
   This command temporarily modifies the `PATH` environment variable for the current terminal session.
   To make this change permanent, you'll need to add the export command to your shell's startup file, such as `.bashrc`, `.bash_profile`, or `.zshrc`, depending on your shell and operating system.

   The following command uses `go list -f '{{.Target}}'` to dynamically find the installation path of your Go executable and adds it to your `PATH`:

   ```shell
   export PATH=$PATH:$(go list -f '{{.Target}}')
   ```
   
   > **Note**: This step assumes you've used `go install` to compile and install your executable.
   > `go install` places binaries in the `$GOPATH/bin` directory (or `$GOBIN` if set), which is the path resolved by `go list -f '{{.Target}}'`.
   > If you're managing your Go environment correctly, `$GOPATH/bin` should already be in your `PATH`.
   > However, if you find it's not, or if you've compiled your executable with `go build` and placed it in a custom directory, this command can be adapted to include that directory in your `PATH`.

   > **Important**: Remember, changes made with the `export` command to the `PATH` are temporary and only affect the current terminal session.
   > For a permanent solution, you need to add the command to your shell's startup file as mentioned earlier.

</details>

### Running an Executable

<details>
<summary>Click to expand</summary>

To run an installed executable with options:

```shell
./executable_name -o --option -k=val --key=value
```

Or if the executable has been added to your `PATH`:

```shell
executable_name -o --option -k=val --key=value
```

</details>

### Build or Run a Specific File

<details>
<summary>Click to expand</summary>

Use the `-tags` flag with `go build` or `go run`:

```shell
go build -tags filename .
```

e.g.,

```shell
go build -tags arrays ./04_arrays_and_slices/01_declaration_and_initialization/arrays
```

```shell
go run -tags filename .
```

e.g.,

```shell
go run -tags arrays ./04_arrays_and_slices/01_declaration_and_initialization/arrays
```

</details>

### Running a Script Without Building

<details>
<summary>Click to expand</summary>

#### Use Case

The ability to run a Go script using `go run` without explicitly compiling it first is invaluable during development.
It allows for rapid testing and iteration, letting developers focus on writing and refining code without worrying about the build process.
This approach is ideal for development environments where quick feedback on code changes is crucial.

```shell
go run script_name.go
```

Understanding the context and purpose behind these actions can significantly enhance your learning experience and effectiveness in Go development.

</details>

---

## Additional Resources

For a more in-depth introduction to Go, consider the official ["Getting Started" tutorial](https://go.dev/doc/tutorial/getting-started).
