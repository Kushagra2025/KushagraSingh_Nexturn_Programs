SELECT 
    e.EmployeeID, 
    e.Name, 
    e.Salary, 
    e.HireDate, 
    d.DepartmentName
FROM 
    Employees e
JOIN 
    Departments d ON e.DepartmentID = d.DepartmentID;
