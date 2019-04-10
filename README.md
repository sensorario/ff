# ff 1.* [![GitHub version](https://badge.fury.io/gh/sensorario%2Fff.svg)](https://github.com/sensorario/ff/releases) [![Go Report Card](https://goreportcard.com/badge/github.com/sensorario/ff)](https://goreportcard.com/report/github.com/sensorario/ff) [![Build Status](https://travis-ci.org/sensorario/ff.svg?branch=master)](https://travis-ci.org/sensorario/ff)

## Install

    go get github.com/sensorario/ff
    env GO111MODULE=on go build -o /usr/local/bin/ff

## Features

The `ff` does not allow command if not allowed according to the git flow. For example, no hotfix/feature branches can be created if current branch is an hotfix/feature branch in turn.

An hotfix/feature branch can be created only from master.

In case of LTS (a minor version) after each new tag merge updates into master to keep development version updated with all hotfixes.

Logs are stored in .git/logger.log file

## Commands

 - commit
 - complete
 - feature
 - help
 - hotfix
 - bugfix
 - publish
 - refactor
 - reset
 - status

## Autocompletion

Append the following lines in your .bash_profile file:

    _ff='commit complete feature help hotfix bugfix publish refactor reset status' && complete -W "${_ff}" 'ff'
