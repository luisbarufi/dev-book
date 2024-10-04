$('#new-post').on('submit', createPost);
$('.like-post').on('click', likePost);

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

  const elementClicked = $(event.target).closest('div').parent().closest('div')
  const postId = elementClicked.data('post-id');

  elementClicked.prop('disabled', true);

  $.ajax({
    url: `/posts/${postId}/like`,
    method: `POST`,
  }).done(function() {
    const counterLikes = elementClicked.find('span');
    const totalLikes = parseInt(counterLikes.text());

    counterLikes.text(totalLikes + 1);
  }).fail(function(err) {
    alert("Erro ao curtir publicação!");
  }).always(function() {
    elementClicked.prop('disabled', false);
  });
}
