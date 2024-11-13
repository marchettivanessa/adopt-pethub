function logout() {
    localStorage.removeItem('authToken');
    window.location.href = 'login.html';
}