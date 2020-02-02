## leetcode 176 第二高的薪水
select ifnull(
    (select distinct Salary
    from Employee
    order by Salary desc
    limit 1 offset 1),null
) as SecondHighestSalary;