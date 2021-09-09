#!/bin/bash

go test $(go list ./... | grep -v ./api/)