$('#unfollow').on('click', unfollow);
$('#follow').on('click', follow);
$('#edit-user').on('submit', edit);
$('#update-password').on('submit', updatePassword);
$('#delete-user').on('click', deleteUser);

function unfollow() {
  const userId = $(this).data('user-id');
  $(this).prop('disabled', true);

  $.ajax({
    url: `/users/${userId}/unfollow`,
    method: 'POST'
  }).done(function() {
    window.location = `/users/${userId}`;
  }).fail(function() {
    toastr.error('Oops! Erro ao parar de seguir o usuário!');
    $('#unfollow').prop('disabled', false);
  });
}

function follow() {
  const userId = $(this).data('user-id');
  $(this).prop('disabled', true);

  $.ajax({
    url: `/users/${userId}/follow`,
    method: 'POST'
  }).done(function() {
    window.location = `/users/${userId}`;
  }).fail(function() {
    toastr.error('Oops! Erro ao seguir o usuário!');
    $('#follow').prop('disabled', false);
  });
}

function edit(event) {
  event.preventDefault();

  $.ajax({
    url: '/edit-user',
    method: 'PUT',
    data: {
      name: $('#name').val(),
      nick: $('#nick').val(),
      email: $('#email').val(),
    }
  }).done(function() {
    toastr.success('Usuário atualizado com sucesso!');
      setTimeout(function() {
        window.location = "/profile";
      }, 2000);
  }).fail(function() {
    toastr.error('Oops! Erro ao atualizar o usuário!');
  });
}

function updatePassword(event) {
  event.preventDefault();

  if ($('#new-password').val() !== $("#confirm-password").val()) {
    toastr.warning('Oops! As senhas não coincidem!');
  }

  $.ajax({
    url: '/update-password',
    method: 'POST',
    data: {
      currentPassword: $('#current-password').val(),
      newPassword: $('#new-password').val(),
    }
  }).done(function() {
    toastr.success('Senha foi atualizada com sucesso!');
      setTimeout(function() {
        window.location = "/profile";
      }, 2000);
  }).fail(function() {
    toastr.error('Oops! Erro ao atualizar a senha!');
  });
}

function deleteUser() {
  const confirmed = confirm("Tem certeza que deseja apagar a sua conta? Essa é uma ação irreversivel!");

  if (confirmed) {
    $.ajax({
      url: '/delete-user',
      method: 'DELETE',
    }).done(function() {
      alert("Seu usuário foi excluido com sucesso!");
      window.location = "/logout";
    }).fail(function() {
      toastr.error('Oops! Erro ao excluir a usuário!');
    });
  }
}
