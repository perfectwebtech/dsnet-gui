dsnet-gui is a web based gui for [dsnet](https://github.com/naggie)

!TODO add image when public

It's a very simple tool that gives you a self updating web ui about the dsnet controlled wireguard interface.

To make `make`  
To run do `sudo ./dsnet-gui serve`  
Open your browser to `http://127.0.0.1:20080`  

There are plenty of cleanup items left to do on the list

### Dev

This will eventually just be one command

`git clone git@github.com/botto/dsnet-gui.git`  
`cd client`  
`npm i`  
`npm run start`  
New terminal  
`cd dsnet-gui`  
`sudo go run . serve`  
