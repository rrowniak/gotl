#!/bin/bash

GO_COMP=gotip
REP=coverage.out

cd cont || exit 1

if $GO_COMP test -coverprofile=$REP; then
    trap rm $REP
    $GO_COMP tool cover -html=$REP
fi

