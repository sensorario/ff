# FF 1.*

[![Go Report Card](https://goreportcard.com/badge/github.com/sensorario/ff)](https://goreportcard.com/report/github.com/sensorario/ff)
[![Build Status](https://travis-ci.org/sensorario/ff.svg?branch=master)](https://travis-ci.org/sensorario/ff)

## Install Binaries

Download binaries from: https://github.com/sensorario/ff/raw/master/dist/ff
And then make it executable:

    sudo chmod +x ff

Move it into `/usr/local/bin/` folder to execute it from everywhere.

Remember to load `gol` and `color`:

    go get github.com/sensorario/gol
    go get github.com/fatih/color

## Features

The `ff` does not allow command if not allowed according to the git flow. For example, no hotfix/feature branches can be created if current branch is an hotfix/feature branch in turn.

An hotfix/feature branch can be created only from master.

In case of LTS (a minor version) after each new tag merge updates into master to keep development version updated with all hotfixes.

## Personalization

By default log folder is `/var/log/ff` and file is `logger.log`.  To change log destination path you  should define environment variable `FF_LOG_PATH`.  Remember that the program MUST HAVE write permission in the selected folder. Actually is not possible to change file name.

    sudo mkdir /var/log/ff/
    sudo chmod 777 /var/log/ff/

## Available commands

 - ff commit
 - ff complete
 - ff feature
 - ff help
 - ff hotfix
 - ff bugfix
 - ff publish
 - ff refactor
 - ff reset
 - ff status

## Autocompletion

Append the following lines in your .bash_profile file:

    _ff='commit complete feature help hotfix bugfix publish refactor reset status' && complete -W "${_ff}" 'ff'
