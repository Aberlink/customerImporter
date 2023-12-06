## Task

Package `customerimporter` reads from a CSV file and returns a sorted data structure (of your choice) of email domains along with the number of customers with e-mail addresses for each domain. This should be able to be run from the CLI and output the sorted domains to the terminal or to a file. Any errors should be logged (or handled). Performance matters (this is only ~3k lines, but could be 1m lines or run on a small machine).

## Description

When in the project directory:

- To build:
```bash
go build -o <file_name> ./cmd/app 
```

- To run:
```bash
go run ./cmd/app
```

- To run tests:
```bash
go test -v ./...
```

- Docker:
  build:
```bash
docker build -t customerimporter:latest .
```
start container (adjust local paths, both files have to exist when container starts):
```bash
docker run -d --name customerimporter -v "$(pwd)/customers.csv:/app/customers.csv" -v "$(pwd)/sorted_domains.csv:/app/sorted_domains.csv" customerimporter
```
enter into container, use `./customerimporter` to execute scripts. 
```bash
docker exec -it customerimporter /bin/bash
```

Flags:

- `-input`: Input .csv file path where the first row is a header containing column names. "email" is the only one used so far (default "customers.csv").
- `-print`: Print program output (default true).
- `-save`: Save program output to .csv file (default true).
- `-sortby`: Choose to sort domains by 'count' (amount of customers) or 'domain' (alphabetic) (default "count").
- `-output`: Output .csv file path (default "sorted_domains.csv").

To make the code easy to scale, "handlers" might be defined to deal with each column. So far only 'email' has been added, but to work with other columns, it is enough to extend handlers map in `pkg/handlers.go`. Logic to output additional data might be created as independent code and called inside `cmd/main.go`, so you can extend without editing existing logic.


Further improvements are possible, like:

- Use DNS to validate if the mailing domain is correct.
- If the input file header is not present, try to find the email column by comparing its values with the template.
- Output invalid domains/emails (separated file?).
- Add parsing to other columns, like gender or IP.


