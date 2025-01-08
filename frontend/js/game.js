// Проверка и обновление токена
async function checkToken() {
    try {
        const response = await fetch('http://localhost:8080/jwt/checktoken', {
            method: 'GET',
            credentials: 'include',
        });
        if (!response.ok) {
            throw new Error('Token is invalid or missing');
        }
        const data = await response.json();
        console.log('Token is valid:', data);
        document.getElementById('username').textContent = data.name || "Пользователь";
    } catch (error) {
        console.error('Error checking token:', error);
        await refreshToken();
    }
}

async function refreshToken() {
    try {
        const response = await fetch('http://localhost:8080/refresh/refresh', {
            method: 'GET',
            credentials: 'include',
        });
        if (!response.ok) {
            throw new Error('Failed to refresh token');
        }
        const data = await response.json();
        console.log('Token refreshed:', data);
        document.getElementById('username').textContent = data.name || "Пользователь";
    } catch (error) {
        console.error('Error refreshing token:', error);
        window.location.href = 'auth.html';
    }
}

// Эффект перелива для кнопки
const button = document.querySelector('.start-button');
button.addEventListener('mousemove', (e) => {
    const rect = button.getBoundingClientRect();
    const x = e.clientX - rect.left;
    const y = e.clientY - rect.top;
    button.style.background = `radial-gradient(circle at ${x}px ${y}px, #87CEEB, #4169E1)`;
});

button.addEventListener('mouseleave', () => {
    button.style.background = 'linear-gradient(45deg, #4169E1, #87CEEB)';
});

function createBomb() {
    const bomb = document.createElement('div');
    bomb.textContent = '💣';
    bomb.style.position = 'fixed';
    bomb.style.fontSize = '24px';
    bomb.style.left = Math.random() * window.innerWidth + 'px';
    bomb.style.top = '-50px';
    bomb.style.transition = 'transform 1s';
    bomb.style.transform = 'rotate(0deg)';
    bomb.style.cursor = 'pointer';
    document.body.appendChild(bomb);

    // Добавляем обработчик клика для взрыва
    bomb.addEventListener('click', () => {
        explodeBomb(bomb);
    });

    const speed = 2 + Math.random() * 3;
    const rotation = Math.random() * 360;
    const horizontal = (Math.random() - 0.5) * 2;
    let isExploded = false;

    function animate() {
        if (isExploded) return;
        
        const top = parseFloat(bomb.style.top);
        const left = parseFloat(bomb.style.left);
        
        if (top > window.innerHeight) {
            document.body.removeChild(bomb);
            createBomb();
            return;
        }

        bomb.style.top = (top + speed) + 'px';
        bomb.style.left = (left + horizontal) + 'px';
        bomb.style.transform = `rotate(${rotation}deg)`;
        requestAnimationFrame(animate);
    }

    function explodeBomb(bomb) {
        isExploded = true;
        bomb.textContent = '💥';
        bomb.style.fontSize = '48px';
        bomb.style.transition = 'all 0.3s ease-out';
        bomb.style.transform = 'scale(2)';
        bomb.style.opacity = '0';
        
        setTimeout(() => {
            document.body.removeChild(bomb);
            createBomb();
        }, 300);
    }

    animate();
}



for (let i = 0; i < 5; i++) {
    createBomb();
}

window.onload = checkToken;