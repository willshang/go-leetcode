# Bash

* [Bash](#bash)
    * [193.有效电话号码(4)](#193有效电话号码4)
    * [195. 第十行(4)](#195-第十行4)

# Easy


## 193.有效电话号码(4)

- 题目

```
给定一个包含电话号码列表（一行一个电话号码）的文本文件 file.txt，
写一个 bash 脚本输出所有有效的电话号码。

你可以假设一个有效的电话号码必须满足以下两种格式： 
(xxx) xxx-xxxx 或 xxx-xxx-xxxx。（x 表示一个数字）

你也可以假设每行前后没有多余的空格字符。
示例:
假设 file.txt 内容如下：
987-123-4567
123 456 7890
(123) 456-7890

你的脚本应当输出下列有效的电话号码：
987-123-4567
(123) 456-7890
```

- 解题思路分析

| No.  | 思路                |
| ---- | ------------------- |
| 01   | 正则表达式_cat_grep |
| 02   | 正则表达式_grep     |
| 03   | 正则表达式_awk      |
| 04   | 正则表达式_grep     |

```bash
# 正则表达式_cat_grep
# (xxx) xxx-xxxx 或 xxx-xxx-xxxx
# ^\([0-9]{3}\) [0-9]{3}-[0-9]{4}$ (xxx) xxx-xxxx
# ^[0-9]{3}-[0-9]{3}-[0-9]{4}$ xxx-xxx-xxxx
# grep -P 匹配正则
cat file.txt | grep -P "^\([0-9]{3}\) [0-9]{3}-[0-9]{4}$|^[0-9]{3}-[0-9]{3}-[0-9]{4}$"

# (xxx) xxx-xxxx 或 xxx-xxx-xxxx
grep -P "^([0-9]{3}-|\([0-9]{3}\) )[0-9]{3}-[0-9]{4}$" file.txt
grep -E "^([0-9]{3}-|\([0-9]{3}\) )[0-9]{3}-[0-9]{4}$" file.txt

#
awk "/^([0-9]{3}-|\([0-9]{3}\) )[0-9]{3}-[0-9]{4}$/" file.txt

#
grep -P "^(\d{3}-|\(\d{3}\) )\d{3}-\d{4}$" file.txt
```

## 195. 第十行(4)

- 题目

```
给定一个文本文件 file.txt，请只打印这个文件中的第十行。

示例:
假设 file.txt 有如下内容：
Line 1
Line 2
Line 3
Line 4
Line 5
Line 6
Line 7
Line 8
Line 9
Line 10

你的脚本应当显示第十行：
Line 10
说明:
1. 如果文件少于十行，你应当输出什么？
2. 至少有三种不同的解法，请尝试尽可能多的方法来解题。
```

- 解题思路分析

| No.  | 思路      |
| ---- | --------- |
| 01   | awk NR    |
| 02   | tail head |
| 03   | sed       |

```bash
# 
awk 'NR==10' file.txt

#
tail -n +10 file.txt | head -1

#
sed -n 10p file.txt
```

# Medium

## 192.统计词频(2)

- 题目

```
写一个 bash 脚本以统计一个文本文件 words.txt 中每个单词出现的频率。
为了简单起见，你可以假设：
    words.txt只包括小写字母和 ' ' 。
    每个单词只由小写字母组成。
    单词间由一个或多个空格字符分隔。
示例:假设 words.txt 内容如下：
the day is sunny the the
the sunny is is
你的脚本应当输出（以词频降序排列）：
the 4
is 3
sunny 2
day 1
说明:不要担心词频相同的单词的排序问题，每个单词出现的频率都是唯一的。
    你可以使用一行 Unix pipes 实现吗？
```

- 解题思路

| No.  | 思路  |
| ---- | ----- |
| 01   | xargs |
| 02   | tr    |

```bash
# xargs 分割字符串 -n 1表示每行输出一个
# uniq -c 统计重复次数
# sort -r 降序排序 -n 以数字排序(默认字符)
cat words.txt | xargs -n 1 | sort | uniq -c | sort -nr | awk '{print $2" "$1}'


# 2
cat words.txt | tr -s ' ' '\n' | sort | uniq -c | sort -nr | awk '{print $2" "$1}'
```

## 194.转置文件(1)

- 题目

```
给定一个文件 file.txt，转置它的内容。
你可以假设每行列数相同，并且每个字段由 ' ' 分隔.
示例:假设 file.txt 文件内容如下：
name age
alice 21
ryan 30
应当输出：
name alice ryan
age 21 30
```

- 解题思路

| No.  | 思路 |
| ---- | ---- |
| 01   | awk  |

```bash
awk '{
    for (i=1;i<=NF;i++){
        if (NR==1){
            res[i]=$i
        }
        else{
            res[i]=res[i]" "$i
        }
    }
}END{
    for(i=1;i<=NF;i++){
        print res[i]
    }
}' file.txt
```

