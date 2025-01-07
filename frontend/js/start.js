
const phrases = ['Сапер', 'Онлайн'];
const titleElement = document.querySelector('.title');
let phraseIndex = 0;
let charIndex = 0;
let isDeleting = false;
let typingSpeed = 150;

function type() {
    const currentPhrase = phrases[phraseIndex];
    
    if (isDeleting) {
        titleElement.textContent = currentPhrase.substring(0, charIndex - 1);
        charIndex--;
        typingSpeed = 100;
    } else {
        titleElement.textContent = currentPhrase.substring(0, charIndex + 1);
        charIndex++;
        typingSpeed = 150;
    }

    if (!isDeleting && charIndex === currentPhrase.length) {
        typingSpeed = 2000; // Пауза перед удалением
        isDeleting = true;
    } else if (isDeleting && charIndex === 0) {
        isDeleting = false;
        phraseIndex = (phraseIndex + 1) % phrases.length;
        typingSpeed = 500; // Пауза перед следующим словом
    }

    setTimeout(type, typingSpeed);
}

// Запускаем анимацию при загрузке страницы
window.onload = type;
