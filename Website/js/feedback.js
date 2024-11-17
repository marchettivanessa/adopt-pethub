document.addEventListener('DOMContentLoaded', () => {
    const token = localStorage.getItem('authToken');
    if (!token) {
        alert('Você não está autenticado. Faça login para continuar.');
        window.location.href = './login.html';
        return;
    }

    const usuarioId = localStorage.getItem('usuario_id');
    const animalId = localStorage.getItem('animal_id');

    const usuarioIdInt = parseInt(usuarioId, 10);
    const animalIdInt = parseInt(animalId, 10);

    const mensagemField = document.getElementById('mensagem');
    const avaliacaoField = document.getElementById('avaliacao');

    if (mensagemField && avaliacaoField) {
        document.getElementById('feedbackForm').addEventListener('submit', async function (event) {
            event.preventDefault();
            const avaliacaoNumber = parseInt(avaliacaoField.value,10)
            const feedbackData = {
                mensagem: mensagemField.value,
                avaliacao: avaliacaoNumber,
            };
            if (!isNaN(usuarioIdInt)) {
                feedbackData.usuario_id = usuarioIdInt;
            }
            if (!isNaN(animalIdInt)) {
                feedbackData.animal_id = animalIdInt;
            }

            try {
                const response = await fetch('http://localhost:5802/feedback', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${token}`,
                    },
                    body: JSON.stringify(feedbackData),
                });

                if (response.ok) {
                    alert('Feedback enviado com sucesso!');
                    window.location.href = './principal.html';
                } else {
                    const errorData = await response.json();
                    alert(`Erro ao enviar feedback: ${errorData.message || 'Erro desconhecido.'}`);
                }
            } catch (error) {
                console.error('Erro de rede:', error);
                alert('Erro de rede. Tente novamente mais tarde.');
            }
        });
    } else {
        console.error('Campos de mensagem ou avaliação não encontrados no DOM.');
    }
});
