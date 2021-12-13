#!/bin/bash

GO_COMP=gotip

function fail() {
    echo $1
    exit 1
}

cd cont
for step in build fmt vet "test -cover"; do
    $GO_COMP $step || fail "Step $step failed. Aborting"
done