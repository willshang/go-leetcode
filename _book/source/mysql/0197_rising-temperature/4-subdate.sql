select A.Id as "Id"
from Weather A join Weather B
on A.Temperature > B.Temperature and B.RecordDate = subdate(A.RecordDate, 1)