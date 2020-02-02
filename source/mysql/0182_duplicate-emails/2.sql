# leetcode 182 查找重复的电子邮箱
select Email
from Person
group by Email
having count(Email) > 1;