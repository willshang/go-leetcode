select name as Employee
from employee a
where salary > (select salary from employee where a.managerid = id);