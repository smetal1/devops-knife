## DevOps-knife

This tool consist of all comonly used commands that we use during day-to-day devops work. Its designed using Cobra framework available in go-lang ecosystem.
The following is the project structure as of now:
<img src="https://ik.imagekit.io/m5gtndqap/Devops-knife/DevOps-knife.drawio.png?ik-sdk-version=javascript-1.4.3&updatedAt=1676862004871"/>

## Pre-requisite on Linux/Mac:
1. Clone the repo
2. Install go lang and set up go-path
3. Install cobra-cli

## Execution
### To run the code without building (this will display all available commands that are available in the toolkit):
```
go run main.go -h
```

### To build the executable
```
go build .
```

This will create an executable - Devops-knife

it can be used like normal executable. Refer the diagram:
<img src="https://ik.imagekit.io/m5gtndqap/Devops-knife/GO-BUILD.png?ik-sdk-version=javascript-1.4.3&updatedAt=1676862187844"/>

To use the executable:
<img src="https://ik.imagekit.io/m5gtndqap/Devops-knife/DevOps-Knife-docker.png?ik-sdk-version=javascript-1.4.3&updatedAt=1676862175900"/>

### Sample Output
Executing docker pull command using this toolkit:
<img src="https://ik.imagekit.io/m5gtndqap/Devops-knife/pull.png?ik-sdk-version=javascript-1.4.3&updatedAt=1676862771565"/>