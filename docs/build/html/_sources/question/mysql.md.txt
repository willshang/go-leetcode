# mysql

* [mysql](#mysql)
    * [175.组合两个表(2)](#175组合两个表2)
    * [176.第二高的薪水(3)](#176第二高的薪水3)
    * [181.超过经理收入的员工(3)](#181超过经理收入的员工3)
    * [182.查找重复的电子邮箱(2)](#182查找重复的电子邮箱2)
    * [183.从不订购的客户(3)](#183从不订购的客户3)
    * [193.删除重复的电子邮箱(2)](#193删除重复的电子邮箱2)
    * [197.上升的温度(4)](#197上升的温度4)
    * [595.大的国(2)](#595大的国2)
    * [596.超过5名学生的课(2)](#596超过5名学生的课2)
    * [620.有趣的电影(2)](#620有趣的电影2)
    * [627.交换工资(3)](#627交换工资3)
    
## 175.组合两个表(2)

- 题目

```
SQL架构
Create table Person (PersonId int, FirstName varchar(255), LastName varchar(255))
Create table Address (AddressId int, PersonId int, City varchar(255), State varchar(255))
Truncate table Person
insert into Person (PersonId, LastName, FirstName) values ('1', 'Wang', 'Allen')
Truncate table Address
insert into Address (AddressId, PersonId, City, State) values ('1', '2', 'New York City', 'New York')

表1: Person
+-------------+---------+
| 列名         | 类型     |
+-------------+---------+
| PersonId    | int     |
| FirstName   | varchar |
| LastName    | varchar |
+-------------+---------+
PersonId 是上表主键

表2: Address
+-------------+---------+
| 列名         | 类型    |
+-------------+---------+
| AddressId   | int     |
| PersonId    | int     |
| City        | varchar |
| State       | varchar |
+-------------+---------+
AddressId 是上表主键

编写一个 SQL 查询，满足条件：无论 person 是否有地址信息，都需要基于上述两表提供 person 的以下信息：
FirstName, LastName, City, State
```

- 解题思路

| No.  | 思路               |
| ---- | ------------------ |
| 01   | 考察join的基本使用 |
| 02   | 考察join的基本使用 |

```mysql
select FirstName, LastName, City, State 
from Person left join Address on Person.PersonId = Address.PersonId

#
select A.FirstName, A.LastName, B.City, B.State
from Person A
left join (select distinct PersonId, City, State from Address) B
on A.PersonId=B.PersonId;
```

## 176.第二高的薪水(3)

- 题目

```
SQL架构
Create table If Not Exists Employee (Id int, Salary int)
Truncate table Employee
insert into Employee (Id, Salary) values ('1', '100')
insert into Employee (Id, Salary) values ('2', '200')
insert into Employee (Id, Salary) values ('3', '300')

编写一个 SQL 查询，获取 Employee 表中第二高的薪水（Salary） 。
+----+--------+
| Id | Salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+

例如上述 Employee 表，SQL查询应该返回 200 作为第二高的薪水。
如果不存在第二高的薪水，那么查询应返回 null。
+---------------------+
| SecondHighestSalary |
+---------------------+
| 200                 |
+---------------------+
```

- 解题思路

| No.  | 思路                                 |
| ---- | ------------------------------------ |
| 01   | 把select语句包起来，使空的时候为null |
| 02   | 使用ifnull                           |
| 03   | 先查出最大的，然后查出比最大小的     |

```sql
select(
    select distinct Salary 
    from Employee 
    order by Salary desc
    limit 1 offset 1
) as SecondHighestSalary;

#
select ifnull(
    (select distinct Salary 
    from Employee 
    order by Salary desc
    limit 1 offset 1),null
) as SecondHighestSalary;

# 
select max(Salary) as SecondHighestSalary 
from Employee 
where Salary < (select max(Salary) from Employee);
```

## 181.超过经理收入的员工(3)

- 题目

```
SQL架构
Create table If Not Exists Employee (Id int, Name varchar(255), Salary int, ManagerId int)
Truncate table Employee
insert into Employee (Id, Name, Salary, ManagerId) values ('1', 'Joe', '70000', '3')
insert into Employee (Id, Name, Salary, ManagerId) values ('2', 'Henry', '80000', '4')
insert into Employee (Id, Name, Salary, ManagerId) values ('3', 'Sam', '60000', 'None')
insert into Employee (Id, Name, Salary, ManagerId) values ('4', 'Max', '90000', 'None')

Employee 表包含所有员工，他们的经理也属于员工。
每个员工都有一个 Id，此外还有一列对应员工的经理的 Id。

+----+-------+--------+-----------+
| Id | Name  | Salary | ManagerId |
+----+-------+--------+-----------+
| 1  | Joe   | 70000  | 3         |
| 2  | Henry | 80000  | 4         |
| 3  | Sam   | 60000  | NULL      |
| 4  | Max   | 90000  | NULL      |
+----+-------+--------+-----------+

给定 Employee 表，编写一个 SQL 查询，该查询可以获取收入超过他们经理的员工的姓名。
在上面的表格中，Joe 是唯一一个收入超过他的经理的员工。

+----------+
| Employee |
+----------+
| Joe      |
+----------+
```

- 解题思路

| No.  | 思路                        |
| ---- | --------------------------- |
| 01   | 使用笛卡尔乘积，和方法2一样 |
| 02   | 使用内链接                  |
| 03   | 子查询                      |

```sql
SELECT a.Name AS 'Employee'
FROM Employee AS a, Employee AS b
WHERE a.ManagerId = b.Id AND a.Salary > b.Salary;

#
SELECT a.Name AS 'Employee'
FROM Employee AS a join Employee AS b
on a.ManagerId = b.Id AND a.Salary > b.Salary;

#
select name as Employee 
from employee a 
where salary > (select salary from employee where a.managerid = id);
```

## 182.查找重复的电子邮箱(2)

- 题目

```
SQL架构
Create table If Not Exists Person (Id int, Email varchar(255))
Truncate table Person
insert into Person (Id, Email) values ('1', 'a@b.com')
insert into Person (Id, Email) values ('2', 'c@d.com')
insert into Person (Id, Email) values ('3', 'a@b.com')

编写一个 SQL 查询，查找 Person 表中所有重复的电子邮箱。
示例：
+----+---------+
| Id | Email   |
+----+---------+
| 1  | a@b.com |
| 2  | c@d.com |
| 3  | a@b.com |
+----+---------+

根据以上输入，你的查询应返回以下结果：
+---------+
| Email   |
+---------+
| a@b.com |
+---------+

说明：所有电子邮箱都是小写字母。
```

- 解题思路

| No.  | 思路           |
| ---- | -------------- |
| 01   | 使用临时表     |
| 02   | 使用having子句 |

```sql
select Email from
(
    select Email, count(Email) as num
    from Person 
    Group by Email
) as temp_table 
where num > 1;

//
select Email 
from Person
group by Email
having count(Email) > 1;
```

## 183.从不订购的客户(3)

- 题目

```
SQL架构
Create table If Not Exists Customers (Id int, Name varchar(255))
Create table If Not Exists Orders (Id int, CustomerId int)
Truncate table Customers
insert into Customers (Id, Name) values ('1', 'Joe')
insert into Customers (Id, Name) values ('2', 'Henry')
insert into Customers (Id, Name) values ('3', 'Sam')
insert into Customers (Id, Name) values ('4', 'Max')
Truncate table Orders
insert into Orders (Id, CustomerId) values ('1', '3')
insert into Orders (Id, CustomerId) values ('2', '1')
某网站包含两个表，Customers 表和 Orders 表。编写一个 SQL 查询，找出所有从不订购任何东西的客户。

Customers 表：
+----+-------+
| Id | Name  |
+----+-------+
| 1  | Joe   |
| 2  | Henry |
| 3  | Sam   |
| 4  | Max   |
+----+-------+

Orders 表：
+----+------------+
| Id | CustomerId |
+----+------------+
| 1  | 3          |
| 2  | 1          |
+----+------------+

例如给定上述表格，你的查询应返回：
+-----------+
| Customers |
+-----------+
| Henry     |
| Max       |
+-----------+
```

- 解题思路

| No.  | 思路           |
| ---- | -------------- |
| 01   | 使用not in     |
| 02   | 左连接         |
| 03   | 使用not exists |

```sql
select Customers.Name as Customers
from Customers
where Customers.Id not in (
    select CustomerId from Orders
);

#
select a.Name as Customers
from Customers as a left join Orders as b on a.Id=b.CustomerId
where b.CustomerId is null;

#
select name Customers 
from customers c 
where not exists (
    select 1 from orders o 
    where o.customerid=c.id
)
```

## 193.删除重复的电子邮箱(2)

- 题目

```
编写一个 SQL 查询，来删除 Person 表中所有重复的电子邮箱，重复的邮箱里只保留 Id 最小 的那个。
+----+------------------+
| Id | Email            |
+----+------------------+
| 1  | john@example.com |
| 2  | bob@example.com  |
| 3  | john@example.com |
+----+------------------+
Id 是这个表的主键。
例如，在运行你的查询语句之后，上面的 Person 表应返回以下几行:
+----+------------------+
| Id | Email            |
+----+------------------+
| 1  | john@example.com |
| 2  | bob@example.com  |
+----+------------------+
提示：
    执行 SQL 之后，输出是整个 Person 表。
    使用 delete 语句。
```

- 解题思路

| No.  | 思路                                          |
| ---- | --------------------------------------------- |
| 01   | 使用delete+自连接                             |
| 02   | 使用delete，根据group by和min()查询出最小的id |

```sql
# 
delete p1 from Person p1, Person P2 
where p1.Email = p2.Email and p1.Id > p2.Id

#
delete from Person
where id not in (
    select id from (
        select min(id) as id
        from Person
        group by Email
    ) as temp_table
)
```

## 197.上升的温度(4)

- 题目

```
SQL架构
Create table If Not Exists Weather (Id int, RecordDate date, Temperature int)
Truncate table Weather
insert into Weather (Id, RecordDate, Temperature) values ('1', '2015-01-01', '10')
insert into Weather (Id, RecordDate, Temperature) values ('2', '2015-01-02', '25')
insert into Weather (Id, RecordDate, Temperature) values ('3', '2015-01-03', '20')
insert into Weather (Id, RecordDate, Temperature) values ('4', '2015-01-04', '30')

给定一个 Weather 表，编写一个 SQL 查询，来查找与之前（昨天的）日期相比温度更高的所有日期的 Id。
+---------+------------------+------------------+
| Id(INT) | RecordDate(DATE) | Temperature(INT) |
+---------+------------------+------------------+
|       1 |       2015-01-01 |               10 |
|       2 |       2015-01-02 |               25 |
|       3 |       2015-01-03 |               20 |
|       4 |       2015-01-04 |               30 |
+---------+------------------+------------------+

例如，根据上述给定的 Weather 表格，返回如下 Id:
+----+
| Id |
+----+
|  2 |
|  4 |
+----+
```

- 解题思路

| No.  | 思路                           |
| ---- | ------------------------------ |
| 01   | 自连接和datediff()的使用       |
| 02   | 自连接和adddate()的使用        |
| 03   | 自连接和unix_timestamp()的使用 |
| 04   | 自连接和subdate()的使用        |

```sql
#
select A.Id as "Id"
from Weather A join Weather B
on datediff(A.RecordDate, B.RecordDate) = 1 and A.Temperature > B.Temperature

#
select A.Id as "Id"
from Weather A join Weather B
on A.Temperature > B.Temperature and A.RecordDate = adddate(B.RecordDate, 1) 

#
select A.Id as "Id"
from Weather A join Weather B
on unix_timestamp(A.RecordDate) = unix_timestamp(B.RecordDate) + 86400
and A.Temperature > B.Temperature

#
select A.Id as "Id"
from Weather A join Weather B
on A.Temperature > B.Temperature and B.RecordDate = subdate(A.RecordDate, 1)
```

## 595.大的国(2)

- 题目

```
Create table If Not Exists World (name varchar(255), continent varchar(255), area int, population int, gdp int)
Truncate table World
insert into World (name, continent, area, population, gdp) values ('Afghanistan', 'Asia', '652230', '25500100', '20343000000')
insert into World (name, continent, area, population, gdp) values ('Albania', 'Europe', '28748', '2831741', '12960000000')
insert into World (name, continent, area, population, gdp) values ('Algeria', 'Africa', '2381741', '37100000', '188681000000')
insert into World (name, continent, area, population, gdp) values ('Andorra', 'Europe', '468', '78115', '3712000000')
insert into World (name, continent, area, population, gdp) values ('Angola', 'Africa', '1246700', '20609294', '100990000000')

这里有张 World 表

+-----------------+------------+------------+--------------+---------------+
| name            | continent  | area       | population   | gdp           |
+-----------------+------------+------------+--------------+---------------+
| Afghanistan     | Asia       | 652230     | 25500100     | 20343000      |
| Albania         | Europe     | 28748      | 2831741      | 12960000      |
| Algeria         | Africa     | 2381741    | 37100000     | 188681000     |
| Andorra         | Europe     | 468        | 78115        | 3712000       |
| Angola          | Africa     | 1246700    | 20609294     | 100990000     |
+-----------------+------------+------------+--------------+---------------+

如果一个国家的面积超过300万平方公里，或者人口超过2500万，那么这个国家就是大国家。
编写一个SQL查询，输出表中所有大国家的名称、人口和面积。
例如，根据上表，我们应该输出:
+--------------+-------------+--------------+
| name         | population  | area         |
+--------------+-------------+--------------+
| Afghanistan  | 25500100    | 652230       |
| Algeria      | 37100000    | 2381741      |
+--------------+-------------+--------------+
```

- 解题思路

| No.  | 思路        |
| ---- | ----------- |
| 01   | or的使用    |
| 02   | union的使用 |

```sql
select name, population, area from world
where area > 3000000 or population > 25000000

#
# or具有全表扫描机制
# union具有索引列查询速度快
# Union：对两个结果集进行并集操作，不包括重复行，同时进行默认规则的排序；
# Union All：对两个结果集进行并集操作，包括重复行，不进行排序；
select name, population, area from world where area > 3000000
union
select name, population, area from world where population > 25000000
```

## 596.超过5名学生的课(2)

- 题目

```
Create table If Not Exists courses (student varchar(255), class varchar(255))
Truncate table courses
insert into courses (student, class) values ('A', 'Math')
insert into courses (student, class) values ('B', 'English')
insert into courses (student, class) values ('C', 'Math')
insert into courses (student, class) values ('D', 'Biology')
insert into courses (student, class) values ('E', 'Math')
insert into courses (student, class) values ('F', 'Computer')
insert into courses (student, class) values ('G', 'Math')
insert into courses (student, class) values ('H', 'Math')
insert into courses (student, class) values ('I', 'Math')

有一个courses 表 ，有: student (学生) 和 class (课程)。

请列出所有超过或等于5名学生的课。

例如,表:
+---------+------------+
| student | class      |
+---------+------------+
| A       | Math       |
| B       | English    |
| C       | Math       |
| D       | Biology    |
| E       | Math       |
| F       | Computer   |
| G       | Math       |
| H       | Math       |
| I       | Math       |
+---------+------------+

应该输出:
+---------+
| class   |
+---------+
| Math    |
+---------+
Note:
学生在每个课中不应被重复计算。
```

- 解题思路

| No.  | 思路              |
| ---- | ----------------- |
| 01   | group by + having |
| 02   | group by + 临时表 |

```sql
select class from courses
group by class
having count(distinct student) >= 5

#
select class from 
(select class, count(distinct student) as num from courses group by class) as temp_table
where num >= 5
```

## 620.有趣的电影(2)

- 题目

```
Create table If Not Exists cinema (id int, movie varchar(255), description varchar(255), rating float(2, 1))
Truncate table cinema
insert into cinema (id, movie, description, rating) values ('1', 'War', 'great 3D', '8.9')
insert into cinema (id, movie, description, rating) values ('2', 'Science', 'fiction', '8.5')
insert into cinema (id, movie, description, rating) values ('3', 'irish', 'boring', '6.2')
insert into cinema (id, movie, description, rating) values ('4', 'Ice song', 'Fantacy', '8.6')
insert into cinema (id, movie, description, rating) values ('5', 'House card', 'Interesting', '9.1')

某城市开了一家新的电影院，吸引了很多人过来看电影。
该电影院特别注意用户体验，专门有个 LED显示板做电影推荐，上面公布着影评和相关电影描述。
作为该电影院的信息部主管，您需要编写一个 SQL查询，
找出所有影片描述为非 boring (不无聊) 的并且 id 为奇数 的影片，结果请按等级 rating 排列。
例如，下表 cinema:
+---------+-----------+--------------+-----------+
|   id    | movie     |  description |  rating   |
+---------+-----------+--------------+-----------+
|   1     | War       |   great 3D   |   8.9     |
|   2     | Science   |   fiction    |   8.5     |
|   3     | irish     |   boring     |   6.2     |
|   4     | Ice song  |   Fantacy    |   8.6     |
|   5     | House card|   Interesting|   9.1     |
+---------+-----------+--------------+-----------+
对于上面的例子，则正确的输出是为：
+---------+-----------+--------------+-----------+
|   id    | movie     |  description |  rating   |
+---------+-----------+--------------+-----------+
|   5     | House card|   Interesting|   9.1     |
|   1     | War       |   great 3D   |   8.9     |
+---------+-----------+--------------+-----------+
```

- 解题思路

| No.  | 思路      |
| ---- | --------- |
| 01   | mod的使用 |
| 02   | 常用操作  |

```mysql
select * from cinema
where mod(id,2)=1 and description != 'boring'
order by rating desc

#
select * from cinema
where id%2=1 and description != 'boring'
order by rating desc
```

## 627.交换工资(3)

- 题目

```
create table if not exists salary(id int, name varchar(100), sex char(1), salary int)
Truncate table salary
insert into salary (id, name, sex, salary) values ('1', 'A', 'm', '2500')
insert into salary (id, name, sex, salary) values ('2', 'B', 'f', '1500')
insert into salary (id, name, sex, salary) values ('3', 'C', 'm', '5500')
insert into salary (id, name, sex, salary) values ('4', 'D', 'f', '500')
给定一个 salary 表，如下所示，有 m = 男性 和 f = 女性 的值。
交换所有的 f 和 m 值（例如，将所有 f 值更改为 m，反之亦然）。
要求只使用一个更新（Update）语句，并且没有中间的临时表。

注意，您必只能写一个 Update 语句，请不要编写任何 Select 语句。
例如：
| id | name | sex | salary |
|----|------|-----|--------|
| 1  | A    | m   | 2500   |
| 2  | B    | f   | 1500   |
| 3  | C    | m   | 5500   |
| 4  | D    | f   | 500    |
运行你所编写的更新语句之后，将会得到以下表:
| id | name | sex | salary |
|----|------|-----|--------|
| 1  | A    | f   | 2500   |
| 2  | B    | m   | 1500   |
| 3  | C    | f   | 5500   |
| 4  | D    | m   | 500    |
```

- 解题思路

| No.  | 思路             |
| ---- | ---------------- |
| 01   | update+case      |
| 02   | update+if        |
| 03   | update+ascii互转 |

```mysql
update salary
set sex=
CASE sex
when 'm' then 'f'
else 'm'
END

#
update salary
set sex=if(sex='f','m','f')

#
update salary
set sex=char(ascii('m')+ascii('f')-ascii(sex))
```



