select d.name AS `Department`,a.name AS `Employee`,a.Salary
from(select *,dense_rank() over(partition by DepartmentId order by salary desc) as ranking from Employee) as a,
department as d
where a.DepartmentId = d.Id and ranking<=3
