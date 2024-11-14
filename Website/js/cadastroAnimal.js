function cadastroAnimal(event) {
    event.preventDefault(); // Evita o recarregamento da página

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
        data_cadastro: new Date(),
    };

    console.log(animalData); // Para verificar os dados
    console.log("banana"); // Confirmar que a função foi chamada

    fetch('http://localhost:5802/animais', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${authToken}`, // Add the token for authentication
        },
        body: JSON.stringify(animalData),
    })
    .then(response => {
        if (!response.ok) throw new Error(`Erro: ${response.status}`);
        return response.json();
    })
    .then(data => {
        console.log('Animal registrado:', data);
    })
    .catch(error => {
        console.error('Erro ao registrar animal:', error);
    });
}
