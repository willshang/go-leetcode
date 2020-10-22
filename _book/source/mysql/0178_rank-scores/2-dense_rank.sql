# OVER() 指定分析函数工作的数据窗口大小,这个数据窗口大小可能会随着行的变化而变化
# RANK() 排序相同时会重复,总数不会变
# DENSE_RANK() 排序相同时会重复,总数会减少
# ROW_NUMBER() 会根据顺序计算
select Score, dense_rank() over(order by Score desc) as `Rank`
from Scores