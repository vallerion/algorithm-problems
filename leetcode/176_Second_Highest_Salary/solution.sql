/*
Input:
Employee table:
+----+--------+
| id | salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+
Output:
+---------------------+
| SecondHighestSalary |
+---------------------+
| 200                 |
+---------------------+
 */

select
    Employee.salary as SecondHighestSalary
from (select max(Employee.salary) as salary from Employee) as emax
left join Employee on Employee.salary <= emax.salary
order by Employee.salary desc
limit 1