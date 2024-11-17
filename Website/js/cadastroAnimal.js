document.addEventListener('DOMContentLoaded', () => {
    const token = localStorage.getItem('authToken');
    if (!token) {
        alert('Você não está autenticado. Faça login para continuar.');
        window.location.href = './login.html';
        return;
    }

    function cadastroAnimal(event) {
        event.preventDefault();

        const nome = document.getElementById('nome').value;
        const especie = document.getElementById('especie').value;
        const raca = document.getElementById('raca').value;
        const idade = document.getElementById('idade').value;
        const sexo = document.getElementById('sexo').value;
        const descricao = document.getElementById('descricao').value;

        if (!nome || !especie || !raca || !idade || !sexo || !descricao) {
            alert('Preencha todos os campos obrigatórios');
            return;
        }

        const formData = new FormData();
        formData.append('nome', nome);
        formData.append('especie', especie);
        formData.append('raca', raca);
        formData.append('idade', idade);
        formData.append('sexo', sexo);
        formData.append('descricao', descricao);
        formData.append('vacinado', document.getElementById('vacinado').checked);
        formData.append('castrado', document.getElementById('castrado').checked);
        formData.append('vermifugado', document.getElementById('vermifugado').checked);

        const fotoInput = document.getElementById('foto_url');
        if (fotoInput.files.length > 0) {
            formData.append('foto_url', fotoInput.files[0]);
        }

        fetch('http://localhost:5802/animais', {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`,
            },
            body: formData,
        })
            .then(response => {
                if (!response.ok) {
                    return response.json().then(errData => {
                        throw new Error(errData.error || 'Erro ao cadastrar animal');
                    });
                }
                return response.json();
            })
            .then(data => {
                alert('Animal cadastrado com sucesso!');
            })
            .catch(error => {
                console.error('Erro ao cadastrar animal:', error.message);
                alert(`Erro ao enviar os dados: ${error.message}`);
            });
    }

    document.getElementById('formCadastroAnimal').addEventListener('submit', cadastroAnimal);
});
