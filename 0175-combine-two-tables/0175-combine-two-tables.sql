-- Write your PostgreSQL query statement below
SELECT u.firstName,u.lastName, a.city, a.state FROM Person u LEFT JOIN Address a ON u.personId=a.personId;
