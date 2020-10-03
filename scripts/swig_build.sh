#!/bin/bash

cd dtc
swig -c++ -go -cgo -intgosize 64 -module dtc DTCProtocol.cpp

