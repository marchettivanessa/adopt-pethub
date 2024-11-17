function login(event) {
    event.preventDefault();

    const email = document.getElementById('email').value;
    const senha = document.getElementById('senha').value;

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
            console.log('Token setado no localStorage:');// maintaining this log for debugging purposes
            window.location.href = './principal.html';
        } else {
            console.error('token não recebido');
        }
    })
    .catch(error => {
        console.error('Erro ao fazer login:', error);
        alert('Falha no login. Verifique suas credenciais.');
    });
}
