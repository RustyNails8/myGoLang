@echo off

rem Build exec for headfirst own codes
rem Sumit Das
rem 2020 05 24


set PROJDIR=C:\Users\in10c2\go\src\gitlab.com\RustyNails8
rem set WORKDIR=C:\Users\in10c2\go\src\gitlab.com\RustyNails8\cmd
set WORKDIR=%cd%
set BINDIR=C:\Users\in10c2\go\src\gitlab.com\RustyNails8\bin

rem set CHKDIR=%cd%

rem echo "Directory is " %CHKDIR%
cd %BINDIR%
echo "Running your GO code"
echo .
echo ..
echo ...
go run %WORKDIR%\%1
echo .
echo ..
echo ...
echo "Building your GO code"
go build %WORKDIR%\%1

echo .
echo ..
echo ...
echo "ALL DONE !"
cd %WORKDIR%
