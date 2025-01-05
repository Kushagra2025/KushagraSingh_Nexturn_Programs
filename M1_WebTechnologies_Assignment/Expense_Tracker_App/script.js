// DOM elements
const form = document.getElementById("expense-form");
const expenseTableBody = document.querySelector("#expense-table tbody");
const summarySection = document.getElementById("summary");
const expenseChart = document.getElementById("expense-chart");
const categories = ['Food', 'Travel', 'Shopping', 'Entertainment', 'Utilities'];

// Initialize expenses from localStorage
let expenses = JSON.parse(localStorage.getItem("expenses")) || [];

// Update the UI
function updateUI() {
    // Update Table
    expenseTableBody.innerHTML = '';
    expenses.forEach((expense, index) => {
        const row = document.createElement("tr");
        row.innerHTML = `
            <td>${expense.amount}</td>
            <td>${expense.description}</td>
            <td>${expense.category}</td>
            <td><button class="delete" onclick="deleteExpense(${index})">Delete</button></td>
        `;
        expenseTableBody.appendChild(row);
    });

    // Update Summary
    const categoryTotals = categories.reduce((totals, category) => {
        totals[category] = expenses
            .filter(expense => expense.category === category)
            .reduce((sum, expense) => sum + parseFloat(expense.amount), 0);
        return totals;
    }, {});

    summarySection.innerHTML = categories.map(category => {
        return `<p>${category}: $${categoryTotals[category].toFixed(2)}</p>`;
    }).join('');

    // Update Chart
    const chartData = categories.map(category => categoryTotals[category]);
    new Chart(expenseChart, {
        type: 'pie',
        data: {
            labels: categories,
            datasets: [{
                data: chartData,
                backgroundColor: ['#FF6384', '#36A2EB', '#FFCE56', '#FF5733', '#28a745'],
            }]
        },
    });
}

// Add expense
form.addEventListener("submit", (e) => {
    e.preventDefault();
    const amount = document.getElementById("amount").value;
    const description = document.getElementById("description").value;
    const category = document.getElementById("category").value;

    if (amount && description && category) {
        expenses.push({ amount, description, category });
        localStorage.setItem("expenses", JSON.stringify(expenses));
        form.reset();
        updateUI();
    }
});

// Delete expense
function deleteExpense(index) {
    expenses.splice(index, 1);
    localStorage.setItem("expenses", JSON.stringify(expenses));
    updateUI();
}

// Initial UI update
updateUI();
