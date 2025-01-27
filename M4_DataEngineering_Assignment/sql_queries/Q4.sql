SELECT 
    d.DepartmentName
FROM 
    Departments d
LEFT JOIN 
    Employees e ON d.DepartmentID = e.DepartmentID
WHERE 
    e.EmployeeID IS NULL;
