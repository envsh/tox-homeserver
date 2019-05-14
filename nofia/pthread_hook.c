#include <stdlib.h>
#include <dlfcn.h>
#include <pthread.h>
#include <unistd.h>

typedef struct {
    pthread_t *thread;
    void *(*start_routine) (void *);
    void *arg;
} packedpthargs;

typedef int (*pthread_create_t)(pthread_t *thread, const pthread_attr_t *attr,
                                void *(*start_routine) (void *), void *arg);
pthread_create_t pthread_create_f = 0;

// 在nim中实现，在c中调用
extern int nim_pthread_create(packedpthargs* tra);

// 在c中实现，在nim中调用
void nim_pthread_getinfo(packedpthargs* tra) {
    printf("nim_pthread_getinfo %p\n", tra);
    *(tra->thread) = pthread_self();
    void*(*r)(void*) = tra->start_routine;
    void*arg = tra->arg;
    free(tra);

    void* rv = r(arg);
}


static int thread_owner = 0; // 0: nim-startup, 1: goroutines, 2: nim-x11, 3: nng-worker, 4: others
static char* thread_owner_str = "nim-startup";
void pthread_setowner(int owner) {
    thread_owner = owner;
    switch (owner) {
        case 0:
            break;
        case 1:
            thread_owner_str = "goroutines";
            break;
        case 2:
            thread_owner_str = "nim-x11";
            break;
        case 3:
            thread_owner_str = "nng-worker";
            break;
        case 4:
            thread_owner_str = "others";
            break;
        default:
            break;
    }
}

static int pthread_hook_cnter = 0;
int pthread_create(pthread_t *thread, const pthread_attr_t *attr,
                   void *(*start_routine) (void *), void *arg) {
    packedpthargs* tra = (packedpthargs*)calloc(1, sizeof(packedpthargs));
    printf("catched pthread_create %d %s, t=%p, arg=%p, tra=%p, f=%p\n",
           pthread_hook_cnter, thread_owner_str, thread, arg, tra, pthread_create_f);
    tra->thread = thread;
    // tra->attr = attr;
    tra->start_routine = start_routine;
    tra->arg = arg;
    if (thread_owner == 1 || thread_owner == 3) { // for goroutines
        pthread_hook_cnter ++;
        if (pthread_hook_cnter%2== 0) { // 从nim调用而来
            pthread_create_f(thread, attr, start_routine, arg);
        }else{ // 直接从c中某处调用
            nim_pthread_create(tra);
        }
        usleep(100);
    } else {
        pthread_create_t f = (pthread_create_t)dlsym(RTLD_NEXT, "pthread_create");
        return f(thread, attr, start_routine, arg);
    }
    return 0;
}

extern int __pthread_create(pthread_t *thread, const pthread_attr_t *attr,
                            void *(*start_routine) (void *), void *arg);

void init_pthread_hook() {
    pthread_create_f = (pthread_create_t)dlsym(RTLD_NEXT, "pthread_create");
    printf("pthread_create_f=%p\n", pthread_create_f);
}
