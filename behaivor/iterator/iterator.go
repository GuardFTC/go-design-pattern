// Package iterator @Author:冯铁城 [17615007230@163.com] 2025-09-26 14:27:18
package iterator

import "fmt"

// Song 音乐
type Song struct {
	Name string
}

func (s *Song) Play() {
	fmt.Printf("播放音乐：%s\n", s.Name)
}

// NewSong 创建音乐
func NewSong(name string) *Song {
	return &Song{
		Name: name,
	}
}

// SongIterator 音乐迭代器
type SongIterator interface {
	HasNext() bool
	Next() *Song
}

// AscSongIterator 升序音乐迭代器
type AscSongIterator struct {
	Songs []*Song
	Index int
}

// HasNext 迭代器是否有下一个元素
func (a *AscSongIterator) HasNext() bool {
	if a.Index < len(a.Songs) {
		return true
	}
	return false
}

// Next 获取下一个元素
func (a *AscSongIterator) Next() *Song {
	if a.HasNext() {
		song := a.Songs[a.Index]
		a.Index++
		return song
	}
	return nil
}

// NewAscSongIterator 创建升序音乐迭代器
func NewAscSongIterator(songs []*Song) SongIterator {
	return &AscSongIterator{
		Songs: songs,
		Index: 0,
	}
}

// DescSongIterator 降序音乐迭代器
type DescSongIterator struct {
	Songs []*Song
	Index int
}

// HasNext 是否有下一个
func (d *DescSongIterator) HasNext() bool {
	if d.Index >= 0 {
		return true
	}
	return false
}

// Next 获取下一个元素
func (d *DescSongIterator) Next() *Song {
	if d.HasNext() {
		song := d.Songs[d.Index]
		d.Index--
		return song
	}
	return nil
}

// NewDescSongIterator 创建降序音乐迭代器
func NewDescSongIterator(songs []*Song) SongIterator {
	return &DescSongIterator{
		Songs: songs,
		Index: len(songs) - 1,
	}
}

// SongCollection 音乐集合
type SongCollection interface {
	GetSongIterator() SongIterator
}

// AscSongCollection 正序播放音乐歌单
type AscSongCollection struct {
	Songs []*Song
}

// GetSongIterator 获取正序播放音乐迭代器
func (a *AscSongCollection) GetSongIterator() SongIterator {
	return NewAscSongIterator(a.Songs)
}

// NewAscSongCollection 创建正序播放音乐歌单
func NewAscSongCollection(songs []*Song) SongCollection {
	return &AscSongCollection{
		Songs: songs,
	}
}

// DescSongCollection 逆序播放音乐歌单
type DescSongCollection struct {
	Songs []*Song
}

// GetSongIterator 获取逆序播放音乐迭代器
func (d *DescSongCollection) GetSongIterator() SongIterator {
	return NewDescSongIterator(d.Songs)
}

// NewDescSongCollection 创建逆序播放音乐歌单
func NewDescSongCollection(songs []*Song) SongCollection {
	return &DescSongCollection{
		Songs: songs,
	}
}
