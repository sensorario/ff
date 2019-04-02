# FF 1.*

## Install Binaries

Download binaries from: https://github.com/sensorario/ff/raw/master/dist/ff
And then make it executable:

    sudo chmod +x ff

Move it into `/usr/local/bin/` folder to execute it from everywhere.

## Personalization

By default log folder is `/var/log/ff` and file is `logger.log`.  To change log destination path you  should define environment variable `FF_LOG_PATH`.  Remember that the program MUST HAVE write permission in the selected folder. Actually is not possible to change file name.

    sudo mkdir /var/log/ff/
    sudo chmod 777 /var/log/ff/

## Commands

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
