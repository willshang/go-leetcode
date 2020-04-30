# or具有全表扫描机制
# union具有索引列查询速度快
# Union：对两个结果集进行并集操作，不包括重复行，同时进行默认规则的排序；
# Union All：对两个结果集进行并集操作，包括重复行，不进行排序；

select name, population, area from world where area > 3000000
union
select name, population, area from world where population > 25000000