$('#form-registration').on('submit', createUser);

function createUser(event) {
  event.preventDefault();

  if ($('#password').val() != $('#password-confirmation').val()) {
    toastr.warning('As senhas não coincidem!');

    return 
  }

  $.ajax({
    url: '/users',
    method: 'POST',
    data: {
      name: $('#name').val(),
      email: $('#email').val(),
      nick: $('#nick').val(),
      password: $('#password').val()
    }
  }).done(function() {
    toastr.success('Usuário cadastrado com sucesso!');

    return $.ajax({
      url: '/login',
      method: 'POST',
      data: {
        email: $('#email').val(),
        password: $('#password').val(),
      }
    }).done(function() {
      window.location = "/home";
    }).fail(function() {
      toastr.error('Erro ao cadastrar usuário!');
    });

  }).fail(function() {
    toastr.error('Erro ao cadastrar usuário!');
  });
}
