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