package main

import "fmt"

type song struct {
	name     string
	artist   string
	next     *song
	previous *song
}

type playlist struct {
	name       string
	head       *song
	nowPlaying *song
	tail       *song
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
func (p *playlist) forward() {
	if p.nowPlaying == nil {
		p.startPlaying()
		return
	}

	p.nowPlaying = p.nowPlaying.next
}

func (p *playlist) backward() {
	if p.nowPlaying == nil {
		p.startPlaying()
		return
	}

	p.nowPlaying = p.nowPlaying.previous
}

func (p *playlist) addSong(name string, artist string) error {
	newSong := &song{
		name:   name,
		artist: artist,
	}

	if p.head == nil {
		p.head = newSong
		p.tail = newSong
	} else {
		currentSong := p.tail
		currentSong.next = newSong
		newSong.previous = p.tail
	}
	p.tail = newSong
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
	p.addSong("August", "Taylor Swift")
	p.showAllSongs()

	fmt.Println("start playing")
	p.startPlaying()
	fmt.Print("now playing :")
	p.currentSong()

	p.forward()
	fmt.Print("now playing :")
	p.currentSong()

	p.backward()
	fmt.Print("now playing :")
	p.currentSong()
}
