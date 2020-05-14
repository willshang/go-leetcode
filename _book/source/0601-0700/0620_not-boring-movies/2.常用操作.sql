# leetcode620_有趣的电影
select * from cinema
where id%2=1 and description != 'boring'
order by rating desc