document.addEventListener('DOMContentLoaded', () => {
    // Obtém o usuario_id do local storage
    const usuarioId = localStorage.getItem('usuario_id');
    if (usuarioId) {
        document.getElementById('usuario_id').value = usuarioId;
    } else {
        console.error("usuario_id não encontrado no localStorage.");
    }

    // Obtém o animal_id da URL, se estiver presente
    const urlParams = new URLSearchParams(window.location.search);
    const animalId = urlParams.get('animal_id');
    if (animalId) {
        document.getElementById('animal_id').value = animalId;
    } else {
        console.error("animal_id não encontrado na URL.");
    }
});

document.getElementById('feedbackForm').addEventListener('submit', async function (event) {
    event.preventDefault();

    const feedbackData = {
        usuario_id: document.getElementById('usuario_id').value,
        animal_id: document.getElementById('animal_id').value,
        mensagem: document.getElementById('mensagem').value,
        avaliacao: document.getElementById('avaliacao').value,
    };

    try {
        const response = await fetch('http://localhost:5802/feedback', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`,  // Certifique-se de que 'token' está definido
            },
            body: JSON.stringify(feedbackData),
        });

        if (response.ok) {
            alert('Feedback enviado com sucesso!');
            window.location.href = './principal.html';
        } else {
            alert('Erro ao enviar feedback.');
        }
    } catch (error) {
        console.error('Erro de rede:', error);
        alert('Erro de rede. Tente novamente mais tarde.');
    }
});
