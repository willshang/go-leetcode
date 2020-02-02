select name Customers
from customers c
where not exists (
    select 1 from orders o
    where o.customerid=c.id
)