select A.Id as "Id"
from Weather A join Weather B
on A.RecordDate = adddate(B.RecordDate, 1) and A.Temperature > B.Temperature

select A.Id as "Id"
from Weather A join Weather B
on A.Temperature > B.Temperature and A.RecordDate = adddate(B.RecordDate, 1)
