# Phone Number Normalizer

>Implementation of Phone Number Normalizer from **[Gophercises](https://courses.calhoun.io/courses/cor_gophercises)**  by Jon Calhoun, including the bonus section.

A program that will iterate through a database and normalize all of the phone numbers in the DB. After normalizing all of the data, there might be duplicates, remove those duplicates and keep just one entry in our database.

While the primary goal is to create a program that normalizes the database, this also covers how to setup the database and write entries to it using Go as well. This is done intentionally to explore how to use Go when interacting with an SQL database.

There are many ways to use SQL in Go. Rather than just picking one.

- Writing raw SQL and using the **database/sql** package in the standard library
- Using the very popular **sqlx** third party package, which is basically an extension of Go’s sql package.
- Using a relatively minimalistic ORM (such as **gorm**)

We will also explore some basic string manipulation techniques to normalize our phone numbers, but the primary focus here is going to be on the SQL. We will start by creating a database along with a `phone_numbers` table. Inside that table we want to add the following entries:

``` text
1234567890
123 456 7891
(123) 456 7892
(123) 456-7893
123-456-7894
123-456-7890
1234567892
(123)456-7892
```

Once you have the entries created, our next step is to learn how to iterate over entries in the database using Go code. With this we should be able to retrieve every number so we can start normalizing its contents.

Once you have all the data in the DB, our next step is to normalize the phone number. We are going to update all of our numbers so that they match the format:

``` text
##########
```

That is, we are going to remove all formatting and only store the digits. When we want to display numbers later we can always format them, but for now we only need the digits.

Once you written code that will successfully take a number in with any format and return the same number in the proper format we are going to use an `UPDATE` to alter the entries in the database. If the value we are inserting into our database already exists (it is a duplicate), we will instead be deleting the original entry.

When your program is done, your database entries should look like this (the order is irrelevant, but duplicates should be removed):

``` text
1234567890
1234567891
1234567892
1234567893
1234567894
```

**Bonus:**

As an extension to this, we could also delete all the phone numbers which do not conform to the **North American Numbering Plan (NANP)**. It's a telephone numbering system used by many countries in North America like the United States, Canada or Bermuda. All NANP-countries share the same international country code: 1.

NANP numbers are ten-digit numbers consisting of a three-digit Numbering Plan Area code, commonly known as area code, followed by a seven-digit local number. The first three digits of the local number represent the exchange code, followed by the unique four-digit number which is the subscriber number.

The format is usually represented as:

``` text
(NXX)-NXX-XXXX
```

where N is any digit from 2 through 9 and X is any digit from 0 through 9.

For example, the inputs

``` text
+1 (613)-995-0253
613-995-0253
1 613 995 0253
613.995.0253
```

should all produce the output below from a Number() function.

``` text
6139950253
```

For more information refer: **[Phone Number](https://exercism.io/tracks/go/exercises/phone-number/solutions/c777be4b16d44f4da9f5f702836a8470)** on Exercism

**Additional Resources:**

- [Using PostgreSQL with Go](https://www.calhoun.io/using-postgresql-with-go/)
- [Connecting to a PostgreSQL database with Go's database/sql package](https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/)

**Run Commands:**

- go run main.go
- go build .
- .\phone-number

For testing and benchmarking the normalization function

- go test github.com/hauntarl/phone-number/normalize
- go test -benchmem -run=^$ -bench=BenchmarkNumber github.com/hauntarl/phone-number/normalize

**Features:**

- making use of transactional database
- writing regular expressions to normalize and validate phone numbers
- performing table-driven testing and benchmarking for normalizer function
- implementing basic CRUD operations using SQLite driver

**Packages explored:**

- regexp - to compile regular expression in Go objects
- testing - to perform unit testing and benchmarking
- os - to create a new database file
- math/rand - to shuffle the list of numbers
- path/filepath - to get the absolute path of database file
- database/sql - a common interface for users to communicate with different SQL driver providers
- **[modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite)** - A SQLite driver which implements the database/sql interface

**Output:**

``` terminal
2021/06/30 06:43:15 Database Created: (location=D:\gophercises\phone-number\phone-numbers.db)
2021/06/30 06:43:15 Table Created: (name=phone_numbers)
2021/06/30 06:43:15 Inserted 123 456 7891        : (uid=1  , rows affected=1)
2021/06/30 06:43:15 Inserted 321234567890        : (uid=2  , rows affected=1)
2021/06/30 06:43:15 Inserted 123456789           : (uid=3  , rows affected=1)
2021/06/30 06:43:16 Inserted (123)456-7892       : (uid=4  , rows affected=1)
2021/06/30 06:43:16 Inserted 1234567892          : (uid=5  , rows affected=1)
2021/06/30 06:43:16 Inserted 123-@:!-7890        : (uid=6  , rows affected=1)
2021/06/30 06:43:16 Inserted 1 (023) 456-7890    : (uid=7  , rows affected=1)
2021/06/30 06:43:16 Inserted (223) 156-7890      : (uid=8  , rows affected=1)
2021/06/30 06:43:16 Inserted 12234567890         : (uid=9  , rows affected=1)
2021/06/30 06:43:16 Inserted (223) 056-7890      : (uid=10 , rows affected=1)
2021/06/30 06:43:16 Inserted 123-456-7894        : (uid=11 , rows affected=1)
2021/06/30 06:43:16 Inserted 1 (223) 056-7890    : (uid=12 , rows affected=1)
2021/06/30 06:43:16 Inserted (023) 456-7890      : (uid=13 , rows affected=1)
2021/06/30 06:43:16 Inserted 22234567890         : (uid=14 , rows affected=1)
2021/06/30 06:43:16 Inserted (223) 456-7890      : (uid=15 , rows affected=1)
2021/06/30 06:43:16 Inserted 123-456-7890        : (uid=16 , rows affected=1)
2021/06/30 06:43:17 Inserted (123) 456-7890      : (uid=17 , rows affected=1)
2021/06/30 06:43:17 Inserted (123) 456-7893      : (uid=18 , rows affected=1)
2021/06/30 06:43:17 Inserted (123) 456 7892      : (uid=19 , rows affected=1)
2021/06/30 06:43:17 Inserted +1 (223) 456-7890   : (uid=20 , rows affected=1)
2021/06/30 06:43:17 Inserted 223.456.7890        : (uid=21 , rows affected=1)
2021/06/30 06:43:17 Inserted 223 456   7890      : (uid=22 , rows affected=1)
2021/06/30 06:43:17 Inserted 1 (223) 156-7890    : (uid=23 , rows affected=1)
2021/06/30 06:43:17 Inserted 1 (123) 456-7890    : (uid=24 , rows affected=1)
2021/06/30 06:43:17 Inserted 123-abc-7890        : (uid=25 , rows affected=1)
Phone Numbers Table...
uid:   1, number: '123 456 7891'
uid:   2, number: '321234567890'
uid:   3, number: '123456789'
uid:   4, number: '(123)456-7892'
uid:   5, number: '1234567892'
uid:   6, number: '123-@:!-7890'
uid:   7, number: '1 (023) 456-7890'
uid:   8, number: '(223) 156-7890'
uid:   9, number: '12234567890'
uid:  10, number: '(223) 056-7890'
uid:  11, number: '123-456-7894'
uid:  12, number: '1 (223) 056-7890'
uid:  13, number: '(023) 456-7890'
uid:  14, number: '22234567890'
uid:  15, number: '(223) 456-7890'
uid:  16, number: '123-456-7890'
uid:  17, number: '(123) 456-7890'
uid:  18, number: '(123) 456-7893'
uid:  19, number: '(123) 456 7892'
uid:  20, number: '+1 (223) 456-7890'
uid:  21, number: '223.456.7890'
uid:  22, number: '223 456   7890   '
uid:  23, number: '1 (223) 156-7890'
uid:  24, number: '1 (123) 456-7890'
uid:  25, number: '123-abc-7890'
total rows: 25
----------------------
2021/06/30 06:43:17 Normalizing all the phone numbers...
2021/06/30 06:43:17 Deleted 123 456 7891        : (uid=1  , rows affected=1), Reason: area code cannot start with one
2021/06/30 06:43:17 Deleted 321234567890        : (uid=2  , rows affected=1), Reason: more than 11 digits
2021/06/30 06:43:18 Deleted 123456789           : (uid=3  , rows affected=1), Reason: incorrect number of digits
2021/06/30 06:43:18 Deleted (123)456-7892       : (uid=4  , rows affected=1), Reason: area code cannot start with one
2021/06/30 06:43:18 Deleted 1234567892          : (uid=5  , rows affected=1), Reason: area code cannot start with one
2021/06/30 06:43:18 Deleted 123-@:!-7890        : (uid=6  , rows affected=1), Reason: punctuations not permitted
2021/06/30 06:43:18 Deleted 1 (023) 456-7890    : (uid=7  , rows affected=1), Reason: area code cannot start with zero
2021/06/30 06:43:18 Deleted (223) 156-7890      : (uid=8  , rows affected=1), Reason: exchange code cannot start with one
2021/06/30 06:43:18 Updated 12234567890 to 2234567890: (uid=9  , rows affected=1)
2021/06/30 06:43:19 Deleted (223) 056-7890      : (uid=10 , rows affected=1), Reason: exchange code cannot start with zero
2021/06/30 06:43:19 Deleted 123-456-7894        : (uid=11 , rows affected=1), Reason: area code cannot start with one
2021/06/30 06:43:19 Deleted 1 (223) 056-7890    : (uid=12 , rows affected=1), Reason: exchange code cannot start with zero
2021/06/30 06:43:19 Deleted (023) 456-7890      : (uid=13 , rows affected=1), Reason: area code cannot start with zero
2021/06/30 06:43:19 Deleted 22234567890         : (uid=14 , rows affected=1), Reason: 11 digits must start with 1
2021/06/30 06:43:19 Deleted (223) 456-7890      : (uid=15 , rows affected=1), Reason: duplicate phone number
2021/06/30 06:43:19 Deleted 123-456-7890        : (uid=16 , rows affected=1), Reason: area code cannot start with one
2021/06/30 06:43:19 Deleted (123) 456-7890      : (uid=17 , rows affected=1), Reason: area code cannot start with one
2021/06/30 06:43:19 Deleted (123) 456-7893      : (uid=18 , rows affected=1), Reason: area code cannot start with one
2021/06/30 06:43:19 Deleted (123) 456 7892      : (uid=19 , rows affected=1), Reason: area code cannot start with one
2021/06/30 06:43:19 Deleted +1 (223) 456-7890   : (uid=20 , rows affected=1), Reason: duplicate phone number
2021/06/30 06:43:20 Deleted 223.456.7890        : (uid=21 , rows affected=1), Reason: duplicate phone number
2021/06/30 06:43:20 Deleted 223 456   7890      : (uid=22 , rows affected=1), Reason: duplicate phone number
2021/06/30 06:43:20 Deleted 1 (223) 156-7890    : (uid=23 , rows affected=1), Reason: exchange code cannot start with one
2021/06/30 06:43:20 Deleted 1 (123) 456-7890    : (uid=24 , rows affected=1), Reason: area code cannot start with one
2021/06/30 06:43:20 Deleted 123-abc-7890        : (uid=25 , rows affected=1), Reason: letters not permitted
Formatted Numbers...
uid:   9, number: '(223) 456-7890'
total rows: 1
--------------------
2021/06/30 06:43:20 Database Closed
```
