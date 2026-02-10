#include <iostream>
#include "../include/playlist.h"

int main(){
    std::cout<<"Hello world"<< std::endl;
    doubleLinkedList* playlist = nullptr;


    // Добавляем треки
    playlist->addTrack(playlist, "Song1", "Artist1", 180);
    playlist->addTrack(playlist, "Song2", "Artist2", 200);
    playlist->addTrack(playlist, "Song3", "Artist3", 150);

    // Выводим плейлист
    std::cout << "Playlist forward:\n";
    playlist->printPlayListForward(playlist);

    std::cout << "Playlist backward:\n";
    playlist->printPlayListBackward(playlist);

    // Ищем трек
    playlist->findTrack(playlist, "Song2");

    // Общая длительность
    std::cout << "Total duration: " << playlist->getTotalDuration(playlist) << " sec\n";

    playlist->shufflePlaylist(playlist);
    std::cout << "After shuffling:\n";
    playlist->printPlayListForward(playlist);


    int fd = open("C:\\Users\\Григорий\\Desktop\\courses\\4day\\out\\playlist.txt", O_WRONLY | O_CREAT | O_TRUNC, 0644);
    if (fd != -1) {
        playlist->fileSavePlaylist(playlist, fd);
        close(fd);
    }

    while (playlist != nullptr) {
        doubleLinkedList* temp = playlist;
        playlist = playlist->next;
        delete temp;
    }



    // Читаем плейлист из файла
    int fdR = open("C:\\Users\\Григорий\\Desktop\\courses\\4day\\include\\srcPlaylists.txt", O_RDONLY);
    if (fdR != -1) {
        playlist->fileReadPlaylist(playlist, fdR);
        close(fdR);
    } else {
        std::cout << "Failed to open file.\n";
        return 1;
    }

    // Выводим плейлист
    std::cout << "Playlist after reading from file:\n";
    playlist->printPlayListForward(playlist);
    
    while (playlist != nullptr) {
        doubleLinkedList* temp = playlist;
        playlist = playlist->next;
        delete temp;
    }
}

