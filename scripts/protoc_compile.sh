#!/bin/bash

#protoc -I external/DTCProtocol --go_out=./dtc/ ./external/DTCProtocol/DTCProtocol.proto
protoc --go_out=./ ./external/DTCProtocol/DTCProtocol.proto

