document.getElementById('register-form').addEventListener('submit', async (event) => {
    event.preventDefault(); // Impede o formulário de recarregar a página

    const userData = {
        nome: document.getElementById('name').value,
        email: document.getElementById('email').value,
        senha: document.getElementById('password').value,
        telefone: document.getElementById('phone').value,
        endereco: document.getElementById('address').value
    };

    try {
        const response = await fetch('http://localhost:5802/usuarios', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(userData)
        });

        if (response.ok) {
            alert('Cadastro realizado com sucesso!');
            window.location.href = 'login.html'; // Redireciona para a página de login após o cadastro
        } else {
            alert('Erro ao realizar cadastro.');
        }
    } catch (error) {
        console.error('Erro de rede:', error);
        alert('Erro de rede. Tente novamente mais tarde.');
    }
});
