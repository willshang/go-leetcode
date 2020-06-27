select id,
sum(if(`month` = 'Jan', revenue, null)) as "Jan_Revenue",
sum(if(`month` = 'Feb', revenue, null)) as "Feb_Revenue",
sum(if(`month` = 'Mar', revenue, null)) as "Mar_Revenue",
sum(if(`month` = 'Apr', revenue, null)) as "Apr_Revenue",
sum(if(`month` = 'May', revenue, null)) as "May_Revenue",
sum(if(`month` = 'Jun', revenue, null)) as "Jun_Revenue",
sum(if(`month` = 'Jul', revenue, null)) as "Jul_Revenue",
sum(if(`month` = 'Aug', revenue, null)) as "Aug_Revenue",
sum(if(`month` = 'Sep', revenue, null)) as "Sep_Revenue",
sum(if(`month` = 'Oct', revenue, null)) as "Oct_Revenue",
sum(if(`month` = 'Nov', revenue, null)) as "Nov_Revenue",
sum(if(`month` = 'Dec', revenue, null)) as "Dec_Revenue"
from Department
group by id;