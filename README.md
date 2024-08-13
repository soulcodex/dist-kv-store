<table align="center">
    <tr style="text-align: center;">
        <td align="center" width="9999">
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

If errors during the process

```shell
go mod tidy -compat=1.22
```