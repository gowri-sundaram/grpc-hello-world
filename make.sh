#!/bin/bash

protoc src/greet/greetpb/greet.proto --go_out=plugins=grpc:.