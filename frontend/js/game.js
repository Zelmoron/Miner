
async function checkToken() {
    try {
        const response = await fetch('http://localhost:8080/jwt/checktoken', {
            method: 'GET',
            credentials: 'include', // Добавлено для отправки куки
        });

        if (!response.ok) {
            // Если статус ответа не OK, выбрасываем ошибку
            throw new Error('Token is invalid or missing');
        }

        const data = await response.json();
        console.log('Token is valid:', data);
        
        // Обновляем имя пользователя в интерфейсе
        document.getElementById('username').textContent = data.name || "Пользователь"; // Устанавливаем имя пользователя

    } catch (error) {
        console.error('Error checking token:', error);
        // Если произошла ошибка при проверке токена, вызываем функцию обновления токена
        await refreshToken();
    }
}

// Функция для обновления токена
async function refreshToken() {
    try {
        const response = await fetch('http://localhost:8080/refresh/refresh', {
            method: 'GET',
            credentials: 'include', // Добавлено для отправки куки
        });

        if (!response.ok) {
            // Если статус ответа не OK, выбрасываем ошибку
            throw new Error('Failed to refresh token');
        }

        const data = await response.json();
        console.log('Token refreshed:', data);
        
        // Обновляем имя пользователя в интерфейсе
        document.getElementById('username').textContent = data.name || "Пользователь"; // Устанавливаем имя пользователя

    } catch (error) {
        console.error('Error refreshing token:', error);
        // Если обновление токена не удалось, перенаправляем пользователя на страницу аутентификации
        window.location.href = 'auth.html';
    }
}

// Добавляем эффект перелива для кнопки
const button = document.querySelector('.start-button');

button.addEventListener('mousemove', (e) => {
    const rect = button.getBoundingClientRect();
    const x = e.clientX - rect.left;
    const y = e.clientY - rect.top;
    
    button.style.background = `
        radial-gradient(circle at ${x}px ${y}px, 
        #87CEEB,
        #4169E1)
    `;
});

button.addEventListener('mouseleave', () => {
    button.style.background = 'linear-gradient(45deg, #4169E1, #87CEEB)';
});
document.querySelector('.start-button').addEventListener('click', function() {
    window.location.href = 'play.html';
  });
// Вызываем функцию проверки токена при загрузке страницы
window.onload = checkToken;
