# lumberjack
Logging library for Go. I didn't like any of the logging frameworks so decided to make my own. Time will tell if this was a good idea.

MakeLogs generates a set of logs for different levels. Each level enables all
sets of logs for each level lower than the selected level.
0 FATAL -- Abandon all hope, log the event, and put an end to the farce
1 ERROR -- Errors encountered during runtime
2 INFO -- General information
3 DEBUG -- An enhanced set of log collection to be used during debugging

```
logger := lumberjack.MakeLogs(3, &buf)
logger.DEBUG("Here is all the available logging info")
```
