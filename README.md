# goScheduler

### Configuration
``` bash
$ git clone git@github.com:chentda/goScheduler.git && cd goscheduler/config/
$ vim example.toml # fill out name of bash script and time you want to schedule it to run
$ mv example.toml ../config.toml 
```

### To Test
``` base
$ make run
```

### To Run As Binary
```
$ make compile
$ ./goScheduler # goScheduler is the name of the executable - could be different for you if you change it in the Makefile
```
