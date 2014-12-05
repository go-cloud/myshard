#!/bin/bash

#Test godep install
# godep path > /dev/null 2>&1
# if [ "$?" = 0 ]; then
#     GOPATH=`godep path`
#     # https://github.com/tools/godep/issues/60
#     # have to rm Godeps/_workspace first, then restore
#     rm -rf $GOPATH
#     godep restore
#     exit 0
# fi
