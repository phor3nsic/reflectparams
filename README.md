# ReflectParams

*Check if param reflect in body of response*

## Install

```
go install github.com/phor3nsic/reflectparams@latest
```
## Usage

Help
```
➜ reflectparams -h

  -i string
      Inject special chars on parameter...
  -l string
      List Of Urls
  -p string
      Param to check
  -x string
      Replay proxy requests
```

Chek of -l option
```
➜ reflectparams -l /tmp/urls -p FUZZ
```

Or check of stdin

```
➜ cat /tmp/urls | reflectparams -p FUZZ
```

## Inject in reflection and replay to proxy:

You can inject and special chars and replay in your proxy:
```bash
➜ refrecparams -l urls.txt -i "<svg>" -x http://localhost:8080

➜ echo https://public-firing-range.appspot.com/reflected/parameter/body?q=FUZZ | refrecparams -l urls.txt -i "<svg>" -x http://localhost:8080
```

### Hint:

```bash
echo public-firing-range.appspot.com | gauplus -subs | grep = | qsreplace FUZZ| reflectparams -p FUZZ
```
