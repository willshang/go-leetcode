# DATEDIFF() 函数返回两个日期之间的天数。
# DATEDIFF('2007-12-30','2007-12-29');
# 1

# leetcode197_上升的温度
select A.Id as "Id"
from Weather A join Weather B
on datediff(A.RecordDate, B.RecordDate) = 1 and A.Temperature > B.Temperature