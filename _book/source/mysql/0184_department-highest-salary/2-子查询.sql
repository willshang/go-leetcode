select Department.Name as `Department`, Employee.Name as `Employee`, Salary
from Employee join Department on Employee.DepartmentId = Department.Id
where Employee.Salary = (
select max(Salary) from Employee where Employee.DepartmentId = Department.Id
)