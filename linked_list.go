package main

import "fmt"

type song struct {
	name   string
	artist string
	next   *song
}

type playlist struct {
	name       string
	head       *song
	nowPlaying *song
}

func createPlaylist(name string) *playlist {
	return &playlist{
		name: name,
	}
}

func (p *playlist) startPlaying() {
	p.nowPlaying = p.head
}

func (p playlist) currentSong() {
	fmt.Printf("%+v\n", p.nowPlaying)
}
func (p *playlist) skip() {
	if p.nowPlaying == nil {
		p.nowPlaying = p.head
		return
	}

	nowPlaying := p.nowPlaying
	nowPlaying = nowPlaying.next
	p.nowPlaying = nowPlaying
}

func (p *playlist) addSong(name string, artist string) error {
	newSong := &song{
		name:   name,
		artist: artist,
	}

	if p.head == nil {
		p.head = newSong
	} else {
		currentSong := p.head
		for currentSong.next != nil {
			currentSong = currentSong.next
		}
		currentSong.next = newSong
	}

	return nil
}

func (p *playlist) showAllSongs() {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("no song in the playlist :", p.name)
		return
	}

	fmt.Printf("%+v\n", *currentNode)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}
}

func main() {
	p := createPlaylist("this is taylor swift")
	p.showAllSongs()
	p.addSong("All Too Well", "Taylor Swift")
	p.addSong("Enchanted", "Taylor Swift")
	p.showAllSongs()

	fmt.Println("start playing")
	p.startPlaying()
	fmt.Print("now playing :")
	p.currentSong()

	p.skip()
	fmt.Print("now playing :")
	p.currentSong()
}
