#include "../include/playlist.h"

void doubleLinkedList::addTrack(doubleLinkedList *&head, std::string title, std::string author, int duration)
{
    doubleLinkedList *newTrack = new doubleLinkedList;
    newTrack->trackInfo.title = title;
    newTrack->trackInfo.author = author;
    newTrack->trackInfo.duration = duration;
    newTrack->next = nullptr;

    if (head == nullptr)
    {
        newTrack->prev = nullptr;
        head = newTrack;
        return;
    }

    doubleLinkedList *current = head;
    while (current->next != nullptr)
    {
        current = current->next;
    }
    current->next = newTrack;
    newTrack->prev = current;
}

void doubleLinkedList::removeTrack(doubleLinkedList *&head, std::string title)
{
    doubleLinkedList *current = head;
    while (current != nullptr && current->trackInfo.title != title)
    {
        current = current->next;
    }
    if (current == nullptr)
        return; // Трек не найден

    if (current->prev != nullptr)
    {
        current->prev->next = current->next;
    }
    else
    {
        head = current->next;
    }
    if (current->next != nullptr)
    {
        current->next->prev = current->prev;
    }
    delete current;
}

void doubleLinkedList::findTrack(doubleLinkedList *&head, std::string title)
{
    doubleLinkedList *current = head;
    while (current != nullptr)
    {
        if (current->trackInfo.title == title)
        {
            std::cout << "Track found: " << current->trackInfo.title << " by "
                      << current->trackInfo.author << ", Duration: " << current->trackInfo.duration << " sec\n";
            return;
        }
        current = current->next;
    }
    std::cout << "Track \"" << title << "\" not found.\n";
}
void doubleLinkedList::printPlayListForward(doubleLinkedList *&head)
{
    doubleLinkedList *current = head;
    if (current == nullptr)
    {
        std::cout << "Playlist is empty.\n";
        return;
    }
    while (current != nullptr)
    {
        std::cout << current->trackInfo.title << " by " << current->trackInfo.author
                  << ", " << current->trackInfo.duration << " sec\n";
        current = current->next;
    }
}

void doubleLinkedList::printPlayListBackward(doubleLinkedList *&head)
{
    doubleLinkedList *current = head;
    if (current == nullptr)
    {
        std::cout << "Playlist is empty.\n";
        return;
    }

    while (current->next != nullptr)
    {
        current = current->next;
    }

    while (current != nullptr)
    {
        std::cout << current->trackInfo.title << " by " << current->trackInfo.author
                  << ", " << current->trackInfo.duration << " sec\n";
        current = current->prev;
    }
}

int doubleLinkedList::getTotalDuration(doubleLinkedList *&head)
{
    doubleLinkedList *current = head;
    int dur = 0;
    while (current != nullptr)
    {
        dur += current->trackInfo.duration;
        current = current->next;
    }
    return dur;
}

void doubleLinkedList::insertAfter(doubleLinkedList *&head, std::string titleAfterInsert, std::string title, std::string author, int duration)
{
    doubleLinkedList *current = head;
    while (current != nullptr && current->trackInfo.title != titleAfterInsert)
    {
        current = current->next;
    }
    if (current == nullptr)
    {
        std::cout << "Track \"" << title << "\" not found for insertion.\n";
        return;
    }

    doubleLinkedList *newTrack = new doubleLinkedList;

    newTrack->trackInfo.title = title;
    newTrack->trackInfo.author = author;
    newTrack->trackInfo.duration = duration;

    newTrack->next = current->next;
    newTrack->prev = current;
    if (current->next != nullptr)
    {
        current->next->prev = newTrack;
    }
    current->next = newTrack;
}
void doubleLinkedList::shufflePlaylist(doubleLinkedList *&head)
{
    if (head == nullptr || head->next == nullptr)
        return;

    int count = 0;
    doubleLinkedList *current = head;

    while (current != nullptr)
    {
        count++;
        current = current->next;
    }

    doubleLinkedList **nodes = new doubleLinkedList *[count];
    current = head;
    for (int i = 0; i < count; i++)
    {
        nodes[i] = current;
        current = current->next;
    }


    srand(time(nullptr));
    for (int i = count - 1; i > 0; i--)
    {
        int j = rand() % (i + 1);
        doubleLinkedList *temp = nodes[i];
        nodes[i] = nodes[j];
        nodes[j] = temp;
    }

    head = nodes[0];
    head->prev = nullptr;
    for (int i = 0; i < count - 1; i++) {
        nodes[i]->next = nodes[i + 1];
        nodes[i + 1]->prev = nodes[i];
    }
    nodes[count - 1]->next = nullptr;

    delete[] nodes;
}
void doubleLinkedList::fileSavePlaylist(doubleLinkedList *head, int fd)
{
    doubleLinkedList* current = head;
    while (current != nullptr) {
        std::string trackData = current->trackInfo.title + " by " + current->trackInfo.author +
                                ", " + std::to_string(current->trackInfo.duration) + " sec\n";
        write(fd, trackData.c_str(), trackData.length());
        current = current->next;
    }
}


void doubleLinkedList::fileReadPlaylist(doubleLinkedList*& head, int fd) {
    char buffer[1024]; // Буфер для чтения
    std::string line;
    ssize_t bytesRead;

    // Читаем файл побайтно
    while ((bytesRead = read(fd, buffer, sizeof(buffer) - 1)) > 0) {
        buffer[bytesRead] = '\0'; // Завершаем строку
        line += buffer; // Добавляем прочитанные данные к строке

        // Обрабатываем строки
        size_t pos = 0;
        while ((pos = line.find('\n')) != std::string::npos) {
            std::string trackLine = line.substr(0, pos);
            line.erase(0, pos + 1);

            // Парсим строку: "название by автор, длительность sec"
            size_t byPos = trackLine.find(" by ");
            size_t commaPos = trackLine.find(", ");
            size_t secPos = trackLine.find(" sec");

            if (byPos == std::string::npos || commaPos == std::string::npos || secPos == std::string::npos) {
                continue; 
            }

            std::string title = trackLine.substr(0, byPos);
            std::string author = trackLine.substr(byPos + 4, commaPos - (byPos + 4));
            std::string durationStr = trackLine.substr(commaPos + 2, secPos - (commaPos + 2));
            int duration = std::stoi(durationStr);

            // Добавляем трек в плейлист
            addTrack(head, title, author, duration);
        }
    }

    // Обрабатываем остаток строки, если он есть
    if (!line.empty()) {
        size_t byPos = line.find(" by ");
        size_t commaPos = line.find(", ");
        size_t secPos = line.find(" sec");

        if (byPos != std::string::npos && commaPos != std::string::npos && secPos != std::string::npos) {
            std::string title = line.substr(0, byPos);
            std::string author = line.substr(byPos + 4, commaPos - (byPos + 4));
            std::string durationStr = line.substr(commaPos + 2, secPos - (commaPos + 2));
            int duration = std::stoi(durationStr);

            addTrack(head, title, author, duration);
        }
    }
}