function logout() {
    alert('Você foi deslogado com sucesso!');
    localStorage.removeItem('authToken');
    window.location.href = './login.html';
}
