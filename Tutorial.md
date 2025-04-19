# Tutorial
A comprehensive guide on how to use cs-client.

## Setup
First we setup a basic blockchain, let's call it `github.com/a/niceChain`:
```shell
ignite scaffold chain github.com/a/niceChain
```

After that we also want to add some life to our blockchain by adding a `testItem` CRUD data store:
```shell
cd ./niceChain
ignite scaffold list testItem -y
```

Then we want to build and init our chain:
```shell
ignite chain init
```

We also want to make sure the cs-client app is installed:
```shell
ignite app install github.com/lxgr-linux/ignite-app-cs-client
```

## Generating
Now we want to generate our C# client by running:
```shell
ignite generate cs-client -y -o ./cs-client
```

This creates a `cs-client` folder with the following structure:
```
├── LICENSE
├── Nicechain
│   ├── Nicechain
│   │   ├── Genesis.pb.cs
│   │   ├── Module
│   │   │   └── Module.pb.cs
│   │   ├── Params.pb.cs
│   │   ├── QueryGrpc.pb.cs
│   │   ├── Query.pb.cs
│   │   ├── TestItem.pb.cs
│   │   ├── TxClient.pb.cs
│   │   ├── TxGrpc.pb.cs
│   │   └── Tx.pb.cs
│   ├── QueryClient.cs
│   └── TxClient.cs
└── Nicechain.csproj
```
The most important files here are `QueryClient.cs` and `TxClient.cs`

Let's go in there and look if it builds:
```shell
cd cs-client
dotnet build
```

## Usage
For usage of that generated code take a look into [`integration/static/testApp`](integration/static/testApp).
To try out the `testApp` with you newly generated code change the
```xml
<ProjectReference Include="../outpath/Nicechain.csproj"/>
```
to your C# client output path an run
```shell
dotnet run
```
