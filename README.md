<table align="center">
    <tr style="text-align: center; border-bottom: transparent">
        <td align="center" width="9999">
            <img src="./etc/raft-diagram.png" width="450px" alt="Project icon" style="margin: 10px auto; display: block">
            <h1 style="color: black; margin-top: 0">Distributed Key Value Store</h1>
            <p style="color: black">A distributed key-value store implementation using the raft consensus algorithm.</p>
        </td>
    </tr>
</table>

You need to have **Go v1.22.3** installed in your machine.
In Mac OS X is as easy as executing [Brew](https://brew.sh/): `brew install go`

You'll need also an IDE, take a look to [Goland](https://www.jetbrains.com/go/)
or [Visual Studio Code](https://code.visualstudio.com/).

[rpcbind](https://www.unix.com/man-page/osx/8/rpcbind/) is also required to run the application.
In Mac OS X is as easy as executing [Brew](https://brew.sh/): `brew install rpcbind`

1. Now is time to download modules

```shell
go mod tidy
```

*If errors during the process, try to run the following command:*

```shell
go mod tidy -compat=1.22
```

2. Prepare the environment to be able to run

```shell
make setup
```

3. To initialize a single node cluster, run the following command:

```shell
make single-node
```

4. To initialize a minimal multi-node cluster with 3 nodes, run the following command:

```shell
make master-node
make slave-node-1
make slave-node-2
```

5. To run benchmarks, run the following command:

```shell
make run-benchmark
```

6. To run acceptance tests, run the following command:

```shell
make run-acceptance-tests
```