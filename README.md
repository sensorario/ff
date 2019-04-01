# FF 1.*

## Install Binaries

Download binaries from: https://github.com/sensorario/ff/raw/master/dist/ff
And then make it executable:

    sudo chmod +x ff

Move it into `/usr/local/bin/` folder to execute it from everywhere.

## Configuration

    sudo mkdir /var/log/ff/
    sudo chmod 777 /var/log/ff/

## Description

A git wrapper.

### Help command

 - ff help

### Clean working directory and stage

 - ff reset

### Open and close feature branch

 - ff feature
 - ff feature complete

### Open and close hotfix branch

 - ff hotfix
 - ff hotfix complete

### Complete branch

 - ff complete

### Commit

 - ff commit

### Status

 - ff status

### Publish

 - ff publish
