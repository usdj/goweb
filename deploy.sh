#!/bin/bash

kill -9 $(pgrep webserver)
cd ~/goweb/
git pull https://github.com/usdj/goweb.git
cd webserver/
./webserver & 
