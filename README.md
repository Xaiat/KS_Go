# KS_Go 狂神说Go语言
KS_Go is a project for notes on learning the Go language, inspired by KuangStudy ('Kuangshen Talks').
KS_Go是学习狂学（狂神说）Go语言的笔记项目。
## 1.KS_Go 项目子目录
## 2.KS_Go 笔记文件夹：KS_Go_MD

# 0.Go mod init 模块初始化
分2个步骤：
1.在项目子目录内模块初始化：go mod init XXX
2.在项目根目录工作区初始化：go work init && go work use ./项目子目录

# ==== Go 项目多模块初始化 ====
## ==== 常用精简指令 ====
```sh
$ mkdir -p 1.XXXX && cd 1.XXXX
$ go mod init XXXX
$ cd .. && go work use ./1.XXXX
```

## ==== Go 项目需求 ====
在项目根目录 KS_Go 操作初始化
项目根目录：/Users/rolex/CodeLangs/Golang/KS_Go
项目子目录1：/Users/rolex/CodeLangs/Golang/KS_Go/1.Go_Hello_World
项目子目录2：/Users/rolex/CodeLangs/Golang/KS_Go/2.RAM_Printf
项目子目录3：/Users/rolex/CodeLangs/Golang/KS_Go/3.XXXXXXX
## 1. ==== 初始化每个子目录为独立的模块 ====
在每个子目录中执行 go mod init：
```sh
$ cd /Users/rolex/CodeLangs/Golang/KS_Go/1.Go_Hello_World
$ go mod init 1.Go_Hello_World

$ cd /Users/rolex/CodeLangs/Golang/KS_Go/2.RAM_Printf
$ go mod init 2.RAM_Printf

$ cd /Users/rolex/CodeLangs/Golang/KS_Go/3.XXXXXXX
$ go mod init 3.XXXXXXX
```
## 2. ==== 创建 go.work 文件 ====
```sh
$ cd /Users/rolex/CodeLangs/Golang/KS_Go
$ go work init
$ vim go.work
```
==== Begin ====
go 1.21.6

use (
    ./1.Go_Hello_World
    ./2.RAM_Printf
    ./3.XXXXXXX
)
==== End ====
## 3. ==== 新增项目子目录 ====
### 在子项目子目录内go mod init 模块初始化
```sh
$ cd /Users/rolex/CodeLangs/Golang/KS_Go/4.AAAAAAA/
$ go mod init AAAAAAA
```
### 在子项目目录内go work use 工作区添加新增项目子目录
```sh
$ cd /Users/rolex/CodeLangs/Golang/KS_Go/
$ go work use ./4.AAAAAAA
```

# ==== GoModule Go 模块是什么 ====
Go 模块是 Go 1.11 版本引入的一种依赖管理系统，它提供了一种方便的方式来处理 Go 代码的依赖问题。模块是相关 Go 包的集合，它们一起被版本化和共享。以下是 Go 模块的一些关键特性和使用方式的详细介绍：

### 关键特性

1. **版本控制**：模块支持版本控制，这意味着您可以指定依赖特定版本的外部模块。这有助于确保项目的一致性和可重复构建。

2. **去中心化**：模块可以托管在任何地方，不再强制要求放在特定的源代码仓库中，如 GitHub、GitLab 等。

3. **易于管理**：每个模块包含一个 `go.mod` 文件，其中列出了模块的依赖项及其版本。这使得依赖管理变得更加明确和简单。

4. **不依赖于 GOPATH**：模块化的项目可以放在文件系统的任何地方，不再需要放在 `GOPATH` 目录下。

### 使用 Go 模块的步骤

1. **初始化模块**：
   - 在项目根目录下运行 `go mod init [module-name]` 来创建一个新的模块。这里的 `[module-name]` 是您为模块选择的名字，通常是代码仓库的路径，比如 `github.com/username/projectname`。

2. **添加依赖项**：
   - 当您导入新的包并运行 `go build` 或 `go test` 时，Go 会自动添加所需的依赖到 `go.mod` 文件。
   - 您也可以使用 `go get [package]` 显式添加依赖。

3. **升级和降级依赖**：
   - 使用 `go get [package]@[version]` 来升级或降级到特定版本的依赖。

4. **整理依赖**：
   - 运行 `go mod tidy` 来移除不再需要的依赖，并添加间接依赖到 `go.mod` 文件。

5. **查看依赖信息**：
   - 使用 `go list -m all` 查看当前模块的所有依赖。
   - 使用 `go mod graph` 查看依赖图。

6. **构建和测试**：
   - 使用 `go build` 和 `go test` 像往常一样构建和测试您的代码。Go 将自动使用 `go.mod` 中指定的依赖版本。

7. **版本控制**：
   - 将 `go.mod` 和 `go.sum` 文件加入到您的版本控制系统（如 Git）。`go.sum` 文件包含每个依赖的预期加密哈希，用于确保依赖的完整性。

### 总结

Go 模块提供了一种强大且灵活的方式来管理 Go 项目的依赖，使项目更加容易维护、共享和协作。它标志着 Go 语言在包管理方面的重大进步。

# ==== 问题1 ====
如果我在/Users/rolex/CodeLangs/Golang路径下分别新建了3目录A 、B 和 C，分别放置不同的Go程序项目，那么是否我需要在这ABC三个路径下分别做一次go mod init ?

# ==== 回答1 ====
是的，您需要在 `/Users/rolex/CodeLangs/Golang` 路径下的每个项目目录（A、B 和 C）中分别执行一次 `go mod init` 命令。这样做是为了确保每个项目都有自己的 Go 模块配置，从而可以独立管理各自的依赖。

对于每个项目，您可以按照以下步骤进行：

1. **进入每个项目的根目录**：
   - 对于项目 A，运行 `cd /Users/rolex/CodeLangs/Golang/A`。
   - 对于项目 B，运行 `cd /Users/rolex/CodeLangs/Golang/B`。
   - 对于项目 C，运行 `cd /Users/rolex/CodeLangs/Golang/C`。

2. **初始化 Go 模块**：
   - 在每个项目的根目录下运行 `go mod init [module-name]`。
   - `[module-name]` 应该是每个项目的唯一标识，通常是代码仓库的 URL。如果您尚未决定或不打算将项目放在远程仓库，您可以使用任何有意义的名称，例如 `go mod init A`、`go mod init B` 和 `go mod init C`。

这样，每个项目都将拥有自己的 `go.mod` 文件，您可以在这些文件中定义各自项目的依赖关系。这种做法保证了每个项目依赖的独立性和项目配置的清晰度，使得项目更容易维护和管理。

# ==== 问题2 ====
如果我设置go mod init KS_Go了以后，那么后面怎么删除这个设置呢？

# ==== 回答2 ====
如果您已经使用 `go mod init KS_Go` 初始化了一个 Go 模块，但之后想要删除这个设置，基本上您需要做的就是删除生成的 `go.mod` 文件（以及通常与之一起创建的 `go.sum` 文件，如果存在的话）。这些文件位于您的项目的根目录中。

以下是删除 Go 模块设置的步骤：

1. **定位到项目根目录**：
   - 使用命令行定位到包含 `go.mod` 文件的目录。例如：
     ```bash
     cd /path/to/your/project
     ```

2. **删除 `go.mod` 和 `go.sum` 文件**：
   - 使用文件系统的命令或文件浏览器删除 `go.mod` 和 `go.sum` 文件。在命令行中，您可以使用：
     ```bash
     rm go.mod go.sum
     ```
   - 这个操作会将您的项目从模块系统中移除。

请注意，在删除这些文件之后，您的项目将不再作为一个 Go 模块来处理，这意味着之前通过模块管理的依赖和版本信息将会丢失。如果您的项目位于 `GOPATH` 外，且不是作为一个模块来管理，那么在 Go 1.11 及以上版本中，您可能需要手动管理依赖。

如果您的项目是在版本控制下（如 Git），并且您想要从项目历史中彻底移除模块设置，您还需要提交这些删除的变更。例如，使用 Git 的话，您可以执行：

```bash
git rm go.mod go.sum
git commit -m "Remove Go module files"
```

最后，如果您计划将来重新将项目转换为 Go 模块，可以随时再次运行 `go mod init` 命令来创建新的 `go.mod` 文件。


# ==== 问题3 ====
go mod init 是不是有点类似git init的原理？

# ==== 回答3 ====
是的，`go mod init` 与 `git init` 在一定程度上是相似的，尤其是在它们的目的和一些基本概念上。以下是它们之间的一些相似之处：

1. **初始化项目**：
   - `git init` 用于初始化一个新的 Git 仓库。它创建一个 `.git` 目录，这个目录包含所有的 Git 仓库元数据，如分支、标签、配置等。
   - `go mod init` 用于初始化一个新的 Go 模块。它创建一个 `go.mod` 文件，这个文件包含了模块的名称和依赖信息。

2. **为项目设定框架**：
   - 在使用 `git init` 之后，您可以开始跟踪文件的版本历史。
   - 使用 `go mod init` 后，您可以开始管理项目的依赖，并确保项目在不同环境中的一致性。

3. **项目的本地作用**：
   - `git init` 的作用局限于本地仓库，直到您决定将其推送到远程仓库。
   - `go mod init` 同样主要影响本地项目，并且不依赖于任何外部源，除非您添加了外部依赖。

尽管 `go mod init` 和 `git init` 在概念上类似，它们服务于不同的目的：`git init` 用于版本控制，而 `go mod init` 用于依赖管理。但它们都是在项目初期设置项目结构的重要工具。


# ==== 问题4 ====
在用git推送到GitHu上做项目版本管理，是否需要在.gitignore排除文件中把go.mod和go.sum两个文件列进去，防止被同步到GitHub上去呢？

# ==== 回答4 ====
不需要将 `go.mod` 和 `go.sum` 文件加入到 `.gitignore` 中以排除它们。事实上，这两个文件应该被包含在您的 Git 仓库中，并且与您的项目一起推送到 GitHub 上。这是因为 `go.mod` 和 `go.sum` 文件在 Go 项目中扮演着重要的角色：

1. **`go.mod` 文件**：这个文件包含了您的项目所依赖的模块及其版本。这是 Go 模块系统用于确保项目依赖一致性的重要部分。当其他人或者自动化系统（如持续集成/持续部署系统）检出您的代码时，`go.mod` 文件会确保他们能获取到正确版本的依赖。

2. **`go.sum` 文件**：这个文件包含了您依赖的每个模块的特定版本的加密哈希。这有助于确保这些依赖没有被意外更改或篡改，提供了一个额外的安全层。

将这两个文件包含在您的 Git 仓库中，对于确保项目的构建可重复性和依赖的完整性是非常重要的。这样，无论是谁克隆您的仓库，都能确保他们获取到与您相同的依赖，从而能够无差错地构建和运行您的项目。












