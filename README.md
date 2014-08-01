It's a simple script that I created to parse the postgres log file and insert into a table. It extracts "username", "database", "duration", "action", "table_name", "sql", "created_at" from log file.

# How to use

first, clone the repo

`git clone git@github.com:marciotrindade/pgsqlog.git`

Changed connection information

`const stringPostgresConnection "user=emailmarketing dbname=psqlog sslmode=disable"`

Run the schema.sql in your databse to create the table log that will use to script.

Then buid the program

'go build'


And run the program to read the file, parse and insert into database.

`./pgsqlog PATH_OF_YOUR_FILE`
