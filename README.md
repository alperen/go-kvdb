# Simple Key-value Database with Go
This is a basic implementation of key-value database that written in Go language.

The project lies under understanding key-value databases. It's not a production-ready code or not an alternative project to Redis or Memcache. 

It's a good kickstart project to understand the following
* Use network-level communication. Parse and send binary json over TCP.
* Service to multiples clients at the same time
* Mutex locks
* Parse the received flags
* IO Basics
* Implementation of the protocol
* Use built-in packages in Go. (The program does not require any 3th party package.)


## About Specifications
Specification of the program has sourced from [hipo/backend-challenges/kvdb](https://github.com/Hipo/backend-challenges/tree/master/kvdb)

## Clients
Clients should open a TTL connection to the port that the program runs. All communication (sending-receiving) will be in binary formatted JSON. Here is a [example client code](https://gist.github.com/alperen/84f921994f0b61f914b281f6638c7aec) that written in Go.

## Commands
There are eight commands to operate the database to do `CRUD`.

### SET
Creates a key-value pair in the database. If the key exists already in the database, the newest value will rewrite on key.

```json
{
    "command": "SET",
    "args": {
        "key": "foo",
        "value": "some string value"
    }
}
```
This commands in binary create a key that named `foo` and a value that content with `some string value`

if you'll pass an argument that named `ttl`, you define a TTL property to the key that defined. Your value will expire in seconds that defined in `ttl` value. You can change that wish with using `EXPIRE` as a command. Refer: EXPIRE
The database only holds string values. If you use `DECR` or `INCR` commands the database will try to format your value in `float` or `int` before executing your command. 

Also `SET` command calculates the database size before the add your request. If your key-value pair will make overflow the database size, your request will be deferred. The database has a max size or it can absolut database that makes it unlimited. You can change it with `-max-mem-size` argument. The argument changes the database maximum size in bytes, zero value means no-limits.

### GET
Get a value that related with received key. If the key doesn't exist in the database response error message.

```json
{
    "command": "GET",
    "args": {
        "key": "foo"
    }
}
```
Reflects the value in `result` array in the response.

### Delete
Deletes the value related with the received key. Always return `OK` response.

```json
{
    "command": "DELETE",
    "args": {
        "key": "foo"
    }
}
```

### INCR
The method increases once the value related with the received key. The holding value could be parsable to `float64` or `int` otherwise throws parse error.

```json
{
    "command": "INCR",
    "args": {
        "key": "x"
    }
}
```

### DECR
Works as the same as `INCR` method. But decreases once the value.

```json
{
    "command": "DECR",
    "args": {
        "key": "x"
    }
}
```

### EXPIRE
Changes or sets TTL value the received key. The `expire` defines the expiration in seconds.
```json
{
    "command": "EXPIRE",
    "args": {
        "key": "foo",
        "ttl": "60"
    }
}
```
### TTL
TTL command responds to the remaining seconds to expire key that received.

```json
{
    "command": "TTL",
    "args": {
        "key": "foo"
    }
}
```

### PING
Just a basic implementation to test the connection. Reflects `PONG` data.

```json
{
    "command": "PING"
}
```

```
go run main.go --help
  -detach
        Prints nothing to screen
  -file string
        Refers to database's location on the disk. Should be existed file.
  -max-mem-size int
        Sets the maximum size of database. Server does not accepts new entries while maximum size is hanging. Default 0 means no limits.
  -panics
        Shows panics.
  -persist-to-disk int
        Server stores the database in every given minutes. (default 60)
  -port string
        Sets serving port. The given port number should be free for communication (default "6379")
  -refresh-rate int
        Sets screen refresh rate in seconds. (default 1)
```