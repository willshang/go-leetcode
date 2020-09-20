CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  RETURN (
      # Write your MySQL query statement below.
      select distinct e.Salary from Employee e
      where (select count(distinct Salary) from Employee where Salary > e.Salary) = N-1
  );
END