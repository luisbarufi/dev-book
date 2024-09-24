$('#form-registration').on('submit', createUser);

function createUser(event) {
  event.preventDefault();

  if ($('#password').val() != $('#password-confirmation').val()) {
    alert("A senhas não coincidem!");

    return 
  }

  $.ajax({
    url: "/users",
    method: 'POST',
    data: {
      name: $('#name').val(),
      email: $('#email').val(),
      nick: $('#nick').val(),
      password: $('#password').val()
    }
  }).done(function() {
    alert("Usuário cadastrado com sucesso!");
  }).fail(function(err) {
    console.log(err);
    alert("Erro ao cadastrar usuário!");
  });
}
