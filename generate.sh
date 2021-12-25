#!/bin/bash

protoc prime/primepb/prime.proto --go_out=plugins=grpc:.