## A basic tool written in Go to check website status 

### Features
* Built using [urfave/cli/v2](https://github.com/urfave/cli) for flag parsing.
* Fast TCP connectivity checks with a 5-second timeout.
* Displays local and remote network addresses on successful connections.
  
  ---
### Installation & Usage Guide
```
go install https://github.com/AY88o/healthChecker-go.git
```

**Display Help Menu**
```
./healthchecker --help
```
Output:
```
NAME:
   healthChecker - a simple tool to check status of websites

USAGE:
   healthChecker [global options] command [command options] [arguments...]

GLOBAL OPTIONS:
   --domain value, -d value  domain name to check
   --port value, -p value    Port number to check (default: "80")
   --help, -h                show help
```


