const API_BASE_URL = 'http://localhost:8080/api';

document.addEventListener('DOMContentLoaded', () => {
    const depositForm = document.getElementById('deposit-form');
    const depositMessage = document.getElementById('deposit-message');

    const betForm = document.getElementById('bet-form');
    const betResult = document.getElementById('bet-result');

    const refreshLeaderboardBtn = document.getElementById('refresh-leaderboard');
    const leaderboardBody = document.getElementById('leaderboard-body');

    // Handle Deposit
    depositForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const username = document.getElementById('username').value.trim();
        const amount = parseInt(document.getElementById('deposit-amount').value);

        const response = await fetch(`${API_BASE_URL}/user/deposit`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ username, amount })
        });

        const data = await response.json();
        if (response.ok) {
            depositMessage.textContent = `Deposit Successful! New Balance: ${data.balance}`;
            depositMessage.style.color = 'lightgreen';
        } else {
            depositMessage.textContent = `Error: ${data}`;
            depositMessage.style.color = 'lightcoral';
        }
    });

    // Handle Bet
    betForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const username = document.getElementById('bet-username').value.trim();
        const betAmount = parseInt(document.getElementById('bet-amount').value);
        const betColor = document.getElementById('bet-color').value;

        const response = await fetch(`${API_BASE_URL}/roulette/bet`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ username, bet_amount: betAmount, bet_color: betColor })
        });

        const data = await response.json();
        if (response.ok) {
            const resultText = `
                Spin Result: ${data.spin_color.toUpperCase()} |
                Bet Result: ${data.result.toUpperCase()} |
                New Balance: ${data.balance}
            `;
            betResult.textContent = resultText;
            betResult.style.color = data.result === 'win' ? 'lightgreen' : 'lightcoral';
        } else {
            betResult.textContent = `Error: ${data}`;
            betResult.style.color = 'lightcoral';
        }
    });

    // Fetch Leaderboard
    const fetchLeaderboard = async () => {
        const response = await fetch(`${API_BASE_URL}/leaderboard`);
        const data = await response.json();
        if (response.ok) {
            leaderboardBody.innerHTML = '';
            data.forEach((user, index) => {
                const row = document.createElement('tr');
                row.innerHTML = `
                    <td>${index + 1}</td>
                    <td>${user.username}</td>
                    <td>${user.balance}</td>
                `;
                leaderboardBody.appendChild(row);
            });
        } else {
            leaderboardBody.innerHTML = `<tr><td colspan="3">Failed to load leaderboard</td></tr>`;
        }
    };

    // Initial Leaderboard Fetch
    fetchLeaderboard();

    // Refresh Leaderboard on Button Click
    refreshLeaderboardBtn.addEventListener('click', fetchLeaderboard);
});