# leetcode 183 从不订购的客户
select Customers.Name as Customers
from Customers
where Customers.Id not in (
    select CustomerId from Orders
);