# Write your MySQL query statement below
select 
    * 
from 
    patients 
where 
    conditions rlike '^DIAB1|.*\\sDIAB1';
