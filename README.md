# goScheduler

### When you would use this scheduler instead of cron
You would use this scheduler on a team machine where the system administrators have prevented the use of cron. Valid reasons why they prevent the use of cron include:
* Jobs overlapping.
* Dataset that job processes becomes too large.
* Limits to the number of open file descriptors, threads and memory usage on a machine
* Other race condition problems since users assume the entire system's resources are at their disposal.

**Beware: Due to above reasons, use this simple scheduler at your own risk as you could get in trouble with system administrators!** I personally used it to refresh application data where the dataset was small and posed no risks.

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
You would normally run this in a terminal multiplexer like [tmux](https://github.com/tmux/tmux/wiki) or [screen](https://linuxize.com/post/how-to-use-linux-screen/)
```
$ make compile
$ ./goScheduler # goScheduler is the name of the executable - could be different for you if you change it in the Makefile
```
### Limitations
* Only schedules bash scripts, so if you need to run something else like a python script, just called it in the bash script.
* Only schedules 1 bash script currently, so you will have to duplicate the executable if you want to run multiple bash scripts.
* Only runs Mon-Fri. 
