#include <sys/socket.h>
#include <netinet/in.h>
#include <unistd.h>
#include <cstring>

int main() {
    int s = socket(AF_INET, SOCK_STREAM, 0);
    sockaddr_in addr{AF_INET, htons(8888), INADDR_ANY};
    bind(s, (sockaddr*)&addr, sizeof(addr));
    listen(s, 10);
    while(1) {
        int c = accept(s, 0, 0);
        char buf[1024];
        read(c, buf, 1024);
        write(c, "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\n\r\nHello World!", 57);
        close(c);
    }
}
