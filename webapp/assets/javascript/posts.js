$('#new-post').on('submit', createPost);
$(document).on('click', '.like-post', likePost);
$(document).on('click', '.dislike-post', dislikePost);
$('#update-post').on('click', updatePost);
$('.delete-post').on('click', deletePost);

function createPost(event) {
  event.preventDefault();

  $.ajax({
    url: "/posts",
    method: 'POST',
    data: {
      title: $('#title').val(),
      content: $('#content').val()
    }
  }).done(function() {
    window.location = "/home";
  }).fail(function(err) {
    alert("Erro ao criar a publicação!");
  });
}

function likePost(event) {
  event.preventDefault();

  const elementClicked = $(event.target);
  const postId = elementClicked.closest('div').parent().closest('div').data('post-id');

  elementClicked.prop('disabled', true);

  $.ajax({
    url: `/posts/${postId}/like`,
    method: `POST`,
  }).done(function() {
    const counterLikes = elementClicked.next('span');
    const totalLikes = parseInt(counterLikes.text());

    counterLikes.text(totalLikes + 1);

    elementClicked.addClass('dislike-post');
    elementClicked.addClass('text-danger');
    elementClicked.removeClass('like-post');
  }).fail(function(err) {
    alert("Erro ao curtir publicação!");
  }).always(function() {
    elementClicked.prop('disabled', false);
  });
}

function dislikePost(event) {
  event.preventDefault();

  const elementClicked = $(event.target);
  const postId = elementClicked.closest('div').parent().closest('div').data('post-id');

  elementClicked.prop('disabled', true);

  $.ajax({
    url: `/posts/${postId}/dislike`,
    method: `POST`,
  }).done(function() {
    const counterLikes = elementClicked.next('span');
    const totalLikes = parseInt(counterLikes.text());

    counterLikes.text(totalLikes - 1);

    elementClicked.removeClass('dislike-post');
    elementClicked.removeClass('text-danger');
    elementClicked.addClass('like-post');
  }).fail(function(err) {
    alert("Erro ao descurtir publicação!");
  }).always(function() {
    elementClicked.prop('disabled', false);
  });
}

function updatePost() {
  $(this).prop('disabled', true);

  const postId = $(this).data('post-id');
  
  $.ajax({
    url: `/posts/${postId}`,
    method: 'PUT',
    data: {
      title: $('#title').val(),
      content: $('#content').val()
    }
  }).done(function() {
    alert("Publicação editada com sucesso!");
  }).fail(function() {
    alert("Erro ao editar a publicação!");
  }).always(function() {
    $('#update-post').prop('disabled', false);
  });
}

function deletePost(event) {
  event.preventDefault();

  const elementClicked = $(event.target);
  const post = elementClicked.closest('div').parent().closest('div')
  const postId = post.data('post-id');

  elementClicked.prop('disabled', true);

  $.ajax({
    url: `/posts/${postId}`,
    method: 'DELETE',
  }).done(function() {
    post.fadeOut('slow', function() {
      $(this).remove()
    });
  }).fail(function() {
    alert("Erro ao excluir a publicação!");
  }).always(function() {
    elementClicked.prop('disabled', true);
  });
}
