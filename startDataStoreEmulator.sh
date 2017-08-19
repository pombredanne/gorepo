#!/bin/bash

gcloud beta emulators datastore start --data-dir=/tmp/gorepo-datastore --host-port=localhost:10101 --no-legacy
