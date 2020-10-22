
# leetcode184_部门工资最高的员工
select Department.Name as `Department`, Employee.Name as `Employee`, Salary
from Employee join Department on Employee.DepartmentId = Department.Id
where (Employee.DepartmentId, Salary) in
(select DepartmentId, max(Salary) from Employee group by DepartmentId)