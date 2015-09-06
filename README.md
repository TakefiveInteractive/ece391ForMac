# ECE391 Auto Setup For Mac Users


This is a repo created for ece391, since officially there is no setup guide provided for OS X. 

Fully tested on OS X 10.10 :) 

For OS X 10.11 User, please ```make``` under ```cowhacker``` and then move ```cowhacker/bin/macosx/*/cowhacker{the one you just compiled}``` under ```cowhacker/```
## Usage

```fish
cd ~
git clone "https://github.com/Tedko/ece391ForMac.git"
```
then run 

```fish
cd ece391ForMac
./start
```

then you can use the command ```ece391.dev``` to start the dev machine


if you prefer to use your default terminal or iTerm, you can then use
```ece391.ssh```after your dev machine successfully boot.


It you've down with your setting or MP, you can then run 

```
./sync
```

under ~/ece391ForMac dir

You can use ```./mount391``` and ```./unmount391``` for shortcuts for mounting Shared Volumes :)



**Special Thanks for our TA Fei Deng**
