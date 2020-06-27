# leetcode1179_重新格式化部门表
select id,
sum(case month when 'Jan' then revenue end) Jan_Revenue,
sum(case month when 'Feb' then revenue end) Feb_Revenue,
sum(case month when 'Mar' then revenue end) Mar_Revenue,
sum(case month when 'Apr' then revenue end) Apr_Revenue,
sum(case month when 'May' then revenue end) May_Revenue,
sum(case month when 'Jun' then revenue end) Jun_Revenue,
sum(case month when 'Jul' then revenue end) Jul_Revenue,
sum(case month when 'Aug' then revenue end) Aug_Revenue,
sum(case month when 'Sep' then revenue end) Sep_Revenue,
sum(case month when 'Oct' then revenue end) Oct_Revenue,
sum(case month when 'Nov' then revenue end) Nov_Revenue,
sum(case month when 'Dec' then revenue end) Dec_Revenue
from department
group by id;