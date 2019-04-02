# FF 1.*

## Install Binaries

Download binaries from: https://github.com/sensorario/ff/raw/master/dist/ff
And then make it executable:

    sudo chmod +x ff

Move it into `/usr/local/bin/` folder to execute it from everywhere.

Remember to load `gol`:

    go get github.com/sensorario/gol

## Personalization

By default log folder is `/var/log/ff` and file is `logger.log`.  To change log destination path you  should define environment variable `FF_LOG_PATH`.  Remember that the program MUST HAVE write permission in the selected folder. Actually is not possible to change file name.

    sudo mkdir /var/log/ff/
    sudo chmod 777 /var/log/ff/

## Available commands

 - ff help
 - ff reset
 - ff feature
 - ff feature complete
 - ff hotfix
 - ff hotfix complete
 - ff complete
 - ff commit
 - ff status
 - ff publish
