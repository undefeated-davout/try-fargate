#include <sys/syscall.h>
#include <unistd.h>

#ifndef DOCKER_IMAGE
#define DOCKER_IMAGE "hello-world-test"
#endif

#ifndef DOCKER_GREETING
#define DOCKER_GREETING "こんにちは、Docker runが成功しました！"
#endif

#ifndef DOCKER_ARCH
#define DOCKER_ARCH "amd64"
#endif

const char message[] =
    "挨拶　　　　　　　　：" DOCKER_GREETING "\n"
    "Docker image名　　　：" DOCKER_IMAGE "\n"
    "Dokcerアーキテクチャ：" DOCKER_ARCH "\n";

int main()
{
  syscall(SYS_write, STDOUT_FILENO, message, sizeof(message) - 1);
  return 0;
}
