if (!localStorage.getItem('authToken')) {
    alert('Você não está autenticado. Faça login para continuar.');
    window.location.href = './login.html';
} else {
    console.log('Token encontrado:', localStorage.getItem('authToken'));
}
