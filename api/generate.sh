#!/bin/bash
protoc -I . --go_out=plugins=grpc:. --go_opt=paths=source_relative conn/conn.proto
protoc -I . --go_out=plugins=grpc:. --go_opt=paths=source_relative logic/logic.proto
protoc -I . --go_out=plugins=grpc:. --go_opt=paths=source_relative job/job.proto