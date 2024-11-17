document.getElementById('animalForm').addEventListener('submit', function(event) {
  event.preventDefault(); // Evita o envio do formulário, vamos fazer via JS

  const animal = {
    nome: document.getElementById('nome').value,
    tipo: document.getElementById('tipo').value,
    idade: document.getElementById('idade').value
  };

  insertAnimal(animal);
});

function fetchAnimals() {
    async function fetchAnimals() {
        try {
            const response = await fetch('http://localhost:5802/animais');
            if (!response.ok) {
                throw new Error('Erro ao buscar os animais');
            }
            const animals = await response.json();
            displayAnimals(animals);
        } catch (error) {
            console.error('Erro:', error);
        }
    }
}

function displayAnimals(animals) {
        const container = document.querySelector('.row'); // Seleciona o container dos cards
        container.innerHTML = ''; // Limpa o conteúdo atual

        animals.forEach(animal => {
            const card = document.createElement('div');
            card.className = 'col-md-4';
            card.innerHTML = `
                <div class="animal-card">
                    <img src="${animal.foto || './Imagens/default.jpg'}" alt="${animal.nome}" class="img-fluid">
                    <h5 class="mt-2">${animal.nome}</h5>
                    <p>Espécie: ${animal.especie}</p>
                    <p>Idade: ${animal.idade} anos</p>
                    <p>Cor: ${animal.cor}</p>
                </div>
            `;
            container.appendChild(card);
        });
    }

// Função para atualizar o grid de animais na página
function updateAnimalGrid(animais) {
  const grid = document.getElementById('animalGrid');
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

  function renderAnimals(animals) {
        const container = document.querySelector('.row');
        container.innerHTML = ''; // Limpa o conteúdo existente

        animals.forEach(animal => {
            const card = document.createElement('div');
            card.className = 'col-md-4';
            card.innerHTML = `
                <div class="animal-card">
                    <img src="${animal.image || './Imagens/default.png'}" alt="${animal.nome}" style="width: 100%; height: auto;">
                    <h5 class="mt-2">${animal.nome}</h5>
                </div>
            `;
            container.appendChild(card);
        });
    }


// Chama a função ao carregar a página
    document.addEventListener('DOMContentLoaded', fetchAnimals);
