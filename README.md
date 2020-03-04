This is a basic implementation of key-value database that written in Go language. Basically stores your data with a key in memory.
The database runs in your memory so if you terminate the process your data may be fly. But if you specify a file path with `--file` argument you'll get the last snapshot your database to file after termination. Also you can put your data back from a file with using `-retro` argument. The back-up file should be valid JSON string that order in binary. You can change the process arguments by your wish. There are nine arguments to modify database system. You check them with run `go run main.go --help`

The project lies under understanding key-value databases. It's not a production-ready code or not an alternative project to Redis or Memcache. 

It's a good kickstart project to understand the following
* Use network-level communication. Marshal/Unmarshal data which comes from TCP
* Service to multiples clients at the same time
* Mutex locks
* Parse the received flags
* IO Basics
* Implementation of protocol

There are eight commands to operate the database to do `CRUD`.

## SET
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
This commands in binary creates a key that named `foo` and a value that content with `some string value`

if you'll pass an argument that named `ttl`, you define a TTL property to the key that defined. Your value will be expire in seconds that defined in `ttl` value. You can change that wish with using `EXPIRE` as a command. Refer: EXPIRE
The database only holds string values. If you use `DECR` or `INCR` commands the database will try to format your value in `float` or `int` before executing your command. 

Also `SET` command calculates the database size before the add your request. If your key-value pair will make overflow the database size, your request will be deferred. The database has a max size or it can be absolut database that makes it unlimited. You can change it with `-max-mem-size` argument. The argument changes the database maximum size in bytes, zero value means not limits.

## GET
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

## Delete
Deletes the value that related with received key. Always return `OK` response.

```json
{
    "command": "DELETE",
    "args": {
        "key": "foo"
    }
}
```

## INCR
The method increases once the value with related with received key. The holding value could be parsable to `float64` or `int` otherwise throws parse error.

```json
{
    "command": "INCR",
    "args": {
        "key": "x"
    }
}
```

## DECR
Works as the same as `INCR` method. But decreases once the value.

```json
{
    "command": "DECR",
    "args": {
        "key": "x"
    }
}
```

## EXPIRE
Changes or sets TTL value the recived key. The `expire` defines the expiration in seconds.
```json
{
    "command": "EXPIRE",
    "args": {
        "key": "foo",
        "ttl": "60"
    }
}
```
## TTL
TTL command responds to the remaining seconds to expire key that received.

```json
{
    "command": "TTL",
    "args": {
        "key": "foo"
    }
}
```

## PING
Just a basic implementation that to test connection. Reflects `PONG` data.

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
  -flush-to-disk
        Fluhes whole data into disk when database is full then deletes the data
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
  -retro
        Retrieves data that storing in disk into memory (default true)
```