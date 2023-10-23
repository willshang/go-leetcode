
# leetcode185_部门工资前三高的所有员工
select Department.Name as `Department`, Employee.Name as `Employee`, Salary
from Employee join Department on Employee.DepartmentId = Department.Id
where 3 >
(select count(distinct e2.Salary) from Employee e2
where e2.Salary > Employee.Salary and e2.DepartmentId = Employee.DepartmentId )