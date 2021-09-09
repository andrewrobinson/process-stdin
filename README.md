
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
File -> open -> select directory

VSCode can show you multiple folders and go projects, but for running a main.go
and making it happy with packages I find it is best to have a window per folder with a go.mod in it


## Then I started typing into the files in vscode

This is your simplest program, and how to run it

https://gobyexample.com/hello-world

Then I found this:

https://yourbasic.org/golang/read-file-line-by-line/

## to build and run

```

go build
cat testfile.txt | ./process-stdin
or 
cat testfile.txt | ./process-stdin > output.txt


or if you want to use go run and not produce a binary


cat testfile.txt | go run main.go
or
cat testfile.txt | go run main.go > output.txt


```

## if you want to extract code to a package

This is what I was fighting with in East London
Fiddle a bit and see what still works, maybe not all needed, but the way I have it:

a) the directory is util and it is "package util"
b) my func GetWhatIShouldAddOntoTheLine starts with an uppercase so it is visible externally in main. Any other private funcs just start them with lowercase
c) you HAVE to import it in main using "github.com/andrewrobinson/process-stdin/util"
vscode is useless at adding this for you and removes it if unused when saving
I always have to look this up from old projects!

d) the "github.com/andrewrobinson/process-stdin" there comes from

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


```

git remote add origin git@github.com:andrewrobinson/process-stdin.git
git add .
git commit -m 'wip'
git push
git push --set-upstream origin master



