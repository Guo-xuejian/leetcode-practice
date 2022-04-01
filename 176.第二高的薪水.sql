# Write your MySQL query statement below
# SELECT DISTINCT
#     Salary AS SecondHighestSalary
# FROM
#     Employee
# ORDER BY Salary DESC
# LIMIT 1, 1;
select
    (
        select distinct Salary
        from Employee
        order by Salary desc
        limit 1,1
    ) as SecondHighestSalary;