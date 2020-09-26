select A.Id as "Id"
from Weather A join Weather B
on unix_timestamp(A.RecordDate) = unix_timestamp(B.RecordDate) + 86400
and A.Temperature > B.Temperature