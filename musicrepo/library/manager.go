package library

import (
	"errors"
)

type Music struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	musics []Music
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]Music, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (musicResult *Music, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range.")
	}

	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *Music {
	if len(m.musics) == 0 {
		return nil
	}

	for _, n := range m.musics {
		if n.Name == name {
			return &n
		}
	}

	return nil
}

func (m *MusicManager) Add(music *Music) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *Music {
	if index < 0 || index >= len(m.musics) {
		return nil
	}

	removedMusic := &m.musics[index]

	if index < len(m.musics)-1 {
		m.musics = append(m.musics[:index-1],m.musics[index+1:]...)
	}else if index == 0 {
		m.musics = make([]Music,0)
	}else{
		m.musics = m.musics[:index-1]
	}

	return removedMusic
}

func (m *MusicManager) RemoveByName(name string) *Music{
	index := -1
	isFind := false
		
	for _, n := range m.musics {
		index ++
		if n.Name == name {			
			isFind = true	
			break
		}		
	}
	
	if isFind == false{
		return nil
	}
	
	return m.Remove(index)
	
}
