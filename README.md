
## How I started up this project from the terminal

```

mkdir process-stdin
cd process-stdin
go mod init github.com/andrewrobinson/process-stdin
touch main.go
mkdir util
touch util/util.go
touch README.md
git init


```

## Then I opened up the process-stdin folder in vscode, in its own window

File -> new window
File -> open -> selec directory

## Then I started typing into the files in vscode

This is your simplest program, and how to run it

https://gobyexample.com/hello-world

the executable produced via go build is more suitable for piping stdin into it

Then I found this:

https://yourbasic.org/golang/read-file-line-by-line/

## to build and run

```

go build
cat testfile.txt | ./process-stdin

or 

cat testfile.txt | ./process-stdin > output.txt


```

## if you want to extract code to a package

This is what I was fighting with in East London
Fiddle a bit and see what still works but the way I have it:

a) the directory is util and it is package util
b) my func GetWhatIShouldAddOntoTheLine starts with an uppercase so it is visible externally in main
c) you HAVE to import it in main using "github.com/andrewrobinson/process-stdin/util"
d) the github.com/andrewrobinson/process-stdin comes from

module github.com/andrewrobinson/process-stdin

in the go.mod

which came from me running 

go mod init github.com/andrewrobinson/process-stdin


(So for your project, do all of the above, but change process-stdin to your project name and use your github / some unique fake path)

## and to add it to git

create a repo on github, and it tells you:

…or push an existing repository from the command line
git remote add origin git@github.com:andrewrobinson/process-stdin.git
git branch -M main
git push -u origin main

BUT 1st init it locally as a git repo

```
git init

