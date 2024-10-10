package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requests"
)

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Nick      string    `json:"nick"`
	CreatedAt time.Time `json:"created_at"`
	Followers []User    `json:"followers"`
	Following []User    `json:"following"`
	Posts     []Post    `json:"posts"`
}

// UserData makes 4 API requests to assemble the user data.
func UserData(userId uint64, r *http.Request) (User, error) {
	channelUser := make(chan User)
	channelFollowers := make(chan []User)
	channelFollowing := make(chan []User)
	channelPosts := make(chan []Post)

	go FetchUserData(channelUser, userId, r)
	go FetchFollowers(channelFollowers, userId, r)
	go FetchFollowing(channelFollowing, userId, r)
	go FetchPosts(channelPosts, userId, r)

	var (
		user      User
		followers []User
		following []User
		posts     []Post
	)

	for i := 0; i < 4; i++ {
		select {
		case userLoaded := <-channelUser:
			if userLoaded.ID == 0 {
				return User{}, errors.New("erro ao buscar o usuário")
			}

			user = userLoaded

		case followersLoaded := <-channelFollowers:
			if followersLoaded == nil {
				return User{}, errors.New("erro ao buscar o seguidores")
			}

			followers = followersLoaded

		case followingLoaded := <-channelFollowing:
			if followingLoaded == nil {
				return User{}, errors.New("erro ao buscar quem o usuário está seguindo")
			}

			following = followingLoaded

		case postsLoaded := <-channelPosts:
			if postsLoaded == nil {
				return User{}, errors.New("erro ao buscar as publicações")
			}

			posts = postsLoaded
		}
	}

	user.Followers = followers
	user.Following = following
	user.Posts = posts

	return user, nil
}

func FetchUserData(channel chan<- User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.ApiUrl, userId)

	response, err := requests.SendAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		channel <- User{}
		return
	}

	channel <- user
}

func FetchFollowers(channel chan<- []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.ApiUrl, userId)

	response, err := requests.SendAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if err := json.NewDecoder(response.Body).Decode(&followers); err != nil {
		channel <- nil
		return
	}

	if followers == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- followers
}

func FetchFollowing(channel chan<- []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.ApiUrl, userId)

	response, err := requests.SendAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var following []User
	if err := json.NewDecoder(response.Body).Decode(&following); err != nil {
		channel <- nil
		return
	}

	channel <- following
}

func FetchPosts(channel chan<- []Post, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/posts", config.ApiUrl, userId)

	response, err := requests.SendAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var posts []Post
	if err := json.NewDecoder(response.Body).Decode(&posts); err != nil {
		channel <- nil
		return
	}

	if posts == nil {
		channel <- make([]Post, 0)
		return
	}

	channel <- posts
}
