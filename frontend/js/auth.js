function toggleForm() {
    document.querySelector('.container').classList.toggle('show-register');
    // Скрываем сообщения об ошибках при переключении форм
    document.getElementById('loginError').style.display = 'none';
    document.getElementById('registerError').style.display = 'none';
}

function showError(elementId, message) {
    const errorElement = document.getElementById(elementId);
    errorElement.textContent = message;
    errorElement.style.display = 'block';
}

// Обработка формы входа
document.getElementById('loginForm').addEventListener('submit', async function (e) {
    e.preventDefault();

    const formData = new FormData(this);
    const data = {
        email: formData.get('email'),
        password: formData.get('password'),
    };

    try {
        const response = await fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include', // Добавьте этот параметр для включения куки
            body: JSON.stringify(data),
        });

        if (response.ok) {
            const result = await response.json();
            console.log('Успешный вход:', result);
            // Перенаправление на game.html
            window.location.href = 'game.html';
        } else {
            showError('loginError', 'Неправильный email или пароль');
        }
    } catch (error) {
        console.error('Ошибка при отправке запроса:', error);
        showError('loginError', 'Ошибка при попытке входа');
    }
});

// Обработка формы регистрации
document.getElementById('registerForm').addEventListener('submit', async function (e) {
    e.preventDefault();

    const formData = new FormData(this);
    const data = {
        name: formData.get('name'),
        email: formData.get('email'),
        password: formData.get('password'),
        confirmPassword: formData.get('confirmPassword'),
    };

    if (data.password !== data.confirmPassword) {
        showError('registerError', 'Пароли не совпадают');
        return;
    }

    try {
        const response = await fetch('http://localhost:8080/registration', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include', // Добавьте этот параметр для включения куки
            body: JSON.stringify(data),
        });

        if (response.ok) {
            const result = await response.json();
            console.log('Успешная регистрация:', result);
            // Перенаправление на game.html
            window.location.href = 'game.html';
        } else {
            const errorData = await response.json();
            showError('registerError', 'Пользователь с таким email уже существует');
        }
    } catch (error) {
        console.error('Ошибка при отправке запроса:', error);
        showError('registerError', 'Ошибка при регистрации');
    }
});
