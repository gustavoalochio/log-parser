# log-parser

This project is reponsible for parser a log file, the log file was generated by a Quake 3 Arena server, including a great deal of information of every match.

Simulating a integration between script and log output, the script use stdin to get the information.

With the informations above

### Dependencies

To execute the script is required `Golang`

### Execution

Execute script command in terminal:
```bash 
make start < log/qgames.log
```
### Tests

Execute the tests with command:

```bash 
make test
```