# Write your MySQL query statement below
select customer_id, count(customer_id) count_no_trans
from 
    Visits v 
left join 
    Transactions t
on v.visit_id = t.visit_id
where amount is null
group by customer_id;
