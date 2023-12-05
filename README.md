Task:

Package customerimporter reads from a CSV file and returns a sorted (data
structure of your choice) of email domains along with the number of customers
with e-mail addresses for each domain. This should be able to be ran from the
CLI and output the sorted domains to the terminal or to a file. Any errors
should be logged (or handled). Performance matters (this is only ~3k lines,
but could be 1m lines or run on a small machine).

Description
when in project directory:
    to run:
        - go run cmd/app/main.go
    to build:
        - go build cmd/app/main.go
        - ./main

flags:
    -input
            input .csv file path where first row is a header containing columns names. 
            "email" is only one used so far (default "customers.csv")
    -print
            Print program output (default true)
    -save
            save program output to .csv file (default true)
    -sortby
            Choose to sort domains by 'count'(amount of custommers)
            or 'domain'(alphabetic) (default "count")
    -output 
            output .csv file path (default "sorted_domains.csv")

    To make code easy to scale "handlers" might be defined to deal with each column.
So far only 'email' has been added, but to work with other colum it is enough to
extend handlers map in pkg/handlers.go. Logic to output aditional data might be 
created as independend code and called inside cmd/main.go, so You can extend without
editing existing logic.

TODO:
    -tests
    -docks

Further improvmenets are possible, like:
    -use DNS to validate if mailing domain is correct
    -if inputfile header is not present try to find email coumn by comparing its
        values with template
    -output invalid domains/email (separed file?)
    -add parsing to other columns, like gender or IP
