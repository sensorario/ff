# FF 1.*

## Install Binaries

Download binaries from: https://github.com/sensorario/ff/raw/master/dist/ff
And then make it executable:

    sudo chmod +x ff

Move it into `/usr/local/bin/` folder to execute it from everywhere.

Remember to load `gol` and `color`:

    go get github.com/sensorario/gol
    go get github.com/fatih/color

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
 - ff publish
 - ff refactor
 - ff reset
 - ff status
