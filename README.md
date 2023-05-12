![License](https://img.shields.io/badge/license-sushiware-red)
![Issues open](https://img.shields.io/github/issues/crashbrz/couchdumper)
![GitHub pull requests](https://img.shields.io/github/issues-pr-raw/crashbrz/couchdumper)
![GitHub closed issues](https://img.shields.io/github/issues-closed-raw/crashbrz/couchdumper)
![GitHub last commit](https://img.shields.io/github/last-commit/crashbrz/couchdumper)

# CouchDumper 
CouchDumper is a tool designed to leverage the capabilities of CouchDb instances without authentication. With CouchDumper, you can effortlessly retrieve comprehensive database dumps from vulnerable CouchDb installations.

CouchDumper takes advantage of the lack of authentication in CouchDb instances. By utilizing it, you can effortlessly initiate the dumping process and retrieve the entire contents of the target database.

Usage example:
```
┌──(crash㉿Anubis)-[~]
└─$ go run ./couchdumper.go -u http://10.10.10.10 -p 6000 | jq

 ```
- The above command will try to retrieve all the content from all databases
- The output is default set to integrate with te jq command. For verbose output set: -j=false 

### Usage/Help ###
Run go run ./couchdumper.go -h to see all options. Also, you can contact me (@crashbrz) on Twitter<br>

### Installation ###
Clone the repository in the desired location.<br>
Install the jq command for better output view.<br>

### License ###
CouchDumper is licensed under the SushiWare license. Check [docs/license.txt](docs/license.txt) for more information.
 
### Go Version ###
Tested on:<br>
Go 1.19.8
