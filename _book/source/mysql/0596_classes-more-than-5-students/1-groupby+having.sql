# leetcode596_超过5名学生的课
select class from courses
group by class
having count(distinct student) >= 5