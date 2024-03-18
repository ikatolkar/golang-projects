# go-diy-database

Database from scratch

JSON database similar to MongoDB

keeps records in a collection (directory) as json (files)

# Features
## Driver not ORM
one database driver can be used to perform operations
one driver per collection
any program using this can import this package, create a driver and use that to interact with DB

## Data integrity using mutexes
Similar to cockroach DB
Driver contains a map of mutexes for each record
On every write and delete, mutex must be acquired

## Operations
1. Write
2. Read/ReadAll
3. Delete

