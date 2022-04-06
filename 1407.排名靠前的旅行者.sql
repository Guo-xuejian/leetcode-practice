# Write your MySQL query statement below
select distinct name, if(sum(distance) is null,0,sum(distance)) as travelled_distance

from 
    Users as u left join Rides as r
    on u.id = r.user_id

group by name
order by travelled_distance desc, name;