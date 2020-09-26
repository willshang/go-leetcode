select s1.id, coalesce(s2.student, s1.student) AS student
from seat s1 left join seat s2 on ((s1.id+1)^1)-1=s2.id
order by s1.id