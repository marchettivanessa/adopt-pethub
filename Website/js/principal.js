document.getElementById('animalForm').addEventListener('submit', function(event) {
  event.preventDefault(); // Evita o envio do formulário, vamos fazer via JS

  const animal = {
    nome: document.getElementById('nome').value,
    tipo: document.getElementById('tipo').value,
    idade: document.getElementById('idade').value
  };

  insertAnimal(animal);
});

// Função para inserir um novo animal no backend
function insertAnimal(animal) {
  fetch('http://localhost:5802/animais', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(animal) // Envia os dados do novo animal
  })
  .then(response => {
    if (!response.ok) {
      throw new Error('Erro ao inserir animal');
    }
    return response.json();
  })
  .then(data => {
    console.log('Animal inserido com sucesso:', data);
    fetchAnimais(); // Atualiza a lista de animais na tela
  })
  .catch(error => {
    console.error('Erro:', error);
  });
}

// Função para buscar os animais
function fetchAnimais() {
  fetch('http://localhost:5802/animais', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
  })
  .then(response => response.json())
  .then(animais => {
    // Chama função para atualizar o grid com os animais
    updateAnimalGrid(animais);
  })
  .catch(error => console.error('Erro ao buscar animais:', error));
}

// Função para atualizar o grid de animais na página
function updateAnimalGrid(animais) {
  const grid = document.getElementById('animalGrid'); // Supondo que você tenha um grid com id animalGrid
  grid.innerHTML = ''; // Limpa o grid

  animais.forEach(animal => {
    const row = document.createElement('tr');
    row.innerHTML = `
      <td>${animal.nome}</td>
      <td>${animal.tipo}</td>
      <td>${animal.idade}</td>
    `;
    grid.appendChild(row);
  });
}

// Chama a função ao carregar a página para exibir os animais
document.addEventListener('DOMContentLoaded', fetchAnimais);
