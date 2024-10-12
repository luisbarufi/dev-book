$('#unfollow').on('click', unfollow);
$('#follow').on('click', follow);
$('#edit-user').on('submit', edit);

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
