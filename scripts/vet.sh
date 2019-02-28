#!/bin/bash

set -euo pipefail

echo "Checking for common Go mistakes"
go vet -printfuncs=Debug,Debugf,Debugln,Info,Infof,Infoln,Notice,Noticef,Noticeln,Error,Errorf,Errorln,Warning,Warningf,Warningln,Critical,Criticalf,Criticalln ./...