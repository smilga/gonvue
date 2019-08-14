# gonvue
Spent too much time creating boilerplate code for some small projects that uses GO and Vue!?
So there goes working go web server with vue.

## Getting Started
Install frontend dependencies
```
cd frontend
npm install
```
There are two shell scripts to make it simple
### Development
To start go web server and webpack dev server both with live realod
```
./dev.sh
```
### Production
Production version will be under './build' folder
```
./prod.sh
```
To launch builded project
```
cd build
./server -port=8888
```
