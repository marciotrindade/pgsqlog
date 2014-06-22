It's a simple script that I crette to parse the postgres log file and insert into a table. Its extract "username", "database", "duration", "action", "table_name", "sql", "created_at" from log file.

# How to use

first clone the repo

`git clone git@github.com:marciotrindade/pgsqlog.git`

Changed connection information

`const stringPostgresConnection "user=emailmarketing dbname=psqlog sslmode=disable"`

Run the schema.sql in your databse to create the table log that will use to script.

And run the go to read the file, parse and insert into database.

`go run main.go PATH_OF_YOUR_FILE`

