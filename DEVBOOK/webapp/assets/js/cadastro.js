$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(evento){
    evento.preventDefault();
    console.log("Dentro da função usuário")

    if ( $('#senha').val() != $('#confirmar-senha').val() ){
        alert("As senhas não coincidem!")
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
            senha: $('#senha').val()
        }
    }).done(function() { // statusCode maior que 200
        alert("Usuário cadastro com sucesso!");
    }).fail(function(erro) { 
        console.log(erro);// statusCode maior que 400
        alert("Erro ao cadastrar o usuário!");
    });
}