select Email from
(
    select Email, count(Email) as num
    from Person
    Group by Email
) as temp_table
where num > 1;