# Write your MySQL query statement below
select
    id,
    case
        when tree.id in (select atree.id from tree atree where atree.p_id is null)
            then 'Root'
        when tree.id in (select atree.p_id from tree atree)
            then 'Inner'
        else 'Leaf'
    end as type
from tree
order by id;