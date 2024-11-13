function cadastroAnimal() {
    const animalData = {
        nome: document.getElementById('nome').value,
        especie: document.getElementById('especie').value,
        raca: document.getElementById('raca').value,
        idade: parseInt(document.getElementById('idade').value),
        sexo: document.getElementById('sexo').value,
        vacinado: document.getElementById('vacinado').checked,
        vermifugado: document.getElementById('vermifugado').checked,
        castrado: document.getElementById('castrado').checked,
        descricao: document.getElementById('descricao').value,
        status_adocao: document.getElementById('statusAdocao').value,
        data_resgate: document.getElementById('dataResgate').value ? new Date(document.getElementById('dataResgate').value) : null,
        data_cadastro: new Date(), // Current date and time
    };

    // Example of how to make the POST request
    fetch('http://localhost:2000/animais', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${authToken}`, // Add the token for authentication
        },
        body: JSON.stringify(animalData),
    })
    .then(response => response.json())
    .then(data => {
        console.log('Animal registrado:', data);
    })
    .catch(error => {
        console.error('Erro ao registrar animal:', error);
    });
}
