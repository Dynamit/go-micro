# MySQL

This package is intended to allow Go applications to share a MySQL connection.

## Environment Variables

The following environment variables must be set in order for the MySQL package to locate the MySQL server. These are typically set my labeling your MySQL docker container "mysql" and linking it to your go application container.

```
MYSQL_PORT_3306_TCP_ADDR
MYSQL_PORT_3306_TCP_PORT
MYSQL_ENV_MYSQL_USER
MYSQL_ENV_MYSQL_PASSWORD
MYSQL_ENV_MYSQL_DATABASE
```

## Usage

First, you must import the package into your project.

```
import github.com/dynamit/go-micro/mysql
```

To open (or reuse) a MySQL connection:

```
mysql.Open()
```

To close a MySQL connection:

```
mysql.Close()
```