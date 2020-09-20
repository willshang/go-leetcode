
# leetcode262_行程和用户
select t.Request_at as `Day`, Round(avg(t.Status!='completed'),2) as `Cancellation Rate`
from Trips as t
join Users as u1 on t.Client_Id=u1.Users_Id and u1.Banned='No'
join Users as u2 on t.Driver_Id=u2.Users_Id and u2.Banned='No'
where t.Request_at between '2013-10-01' and '2013-10-03'
group by t.Request_at