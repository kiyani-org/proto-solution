# proto-solution
A demo app for managing protos

# Requirements

Language specific requirements are mentioned below. In general, you will need `protoc`. 

```
  brew install protobuf
```

## Go

Must have `go` installed.

```
  brew install go
```

Must have `protoc-gen-go` installed.

```
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```

## Node (Typescript)

Must have `node` installed

```
  brew install node
```

# Generate protos

Make changes to protos in the `protos` directory and run the script

```
  chmod +x ./scripts/generate.sh
  ./scripts/generate.sh
```
