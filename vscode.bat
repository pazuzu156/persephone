@echo off
set GOROOT=C:\Go
set GOPATH=%USERPROFILE%\go
set MINGW=C:\msys64\mingw64
set CODEPATH=%LOCALAPPDATA%\Programs\Microsoft VS Code
set WIN=%WINDIR%\System32
set GIT=%PROGRAMFILES%\git
set PATH=%WIN%;%GOROOT%\bin;%GOPATH%\bin;%MINGW%\bin;%GIT%\bin;%CODEPATH%
call code %~dp0
