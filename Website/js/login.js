function login(event) {
    event.preventDefault(); // Evita o comportamento padrão do formulário

    const email = document.getElementById('email').value;
    const senha = document.getElementById('senha').value;

    // Enviar a requisição de login para o backend
    fetch('http://localhost:5802/login', {
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
        console.log(data);

        if (data.token) {
            const authToken = data.token;
            const usuarioId = data.usuario_id;
            localStorage.setItem('authToken', authToken);
            localStorage.setItem('usuario_id', usuarioId);
            console.log('Token:', authToken);
            console.log('Usuário:', data.usuario_id);
            console.log('Usuário:', usuarioId);
             alert('authToken settado');
            alert('Token setado com sucesso: ' + data.token);
            console.log('Token setado no localStorage:', localStorage.getItem('authToken'));
            window.location.href = './principal.html'; // Redireciona para a página principal
        } else {
            console.error('token não recebido');
        }
    })
    .catch(error => {
        console.error('Erro ao fazer login:', error);
        alert('Falha no login. Verifique suas credenciais.');
    });
}
