# ChangeMackerelMonitor

## Description
Change host monitoring status [mackerel](https://mackerel.io).

```
working -> standby
standby -> working
```

## How to
Clone this repository and `go build` on your machine.

- List all hosts.

```
mmonitor
```

- List specifiled hosts.

```
mmonitor -service xxxxx -role xxxxxx
```

- Change monitor status to working

```
mmonitor -service xxxx -role xxxxxx -type working
```

- Change monitor status to standby

```
mmonitor -service xxxx -role xxxxxx -type standby
```
