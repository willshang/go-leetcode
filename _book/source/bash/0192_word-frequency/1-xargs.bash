#


# leetcode192_统计词频
# xargs 分割字符串 -n 1表示每行输出一个
# uniq -c 统计重复次数
# sort -r 降序排序 -n 以数字排序(默认字符)
cat words.txt | xargs -n 1 | sort | uniq -c | sort -nr | awk '{print $2" "$1}'