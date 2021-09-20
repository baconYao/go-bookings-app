#!/bin/bash

# Build the project and execute it if only if build success
go build -o bookings cmd/web/*.go && ./bookings