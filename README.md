# dfh
Datetime base file helper.

## Usage

```
$ ./dfh
NAME:
   dfh - Datetime base file helper. Easy to remove/find files by datetime or time duration.

USAGE:
   dfh [global options] command [command options] [arguments...]

VERSION:
   v0.1.0

COMMANDS:
     remove, rm  Remove files.
     find, f     Find files.
     help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --timetemplate value  Parse time template. (default: "2006-01-02T15:04:05Z07:00")
   --help, -h            show help
   --version, -v         print the version
```

### Find file

```
$ ./dfh find -h
NAME:
   dfh find - Find files.

USAGE:
   dfh find [command options] [arguments...]

OPTIONS:
   --duration value, -d value  How many time duration ago? eg. 24h
   --basetime value, -b value  Datetime base line.
   --recursive, -r             Use recursive.
   --include, -i               Include time range. Default exclude.
   --force, -f                 Force, without asking.
   --dir, -D                   Search directory.
   --dironly, --DO             Search directory only.
```

#### Example

* Find files's modify date in 10 hours in `Download` folder.
```
$ ./dfh find -i -d 10h ~/Downloads
```

* File files's last modify data 10 days ago in `Download` folder.
```
./dfh find -d 240h ~/Downloads
```

### Remove file

```
$ ./dfh rm -h
NAME:
   dfh remove - Remove files.

USAGE:
   dfh remove [command options] [arguments...]

OPTIONS:
   --duration value, -d value  How many time duration ago? eg. 24h
   --basetime value, -b value  Datetime base line.
   --recursive, -r             Use recursive.
   --include, -i               Include time range. Default exclude.
   --force, -f                 Force, without asking.
   --dir, -D                   Search directory.
   --dironly, --DO             Search directory only.
```

### Example

* Find files's modify date in 10 hours in `Download` folder.
```
$ ./dfh find -i -d 10h ~/Downloads
```

* Find files's last modify data 10 days ago in `Download` folder.
```
./dfh find -d 240h ~/Downloads
```

* Remove files and dictionary 10 days ago in `Download` folder.
```
./dfh rm -d 240h -D ~/Downloads
```

* Remove files and dictionary recursively 10 days ago in `Download` folder.
```
./dfh rm -d 240h -D ~/Downloads
```

* Remove dictionary only recursively 10 days ago in `Download` folder.
```
./dfh rm -d 240h -DO ~/Downloads
```
