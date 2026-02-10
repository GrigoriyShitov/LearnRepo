#ifndef PLAYLIST_H
#define PLAYLIST_H
#include <iostream>
#include <fcntl.h>

struct track {
    std::string title;
    std::string author;
    int duration;
    track(): title(""),author(""),duration(0){}
};

struct doubleLinkedList
{
    track trackInfo;
    struct doubleLinkedList *prev;
    struct doubleLinkedList *next;
    void addTrack(doubleLinkedList*& head, std::string title, std::string author, int duration);
    void removeTrack(doubleLinkedList*& head, std::string title);
    void findTrack(doubleLinkedList*& head, std::string title);
    void printPlayListForward(doubleLinkedList*& head);
    void printPlayListBackward(doubleLinkedList*& head);
    int getTotalDuration(doubleLinkedList*& head);
    void insertAfter(doubleLinkedList*& head, std::string titleAfterInsert,std::string title, std::string author, int duration);
    void shufflePlaylist(doubleLinkedList*& head);
    void fileSavePlaylist(doubleLinkedList* head, int fd);
    void fileReadPlaylist(doubleLinkedList*& head, int fd);

    doubleLinkedList() : trackInfo(), prev(nullptr), next(nullptr) {}
};


#endif