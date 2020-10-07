#!/bin/bash

protoc -I external/DTCProtocol --go_out=../ ./external/DTCProtocol/DTCProtocol.proto

