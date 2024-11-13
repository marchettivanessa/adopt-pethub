function login(event) {
    event.preventDefault(); // Evita o comportamento padrão do formulário

    const email = document.getElementById('email').value;
    const senha = document.getElementById('senha').value;

    // Enviar a requisição de login para o backend
    fetch('http://localhost:2000/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            email: email,
            senha: senha,
        }),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Falha no login');
        }
        return response.json();
    })
    .then(data => {
        const authToken = data.token; // Captura o token da resposta
        console.log('Token:', authToken);
        localStorage.setItem('authToken', authToken); // Armazena o token no localStorage
        window.location.href = '/principal'; // Redireciona para a página principal
    })
    .catch(error => {
        console.error('Erro ao fazer login:', error);
        alert('Falha no login. Verifique suas credenciais.');
    });
}
