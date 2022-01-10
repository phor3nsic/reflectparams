# ReflectParams

*Check if param reflect in body of response*

## Install

```
go install github.com/phor3nsic/reflectparams
```
## Usage

Help
```
reflectparams -h

  -l string
    	List Of Urls
  -p string
    	Param to check
```

Chek of -l option
```
reflectparams -l /tmp/urls -p FUZZ
```

Or check of stdin

```
cat /tmp/urls | reflectparams -p FUZZ
```

### Hint:

*Use output of paramspider with this app!*