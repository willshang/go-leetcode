select distinct stadium.id, visit_date, people
from stadium,(select t1.id from stadium t1,stadium t2,stadium t3 where
    t1.id - t2.id = 1 and
    t2.id - t3.id = 1 and
    t1.people >=100 and t2.people >=100 and t3.people >=100
) t
where (t.id - stadium.id) between 0 and 2