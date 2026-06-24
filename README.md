<h1 align="center">
  ReflectParams
  <br>
</h1>

<h4 align="center">Check if the param reflects in the body of the response!</h4>

<h6 align="center"> Coded with 💙 by ph0r3nsic </h6>

<p align="center">

<br>
  <!--Tweet button-->
  <a href="https://twitter.com/intent/tweet?text=reflectparams%20-%20Check%20if%20the%20param%20reflects%20in%20the%20body%20of%20the%20response!%20https%3A%2F%2Fgithub.com%2Fphor3nsic%2Freflectparams%20%23bash%20%23xss%20%23bugbounty%20%23bugbountytips%20%23infosec" target="_blank">Share on Twitter!
  </a>
</p>

<p align="center">
  <a href="#install-">Install</a> •
  <a href="#examples-">Examples</a> •
  <a href="#contributing-">Contributing</a> •
  <a href="#license-">License</a>
</p>

Install 📡
----------

### Installation

```console
go install github.com/phor3nsic/reflectparams@latest
```

Examples 💡
----------

### Help
```
➜  ~ reflectparams -h
Usage of reflectparams:
  -c int
    	Concorrency value. (default 5)
  -i string
    	Inject special chars on parameter. Ex: ['"><s>s]
  -l string
    	List of urls. Ex: /tmp/urls.txt
  -p string
    	Param to check. Ex: FUZZ
  -x string
    	Replay proxy requests. Ex: http://localhost:8080

```

Check with the -l option
```
➜ reflectparams -l /tmp/urls -p FUZZ
```

Or check from stdin

```
➜ cat /tmp/urls | reflectparams -p FUZZ
```

### Inject in reflection and replay to proxy:

You can inject and special chars and replay in your proxy:
```bash
➜ reflectparams -l urls.txt -i "'\"><svg/onload=alert(1)>" -x http://localhost:8080

➜ echo 'https://public-firing-range.appspot.com/reflected/parameter/body?q=FUZZ' | reflectparams -p FUZZ -i "'\"><svg/onload=alert(1)>" -x http://localhost:8080
```

### Hint:

```bash
echo public-firing-range.appspot.com | gau -subs | grep = | qsreplace FUZZ| reflectparams -p FUZZ
```

Contributing 🛠
-------

Just open an [issue](https://github.com/phor3nsic/reflectparams/issues) / [pull request](https://github.com/phor3nsic/reflectparams/pulls).

License 📝
-------

This repository is under [MIT License](https://github.com/phor3nsic/reflectparams/blob/master/LICENSE).  
[ph0r3nsic@wearehackerone.com](mailto:ph0r3nsic@wearehackerone.com) to contact me.

