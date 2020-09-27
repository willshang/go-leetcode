mysql

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
    * [1179.重新格式化部门表(2)](#1179重新格式化部门表2)

    

# Easy

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

## 1179.重新格式化部门表(2)

- 题目

```
部门表 Department：
Create table If Not Exists Department (id int, revenue int, month varchar(5))
Truncate table Department
insert into Department (id, revenue, month) values ('1', '8000', 'Jan')
insert into Department (id, revenue, month) values ('2', '9000', 'Jan')
insert into Department (id, revenue, month) values ('3', '10000', 'Feb')
insert into Department (id, revenue, month) values ('1', '7000', 'Feb')
insert into Department (id, revenue, month) values ('1', '6000', 'Mar')
+---------------+---------+
| Column Name   | Type    |
+---------------+---------+
| id            | int     |
| revenue       | int     |
| month         | varchar |
+---------------+---------+
(id, month) 是表的联合主键。
这个表格有关于每个部门每月收入的信息。
月份（month）可以取下列值 ["Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"]。
编写一个 SQL 查询来重新格式化表，使得新的表中有一个部门 id 列和一些对应 每个月 的收入（revenue）列。
查询结果格式如下面的示例所示：
Department 表：
+------+---------+-------+
| id   | revenue | month |
+------+---------+-------+
| 1    | 8000    | Jan   |
| 2    | 9000    | Jan   |
| 3    | 10000   | Feb   |
| 1    | 7000    | Feb   |
| 1    | 6000    | Mar   |
+------+---------+-------+
查询得到的结果表：
+------+-------------+-------------+-------------+-----+-------------+
| id   | Jan_Revenue | Feb_Revenue | Mar_Revenue | ... | Dec_Revenue |
+------+-------------+-------------+-------------+-----+-------------+
| 1    | 8000        | 7000        | 6000        | ... | null        |
| 2    | 9000        | null        | null        | ... | null        |
| 3    | null        | 10000       | null        | ... | null        |
+------+-------------+-------------+-------------+-----+-------------+
注意，结果表有 13 列 (1个部门 id 列 + 12个月份的收入列)。
```

- 解题思路

| No.  | 思路                                |
| ---- | ----------------------------------- |
| 01   | group by+case when then end+sum/max |
| 02   | group by+if+sum/max                 |

```mysql
select id,
sum(case month when 'Jan' then revenue end) Jan_Revenue,
sum(case month when 'Feb' then revenue end) Feb_Revenue,
sum(case month when 'Mar' then revenue end) Mar_Revenue,
sum(case month when 'Apr' then revenue end) Apr_Revenue,
sum(case month when 'May' then revenue end) May_Revenue,
sum(case month when 'Jun' then revenue end) Jun_Revenue,
sum(case month when 'Jul' then revenue end) Jul_Revenue,
sum(case month when 'Aug' then revenue end) Aug_Revenue,
sum(case month when 'Sep' then revenue end) Sep_Revenue,
sum(case month when 'Oct' then revenue end) Oct_Revenue,
sum(case month when 'Nov' then revenue end) Nov_Revenue,
sum(case month when 'Dec' then revenue end) Dec_Revenue
from department
group by id;

#
select id,
sum(if(`month` = 'Jan', revenue, null)) as "Jan_Revenue",
sum(if(`month` = 'Feb', revenue, null)) as "Feb_Revenue",
sum(if(`month` = 'Mar', revenue, null)) as "Mar_Revenue",
sum(if(`month` = 'Apr', revenue, null)) as "Apr_Revenue",
sum(if(`month` = 'May', revenue, null)) as "May_Revenue",
sum(if(`month` = 'Jun', revenue, null)) as "Jun_Revenue",
sum(if(`month` = 'Jul', revenue, null)) as "Jul_Revenue",
sum(if(`month` = 'Aug', revenue, null)) as "Aug_Revenue",
sum(if(`month` = 'Sep', revenue, null)) as "Sep_Revenue",
sum(if(`month` = 'Oct', revenue, null)) as "Oct_Revenue",
sum(if(`month` = 'Nov', revenue, null)) as "Nov_Revenue",
sum(if(`month` = 'Dec', revenue, null)) as "Dec_Revenue"
from Department
group by id;
```

# Medium

## 177.第N高的薪水(2)

- 题目

```
编写一个 SQL 查询，获取 Employee 表中第 n 高的薪水（Salary）。
+----+--------+
| Id | Salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+
例如上述 Employee 表，n = 2 时，应返回第二高的薪水 200。如果不存在第 n 高的薪水，那么查询应返回 null。
+------------------------+
| getNthHighestSalary(2) |
+------------------------+
| 200                    |
+------------------------+
```

- 解题思路

| No.  | 思路     |
| ---- | -------- |
| 01   | 单表查询 |
| 02   | 子查询   |

```mysql
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  SET N := N-1;
  RETURN (
      # Write your MySQL query statement below.
      select Salary from Employee group by Salary order by Salary desc limit N, 1
  );
END

# 2
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  RETURN (
      # Write your MySQL query statement below.
      select distinct e.Salary from Employee e
      where (select count(distinct Salary) from Employee where Salary > e.Salary) = N-1
  );
END
```

## 178.分数排名(2)

- 题目

```
SQL架构
Create table If Not Exists Scores (Id int, Score DECIMAL(3,2))
Truncate table Scores
insert into Scores (Id, Score) values ('1', '3.5')
insert into Scores (Id, Score) values ('2', '3.65')
insert into Scores (Id, Score) values ('3', '4.0')
insert into Scores (Id, Score) values ('4', '3.85')
insert into Scores (Id, Score) values ('5', '4.0')
insert into Scores (Id, Score) values ('6', '3.65')

编写一个 SQL 查询来实现分数排名。
如果两个分数相同，则两个分数排名（Rank）相同。请注意，平分后的下一个名次应该是下一个连续的整数值。
换句话说，名次之间不应该有“间隔”。
+----+-------+
| Id | Score |
+----+-------+
| 1  | 3.50  |
| 2  | 3.65  |
| 3  | 4.00  |
| 4  | 3.85  |
| 5  | 4.00  |
| 6  | 3.65  |
+----+-------+
例如，根据上述给定的 Scores 表，你的查询应该返回（按分数从高到低排列）：
+-------+------+
| Score | Rank |
+-------+------+
| 4.00  | 1    |
| 4.00  | 1    |
| 3.85  | 2    |
| 3.65  | 3    |
| 3.65  | 3    |
| 3.50  | 4    |
+-------+------+
重要提示：对于 MySQL 解决方案，如果要转义用作列名的保留字，可以在关键字之前和之后使用撇号。例如 `Rank`
```

- 解题思路

| No.  | 思路       |
| ---- | ---------- |
| 01   | 查询       |
| 02   | dense_rank |

```sql
select a.Score as Score,
(select count(distinct b.Score) from Scores b where b.Score >= a.Score) as `Rank`
from Scores a
order by a.Score desc

# 2
# OVER() 指定分析函数工作的数据窗口大小,这个数据窗口大小可能会随着行的变化而变化
# RANK() 排序相同时会重复,总数不会变
# DENSE_RANK() 排序相同时会重复,总数会减少
# ROW_NUMBER() 会根据顺序计算
select Score, dense_rank() over(order by Score desc) as `Rank`
from Scores
```

## 180.连续出现的数字(1)

- 题目

```
SQL架构
Create table If Not Exists Logs (Id int, Num int)
Truncate table Logs
insert into Logs (Id, Num) values ('1', '1')
insert into Logs (Id, Num) values ('2', '1')
insert into Logs (Id, Num) values ('3', '1')
insert into Logs (Id, Num) values ('4', '2')
insert into Logs (Id, Num) values ('5', '1')
insert into Logs (Id, Num) values ('6', '2')
insert into Logs (Id, Num) values ('7', '2')
编写一个 SQL 查询，查找所有至少连续出现三次的数字。
+----+-----+
| Id | Num |
+----+-----+
| 1  |  1  |
| 2  |  1  |
| 3  |  1  |
| 4  |  2  |
| 5  |  1  |
| 6  |  2  |
| 7  |  2  |
+----+-----+
例如，给定上面的 Logs 表， 1 是唯一连续出现至少三次的数字。
+-----------------+
| ConsecutiveNums |
+-----------------+
| 1               |
+-----------------+
```

- 解题思路

| No.  | 思路    |
| ---- | ------- |
| 01   | 3表连接 |

```sql
select distinct l1.Num as ConsecutiveNums
from Logs l1, Logs l2, Logs l3
where l1.Id = l2.Id-1 and l2.Id = l3.Id-1 and l1.Num = l2.Num and l2.Num = l3.Num
```

## 184.部门工资最高的员工(2)

- 题目

```
SQL架构
Create table If Not Exists Employee (Id int, Name varchar(255), Salary int, DepartmentId int)
Create table If Not Exists Department (Id int, Name varchar(255))
Truncate table Employee
insert into Employee (Id, Name, Salary, DepartmentId) values ('1', 'Joe', '70000', '1')
insert into Employee (Id, Name, Salary, DepartmentId) values ('2', 'Jim', '90000', '1')
insert into Employee (Id, Name, Salary, DepartmentId) values ('3', 'Henry', '80000', '2')
insert into Employee (Id, Name, Salary, DepartmentId) values ('4', 'Sam', '60000', '2')
insert into Employee (Id, Name, Salary, DepartmentId) values ('5', 'Max', '90000', '1')
Truncate table Department
insert into Department (Id, Name) values ('1', 'IT')
insert into Department (Id, Name) values ('2', 'Sales')

Employee 表包含所有员工信息，每个员工有其对应的 Id, salary 和 department Id。
+----+-------+--------+--------------+
| Id | Name  | Salary | DepartmentId |
+----+-------+--------+--------------+
| 1  | Joe   | 70000  | 1            |
| 2  | Jim   | 90000  | 1            |
| 3  | Henry | 80000  | 2            |
| 4  | Sam   | 60000  | 2            |
| 5  | Max   | 90000  | 1            |
+----+-------+--------+--------------+

Department 表包含公司所有部门的信息。
+----+----------+
| Id | Name     |
+----+----------+
| 1  | IT       |
| 2  | Sales    |
+----+----------+

编写一个 SQL 查询，找出每个部门工资最高的员工。
对于上述表，您的 SQL 查询应返回以下行（行的顺序无关紧要）。
+------------+----------+--------+
| Department | Employee | Salary |
+------------+----------+--------+
| IT         | Max      | 90000  |
| IT         | Jim      | 90000  |
| Sales      | Henry    | 80000  |
+------------+----------+--------+
解释：Max 和 Jim 在 IT 部门的工资都是最高的，Henry 在销售部的工资最高。
```

- 解题思路

| No.  | 思路    |
| ---- | ------- |
| 01   | join+in |
| 02   | 子查询  |

```sql
select Department.Name as `Department`, Employee.Name as `Employee`, Salary
from Employee join Department on Employee.DepartmentId = Department.Id
where (Employee.DepartmentId, Salary) in
(select DepartmentId, max(Salary) from Employee group by DepartmentId)

# 2
select Department.Name as `Department`, Employee.Name as `Employee`, Salary
from Employee join Department on Employee.DepartmentId = Department.Id
where Employee.Salary = (
select max(Salary) from Employee where Employee.DepartmentId = Department.Id
)
```

## 626.换座位(2)

- 题目

```
SQL架构
Create table If Not Exists seat(id int, student varchar(255))
Truncate table seat
insert into seat (id, student) values ('1', 'Abbot')
insert into seat (id, student) values ('2', 'Doris')
insert into seat (id, student) values ('3', 'Emerson')
insert into seat (id, student) values ('4', 'Green')
insert into seat (id, student) values ('5', 'Jeames')

小美是一所中学的信息科技老师，她有一张 seat 座位表，平时用来储存学生名字和与他们相对应的座位 id。
其中纵列的 id 是连续递增的
小美想改变相邻俩学生的座位。
你能不能帮她写一个 SQL query 来输出小美想要的结果呢？
示例：
+---------+---------+
|    id   | student |
+---------+---------+
|    1    | Abbot   |
|    2    | Doris   |
|    3    | Emerson |
|    4    | Green   |
|    5    | Jeames  |
+---------+---------+
假如数据输入的是上表，则输出结果如下：
+---------+---------+
|    id   | student |
+---------+---------+
|    1    | Doris   |
|    2    | Abbot   |
|    3    | Green   |
|    4    | Emerson |
|    5    | Jeames  |
+---------+---------+
注意：如果学生人数是奇数，则不需要改变最后一个同学的座位。
```

- 解题思路

| No.  | 思路      |
| ---- | --------- |
| 01   | case when |
| 02   | coalesce  |

```sql
select (case
    when mod(id,2) !=0 and total != id then id+1
    when mod(id,2) !=0 and total = id then id
    else id-1 end) as id, student
from  seat, ( select count(*) as total from seat) as seat_counts
order by id asc 

# 2
select s1.id, coalesce(s2.student, s1.student) AS student
from seat s1 left join seat s2 on ((s1.id+1)^1)-1=s2.id
order by s1.id
```

# Hard

## 185.部门工资前三高的所有员工(2)

- 题目

```
SQL架构
Create table If Not Exists Employee (Id int, Name varchar(255), Salary int, DepartmentId int)
Create table If Not Exists Department (Id int, Name varchar(255))
Truncate table Employee
insert into Employee (Id, Name, Salary, DepartmentId) values ('1', 'Joe', '85000', '1')
insert into Employee (Id, Name, Salary, DepartmentId) values ('2', 'Henry', '80000', '2')
insert into Employee (Id, Name, Salary, DepartmentId) values ('3', 'Sam', '60000', '2')
insert into Employee (Id, Name, Salary, DepartmentId) values ('4', 'Max', '90000', '1')
insert into Employee (Id, Name, Salary, DepartmentId) values ('5', 'Janet', '69000', '1')
insert into Employee (Id, Name, Salary, DepartmentId) values ('6', 'Randy', '85000', '1')
insert into Employee (Id, Name, Salary, DepartmentId) values ('7', 'Will', '70000', '1')
Truncate table Department
insert into Department (Id, Name) values ('1', 'IT')
insert into Department (Id, Name) values ('2', 'Sales')
Employee 表包含所有员工信息，每个员工有其对应的工号 Id，
姓名 Name，工资 Salary 和部门编号 DepartmentId 。
+----+-------+--------+--------------+
| Id | Name  | Salary | DepartmentId |
+----+-------+--------+--------------+
| 1  | Joe   | 85000  | 1            |
| 2  | Henry | 80000  | 2            |
| 3  | Sam   | 60000  | 2            |
| 4  | Max   | 90000  | 1            |
| 5  | Janet | 69000  | 1            |
| 6  | Randy | 85000  | 1            |
| 7  | Will  | 70000  | 1            |
+----+-------+--------+--------------+
Department 表包含公司所有部门的信息。
+----+----------+
| Id | Name     |
+----+----------+
| 1  | IT       |
| 2  | Sales    |
+----+----------+
编写一个 SQL 查询，找出每个部门获得前三高工资的所有员工。例如，根据上述给定的表，查询结果应返回：
+------------+----------+--------+
| Department | Employee | Salary |
+------------+----------+--------+
| IT         | Max      | 90000  |
| IT         | Randy    | 85000  |
| IT         | Joe      | 85000  |
| IT         | Will     | 70000  |
| Sales      | Henry    | 80000  |
| Sales      | Sam      | 60000  |
+------------+----------+--------+
解释：IT 部门中，Max 获得了最高的工资，Randy 和 Joe 都拿到了第二高的工资，Will 的工资排第三。
销售部门（Sales）只有两名员工，Henry 的工资最高，Sam 的工资排第二。
```

- 解题思路

| No.  | 思路           |
| ---- | -------------- |
| 01   | join+子查询    |
| 02   | dense_rank函数 |

```sql
select Department.Name as `Department`, Employee.Name as `Employee`, Salary
from Employee join Department on Employee.DepartmentId = Department.Id
where 3 >
(select count(distinct e2.Salary) from Employee e2 
where e2.Salary > Employee.Salary and e2.DepartmentId = Employee.DepartmentId )

# 2
select d.name AS `Department`,a.name AS `Employee`,a.Salary
from(select *,dense_rank() over(partition by DepartmentId order by salary desc) as ranking from Employee) as a,
department as d
where a.DepartmentId = d.Id and ranking<=3
```

## 262.行程和用户(2)

- 题目

```
SQL架构
Create table If Not Exists Trips (Id int, Client_Id int, Driver_Id int, City_Id int, Status ENUM('completed', 'cancelled_by_driver', 'cancelled_by_client'), Request_at varchar(50))
Create table If Not Exists Users (Users_Id int, Banned varchar(50), Role ENUM('client', 'driver', 'partner'))
Truncate table Trips
insert into Trips (Id, Client_Id, Driver_Id, City_Id, Status, Request_at) values ('1', '1', '10', '1', 'completed', '2013-10-01')
insert into Trips (Id, Client_Id, Driver_Id, City_Id, Status, Request_at) values ('2', '2', '11', '1', 'cancelled_by_driver', '2013-10-01')
insert into Trips (Id, Client_Id, Driver_Id, City_Id, Status, Request_at) values ('3', '3', '12', '6', 'completed', '2013-10-01')
insert into Trips (Id, Client_Id, Driver_Id, City_Id, Status, Request_at) values ('4', '4', '13', '6', 'cancelled_by_client', '2013-10-01')
insert into Trips (Id, Client_Id, Driver_Id, City_Id, Status, Request_at) values ('5', '1', '10', '1', 'completed', '2013-10-02')
insert into Trips (Id, Client_Id, Driver_Id, City_Id, Status, Request_at) values ('6', '2', '11', '6', 'completed', '2013-10-02')
insert into Trips (Id, Client_Id, Driver_Id, City_Id, Status, Request_at) values ('7', '3', '12', '6', 'completed', '2013-10-02')
insert into Trips (Id, Client_Id, Driver_Id, City_Id, Status, Request_at) values ('8', '2', '12', '12', 'completed', '2013-10-03')
insert into Trips (Id, Client_Id, Driver_Id, City_Id, Status, Request_at) values ('9', '3', '10', '12', 'completed', '2013-10-03')
insert into Trips (Id, Client_Id, Driver_Id, City_Id, Status, Request_at) values ('10', '4', '13', '12', 'cancelled_by_driver', '2013-10-03')
Truncate table Users
insert into Users (Users_Id, Banned, Role) values ('1', 'No', 'client')
insert into Users (Users_Id, Banned, Role) values ('2', 'Yes', 'client')
insert into Users (Users_Id, Banned, Role) values ('3', 'No', 'client')
insert into Users (Users_Id, Banned, Role) values ('4', 'No', 'client')
insert into Users (Users_Id, Banned, Role) values ('10', 'No', 'driver')
insert into Users (Users_Id, Banned, Role) values ('11', 'No', 'driver')
insert into Users (Users_Id, Banned, Role) values ('12', 'No', 'driver')
insert into Users (Users_Id, Banned, Role) values ('13', 'No', 'driver')

Trips 表中存所有出租车的行程信息。每段行程有唯一键 Id，
Client_Id 和 Driver_Id 是 Users 表中 Users_Id 的外键。
Status 是枚举类型，枚举成员为 (‘completed’, ‘cancelled_by_driver’, ‘cancelled_by_client’)。
+----+-----------+-----------+---------+--------------------+----------+
| Id | Client_Id | Driver_Id | City_Id |        Status      |Request_at|
+----+-----------+-----------+---------+--------------------+----------+
| 1  |     1     |    10     |    1    |     completed      |2013-10-01|
| 2  |     2     |    11     |    1    | cancelled_by_driver|2013-10-01|
| 3  |     3     |    12     |    6    |     completed      |2013-10-01|
| 4  |     4     |    13     |    6    | cancelled_by_client|2013-10-01|
| 5  |     1     |    10     |    1    |     completed      |2013-10-02|
| 6  |     2     |    11     |    6    |     completed      |2013-10-02|
| 7  |     3     |    12     |    6    |     completed      |2013-10-02|
| 8  |     2     |    12     |    12   |     completed      |2013-10-03|
| 9  |     3     |    10     |    12   |     completed      |2013-10-03| 
| 10 |     4     |    13     |    12   | cancelled_by_driver|2013-10-03|
+----+-----------+-----------+---------+--------------------+----------+
Users 表存所有用户。每个用户有唯一键 Users_Id。Banned 表示这个用户是否被禁止，
Role 则是一个表示（‘client’, ‘driver’, ‘partner’）的枚举类型。
+----------+--------+--------+
| Users_Id | Banned |  Role  |
+----------+--------+--------+
|    1     |   No   | client |
|    2     |   Yes  | client |
|    3     |   No   | client |
|    4     |   No   | client |
|    10    |   No   | driver |
|    11    |   No   | driver |
|    12    |   No   | driver |
|    13    |   No   | driver |
+----------+--------+--------+
写一段 SQL 语句查出 2013年10月1日 至 2013年10月3日 期间非禁止用户的取消率。
基于上表，你的 SQL 语句应返回如下结果，取消率（Cancellation Rate）保留两位小数。
取消率的计算方式如下：(被司机或乘客取消的非禁止用户生成的订单数量) / (非禁止用户生成的订单总数)
+------------+-------------------+
|     Day    | Cancellation Rate |
+------------+-------------------+
| 2013-10-01 |       0.33        |
| 2013-10-02 |       0.00        |
| 2013-10-03 |       0.50        |
+------------+-------------------+
致谢:非常感谢 @cak1erlizhou 详细的提供了这道题和相应的测试用例。
```

- 解题思路

| No.  | 思路 |
| ---- | ---- |
| 01   | 聚合 |
| 02   | 聚合 |

```sql
select t.Request_at as `Day`, Round(sum(if(t.Status='completed',0,1))/count(t.Status),2) as `Cancellation Rate`
from Trips as t
join Users as u1 on t.Client_Id=u1.Users_Id and u1.Banned='No'
join Users as u2 on t.Driver_Id=u2.Users_Id and u2.Banned='No'
where t.Request_at between '2013-10-01' and '2013-10-03'
group by t.Request_at

# 2
select t.Request_at as `Day`, Round(avg(t.Status!='completed'),2) as `Cancellation Rate`
from Trips as t
join Users as u1 on t.Client_Id=u1.Users_Id and u1.Banned='No'
join Users as u2 on t.Driver_Id=u2.Users_Id and u2.Banned='No'
where t.Request_at between '2013-10-01' and '2013-10-03'
group by t.Request_at
```

## 601.体育馆的人流量(2)

- 题目

```
SQL架构
Create table If Not Exists stadium (id int, visit_date DATE NULL, people int)
Truncate table stadium
insert into stadium (id, visit_date, people) values ('1', '2017-01-01', '10')
insert into stadium (id, visit_date, people) values ('2', '2017-01-02', '109')
insert into stadium (id, visit_date, people) values ('3', '2017-01-03', '150')
insert into stadium (id, visit_date, people) values ('4', '2017-01-04', '99')
insert into stadium (id, visit_date, people) values ('5', '2017-01-05', '145')
insert into stadium (id, visit_date, people) values ('6', '2017-01-06', '1455')
insert into stadium (id, visit_date, people) values ('7', '2017-01-07', '199')
insert into stadium (id, visit_date, people) values ('8', '2017-01-08', '188')

X 市建了一个新的体育馆，每日人流量信息被记录在这三列信息中：序号 (id)、日期 (visit_date)、
人流量 (people)。
请编写一个查询语句，找出人流量的高峰期。高峰期时，至少连续三行记录中的人流量不少于100。
例如，表 stadium：
+------+------------+-----------+
| id   | visit_date | people    |
+------+------------+-----------+
| 1    | 2017-01-01 | 10        |
| 2    | 2017-01-02 | 109       |
| 3    | 2017-01-03 | 150       |
| 4    | 2017-01-04 | 99        |
| 5    | 2017-01-05 | 145       |
| 6    | 2017-01-06 | 1455      |
| 7    | 2017-01-07 | 199       |
| 8    | 2017-01-08 | 188       |
+------+------------+-----------+
对于上面的示例数据，输出为：
+------+------------+-----------+
| id   | visit_date | people    |
+------+------------+-----------+
| 5    | 2017-01-05 | 145       |
| 6    | 2017-01-06 | 1455      |
| 7    | 2017-01-07 | 199       |
| 8    | 2017-01-08 | 188       |
+------+------------+-----------+
提示：每天只有一行记录，日期随着 id 的增加而增加。
    体育馆并不是每天都开放的，所以记录中的日期可能会出现断层。
```

- 解题思路

| No.  | 思路 |
| ---- | ---- |
| 01   | 连接 |
| 02   | 连接 |

```sql
select distinct t1.*
from stadium t1, stadium t2, stadium t3
where t1.people >= 100 and t2.people >= 100 and t3.people >= 100 and
(
    (t1.id-t2.id=1 and t1.id-t3.id=2 and t2.id-t3.id=1) or
    (t2.id-t1.id=1 and t2.id-t3.id=2 and t1.id-t3.id=1) or
    (t3.id-t2.id=1 and t3.id-t1.id=2 and t2.id-t1.id=1)
)
order by t1.id

# 2
select distinct stadium.id, visit_date, people
from stadium,(select t1.id from stadium t1,stadium t2,stadium t3 where 
    t1.id - t2.id = 1 and 
    t2.id - t3.id = 1 and 
    t1.people >=100 and t2.people >=100 and t3.people >=100
) t
where (t.id - stadium.id) between 0 and 2
```

