
# leetcode626_换座位
select (case
    when mod(id,2) !=0 and total != id then id+1
    when mod(id,2) !=0 and total = id then id
    else id-1 end) as id, student
from  seat, ( select count(*) as total from seat) as seat_counts
order by id asc