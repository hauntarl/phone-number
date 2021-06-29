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

**Additional Resources:**

- [Using PostgreSQL with Go](https://www.calhoun.io/using-postgresql-with-go/)
- [Connecting to a PostgreSQL database with Go's database/sql package](https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/)

<!-- 
**Run Commands:**

**Features:**

**Packages explored:**

**Output:**

``` terminal
```
-->
