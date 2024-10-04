$('#new-post').on('submit', createPost);
$(document).on('click', '.like-post', likePost);
$(document).on('click', '.dislike-post', dislikePost);

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
